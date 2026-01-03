package posts

import (
	"database/sql"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

func HandleCreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.CreatePost, "posts.HandleCreatePost")(w, r, db)	
}

func HandleViewPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewPost, "posts.HandleViewPost")(w, r, db)	
}

func HandleViewPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewPosts, "posts.HandleViewPosts")(w, r, db)	
}

func HandleEditPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.EditPost, "posts.HandleEditPost")(w, r, db)	
}

func HandleDeletePost(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.DeletePost, "posts.HandleDeletePost")(w, r, db)
}


