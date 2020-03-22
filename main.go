package main

import (
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/service"
	. "github.com/Dadard29/go-music-researcher/api"
	"github.com/Dadard29/go-music-researcher/controllers"
	"github.com/Dadard29/go-music-researcher/repositories"
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"net/http"
)

var routes = service.RouteMapping{
	"/search/youtube": service.Route{
		Description:   "search in youtube video from user query",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.YoutubeSearchGet,
		},
	},
	"/search/spotify": service.Route{
		Description:   "search in spotify music from user query",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.SpotifySearchGet,
		},
	},
}

// ENV:
// - VERSION: ...
// - CORS_ORIGIN: ... (from dockerfile)

// - HOST_SUB: host where to check the sub token
// - YT_API_KEY: the youtube API key
// - SP_CLIENT_ID: the spotify client ID
// - SP_CLIENT_SECRET: the spotify client secret

const (
	SpotifyClientIdKey     = "SP_CLIENT_ID"
	SpotifyClientSecretKey = "SP_CLIENT_SECRET"
	YoutubeApiKeyKey = "YT_API_KEY"
)

func main() {
	var err error
	Api = API.NewAPI("Music-Researcher",
		"config/config.json", routes, true)

	// init the connectors
	controllers.Sc = subChecker.NewSubChecker(Api.Config.GetEnv("HOST_SUB"))
	err = repositories.SetSpotifyConnector(
		Api.Config.GetEnv(SpotifyClientIdKey),
		Api.Config.GetEnv(SpotifyClientSecretKey))
	Api.Logger.CheckErrFatal(err)

	repositories.SetYoutubeConnector(Api.Config.GetEnv(YoutubeApiKeyKey))

	Api.Service.Start()

	Api.Service.Stop()
}
