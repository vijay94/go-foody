package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"time"
	controller "foody/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HelloWorld).Methods("GET")
	// router.HandleFunc("/rover", controller.GetRovers).Methods("GET")
	// router.HandleFunc("/rover/operate", controller.OperateRovers).Methods("PUT")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}