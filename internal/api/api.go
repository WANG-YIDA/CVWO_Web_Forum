package api

import (
	"encoding/json"
)

type Payload struct {
	Meta json.RawMessage `json:"meta,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

type Response struct {
	Payload   Payload  `json:"payload"`
	Success   bool     `json:"success,omitempty"`
	Error     string   `json:"error,omitempty"`
}
const (
	ErrRetrieveDatabase = "Failed to retrieve database"
	ErrGetFromRequest = "Invalid request in %s"
	ErrDB = "Database query failed in %s"
)
	

