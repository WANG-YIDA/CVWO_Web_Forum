package topics

import (
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

func HandleCreateTopic(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.CreateTopic, "topics.HandleCreateTopic")(w, r)	
}

func HandleViewTopic(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewTopic, "topics.HandleViewTopic")(w, r)	
}

func HandleViewTopics(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.ViewTopics, "topics.HandleViewTopics")(w, r)	
}

func HandleEditTopic(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.EditTopic, "topics.HandleEditTopic")(w, r)	
}

func HandleDeleteTopic(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.DeleteTopic, "topics.HandleDeleteTopic")(w, r)
}


