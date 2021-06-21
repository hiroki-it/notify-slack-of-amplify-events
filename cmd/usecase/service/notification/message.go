package notification

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/detail"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * メッセージを構成します．
 */
type Message struct {
	detail *detail.Detail
	branch *aws_amplify.Branch
}

/**
 * Slackメッセージを構成します．
 */
type SlackMessage struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

/**/
type Attachment struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

/**/
type Block struct {
	Type     string    `json:"type"`
	Text     *Text     `json:"text,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

/**/
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**/
type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**
 * コンストラクタ
 * Messageを作成します．
 */
func NewMessage(detail *detail.Detail, branch *aws_amplify.Branch) *Message {

	return &Message{
		detail: detail,
		branch: branch,
	}
}

/**
 * コンストラクタ
 * Messageを作成します．
 */
func (message *Message) BuildSlackMessage() *SlackMessage {

	// メッセージを構成します．
	return &SlackMessage{
		Channel: os.Getenv("SLACK_CHANNEL_ID"),
		Text:    "検証用dev環境",
		Attachments: []Attachment{
			Attachment{
				Color: message.ColorCode(),
				Blocks: []Block{
					Block{
						Type: "section",
						Text: &Text{
							Type: "mrkdwn",
							Text: "*検証用dev環境*",
						},
					},
					Block{
						Type: "context",
						Elements: []Element{
							Element{
								Type: "mrkdwn",
								Text: fmt.Sprintf(
									"*結果*: %s",
									message.StatusWord(),
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
									message.detail.BranchName,
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
									"*プルリクURL*: https://github.com/hiroki-it/notify-slack-of-amplify-events/compare/%s",
									message.detail.BranchName,
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
									aws.StringValue(message.branch.DisplayName),
									message.detail.AppId,
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
									os.Getenv("AWS_AMPLIFY_REGION"),
									os.Getenv("AWS_AMPLIFY_REGION"),
									message.detail.AppId,
									aws.StringValue(message.branch.DisplayName),
									message.detail.JobId,
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
 * ジョブステータスを表現する文言を返却します．
 */
func (message *Message) StatusWord() string {

	if message.Detail.IsSucceed() {
		return "成功"
	}

	return "失敗"
}

/**
 * ジョブステータスを表現する色を返却します．
 */
func (message *Message) ColorCode() string {

	if message.Detail.IsSucceed() {
		return "#00FF00"
	}

	return "#ff0000"
}
