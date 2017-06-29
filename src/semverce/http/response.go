package http

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(response http.ResponseWriter, resource interface{}) {

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)

	json.NewEncoder(response).Encode(resource)
}