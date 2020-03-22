package spotify

import (
	"fmt"
	"os"
	"testing"
)

func TestNewSpotifyConnector(t *testing.T) {
	s, err := NewConnector(
		os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
		return
	}

	if !s.isTokenValid() {
		t.Error("invalid token")
		return
	}
}

func TestSpotifyConnector_Search(t *testing.T) {
	s, err := NewConnector(
		os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		t.Error(err)
		return
	}

	results, err := s.Search("chilly gonzales crying", searchTypeTrack)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("len results:", len(results.Tracks.Items))
	if len(results.Tracks.Items) > 0 {
		i := results.Tracks.Items[0]
		fmt.Println(i.Name, i.Artists[0].Name)
	}
}
