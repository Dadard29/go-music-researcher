package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	searchUrl = "https://api.spotify.com/v1/search"
)

func (s *Connector) doRequest(url string, httpMethod string, output interface{}) error {
	r, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return err
	}

	if err := s.setHeaders(r); err != nil {
		return err
	}

	resp, err := s.client.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("error processing request, status is " + resp.Status)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, output)
	if err != nil {
		return err
	}

	return nil
}

func (s *Connector) Search(query string, searchType string) (*SearchResponse, error) {
	queryEncoded := url.QueryEscape(query)
	searchTypeEncoded := url.QueryEscape(searchType)
	urlQuery := fmt.Sprintf("%s?q=%s&type=%s", searchUrl, queryEncoded, searchTypeEncoded)

	var respJson SearchResponse
	err := s.doRequest(urlQuery, http.MethodGet, &respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil
}

func (s *Connector) GetArtist(artistId string) (*ArtistResponse, error) {
	urlQuery := "https://api.spotify.com/v1/artists/" + artistId
	var respJson ArtistResponse
	err := s.doRequest(urlQuery, http.MethodGet, &respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil
}

