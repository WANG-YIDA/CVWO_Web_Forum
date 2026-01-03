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
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "success": false,
                "error":   err.Error(), 
            })
            return
        }
        
        json.NewEncoder(w).Encode(response)
    }
}

func GetRoutes(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		// Middleware
		r.Use(func(handler http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE,OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				if req.Method == http.MethodOptions {
					w.WriteHeader(http.StatusOK)
					return
				}
				handler.ServeHTTP(w, req)
			})
		})

		// For testing connection with frontend
		r.Get("/handshake", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") 
    		w.Header().Set("Content-Type", "application/json")
	
    		json.NewEncoder(w).Encode(map[string]string{
				   "status": "Connected",
   			       "message": "Go backend is connected to React!",
    		})
		})

		// Topics Handlers
		r.Post("/topics", CreateRouteHandler(topics.HandleCreateTopic, db))
		r.Get("/topics/{id}", CreateRouteHandler(topics.HandleViewTopic, db))
		r.Get("/topics", CreateRouteHandler(topics.HandleViewTopics, db))
		r.Patch("/topics/{id}", CreateRouteHandler(topics.HandleEditTopic, db))
		r.Delete("/topics/{id}", CreateRouteHandler(topics.HandleDeleteTopic, db))

		// Posts Handlers
		r.Post("/topics/{topic_id}/posts", CreateRouteHandler(posts.HandleCreatePost, db))
		r.Get("/topics/{topic_id}/posts/{post_id}", CreateRouteHandler(posts.HandleViewPost, db))
		r.Get("/topics/{topic_id}/posts", CreateRouteHandler(posts.HandleViewPosts, db))
		r.Patch("/topics/{topic_id}/posts/{post_id}", CreateRouteHandler(posts.HandleEditPost, db))
		r.Delete("/topics/{topic_id}/posts/{post_id}", CreateRouteHandler(posts.HandleDeletePost, db))

		// Comments Handlers
		r.Post("/topics/{topic_id}/posts/{post_id}/comments", CreateRouteHandler(comments.HandleCreateComment, db))
		r.Get("/topics/{topic_id}/posts/{post_id}/comments", CreateRouteHandler(comments.HandleViewComments, db))
		r.Delete("/topics/{topic_id}/posts/{post_id}/comments/{comment_id}", CreateRouteHandler(comments.HandleDeleteComment, db))

		// Authentication Handlers
		r.Post("/auth/login", CreateRouteHandler(auth.HandleLogin, db))
		r.Post("/auth/register", CreateRouteHandler(auth.HandleRegister, db))
	}
}
