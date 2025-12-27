package routes

import (
	"encoding/json"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers/auth"
	"github.com/go-chi/chi/v5"
)

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

		// Users Handlers
		

		// Authentication Handlers
		r.Post("/auth/login", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") 
			w.Header().Set("Content-Type", "application/json")

			response, err := auth.HandleLogin(w, req)
			if err != nil {
        		w.WriteHeader(http.StatusInternalServerError)
        		json.NewEncoder(w).Encode(map[string]interface{}{
            		"success": false,
            		"error": err,
        		})
        		return
    		}

			json.NewEncoder(w).Encode(response)
		})	

		r.Post("/auth/register", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") 
			w.Header().Set("Content-Type", "application/json")

			response, err := auth.HandleRegister(w, req)
			if err != nil {
        		w.WriteHeader(http.StatusInternalServerError)
        		json.NewEncoder(w).Encode(map[string]interface{}{
            		"success": false,
            		"error": err,
        		})
        		return
    		}

			json.NewEncoder(w).Encode(response)
		})	
	}
}
