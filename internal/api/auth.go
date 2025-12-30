package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/dataaccess"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/database"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
	"github.com/pkg/errors"
)

const (
	InvalidUsername = "Invalid Username: must consist of 3-16 alphannumeric characters, _ or -"
)

var validUsernamePattern = regexp.MustCompile("^[a-zA-Z0-9_-]{3,16}$")

func Login(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.Login"))
	}
	defer db.Close()

	//Get username from request
	user := &models.User{}	

	var username string
	err = json.NewDecoder(r.Body).Decode(user)
	username = user.Username
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.Login"))
	}

	// Validation
	valid := validUsernamePattern.MatchString(username)
	if !valid {
		return &models.AuthResult{
			Success: false,
			Error: InvalidUsername,
		}, nil
	}

	// Check if username exists
	user, err = dataaccess.GetUserByUsername(db, username)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
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

func Register(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.Register"))
	}
	defer db.Close()

	//Get username from request
	user := &models.User{}

	var username string
	err = json.NewDecoder(r.Body).Decode(user)
	username = user.Username
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetFromRequest, "api.Register"))
	}

	// Validation
	valid := validUsernamePattern.MatchString(username)
	if !valid {
		return &models.AuthResult{
			Success: false,
			Error: InvalidUsername,
		}, nil
	}

	// Check if username exists
	exist, err := dataaccess.CheckUserExistByUsername(db, username)	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Register"))
    } 

	if exist {
		return &models.AuthResult{
			Success: false,
			Error: fmt.Sprintf("Username taken: %s", username),
		}, nil
	}

	// Create a new user and insert into DB
	t := time.Now()

	res, err := dataaccess.InsertNewUser(db, username, t)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Register"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Register"))
	}

	user.ID = int(id)
	user.CreatedAt = t

	return &models.AuthResult{
		Success: true,
		User: user,
	}, nil
}