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
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

const (
	InvalidTopicName = "Invalid Topic Name: must consist of 3-16 alphannumeric characters, _ or -"
	InvalidDescriptionPattern = `Invalid Description Pattern: must consist of 0-60 characters in a-zA-Z0-9 .,!?'"()_\-`
)

var validTopicNamePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,16}$`)
var validTopicDescriptionPattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{0,60}$`)

func CreateTopic(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.CreateTopic"))
	}
	defer db.Close()
	
	// Get topic name from request 
	topic := &models.Topic{}

	err = json.NewDecoder(r.Body).Decode(topic)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.CreateTopic"))
	}

	// Topic name validation
	valid := validTopicNamePattern.MatchString(topic.Name)
	if !valid {
		return &models.TopicsResult{
			Success: false,
			Error: InvalidTopicName,
		}, nil
	}

	valid = validTopicDescriptionPattern.MatchString(topic.Description)
	if !valid {
		return &models.TopicsResult{
			Success: false,
			Error: InvalidDescriptionPattern,
		}, nil
	}

	// Check if topic, user exists
	exist, err := dataaccess.CheckTopicExistByTopicName(db, topic.Name)	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
    } 

	if exist {
		return &models.TopicsResult{
			Success: false,
			Error: fmt.Sprintf("Topic name taken: %s", topic.Name),
		}, nil
	}

	exist, err = dataaccess.CheckUserExistByUserID(db, topic.UserID)	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
    } 

	if !exist {
		return &models.TopicsResult{
			Success: false,
			Error: fmt.Sprintf("User does not exist: %d", topic.UserID),
		}, nil
	}

	// Create and insert a new topic 
	t := time.Now()

	res, err := dataaccess.InsertNewTopic(db, topic.Name, topic.UserID, topic.Description, t)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
	}

	topic.ID = int(id)	
	topic.CreatedAt = t

	return &models.TopicsResult{
		Success: true,
		Topic: topic,
	}, nil
}

func ViewTopic(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.ViewTopic"))
	}
	defer db.Close()
	
	// Get topic id from request 
	id_str := chi.URLParam(r, "id")
	id, err := strconv.Atoi(id_str)
    if err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.ViewTopic"))
	}

	// Check if topic exists
	topic, err := dataaccess.GetTopicByTopicID(db, id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.TopicsResult{
				Success: false,
				Error: fmt.Sprintf("Topic does not exist: %d", id),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewTopic"))
    }	

	return &models.TopicsResult{
		Success: true,
		Topic: topic,
	}, nil
}

func ViewTopics(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.ViewTopics"))
	}
	defer db.Close()
	
	// Check if any topic exists
	topics, err := dataaccess.GetTopics(db)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.TopicListResult{
				Success: true,
				Topics: nil,
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewTopics"))
    }	

	return &models.TopicListResult{
		Success: true,
		Topics: topics,
	}, nil
}

func EditTopic(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.EditTopic"))
	}
	defer db.Close()

	// Get topic id, new description and user id from request 
	topic_id_str := chi.URLParam(r, "id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
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
            return &models.TopicsResult{
				Success: false,
				Error: fmt.Sprintf("Topic does not exist: %d", topic_id),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditTopic"))
    }	

	// Access control
	if userID != topic.UserID {
		return &models.TopicsResult{
				Success: false,
				Error: fmt.Sprintf("User: %d does not have right to edit this topic: %d", userID, topic_id),
			}, nil	
	}

	// Form inputs validation (only description is editable)
	valid := validTopicDescriptionPattern.MatchString(description)
	if !valid {
		return &models.TopicsResult{
			Success: false,
			Error: InvalidDescriptionPattern,
		}, nil
	}

	// Update and return modified topic
	topic.Description = description
	_, err = dataaccess.UpdateTopicDescription(db, topic_id, description)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditTopic"))	
	}

	return &models.TopicsResult{
		Success: true,
		Topic: topic,
	}, nil
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.DeleteTopic"))
	}
	defer db.Close()

	// Get topic id, user id from request 
	topic_id_str := chi.URLParam(r, "id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.DeleteTopic"))
	}

	type Body struct {
		UserID int `json:"user_id"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.DeleteTopic"))
	}
	userID := body.UserID

	// Check if topic exists
	topic, err := dataaccess.GetTopicByTopicID(db, topic_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.TopicsResult{
				Success: false,
				Error: fmt.Sprintf("Topic does not exist: %d", topic_id),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteTopic"))
    }	

	// Access control
	if userID != topic.UserID {
		return &models.TopicsResult{
				Success: false,
				Error: fmt.Sprintf("User: %d does not have right to delete topic: %d", userID, topic_id),
			}, nil	
	}

	// Delete topic
	res, err := dataaccess.DeleteTopicByTopicID(db, topic_id) 
	rowsAffected, errRA := res.RowsAffected()
	if err != nil || errRA != nil || rowsAffected != 1  {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteTopic"))
	}

	return &models.TopicsResult{
		Success: true,
	}, nil
}
