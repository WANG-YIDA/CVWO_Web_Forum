package models

type Result struct {
	Success string `json:"success"`	
	Error string `json:"error"`
	User User `json:"user"`
}