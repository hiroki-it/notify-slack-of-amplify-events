package slack

import (
	"net/http"
)

/**
 * コンストラクタ
 * SlackClientを作成します．
 */
func NewSlackClient(httpClient *http.Client, url string) *SlackClient {
	return &SlackClient{
		httpClient: httpClient,
		url:        url,
	}
}
