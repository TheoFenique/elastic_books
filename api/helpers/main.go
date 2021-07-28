package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func EncodeData(w http.ResponseWriter, data interface{}) {
	var newRes response
	newRes.Data = data
	json.NewEncoder(w).Encode(newRes)
}

func EncodeError(w http.ResponseWriter, err error) {
	var newRes response
	newRes.Error = err.Error()
	json.NewEncoder(w).Encode(newRes)
}
