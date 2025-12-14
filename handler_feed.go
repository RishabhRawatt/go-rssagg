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

func (apiCfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
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

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("could not create feed: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("could not get feeds :%v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedToFeeds(feeds))
}
