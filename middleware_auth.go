package main

import (
	"fmt"
	"net/http"

	"github.com/rishabhrawatt/rssagg/neon/auth"
	database "github.com/rishabhrawatt/rssagg/neon/db"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIkey(r.Context(), apikey)
		if err != nil {
			responseWithError(w, 404, fmt.Sprintf("User not found :%v", err))
			return
		}
		handler(w, r, user)
	}
}
