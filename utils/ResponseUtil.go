package utils

import (
	"net/http"
	"encoding/json"
)

type GenericResponse struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

func EncodeToJson(response http.ResponseWriter, class interface{}) {

	output, err := json.Marshal(GenericResponse { 200, class })
	if err != nil {
		http.Error(response, err.Error(), 500)
		return
	}
	response.Header().Set("content-type", "application/json")
	response.Write(output)
}