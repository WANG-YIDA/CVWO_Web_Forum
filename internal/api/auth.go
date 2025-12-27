package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
	"github.com/pkg/errors"
)

const (
	ErrRetrieveDatabase = "Failed to retrieve database in %s"
	ErrGetUsernameFromRequest = "Invalid request in %s"
	ErrFindUser = "Failed to find user in %s"
	

	InvalidUsername = "Username must be between 3 and 16 characters long and can only contain alphanumeric characters, '_', or '-'."
)

var validUsernamePattern = regexp.MustCompile("^[a-zA-Z0-9_-]{3,16}$")

func Login(w http.ResponseWriter, r *http.Request) (models.Result, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return models.Result{}, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.Login"))
	}
	defer db.Close()

	//Get username from request
	var body struct {
	    Username string `json:"username"`
	}
	var username string
	err = json.NewDecoder(r.Body).Decode(&body)
	username = body.Username
	if err != nil {
		return models.Result{}, errors.Wrap(err, fmt.Sprintf(ErrGetUsernameFromRequest, "api.Login"))
	}

	// Validation
	valid := validUsernamePattern.MatchString(username)
	if !valid {
		return models.Result{
			Success: false,
			Error: InvalidUsername,
		}, nil
	}

	// Check if username exists
	query := `SELECT id, username FROM users WHERE username = ?`
	user := &models.User{}
	err = db.QueryRow(query, username).Scan(&user.ID, &user.Username)
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return models.Result{
				Success: false,
				Error: fmt.Sprintf("User does not exist: %s", username),
			}, nil
        }
        return models.Result{}, errors.Wrap(err, fmt.Sprintf(ErrFindUser, "api.Login"))
    }

	return models.Result{
		Success: true,
		User: user,
	}, nil

}

func Register(w http.ResponseWriter, r *http.Request) (models.Result, error) {
	// Get DB
	// Validation
	// Check if username exists
	// Create a new user and insert into DB

	return models.Result{}, nil
}