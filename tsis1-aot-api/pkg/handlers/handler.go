package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tsis1-aot-api/api"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to AOT Api made by Uldana🛡️")
}

func Titans(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	titans := api.Titans

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(titans)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
