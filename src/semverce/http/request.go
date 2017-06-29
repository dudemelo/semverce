package http

import (
	"encoding/json"
	"net/http"
)

func JsonRequest(request *http.Request, resource interface{}) error {

	decoder := json.NewDecoder(request.Body)

	if decodedJson := decoder.Decode(resource); decodedJson != nil {
		return decodedJson
	}

	return nil
}