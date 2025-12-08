package main

import "net/http"

// w -- response writer ,r -- pointer to http req

func handlerError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "something went wrong")
}
