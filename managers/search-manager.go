package managers

import (
	"errors"
	"github.com/Dadard29/go-music-researcher/repositories"
	"github.com/Dadard29/go-music-researcher/spotify"
	"github.com/Dadard29/go-music-researcher/youtube"
	"strconv"
)

var limit = 4

func youtubeGetViewsCount(videoId string) int {
	videoResponse, err := repositories.YoutubeGetVideo(videoId)
	if err != nil {
		logger.Warning("error getting video")
		logger.Warning(err.Error())
		return -1
	}

	if len(videoResponse.Items) == 0 {
		logger.Warning("no video found for id " + videoId)
		return -1
	}

	viewsCountStr := videoResponse.Items[0].Statistics.ViewCount
	viewsCount, err := strconv.Atoi(viewsCountStr)
	if err != nil {
		logger.Warning("error converting views count to int")
		logger.Warning(err.Error())
		return -1
	}

	return viewsCount
}

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

		viewsCount := youtubeGetViewsCount(item.ID.VideoID)

		y = append(y, youtube.VideoJson{
			VideoTitle:        item.Snippet.Title,
			VideoId:           item.ID.VideoID,
			VideoThumbnailUrl: item.Snippet.Thumbnails.Default.URL,
			ChannelName:       item.Snippet.ChannelTitle,
			PublishedAt:       item.Snippet.PublishedAt,
			ViewsCount:        viewsCount,
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
			ID: item.ID,
			Title:  item.Name,
			Artist: item.Artists[0].Name,
			Album:  item.Album.Name,
			Date:   item.Album.ReleaseDate,
			Genre:  artist.Genres,
			ImageURL: item.Album.Images[0].URL,
			PreviewURL: item.PreviewURL,
		})

	}

	return s, nil
}
