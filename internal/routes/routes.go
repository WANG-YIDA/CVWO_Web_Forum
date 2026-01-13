package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/auth"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/comments"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/posts"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/topics"
	"github.com/go-chi/chi/v5"
)

func CreateRouteHandler(handlerFunc func(http.ResponseWriter, *http.Request, *sql.DB) (*api.Response, error), db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        response, err := handlerFunc(w, req, db)
        if err != nil {
			response = &api.Response{
				Payload: api.Payload{},
				Success: false,
				Error: err.Error(),
			}
        }
        
        json.NewEncoder(w).Encode(response)
    }
}

func GetRoutes(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		// For testing connection with frontend
		r.Get("/api/handshake", func(w http.ResponseWriter, req *http.Request) {
    		json.NewEncoder(w).Encode(map[string]string{
				   "status": "Connected",
   			       "message": "Go backend is connected to React!",
    		})
		})

		// Topics Handlers
		r.Post("/api/topics", CreateRouteHandler(topics.HandleCreateTopic, db))
		r.Get("/api/topics/{id}", CreateRouteHandler(topics.HandleViewTopic, db))
		r.Get("/api/topics", CreateRouteHandler(topics.HandleViewTopics, db))
		r.Patch("/api/topics/{id}", CreateRouteHandler(topics.HandleEditTopic, db))
		r.Delete("/api/topics/{id}", CreateRouteHandler(topics.HandleDeleteTopic, db))

		// Posts Handlers
		r.Post("/api/topics/{topic_id}/posts", CreateRouteHandler(posts.HandleCreatePost, db))
		r.Get("/api/topics/{topic_id}/posts/{post_id}", CreateRouteHandler(posts.HandleViewPost, db))
		r.Get("/api/topics/{topic_id}/posts", CreateRouteHandler(posts.HandleViewPosts, db))
		r.Patch("/api/topics/{topic_id}/posts/{post_id}", CreateRouteHandler(posts.HandleEditPost, db))
		r.Delete("/api/topics/{topic_id}/posts/{post_id}", CreateRouteHandler(posts.HandleDeletePost, db))

		// Comments Handlers
		r.Post("/api/topics/{topic_id}/posts/{post_id}/comments", CreateRouteHandler(comments.HandleCreateComment, db))
		r.Get("/api/topics/{topic_id}/posts/{post_id}/comments", CreateRouteHandler(comments.HandleViewComments, db))
		r.Delete("/api/topics/{topic_id}/posts/{post_id}/comments/{comment_id}", CreateRouteHandler(comments.HandleDeleteComment, db))

		// Authentication Handlers
		r.Post("/api/auth/login", CreateRouteHandler(auth.HandleLogin, db))
		r.Post("/api/auth/register", CreateRouteHandler(auth.HandleRegister, db))
	}
}
