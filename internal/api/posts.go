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
	InvalidPostTitle = `Invalid Post Title: must consist of 3-50 characters in a-zA-Z0-9 .,!?'"()_\-`
	InvalidPostContent = `Invalid Post Content: must consist of 1-500 characters in a-zA-Z0-9 .,!?'"()_\-`
)

var validPostTitlePattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{3,50}$`)
var validPostContentPattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{1,500}$`)

func CreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get post title, user id, topic_id, content from request 
	post := &models.Post{}

	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.CreatePost"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.CreatePost"))
	}

	// Post title, content, topic, user validation
	valid := validPostTitlePattern.MatchString(post.Title)
	if !valid {
		w.WriteHeader(http.StatusBadRequest) 
		return &models.PostsResult{
			Success: false,
			Error: InvalidPostTitle,
		}, nil
	}
	
	valid = validPostContentPattern.MatchString(post.Content)
	if !valid {
		w.WriteHeader(http.StatusBadRequest) 
		return &models.PostsResult{
			Success: false,
			Error: InvalidPostContent,
		}, nil
	}

	// Check if topic, user exist
	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	exist, err = dataaccess.CheckUserExistByUserID(db, post.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return nil, errors.Wrap(err, fmt.Sprintf("User does not exist: %d", post.UserID))
	}

	// Get username
	username, err := dataaccess.GetUsernameByUserID(db, post.UserID)
		if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))
	}

	// Create and insert a new post 
	t := time.Now()

	res, err := dataaccess.InsertNewPost(db, post.Title, post.UserID, username, topic_id, post.Content, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))
	}

	post.TopicID = topic_id
	post.ID = int(id)	
	post.Author = username
	post.CreatedAt = t

	return &models.PostsResult{
		Success: true,
		Post: post,
	}, nil
}

func ViewPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get post id, topic id from request 
	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.ViewPost"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.ViewPost"))
	}

	// Check if topic, post exist
	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewPost"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	post, err := dataaccess.GetPostByPostIDAndTopicID(db, post_id, topic_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Post: %d does not exist in topic: %d", post_id, topic_id))
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewPost"))
    }	

	return &models.PostsResult{
		Success: true,
		Post: post,
	}, nil
}

func ViewPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get topic id from request 
	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.ViewPosts"))
	}

	// Check if topic exists
	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewPosts"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	// Check if any post exists
	posts, err := dataaccess.GetPostsByTopicID(db, topic_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.PostListResult{
				Success: true,
				Posts: nil,
			}, nil
        }
		w.WriteHeader(http.StatusInternalServerError) 
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewPosts"))
    }	

	return &models.PostListResult{
		Success: true,
		Posts: posts,
	}, nil
}

func EditPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get post id, topic id, user id, title and content from request 
	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.EditPost"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.EditPost"))
	}

	type Body struct {
		UserID int `json:"user_id"`
		Title string `json:"title"`
		Content string `json:"content"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) 
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.EditPost"))
	}
	userID := body.UserID
	title := body.Title
	content := body.Content

	/// Check if topic, post exist
	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditPost"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	post, err := dataaccess.GetPostByPostIDAndTopicID(db, post_id, topic_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Post: %d does not exist in topic: %d", post_id, topic_id))
        }
		w.WriteHeader(http.StatusInternalServerError)
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditPost"))
    }	

	// Access control
	if userID != post.UserID {
		w.WriteHeader(http.StatusForbidden)
		return &models.PostsResult{
				Success: false,
				Error: fmt.Sprint("You don't have the right to edit this post"),
			}, nil	
	}

	// Form inputs validation (title, content, are editable)
	valid := validPostContentPattern.MatchString(content)
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		return &models.PostsResult{
			Success: false,
			Error: InvalidPostContent,
		}, nil
	}

	valid = validPostTitlePattern.MatchString(title)
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		return &models.PostsResult{
			Success: false,
			Error: InvalidPostTitle,
		}, nil
	}
	
	// Update and return modified post
	post.Title = title
	post.Content = content
	_, err = dataaccess.UpdatePost(db, post_id, title, content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditPost"))	
	}

	return &models.PostsResult{
		Success: true,
		Post: post,
	}, nil
}

func DeletePost(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	// Get post id, topic id, user id from request 
	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.DeletePost"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.DeletePost"))
	}

	type Body struct {
		UserID int `json:"user_id"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.DeletePost"))
	}
	userID := body.UserID

	// Check if topic, post exists
	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeletePost"))	
	}

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Topic does not exist: %d", topic_id))
	}

	post, err := dataaccess.GetPostByPostIDAndTopicID(db, post_id, topic_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
		  return nil, errors.Wrap(err, fmt.Sprintf("Post: %d does not exist in topic: %d", post_id, topic_id))
        }
		w.WriteHeader(http.StatusInternalServerError)
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeletePost"))
    }	

	// Access control
	if userID != post.UserID {
		w.WriteHeader(http.StatusForbidden)
		return &models.PostsResult{
				Success: false,
				Error: fmt.Sprint("You don't have the right to delete this post"),
			}, nil	
	}

	// Delete post
	res, err := dataaccess.DeletePostByPostIDTopicID(db, post_id, topic_id) 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
    	return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeletePost"))
	}

	rows, errRA := res.RowsAffected()
	if errRA != nil {
		w.WriteHeader(http.StatusInternalServerError)
    	return nil, errors.Wrap(errRA, fmt.Sprintf(ErrDB, "api.DeletePost"))
	}
	if rows != 1 {
		w.WriteHeader(http.StatusInternalServerError)
    	return nil, errors.Errorf("api.DeletePost: expected to delete 1 row, deleted %d", rows)
	}

	return &models.PostsResult{
		Success: true,
	}, nil
}
