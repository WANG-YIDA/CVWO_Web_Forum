package dataaccess

import (
	"database/sql"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
)

func List(db *sql.DB) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Username: "CVWO",
		},
	}
	return users, nil
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	query := `SELECT id, username, created_at FROM users WHERE username = ?`
	user := &models.User{}
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.CreatedAt)
	return user, err 
}

func CheckUserExistByUsername(db *sql.DB, username string) (bool, error) {
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)`
	err := db.QueryRow(query, username).Scan(&exist)
	return exist, err
}

func InsertNewUser(db *sql.DB, username string, created_at time.Time) (sql.Result, error) {
	query := `INSERT INTO users (username, created_at) VALUES (?, ?)`
	res, err := db.Exec(query, username, created_at)
	return res, err
}