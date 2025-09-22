package main

import (
	"encoding/json"
	"net/http"
)

// respondWithJSON sends a JSON response with status code
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(dat); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}

// respondWithError sends an error message (and optional error details) as JSON
// Signature accepts an error argument so handler calls like:
//
//	respondWithError(w, http.StatusBadRequest, "message", err)
//
// will work.
func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	errorMessage := msg
	if err != nil {
		errorMessage = msg + ": " + err.Error()
	}
	respondWithJSON(w, code, map[string]string{"error": errorMessage})
}
