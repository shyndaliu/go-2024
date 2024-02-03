package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	aotapi "tsis1-aot-api/pkg/aot-api"

	"github.com/gorilla/mux"
)

type Response struct {
	Titans []aotapi.Titan `json:"titans"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": aotapi.Info()})
}

func Titans(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, aotapi.GetTitans())
}
func Titan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["titan"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "poor request")
		return
	}
	titan, err := aotapi.GetTitan(i)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 not found")
		return
	}
	respondWithJSON(w, http.StatusOK, titan)
}
