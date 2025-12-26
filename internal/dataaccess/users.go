package dataaccess

import (
	"database/sql"

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
