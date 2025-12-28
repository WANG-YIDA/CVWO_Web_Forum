package auth

import (
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/handlers"
)

const (
	Login = "auth.HandleLogin"
	Register = "auth.HandleRegister"

	ErrEncodeView              = "Failed to encode data in %s"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.Login, Login)(w, r)	
}

func HandleRegister(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return handlers.CreateAPIHandler(api.Register, Register)(w, r)	
}