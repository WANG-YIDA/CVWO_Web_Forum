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
	InvalidPostTitle = `Invalid Post Title: must consist of 3-50 characters in a-zA-Z0-9 .,!?'"()_\-`
	InvalidPostContent = `Invalid Post Content: must consist of 1-500 characters in a-zA-Z0-9 .,!?'"()_\-`
)

var validPostTitlePattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{3,50}$`)
var validPostContentPattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{1,500}$`)

func CreatePost(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.CreatePost"))
	}
	defer db.Close()
	
	// Get post title, user id, topic_id, content from request 
	post := &models.Post{}

	err = json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.CreatePost"))
	}

	topic_id_str := chi.URLParam(r, "topic_id")
	topic_id, err := strconv.Atoi(topic_id_str)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Invalid topic ID in %s", "api.CreatePost"))
	}

	// Post title, content, topic validation
	valid := validPostTitlePattern.MatchString(post.Title)
	if !valid {
		return &models.PostsResult{
			Success: false,
			Error: InvalidPostTitle,
		}, nil
	}
	
	valid = validPostContentPattern.MatchString(post.Content)
	if !valid {
		return &models.PostsResult{
			Success: false,
			Error: InvalidPostContent,
		}, nil
	}

	exist, err := dataaccess.CheckTopicExistByTopicID(db, topic_id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))	
	}

	if !exist {
		return &models.PostsResult{
			Success: false,
			Error: fmt.Sprintf("Topic does not exist: %d", topic_id),
		}, nil
	}

	// Create and insert a new post 
	t := time.Now()

	res, err := dataaccess.InsertNewPost(db, post.Title, post.UserID, topic_id, post.Content, t)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreatePost"))
	}

	post.TopicID = topic_id
	post.ID = int(id)	
	post.CreatedAt = t

	return &models.PostsResult{
		Success: true,
		Post: post,
	}, nil
}

func ViewPost(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.ViewPost"))
	}
	defer db.Close()
	
	// Get post id from request 
	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.ViewPost"))
	}

	// Check if post exists
	post, err := dataaccess.GetPostByPostID(db, post_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.PostsResult{
				Success: false,
				Error: fmt.Sprintf("Post does not exist: %d", post_id),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewPost"))
    }	

	return &models.PostsResult{
		Success: true,
		Post: post,
	}, nil
}

func EditPost(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.EditPost"))
	}
	defer db.Close()

	// Get post id, topic id, user id, title and content from request 
	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.EditPost"))
	}

	type Body struct {
		UserID int `json:"user_id"`
		Title string `json:"title"`
		Content string `json:"content"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.EditPost"))
	}
	userID := body.UserID
	title := body.Title
	content := body.Content

	// Check if post exists
	post, err := dataaccess.GetPostByPostID(db, post_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.PostsResult{
				Success: false,
				Error: fmt.Sprintf("Post does not exist: %d", post_id),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditPost"))
    }	

	// Access control
	if userID != post.UserID {
		return &models.PostsResult{
				Success: false,
				Error: fmt.Sprintf("User: %d does not have right to edit this post: %d", userID, post_id),
			}, nil	
	}

	// Form inputs validation (title, content, are editable)
	valid := validPostContentPattern.MatchString(content)
	if !valid {
		return &models.PostsResult{
			Success: false,
			Error: InvalidPostContent,
		}, nil
	}

	valid = validPostTitlePattern.MatchString(title)
	if !valid {
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.EditPost"))	
	}

	return &models.PostsResult{
		Success: true,
		Post: post,
	}, nil
}

func DeletePost(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.DeletePost"))
	}
	defer db.Close()

	// Get post id, user id from request 
	post_id_str := chi.URLParam(r, "post_id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.DeletePost"))
	}

	type Body struct {
		UserID int `json:"user_id"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.DeletePost"))
	}
	userID := body.UserID

	// Check if post exists
	post, err := dataaccess.GetPostByPostID(db, post_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.PostsResult{
				Success: false,
				Error: fmt.Sprintf("Post does not exist: %d", post_id),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeletePost"))
    }	

	// Access control
	if userID != post.UserID {
		return &models.PostsResult{
				Success: false,
				Error: fmt.Sprintf("User: %d does not have right to delete post: %d", userID, post_id),
			}, nil	
	}

	// Delete post
	res, err := dataaccess.DeletePostByPostID(db, post_id) 
	rowsAffected, errRA := res.RowsAffected()
	if err != nil || errRA != nil || rowsAffected != 1  {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeletePost"))
	}

	return &models.PostsResult{
		Success: true,
	}, nil
}
