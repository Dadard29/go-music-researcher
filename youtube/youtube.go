package youtube

import "net/http"

type Connector struct {
	apiKey string
	client *http.Client
}

func NewConnector(apiKey string) *Connector {
	return &Connector{
		apiKey: apiKey,
		client: &http.Client{},
	}
}
