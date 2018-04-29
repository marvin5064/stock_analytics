package stockfetch

import (
	"github.com/parnurzeal/gorequest"
)

type manager struct {
	url       string
	apiKey    string
	requester *gorequest.SuperAgent
}
type Manager interface {
	GetData(symbol string) (string, error)
}

func New(url, apiKey string) Manager {
	return &manager{
		url:       url,
		apiKey:    apiKey,
		requester: gorequest.New(),
	}
}
