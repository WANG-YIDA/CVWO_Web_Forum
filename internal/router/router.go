package router

import (
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/routes"
	"github.com/go-chi/chi/v5"
)

func Setup() (chi.Router, error) {
	r := chi.NewRouter()

	err := setUpRoutes(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func setUpRoutes(r chi.Router) error {
	routes, err := routes.GetRoutes()
	if err != nil {
		return err
	}

	r.Group(routes)
	return nil
}
