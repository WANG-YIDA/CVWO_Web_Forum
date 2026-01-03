package comments

import (
	"database/sql"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

func HandleCreateComment(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.CreateComment, "comments.HandleCreateComment")(w, r, db)	
}

func HandleViewComments(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewComments, "comments.HandleViewComments")(w, r, db)	
}

func HandleDeleteComment(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.DeleteComment, "comments.HandleDeleteComment")(w, r, db)
}


