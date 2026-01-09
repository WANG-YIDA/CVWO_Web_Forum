package models

import "time"

type Post struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    TopicID   int       `json:"topic_id"`
    Title     string    `json:"title"`
	Author    string    `json:"author"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

type PostsResult struct {
	Success bool `json:"success"`	
	Error string `json:"error,omitempty"`
	Post *Post `json:"post,omitempty"`
}

type PostListResult struct {
	Success bool `json:"success"`	
	Error string `json:"error,omitempty"`
	Posts *[]Post `json:"posts,omitempty"`
}