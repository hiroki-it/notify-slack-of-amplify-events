package handler

import (
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

/**
 * EventBrdigeのイベントを構成します．
 */
type Request struct {
	Records []struct {
		EventBridge struct {
			Event string `json:"event"`
		}
	}
}

/**
 * EventBrdigeのイベントを構成します．
 */
type AmplifyBranch struct {
	DisplayName string
}

/**
 * AmplifyClientを扱うAPIを構成します．
 */
type AmplifyAPIInterface interface {
	getBranchFromAmplify(api *AmplifyAPIImpl, event Event) (*amplify.GetBranchOutput, error)
}

/**/
type AmplifyAPIImpl struct {
	AmplifyAPIInterface
	Svc *amplify.Client
}

/**
 * Slackに送信するMessageを構成します．
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

/**/
type SlackClientInterface interface {
	buildMessage(event Event, amplifyBranch AmplifyBranch) Message
	jobStatusMessage(jobStatus string) (string, string)
}

/**/
type SlackClientImpl struct {
	SlackClientInterface
}
