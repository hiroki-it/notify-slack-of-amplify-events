package notification

import (
	"net/http"
)

type SlackClientInterface interface {
}

type SlackClient struct {
	httpClient *http.Client
	url        string
}

var _ SlackClientInterface = &SlackClient{}

// NewSlackClient コンストラクタ
func NewSlackClient(httpClient *http.Client, url string) *SlackClient {
	return &SlackClient{
		httpClient: httpClient,
		url:        url,
	}
}
