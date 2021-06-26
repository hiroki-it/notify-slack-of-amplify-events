package notification

import (
	"net/http"
)

type SlackClientInterface interface {
}

type SlackClient struct {
	SlackClientInterface
	httpClient *http.Client
	url        string
}

// NewSlackClient コンストラクタ
func NewSlackClient(httpClient *http.Client, url string) *SlackClient {
	return &SlackClient{
		httpClient: httpClient,
		url:        url,
	}
}
