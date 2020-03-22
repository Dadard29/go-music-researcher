package repositories

import (
	"github.com/Dadard29/go-music-researcher/spotify"
	"github.com/Dadard29/go-music-researcher/youtube"
)

var youtubeConnector *youtube.Connector
var spotifyConnector *spotify.Connector

func SetYoutubeConnector(apiKey string) {
	youtubeConnector = youtube.NewConnector(apiKey)
}

func SetSpotifyConnector(clientId string, clientSecret string) error {
	var err error
	spotifyConnector, err = spotify.NewConnector(clientId, clientSecret)
	return err
}
