package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing Request Body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(r*http.ResponseWriter,status int, v any) {
	
}