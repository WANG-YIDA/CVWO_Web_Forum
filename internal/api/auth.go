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
	ErrRetrieveDatabase = "Failed to retrieve database in %s"
	ErrGetUsernameFromRequest = "Invalid request in %s"
	ErrDB = "Databse query failed in %s"
	

	InvalidUsername = "Invalid Username: must consist of 3-16 alphannumeric characters, _ or -"
)

var validUsernamePattern = regexp.MustCompile("^[a-zA-Z0-9_-]{3,16}$")

func Login(w http.ResponseWriter, r *http.Request) (*models.Result, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.Login"))
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetUsernameFromRequest, "api.Login"))
	}

	// Validation
	valid := validUsernamePattern.MatchString(username)
	if !valid {
		return &models.Result{
			Success: false,
			Error: InvalidUsername,
		}, nil
	}

	// Check if username exists
	user, err := dataaccess.GetUserByUsername(db, username)	
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return &models.Result{
				Success: false,
				Error: fmt.Sprintf("User does not exist: %s", username),
			}, nil
        }
        return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.Login"))
    }

	return &models.Result{
		Success: true,
		User: user,
	}, nil

}

func Register(w http.ResponseWriter, r *http.Request) (*models.Result, error) {
	// Get DB
	db, err := database.GetDB()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, "api.register"))
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
		return nil, errors.Wrap(err, fmt.Sprintf(ErrGetUsernameFromRequest, "api.register"))
	}

	// Validation
	valid := validUsernamePattern.MatchString(username)
	if !valid {
		return &models.Result{
			Success: false,
			Error: InvalidUsername,
		}, nil
	}

	// Check if username exists
	exist, err := dataaccess.CheckUserExistByUsername(db, username)	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.register"))
    } 

	if exist {
		return &models.Result{
			Success: false,
			Error: fmt.Sprintf("Username taken: %s", username),
		}, nil
	}

	// Create a new user and insert into DB
	t := time.Now()

	res, err := dataaccess.InsertNewUser(db, username, t)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.register"))
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrDB, "api.register"))
	}

	user := &models.User{
		ID: int(id), 
		Username: username,
		CreatedAt: t,
	}

	return &models.Result{
		Success: true,
		User: user,
	}, nil
}