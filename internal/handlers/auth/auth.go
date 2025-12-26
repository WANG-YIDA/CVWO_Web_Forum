package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	auth "github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/pkg/errors"
)

const (
	Login = "auth.HandleLogin"
	Register = "auth.HandleRegister"

	SuccessfulLoginMessage     = "Login successfully"
	SuccessfulRegisterMessage  = "Register successfully"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrLogin 				   = "Failed to login in %s"
	ErrRegister				   = "Failed to register in %s"
	ErrEncodeView              = "Failed to encode data in %s"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	result, err := auth.Login(w, r)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, Login))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulLoginMessage},
	}, nil
}

func HandleRegister(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	result, err := auth.Register(w, r)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRegister, Register))
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulRegisterMessage},
	}, nil
}