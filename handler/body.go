package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

// UnmarshalRequestBody read request body and unmarshal to instance, instance must be a reference
// Request body will be close
func UnmarshalRequestBody(r *http.Request, instance interface{}) error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	return json.Unmarshal(body, instance)
}
