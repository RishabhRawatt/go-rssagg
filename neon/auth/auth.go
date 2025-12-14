package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts api from header from http req
// example
// Authorization:ApiKey {insert apiKey here}
func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")

	if val == "" {
		return "", errors.New("no auth info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("wrong auth key")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("wrong auth key or apikey header not found")
	}
	return vals[1], nil
}
