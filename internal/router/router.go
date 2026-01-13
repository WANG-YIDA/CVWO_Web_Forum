package router

import (
	"database/sql"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Setup(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	// CORS Middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, 
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	setUpRoutes(r, db)
	return r
}

func setUpRoutes(r chi.Router, db *sql.DB) {
	routes := routes.GetRoutes(db)
	r.Group(routes)
}