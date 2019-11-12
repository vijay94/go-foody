package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"time"
	controller "foody/controllers"
	"os"
	"fmt"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HelloWorld).Methods("GET")
	// router.HandleFunc("/rover", controller.GetRovers).Methods("GET")
	// router.HandleFunc("/rover/operate", controller.OperateRovers).Methods("PUT")

	// server := &http.Server{
	// 	Handler:      router,
	// 	Addr:         getPort(),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	err := http.ListenAndServe(GetPort(), nil)

	// log.Fatal(server.ListenAndServe())
}

func getPort() string {
 	var port = os.Getenv("PORT")
 	// Set a default port if there is nothing in the environment
 	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}