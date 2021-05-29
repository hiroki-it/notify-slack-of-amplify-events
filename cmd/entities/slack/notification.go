package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
)

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

	// マッピングを元に，構造体をJSONに変換する．
	json, err := json.Marshal(no.slackMessage)

	if err != nil {
		return err
	}

	log := logger.NewLogger()

	log.Info(string(json))

	// リクエストメッセージを定義する．
	request, err := http.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage",
		bytes.NewBuffer(json),
	)

	if err != nil {
		return err
	}

	// ヘッダーを定義する．
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SLACK_API_TOKEN")))

	httpClient := &http.Client{}

	// HTTPリクエストを送信する．
	response, err := httpClient.Do(request)

	if err != nil || response.StatusCode != 200 {
		return err
	}

	// deferで宣言しておき，HTTP通信を必ず終了できるようにする．
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Info(string(body))

	return nil
}
