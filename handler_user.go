package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	database "github.com/rishabhrawatt/rssagg/neon/db"

	"github.com/google/uuid"
)

// w -- response writer ,r -- pointer to http req

func (apiCfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	if params.Name == "" {
		responseWithError(w, 400, fmt.Sprintf("Benam h kya !: %v", errors.New("cannot keep name empty")))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("could not create user: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// apikey, err := auth.GetApiKey(r.Header)
	// if err != nil {
	// 	responseWithError(w, 403, fmt.Sprintf("Auth error %v", err))
	// 	return
	// }
	// user, err := apiCfg.DB.GetUserByAPIkey(r.Context(), apikey)
	// if err != nil {
	// 	responseWithError(w, 404, fmt.Sprintf("User not found :%v", err))
	// 	return
	// }
	responseWithJSON(w, 200, databaseUserToUser(user))

}
