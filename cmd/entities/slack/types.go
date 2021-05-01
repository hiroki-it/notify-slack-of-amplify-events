package slack

import (
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
)

/**
 * Slackに送信するメッセージデータを構成します．
 */
type Message struct {
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
	Type string `json:"type"`
	Text struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"text,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

/**/
type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

/**
 * SlackClientインターフェースを構成します．
 */
type SlackClientInterface interface {
	buildMessage(*eventbridge.EventDetail, *aws_amplify.Branch) Message
	jobStatusMessage(string) (string, string)
	postMessage(Message) error
}

/**
 * SlackClientインターフェースの実装を構成します．
 */
type SlackClient struct {
	SlackClientInterface
}
