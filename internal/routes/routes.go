package routes

import (
	"encoding/json"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/auth"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/comments"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/posts"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/topics"
	"github.com/go-chi/chi/v5"
)

func CreateRouteHandler(handlerFunc func(http.ResponseWriter, *http.Request) (*api.Response, error)) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Content-Type", "application/json")
        
        response, err := handlerFunc(w, req)
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

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
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
		r.Post("/topics", CreateRouteHandler(topics.HandleCreateTopic))
		r.Get("/topics/{id}", CreateRouteHandler(topics.HandleViewTopic))
		r.Patch("/topics/{id}", CreateRouteHandler(topics.HandleEditTopic))
		r.Delete("/topics/{id}", CreateRouteHandler(topics.HandleDeleteTopic))

		// Posts Handlers
		r.Post("/posts", CreateRouteHandler(posts.HandleCreatePost))
		r.Get("/posts/{id}", CreateRouteHandler(posts.HandleViewPost))
		r.Patch("/posts/{id}", CreateRouteHandler(posts.HandleEditPost))
		r.Delete("/posts/{id}", CreateRouteHandler(posts.HandleDeletePost))

		// Comments Handlers
		r.Post("/comments", CreateRouteHandler(comments.HandleCreateComment))
		r.Get("/comments/{id}", CreateRouteHandler(comments.HandleViewComment))
		r.Patch("/comments/{id}", CreateRouteHandler(comments.HandleEditComment))
		r.Delete("/comments/{id}", CreateRouteHandler(comments.HandleDeleteComment))

		// Authentication Handlers
		r.Post("/auth/login", CreateRouteHandler(auth.HandleLogin))
		r.Post("/auth/register", CreateRouteHandler(auth.HandleRegister))
	}
}
