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

/**/
type Event struct {
	Version    string   `json:"version"`
	Id         string   `json:"id"`
	DetailType string   `json:"detail-type"`
	Source     string   `json:"source"`
	Account    string   `json:"account"`
	Time       string   `json:"time"`
	Region     string   `json:"region"`
	Resources  []string `json:"resources"`
	Detail     struct {
		AppId      string `json:"appId"`
		BranchName string `json:"branchName"`
		JobId      string `json:"jobId"`
		JobStatus  string `json:"jobStatus"`
	} `json:"detail"`
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
	Client *amplify.Client
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

/**
 * SlackClientを構成します．
 */
type SlackClientInterface interface {
	buildMessage(event Event, amplifyBranch AmplifyBranch) Message
	jobStatusMessage(jobStatus string) (string, string)
	postMessage(message Message) error
}

/**/
type SlackClientImpl struct {
	SlackClientInterface
}
