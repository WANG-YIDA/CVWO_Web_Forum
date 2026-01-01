package dataaccess

import (
	"database/sql"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
)

func GetPostByPostID(db *sql.DB, post_id int) (*models.Post, error) {
	query := `SELECT * FROM posts WHERE id = ?`
	post := &models.Post{}
	err := db.QueryRow(query, post_id).Scan(&post.ID, &post.UserID, &post.TopicID, &post.Title, &post.Content, &post.CreatedAt)
	return post, err 
}

func InsertNewPost(db *sql.DB, title string, user_id int, topic_id int, content string, created_at time.Time) (sql.Result, error) {
	query := `INSERT INTO posts (title, user_id, topic_id, content, created_at) VALUES (?, ?, ?, ?, ?)`
	res, err := db.Exec(query, title, user_id, topic_id, content, created_at)
	return res, err
}

func UpdatePost(db *sql.DB, post_id int, topic_id int, title string, content string) (sql.Result, error) {
	query := `UPDATE posts SET topic_id = ?, title = ?, content = ? WHERE id = ?`
	res, err := db.Exec(query, topic_id, title, content, post_id)
	return res, err
}

func DeletePostByPostID(db *sql.DB, post_id int) (sql.Result, error) {
	query := `DELETE FROM posts WHERE id = ?`
	res, err := db.Exec(query, post_id)
	return res, err	
}