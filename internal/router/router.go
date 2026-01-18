package router

import (
	"database/sql"
	"os"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func Setup(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	// CORS Middleware
	_ = godotenv.Load(".env")
	frontendOriginURL := os.Getenv("FRONTEND_ORIGIN_URL")
	if frontendOriginURL == "" {
		frontendOriginURL = "http://localhost:3000" // fallback for local dev
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{frontendOriginURL}, 
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