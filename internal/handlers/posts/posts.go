package posts

import (
	"database/sql"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

func HandleCreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.CreatePost, "topics.HandleCreatePost")(w, r, db)	
}

func HandleViewPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewPost, "topics.HandleViewPost")(w, r, db)	
}

func HandleViewPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewPosts, "topics.HandleViewPosts")(w, r, db)	
}

func HandleEditPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.EditPost, "topics.HandleEditPost")(w, r, db)	
}

func HandleDeletePost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.DeletePost, "topics.HandleDeletePost")(w, r, db)
}


