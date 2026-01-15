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
	InvalidCommentContent = `Invalid Comment Content: must consist of 1-250 characters in a-zA-Z0-9 .,!?'"()_\-`
)

var validCommentContentPattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{1,250}$`)

func CreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get comment topic id, post id, user id, content from request 
	comment := &models.Comment{}

	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.CreateComment"))
	}

	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.CreateComment"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.CreateComment"))
	}

	// Check if post, topic, user exist	
	exist, err := dataaccess.CheckUserExistByUserID(db, comment.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateComment"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("User does not exist: %d", comment.UserID))
	}

	exist, err = dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateComment"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	exist, err = dataaccess.CheckPostExistByPostIDTopicID(db, post_id, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateComment"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Post: %d does not exist in topic: %d", post_id, topic_id))
	}

	// Comment content validation
	valid := validCommentContentPattern.MatchString(comment.Content)
	if !valid {
		w.WriteHeader(http.StatusBadRequest) 
		return &models.CommentsResult{
			Success: false,
			Error: InvalidCommentContent,
		}, nil
	}

	// Get username
	username, err := dataaccess.GetUsernameByUserID(db, comment.UserID)
		if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateTopic"))
	}
	
	// Create and insert a new comment 
	t := time.Now()

	res, err := dataaccess.InsertNewComment(db, post_id, comment.UserID, username, comment.Content, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateComment"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateComment"))
	}

	comment.ID = int(id)	
	comment.Author = username
	comment.CreatedAt = t
	comment.PostID = post_id

	return &models.CommentsResult{
		Success: true,
		Comment: comment,
	}, nil
}

func ViewComments(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get post id, topic id from request 
	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.ViewComments"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.ViewComments"))
	}

	// Check if topic, post exist
	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewComments"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound) 
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	exist, err = dataaccess.CheckPostExistByPostIDTopicID(db, post_id, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewComments"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound) 
		  return nil, errors.Wrap(err, fmt.Sprintf("Post: %d does not exist in topic: %d", post_id, topic_id))
	}

	// Check if any comment exists
	comments, err := dataaccess.GetCommentsByPostID(db, post_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.CommentListResult{
				Success: true,
				Comments: nil,
			}, nil
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewComments"))
    }	

	return &models.CommentListResult{
		Success: true,
		Comments: comments,
	}, nil
}

func DeleteComment(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get comment id, topic id, post id, user id from request 
	comment_id_str := chi.URLParam(r, "comment_id")
	comment_id, err := strconv.Atoi(comment_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid comment ID in %s", "api.DeleteComment"))
	}

	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.DeleteComment"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.DeleteComment"))
	}

	type Body struct {
		UserID int `json:"user_id"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.DeleteComment"))
	}
	userID := body.UserID

	// Check if topic, post, comment exist
	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteComment"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	exist, err = dataaccess.CheckPostExistByPostIDTopicID(db, post_id, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteComment"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return nil, errors.Wrap(err, fmt.Sprintf("Post: %d does not exist in topic: %d", post_id, topic_id))
	}

	comment, err := dataaccess.GetCommentByCommentIDAndPostID(db, comment_id, post_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		return nil, errors.Errorf("Comment: %d in post: %d of topic: %d does not exist", comment_id, post_id, topic_id)
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteComment"))
    }	

	// Access control
	if userID != comment.UserID {
		w.WriteHeader(http.StatusForbidden)
		return &models.CommentsResult {
				Success: false,
				Error: fmt.Sprintf("You don't have the right to delete this comment: %d", comment_id),
			}, nil	
	}

	// Delete comment
	res, err := dataaccess.DeleteCommentByCommentIDPostID(db, comment_id, post_id) 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) 
    	return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteComment"))
	}

	rows, errRA := res.RowsAffected()
	if errRA != nil {
		w.WriteHeader(http.StatusInternalServerError) 
    	return nil, errors.Wrap(errRA, fmt.Sprintf(ErrDB, "api.DeleteComment"))
	}
	if rows != 1 {
		w.WriteHeader(http.StatusInternalServerError) 
    	return nil, errors.Errorf("api.DeleteComment: expected to delete 1 row, deleted %d", rows)
	}

	return &models.CommentsResult{
		Success: true,
	}, nil
}
