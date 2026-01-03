package comments

import (
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

func HandleCreateComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.CreateComment, "topics.HandleCreateComment")(w, r)	
}

func HandleViewComments(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewComments, "topics.HandleViewComments")(w, r)	
}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.DeleteComment, "topics.HandleDeleteComment")(w, r)
}


