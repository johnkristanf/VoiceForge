package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(res http.ResponseWriter, status int, val any) error {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)

	return json.NewEncoder(res).Encode(val)
}


