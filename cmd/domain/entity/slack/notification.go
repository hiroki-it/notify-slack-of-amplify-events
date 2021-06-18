package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

/**
 * Slack通知を構成します．
 */
type SlackNotification struct {
	slackClient  *SlackClient
	slackMessage *SlackMessage
}

/**
 * コンストラクタ
 * SlackNotificationを作成します．
 */
func NewSlackNotification(slackClient *SlackClient, slackMessage *SlackMessage) *SlackNotification {
	return &SlackNotification{
		slackClient:  slackClient,
		slackMessage: slackMessage,
	}
}

/**
 * メッセージを送信します．
 */
func (no *SlackNotification) PostMessage() error {

	// マッピングを元に，構造体をJSONに変換します．
	json, err := json.Marshal(no.slackMessage)

	if err != nil {
		return err
	}

	log := logger.NewLogger()

	log.Info(string(json))

	// リクエストメッセージを定義します．
	req, err := http.NewRequest(
		"POST",
		no.slackClient.url,
		bytes.NewBuffer(json),
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

	body, _ := ioutil.ReadAll(res.Body)

	log.Info(string(body))

	return nil
}
