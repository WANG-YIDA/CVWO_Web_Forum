package topics

import (
	"database/sql"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

func HandleCreateTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.CreateTopic, "topics.HandleCreateTopic")(w, r, db)	
}

func HandleViewTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewTopic, "topics.HandleViewTopic")(w, r, db)	
}

func HandleViewTopics(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewTopics, "topics.HandleViewTopics")(w, r, db)	
}

func HandleEditTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.EditTopic, "topics.HandleEditTopic")(w, r, db)	
}

func HandleDeleteTopic(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.DeleteTopic, "topics.HandleDeleteTopic")(w, r, db)
}


