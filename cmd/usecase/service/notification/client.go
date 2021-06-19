package notification

import (
	"net/http"
)

/**
 * SlackClientインターフェースを構成します．
 */
type SlackClientInterface interface {
}

/**
 * SlackClientインターフェースの実装を構成します．
 */
type SlackClient struct {
	SlackClientInterface
	httpClient *http.Client
	url        string
}

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
