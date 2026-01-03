package router

import (
	"database/sql"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/routes"
	"github.com/go-chi/chi/v5"
)

func Setup(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	setUpRoutes(r, db)
	return r
}

func setUpRoutes(r chi.Router, db *sql.DB) {
	routes := routes.GetRoutes(db)
	r.Group(routes)
}