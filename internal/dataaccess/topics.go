package dataaccess

import (
	"database/sql"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
)

func GetTopicByTopicID(db *sql.DB, topic_id int) (*models.Topic, error) {
	query := `SELECT id, name, user_id, description, created_at FROM topics WHERE id = ?`
	topic := &models.Topic{}
	err := db.QueryRow(query, topic_id).Scan(&topic.ID, &topic.Name, &topic.UserID, &topic.Description, &topic.CreatedAt)
	return topic, err
}


func GetTopics(db *sql.DB) (*[]models.Topic, error) {
	query := `SELECT id, name, user_id, description, created_at FROM topics`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	topics := []models.Topic{}

	for rows.Next() {
		topic := models.Topic{}
		err := rows.Scan(&topic.ID, &topic.Name, &topic.UserID, &topic.Description, &topic.CreatedAt)
		if err != nil {
            return nil, err
        }
		topics = append(topics, topic)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }
	return &topics, err 
}

func CheckTopicExistByTopicName(db *sql.DB, topic_name string) (bool, error) {
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM topics WHERE name = ?)`
	err := db.QueryRow(query, topic_name).Scan(&exist)
	return exist, err
}

func CheckTopicExistByTopicID(db *sql.DB, topic_id int) (bool, error) {
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM topics WHERE id = ?)`
	err := db.QueryRow(query, topic_id).Scan(&exist)
	return exist, err
}

func InsertNewTopic(db *sql.DB, topic_name string, user_id int, description string, created_at time.Time) (sql.Result, error) {
	query := `INSERT INTO topics (name, user_id, description, created_at) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(query, topic_name, user_id, description, created_at)
	return res, err
}

func UpdateTopicDescription(db *sql.DB, topic_id int, description string) (sql.Result, error) {
	query := `UPDATE topics SET description = ? WHERE id = ?`
	res, err := db.Exec(query, description, topic_id)
	return res, err
}

func DeleteTopicByTopicID(db *sql.DB, topic_id int) (sql.Result, error) {
	query := `DELETE FROM topics WHERE id = ?`
	res, err := db.Exec(query, topic_id)
	return res, err	
}