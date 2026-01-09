package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WANG-YIDA/CVWO_Web_Forum/internal/api"
	"github.com/pkg/errors"
)

const (
	ErrEncodeView = "Failed to encode data in %s"
)

func CreateAPIHandler(
	apiFunc func(http.ResponseWriter, *http.Request, *sql.DB) (interface{}, error),
	handlerName string,
) func(http.ResponseWriter, *http.Request, *sql.DB) (*api.Response, error) {
	return func(w http.ResponseWriter, r *http.Request, db *sql.DB) (*api.Response, error) {
		result, err := apiFunc(w, r, db)
		if err != nil {
			return nil, err
		}
		
		data, err := json.Marshal(result)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, handlerName))
		}
		
		return &api.Response{
			Payload: api.Payload{
				Data: data,
			},
			Success: true,
		}, nil
	}
}