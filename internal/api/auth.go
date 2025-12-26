package api

import (
	"fmt"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	auth "github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
	"github.com/pkg/errors"
)

const (
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
)

func Login(w http.ResponseWriter, r *http.Request) ([]auth.Result, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.Login"))
	}
	defer db.Close()
	// Validation
	// Check if username exists

	return nil, nil

}

func Register(w http.ResponseWriter, r *http.Request) ([]auth.Result, error) {
	// Get DB
	// Validation
	// Check if username exists
	// Create a new user and insert into DB

	return nil, nil
}