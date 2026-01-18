package dataaccess

import (
	"database/sql"
	"time"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/models"
)

func GetCommentByCommentIDAndPostID(db *sql.DB, comment_id int, post_id int) (*models.Comment, error) {
	query := `SELECT * FROM comments WHERE id = ? AND post_id = ?`
	comment := &models.Comment{}
	err := db.QueryRow(query, comment_id, post_id).Scan(&comment.ID, &comment.PostID, &comment.Author, &comment.UserID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return comment, err
}

func GetCommentsByPostID(db *sql.DB, post_id int) (*[]models.Comment, error) {
	query := `SELECT * FROM comments WHERE post_id = ?`
	rows, err := db.Query(query, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}

	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.Author, &comment.UserID, &comment.Content, &comment.CreatedAt)
		if err != nil {
            return nil, err
        }
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }
	return &comments, err 
}

func InsertNewComment(db *sql.DB, post_id int, user_id int, author string, content string, created_at time.Time) (sql.Result, error) {
	query := `INSERT INTO comments (post_id, user_id, author, content, created_at) VALUES (?, ?, ?, ?, ?)`
	res, err := db.Exec(query, post_id, user_id, author, content, created_at)
	return res, err
}

func DeleteCommentByCommentIDPostID(db *sql.DB, comment_id int, post_id int) (sql.Result, error) {
	query := `DELETE FROM comments WHERE id = ? AND post_id = ?`
	res, err := db.Exec(query, comment_id, post_id)
	return res, err	
}