package utils

import (
	"encoding/json"
	"io"
	"net/http"
)


func DecodeRequestBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), &x)
	if err != nil {
		return err
	}	
	return nil
}