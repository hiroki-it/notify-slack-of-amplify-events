package notification

import (
	"fmt"
	"os"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/entities"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

type Message struct {
	detail *entities.Detail
	branch *aws_amplify.Branch
}

type SlackMessage struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type     string    `json:"type"`
	Text     *Text     `json:"text,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// NewMessage コンストラクタ
func NewMessage(detail *entities.Detail, branch *aws_amplify.Branch) *Message {

	return &Message{
		detail: detail,
		branch: branch,
	}
}

// BuildSlackMessage コンストラクタ
func (message *Message) BuildSlackMessage() *SlackMessage {

	// メッセージを構成します．
	return &SlackMessage{
		Channel: os.Getenv("SLACK_CHANNEL_ID"),
		Text:    "検証用dev環境",
		Attachments: []Attachment{
			{
				Color: message.ColorCode(),
				Blocks: []Block{
					{
						Type: "section",
						Text: &Text{
							Type: "mrkdwn",
							Text: "*検証用dev環境*",
						},
					},
					{
						Type: "context",
						Elements: []Element{
							{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*結果*: %s",
									message.detail.JobStatusType().String(),
								),
							},
						},
					},
					{
						Type: "context",
						Elements: []Element{
							{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*ブランチ名*: %s",
									message.detail.BranchName().Name(),
								),
							},
						},
					},
					{
						Type: "context",
						Elements: []Element{
							{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*プルリクURL*: https://github.com/hiroki-it/notify-slack-of-amplify-events/compare/%s",
									message.detail.BranchName().Name(),
								),
							},
						},
					},
					{
						Type: "context",
						Elements: []Element{
							{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*検証URL*: https://%s.%s.amplifyapp.com",
									*message.branch.DisplayName,
									message.detail.AppId().Id(),
								),
							},
						},
					},
					{
						Type: "context",
						Elements: []Element{
							{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									":amplify: <https://%s.console.aws.amazon.com/amplify/home?region=%s#/%s/%s/%s|*Amplifyコンソール画面はこちら*>",
									os.Getenv("AWS_AMPLIFY_REGION"),
									os.Getenv("AWS_AMPLIFY_REGION"),
									message.detail.AppId().Id(),
									*message.branch.DisplayName,
									message.detail.JobId().Id(),
								),
							},
						},
					},
					{
						Type: "divider",
					},
				},
			},
		},
	}
}

// ColorCode ジョブステータスを表現する色を返却します.
func (message *Message) ColorCode() string {

	if message.detail.JobStatusType().IsSucceed() {
		return "#00FF00"
	}

	return "#ff0000"
}
