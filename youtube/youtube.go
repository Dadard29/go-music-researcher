package youtube

import "net/http"

type Connector struct {
	RequestsDone int
	apiKey       string
	client       *http.Client
}

func NewConnector(apiKey string) *Connector {
	return &Connector{
		RequestsDone: 0,
		apiKey:       apiKey,
		client:       &http.Client{},
	}
}
