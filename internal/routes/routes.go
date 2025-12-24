package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/handshake", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") 
    		w.Header().Set("Content-Type", "application/json")
	
    		json.NewEncoder(w).Encode(map[string]string{
				   "status": "Connected",
   			       "message": "Go backend is connected to React!",
    		})
		})

		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleList(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	}
}
