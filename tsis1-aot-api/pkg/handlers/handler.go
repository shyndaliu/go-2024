package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to AOT Api made by Uldana🛡️")
}

func Titans(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var response Titan

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
