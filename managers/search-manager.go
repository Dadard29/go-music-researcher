package managers

import (
	"errors"
	"github.com/Dadard29/go-music-researcher/repositories"
	"github.com/Dadard29/go-music-researcher/spotify"
	"github.com/Dadard29/go-music-researcher/youtube"
)

var limit = 5

func YoutubeSearchManager(query string) ([]youtube.VideoJson, error) {

	var y []youtube.VideoJson

	r, err := repositories.YoutubeSearch(query)
	if err != nil {
		logger.Error(err.Error())
		return y, errors.New("error searching youtube video")
	}

	y = make([]youtube.VideoJson, 0)
	for i, item := range r.Items {
		if i > limit {
			break
		}

		y = append(y, youtube.VideoJson{
			VideoTitle:        item.Snippet.Title,
			VideoId:           item.ID.VideoID,
			VideoThumbnailUrl: item.Snippet.Thumbnails.Default.URL,
			ChannelName:       item.Snippet.ChannelTitle,
			PublishedAt:       item.Snippet.PublishedAt,
		})
	}

	return y, nil
}

func SpotifySearchManager(query string) ([]spotify.TrackJson, error) {
	var s []spotify.TrackJson

	r, err := repositories.SpotifySearch(query)
	if err != nil {
		logger.Error(err.Error())
		return s, errors.New("error searching spotify music")
	}

	s = make([]spotify.TrackJson, 0)
	for i, item := range r.Tracks.Items {
		if i > limit {
			break
		}

		artistId := item.Artists[0].ID
		artist, err := repositories.SpotifyGetArtist(artistId)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		s = append(s, spotify.TrackJson{
			Title:  item.Name,
			Artist: item.Artists[0].Name,
			Album:  item.Album.Name,
			Date:   item.Album.ReleaseDate,
			Genre:  artist.Genres,
		})

	}

	return s, nil
}
