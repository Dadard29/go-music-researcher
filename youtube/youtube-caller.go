package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (y *Connector) doRequest(url string, httpMethod string, output interface{}) error {
	r, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		return err
	}

	resp, err := y.client.Do(r)
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

func (y *Connector) Search(query string) (*SearchResponse, error) {
	url := "https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&key=%s"
	urlQuery := fmt.Sprintf(url, query, y.apiKey)

	var respJson SearchResponse
	err := y.doRequest(urlQuery, http.MethodGet, &respJson)
	if err != nil {
		return nil, err
	}

	return &respJson, nil

}
