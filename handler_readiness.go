package main

import "net/http"

// w -- response writer ,r -- pointer to http req

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct {
		Message string `json:"message"`
	}{
		Message: "Server is healthy",
	})
}
