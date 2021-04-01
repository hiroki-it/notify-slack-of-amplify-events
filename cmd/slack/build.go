package slack

import (
	"fmt"
	"os"
)

/**
 * Message構造体
 */
type Message struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

/**
 * Attachmentスライスの要素
 */
type Attachment struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

/**
 * Blockスライスの要素
 */
type Block struct {
	Type string `json:"type"`
	Text struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"text,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

/**
 * Elementスライスの要素
 */
type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**
 * Slackに送信するメッセージを構成します．
 */
func buildMessage(event Event, branch Branch) Message {

	status, color := jobStatusMessage(event.Detail.JobStatus)

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
func jobStatusMessage(jobStatus string) (string, string) {

	if jobStatus == "SUCCEED" {
		return "成功", "#00FF00"
	}

	return "失敗", "#ff0000"
}
