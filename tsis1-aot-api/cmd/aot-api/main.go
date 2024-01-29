package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routes")

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)

}
