package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

func createTables(dbConn *sql.DB) {
	tableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS topics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT,
		user_id INTEGER NOT NULL,
		name TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		FOREIGN KEY(user_id) REFERENCES users(id),
	);

	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		topic_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(topic_id) REFERENCES topics(id)
	);

	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err := dbConn.Exec(tableSQL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Tables created successfully")
}

func GetDB() (*Database, error) {
	dbConn, err := sql.Open("sqlite3", "cmd/server/forum.db")
	if err != nil {
		return nil, err
	}
	
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	
	createTables(dbConn)

	return &Database{Conn: dbConn}, nil
}
