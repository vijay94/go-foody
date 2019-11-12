package controllers

import (
	"net/http"
	"foody/utils"
	"foody/responses"
)

func HelloWorld(response http.ResponseWriter, request *http.Request) {

	utils.EncodeToJson(response,responses.StandardResponse { "Success" })
}