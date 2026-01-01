package posts

import (
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

func HandleCreatePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.CreatePost, "topics.HandleCreatePost")(w, r)	
}

func HandleViewPost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewPost, "topics.HandleViewPost")(w, r)	
}

func HandleEditPost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.EditPost, "topics.HandleEditPost")(w, r)	
}

func HandleDeletePost(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.DeletePost, "topics.HandleDeletePost")(w, r)
}


