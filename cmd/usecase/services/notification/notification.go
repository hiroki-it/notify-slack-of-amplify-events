package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

// SlackNotification Slack通知を構成します．
type SlackNotification struct {
	slackClient  *SlackClient
	slackMessage *SlackMessage
}

// NewSlackNotification コンストラクタ
func NewSlackNotification(slackClient *SlackClient, slackMessage *SlackMessage) *SlackNotification {

	return &SlackNotification{
		slackClient:  slackClient,
		slackMessage: slackMessage,
	}
}

// PostMessage メッセージを送信します
func (no *SlackNotification) PostMessage() error {

	// マッピングを元に，構造体をJSONに変換します．
	sm, err := json.Marshal(no.slackMessage)

	if err != nil {
		return err
	}

	log := logger.NewLogger()

	log.Info(string(sm))

	// リクエストメッセージを定義します．
	req, err := http.NewRequest(
		"POST",
		no.slackClient.url,
		bytes.NewBuffer(sm),
	)

	if err != nil {
		return err
	}

	// ヘッダーを定義します．
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SLACK_API_TOKEN")))

	// HTTPリクエストを送信します．
	res, err := no.slackClient.httpClient.Do(req)

	if err != nil || res.StatusCode != 200 {
		return err
	}

	// deferで宣言しておき，HTTP通信を必ず終了できるようにします．
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	log.Info(string(b))

	return err
}
