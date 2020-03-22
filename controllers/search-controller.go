package controllers

import (
	"github.com/Dadard29/go-api-utils/auth"
	"github.com/Dadard29/go-music-researcher/api"
	"github.com/Dadard29/go-music-researcher/managers"
	"net/http"
)

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None
func YoutubeSearchGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	q := r.URL.Query().Get("q")

	resp, err := managers.YoutubeSearchManager(q)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	api.Api.BuildJsonResponse(true, "videos searched", resp, w)
}

func SpotifySearchGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	q := r.URL.Query().Get("q")

	resp, err := managers.SpotifySearchManager(q)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, err.Error(), w)
		return
	}

	api.Api.BuildJsonResponse(true, "music infos retrieved", resp, w)
}
