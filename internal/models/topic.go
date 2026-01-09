package models

import "time"

type Topic struct {
    ID          int       `json:"id"` 
    Description string    `json:"description"` 
    UserID      int       `json:"user_id"`
    Author      string    `json:"author"`
    Name        string    `json:"name"` 
    CreatedAt   time.Time `json:"created_at"`
}

type TopicsResult struct {
	Success bool `json:"success"`	
	Error string `json:"error,omitempty"`
	Topic *Topic `json:"topic,omitempty"`
}

type TopicListResult struct {
	Success bool `json:"success"`	
	Error string `json:"error,omitempty"`
	Topics *[]Topic `json:"topics,omitempty"`
}