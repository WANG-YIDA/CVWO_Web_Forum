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
	InvalidCommentContent = `Invalid Comment Content: must consist of 1-250 characters in a-zA-Z0-9 .,!?'"()_\-`
)

var validCommentContentPattern = regexp.MustCompile(`^[a-zA-Z0-9 .,!?'"()_\-]{1,250}$`)

func CreateComment(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.CreateComment"))
	}
	defer db.Close()
	
	// Get comment post id, user id, content from request 
	comment := &models.Comment{}

	err = json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.CreateComment"))
	}

	// Comment content validation
	valid := validCommentContentPattern.MatchString(comment.Content)
	if !valid {
		return &models.CommentResult{
			Success: false,
			Error: InvalidCommentContent,
		}, nil
	}
	
	// Create and insert a new comment 
	t := time.Now()

	res, err := dataaccess.InsertNewComment(db, comment.PostID, comment.UserID, comment.Content, t)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateComment"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.CreateComment"))
	}

	comment.ID = int(id)	
	comment.CreatedAt = t

	return &models.CommentResult{
		Success: true,
		Comment: comment,
	}, nil
}

func ViewComments(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.ViewComments"))
	}
	defer db.Close()
	
	// Get post id from request 
	post_id_str := chi.URLParam(r, "id")
	post_id, err := strconv.Atoi(post_id_str)
    if err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid post ID in %s", "api.ViewComments"))
	}

	// Check if any comment exists
	comments, err := dataaccess.GetCommentsByPostID(db, post_id)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.CommentListResult{
				Success: true,
				Comment: nil,
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.ViewComments"))
    }	

	return &models.CommentListResult{
		Success: true,
		Comment: comments,
	}, nil
}

func DeleteComment(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.DeleteComment"))
	}
	defer db.Close()

	// Get comment id, post id, user id from request 
	comment_id_str := chi.URLParam(r, "id")
	comment_id, err := strconv.Atoi(comment_id_str)
    if err != nil {
        return nil, errors.Wrap(err, fmt.Sprintf("Invalid comment ID in %s", "api.DeleteComment"))
	}

	type Body struct {
		PostID int `json:"post_id"`
		UserID int `json:"user_id"`
	}
	var body Body

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.DeleteComment"))
	}
	userID := body.UserID
	postId := body.PostID

	// Check if comment exists
	comment, err := dataaccess.GetCommentByCommentIDAndPostID(db, comment_id, postId)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.CommentResult{
				Success: false,
				Error: fmt.Sprintf("Comment in post: %d does not exist: %d", postId, comment_id),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteComment"))
    }	

	// Access control
	if userID != comment.UserID {
		return &models.CommentResult {
				Success: false,
				Error: fmt.Sprintf("User: %d does not have right to delete comment: %d", userID, comment_id),
			}, nil	
	}

	// Delete comment
	res, err := dataaccess.DeleteCommentByCommentID(db, comment_id) 
	rowsAffected, errRA := res.RowsAffected()
	if err != nil || errRA != nil || rowsAffected != 1  {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.DeleteComment"))
	}

	return &models.CommentResult{
		Success: true,
	}, nil
}
