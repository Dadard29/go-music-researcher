package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Connector struct {
	token            string
	dateTokenExpired time.Time
	client           *http.Client
	clientId string
	clientSecret string
}

const (
	getTokenUrl = "https://accounts.spotify.com/api/token"
)

const (
	ArtistType = "artist"
	AlbumType = "album"
	TrackType = "track"
)

func (s *Connector) isTokenValid() bool {
	n := time.Now()
	return n.Before(s.dateTokenExpired)
}

func (s *Connector) setHeaders(r *http.Request) error {
	if !s.isTokenValid() {
		err := s.SetToken()
		if err != nil {
			return err
		}
	}

	r.Header.Set("Authorization", "Bearer " + s.token)
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")

	return nil
}

func (s *Connector) SetToken() error {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	r, err := http.NewRequest(http.MethodPost, getTokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.SetBasicAuth(s.clientId, s.clientSecret)

	resp, err := s.client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respJson GetTokenResponse
	err = json.Unmarshal(respData, &respJson)
	if err != nil {
		return err
	}

	s.token = respJson.AccessToken
	s.dateTokenExpired = time.Now().Add(time.Duration(respJson.ExpiresIn) * time.Second)

	return nil
}

func NewConnector(clientId string, clientSecret string) (*Connector, error) {
	s := &Connector{
		client: &http.Client{},
		clientId: clientId,
		clientSecret: clientSecret,
	}

	if err := s.SetToken(); err != nil {
		return nil, err
	}

	return s, nil
}

