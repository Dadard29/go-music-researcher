package repositories

import (
	"github.com/Dadard29/go-music-researcher/spotify"
	"github.com/Dadard29/go-music-researcher/youtube"
)

func YoutubeSearch(query string) (youtube.SearchResponse, error) {
	var y youtube.SearchResponse
	r, err := youtubeConnector.Search(query)
	if err != nil {
		return y, err
	}

	return *r, nil
}

func YoutubeGetVideo(videoId string) (youtube.GetVideoResponse, error) {
	var y youtube.GetVideoResponse
	r, err := youtubeConnector.GetVideo(videoId)
	if err != nil {
		return y, err
	}

	return *r, err
}

func SpotifySearch(query string) (spotify.SearchResponse, error) {
	var s spotify.SearchResponse
	r, err := spotifyConnector.Search(query, spotify.TrackType)
	if err != nil {
		return s, err
	}

	return *r, nil
}

func SpotifyGetArtist(artistId string) (spotify.ArtistResponse, error) {
	var a spotify.ArtistResponse
	r, err := spotifyConnector.GetArtist(artistId)
	if err != nil {
		return a, err
	}

	return *r, nil
}
