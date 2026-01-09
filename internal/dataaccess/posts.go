package dataaccess

import (
	"database/sql"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
)

func GetPostByPostIDAndTopicID(db *sql.DB, post_id int, topic_id int) (*models.Post, error) {
	query := `SELECT id, user_id, author, topic_id, title, content, created_at FROM posts WHERE id = ? AND topic_id = ?`
	post := &models.Post{}
	err := db.QueryRow(query, post_id, topic_id).Scan(&post.ID, &post.UserID, &post.Author, &post.TopicID, &post.Title, &post.Content, &post.CreatedAt)
	return post, err 
}

func GetPostsByTopicID(db *sql.DB, topic_id int) (*[]models.Post, error) {
	query := `SELECT id, title, user_id, author, topic_id, content, created_at FROM posts WHERE topic_id = ?`
	rows, err := db.Query(query, topic_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.Post{}

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.UserID, &post.Author, &post.TopicID, &post.Content, &post.CreatedAt)
		if err != nil {
            return nil, err
        }
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }
	return &posts, err 
}

func CheckPostExistByPostIDTopicID(db *sql.DB, post_id int, topic_id int) (bool, error) {
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM posts WHERE id = ? AND topic_id = ?)`
	err := db.QueryRow(query, post_id, topic_id).Scan(&exist)
	return exist, err
}

func InsertNewPost(db *sql.DB, title string, user_id int, author string, topic_id int, content string, created_at time.Time) (sql.Result, error) {
	query := `INSERT INTO posts (title, user_id, author, topic_id, content, created_at) VALUES (?, ?, ?, ?, ?)`
	res, err := db.Exec(query, title, user_id, topic_id, author, content, created_at)
	return res, err
}

func UpdatePost(db *sql.DB, post_id int, title string, content string) (sql.Result, error) {
	query := `UPDATE posts SET title = ?, content = ? WHERE id = ?`
	res, err := db.Exec(query, title, content, post_id)
	return res, err
}

func DeletePostByPostIDTopicID(db *sql.DB, post_id int, topic_id int) (sql.Result, error) {
	query := `DELETE FROM posts WHERE id = ? AND topic_id = ?`
	res, err := db.Exec(query, post_id, topic_id)
	return res, err	
}