package models

type Result struct {
	Success bool `json:"success"`	
	Error string `json:"error,omitempty"`
	User *User `json:"user,omitempty"`
}