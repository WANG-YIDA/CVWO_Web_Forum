package models

import "time"

type Comment struct {
    ID        int       `json:"id"`
    PostID    int       `json:"post_id"`
    UserID    int       `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

type CommentsResult struct {
	Success bool `json:"success"`	
	Error string `json:"error,omitempty"`
	Comment *Comment `json:"comment,omitempty"`
}

type CommentListResult struct {
	Success bool `json:"success"`	
	Error string `json:"error,omitempty"`
	Comments *[]Comment `json:"comments,omitempty"`
}