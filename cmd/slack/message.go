package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type string `json:"type"`
	Text struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"text,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type SlackClientInterface interface {
	buildMessage(event Event, branch Branch) Message
	jobStatusMessage(jobStatus string) (string, string)
}

type SlackClientImpl struct {
	SlackClientInterface
}

/**
 * コンストラクタ
 */
func NewSlackClient() *SlackClientImpl {
	return new(SlackClientImpl)
}

/**
 * Slackに送信するメッセージを構成します．
 */
func (slack SlackClientImpl) buildMessage(event Event, branch Branch) Message {

	status, color := slack.jobStatusMessage(event.Detail.JobStatus)

	// メッセージを構成します．
	return Message{
		Channel: os.Getenv("SLACK_CHANNEL_ID"),
		Text:    "検証用dev環境",
		Attachments: []Attachment{
			Attachment{
				Color: color,
				Blocks: []Block{
					Block{
						Type: "section",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: "*検証用dev環境*",
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*結果*: %s",
									status,
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*ブランチ名*: %s",
									event.Detail.BranchName,
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*プルリクURL*: https://github.com/Hiroki-IT/notify_slack_of_amplify_events/compare/%s",
									event.Detail.BranchName,
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*検証URL*: https://%s.%s.amplifyapp.com",
									branch.DisplayName,
									event.Detail.AppId,
								),
							},
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									":amplify: <https://%s.console.aws.amazon.com/amplify/home?region=%s#/%s/%s/%s|*Amplifyコンソール画面はこちら*>",
									event.Region,
									event.Region,
									event.Detail.AppId,
									event.Detail.BranchName,
									event.Detail.JobId,
								),
							},
						},
					},
					Block{
						Type: "divider",
					},
				},
			},
		},
	}
}

/**
 * ジョブ状態を表現するメッセージを返却します．
 */
func (slack SlackClientImpl) jobStatusMessage(jobStatus string) (string, string) {

	if jobStatus == "SUCCEED" {
		return "成功", "#00FF00"
	}

	return "失敗", "#ff0000"
}

/**
 * メッセージを送信します．
 */
func (slack SlackClientImpl) postMessage(message Message) error {

	// マッピングを元に，構造体をJSONに変換する．
	json, err := json.Marshal(message)

	if err != nil {
		return err
	}

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

	client := &http.Client{}

	// HTTPリクエストを送信する．
	response, err := client.Do(request)

	if err != nil || response.StatusCode != 200 {
		return err
	}

	// deferで宣言しておき，HTTP通信を必ず終了できるようにする．
	defer response.Body.Close()

	fmt.Printf("Success: %#v\n", response)

	return nil
}
