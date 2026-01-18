package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/dataaccess"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

const (
	InvalidTopicName = "Invalid Topic Name: must consist of 3-50 alphannumeric characters, _ or - (without spacing)"
	InvalidDescriptionPattern = `Invalid Description Pattern: exceeds max character limit or contains invalid symbol(s)`
)

var validTopicNamePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,50}$`)
var validTopicDescriptionPattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{0,500}$`)

func CreateTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get topic name from request 
	topic := &models.Topic{}

	err := json.NewDecoder(r.Body).Decode(topic)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.CreateTopic"))
	}

	// Topic name validation
	valid := validTopicNamePattern.MatchString(topic.Name)
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		return &models.TopicsResult{
			Success: false,
			Error: InvalidTopicName,
		}, nil
	}

	valid = validTopicDescriptionPattern.MatchString(topic.Description)
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		return &models.TopicsResult{
			Success: false,
			Error: InvalidDescriptionPattern,
		}, nil
	}

	// Check if topic, user exists
	exist, err := dataaccess.CheckTopicExistByTopicName(db, topic.Name)	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
    } 

	if exist {
		w.WriteHeader(http.StatusConflict)
		return &models.TopicsResult{
			Success: false,
			Error: fmt.Sprintf("Topic name taken: %s", topic.Name),
		}, nil
	}

	exist, err = dataaccess.CheckUserExistByUserID(db, topic.UserID)	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
    } 

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return nil, errors.Wrap(err, fmt.Sprintf("User does not exist: %d", topic.UserID))
	}

	// Get username
	username, err := dataaccess.GetUsernameByUserID(db, topic.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
	}

	// Create and insert a new topic 
	t := time.Now()

	res, err := dataaccess.InsertNewTopic(db, topic.Name, topic.UserID, username, topic.Description, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
	}

	topic.ID = int(id)	
	topic.Author = username
	topic.CreatedAt = t

	return &models.TopicsResult{
		Success: true,
		Topic: topic,
	}, nil
}

func ViewTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get topic id from request 
	id_str := chi.URLParam(r, "id")
	id, err := strconv.Atoi(id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.ViewTopic"))
	}

	// Check if topic exists
	topic, err := dataaccess.GetTopicByTopicID(db, id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", id))
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewTopic"))
    }	

	return &models.TopicsResult{
		Success: true,
		Topic: topic,
	}, nil
}

func ViewTopics(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Check if any topic exists
	topics, err := dataaccess.GetTopics(db)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.TopicListResult{
				Success: true,
				Topics: nil,
			}, nil
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewTopics"))
    }	

	return &models.TopicListResult{
		Success: true,
		Topics: topics,
	}, nil
}

func EditTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get topic id, new description and user id from request 
	topic_id_str := chi.URLParam(r, "id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.EditTopic"))
	}

	type Body struct {
		UserID int `json:"user_id"`
		Description string `json:"description"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.EditTopic"))
	}
	userID := body.UserID
	description := body.Description

	// Check if topic exists
	topic, err := dataaccess.GetTopicByTopicID(db, topic_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditTopic"))
    }	

	// Access control
	if userID != topic.UserID {
		w.WriteHeader(http.StatusForbidden)
		return &models.TopicsResult{
				Success: false,
				Error: "You don't have the right to edit this topic",
			}, nil	
	}

	// Form inputs validation (only description is editable)
	valid := validTopicDescriptionPattern.MatchString(description)
	if !valid {
		w.WriteHeader(http.StatusBadRequest) 
		return &models.TopicsResult{
			Success: false,
			Error: InvalidDescriptionPattern,
		}, nil
	}

	// Update and return modified topic
	topic.Description = description
	_, err = dataaccess.UpdateTopicDescription(db, topic_id, description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditTopic"))	
	}

	return &models.TopicsResult{
		Success: true,
		Topic: topic,
	}, nil
}

func DeleteTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get topic id, user id from request 
	topic_id_str := chi.URLParam(r, "id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.DeleteTopic"))
	}

	type Body struct {
		UserID int `json:"user_id"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.DeleteTopic"))
	}
	userID := body.UserID

	// Check if topic exists
	topic, err := dataaccess.GetTopicByTopicID(db, topic_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteTopic"))
    }	

	// Access control
	if userID != topic.UserID {
		w.WriteHeader(http.StatusForbidden)
		return &models.TopicsResult{
				Success: false,
				Error: "You does not have the right to delete this topic",
			}, nil	
	}

	// Delete topic
	res, err := dataaccess.DeleteTopicByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    	return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteTopic"))
	}

	rows, errRA := res.RowsAffected()
	if errRA != nil {
		w.WriteHeader(http.StatusInternalServerError)
    	return nil, errors.Wrap(errRA, fmt.Sprintf(ErrDB, "api.DeleteTopic"))
	}
	if rows != 1 {
		w.WriteHeader(http.StatusInternalServerError)
    	return nil, errors.Errorf("api.DeleteTopic: expected to delete 1 row, deleted %d", rows)
	}

	return &models.TopicsResult{
		Success: true,
	}, nil
}
