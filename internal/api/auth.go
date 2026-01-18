package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/dataaccess"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
	"github.com/pkg/errors"
)

const (
	InvalidUsername = "Invalid Username: must be 3-16 characters long and contain only letters, numbers, spaces, underscores (_), or hyphens (-)"
)

var validUsernamePattern = regexp.MustCompile("^[a-zA-Z0-9_ -]{3,16}$")

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	//Get username from request
	user := &models.User{}	

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.Login"))
	}
	username := user.Username

	// Validation
	valid := validUsernamePattern.MatchString(username)
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		return &models.AuthResult{
			Success: false,
			Error: InvalidUsername,
		}, nil
	}

	// Check if username exists
	user, err = dataaccess.GetUserByUsername(db, username)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
            return &models.AuthResult{
				Success: false,
				Error: fmt.Sprintf("User does not exist: %s", username),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Login"))
    }

	return &models.AuthResult{
		Success: true,
		User: user,
	}, nil
}

func Register(w http.ResponseWriter, r *http.Request, db *sql.DB) (interface{}, error) {
	//Get username from request
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.Register"))
	}
	username := user.Username

	// Validation
	valid := validUsernamePattern.MatchString(username)
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		return &models.AuthResult{
			Success: false,
			Error: InvalidUsername,
		}, nil
	}

	// Check if username exists
	exist, err := dataaccess.CheckUserExistByUsername(db, username)	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Register"))
    } 

	if exist {
		w.WriteHeader(http.StatusConflict)
		return &models.AuthResult{
			Success: false,
			Error: fmt.Sprintf("Username taken: %s", username),
		}, nil
	}

	// Create a new user and insert into DB
	t := time.Now()

	res, err := dataaccess.InsertNewUser(db, username, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Register"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Register"))
	}

	user.ID = int(id)
	user.CreatedAt = t

	return &models.AuthResult{
		Success: true,
		User: user,
	}, nil
}