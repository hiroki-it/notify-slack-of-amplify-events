package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hiroki-it/notify_slack_of_amplify_events/config"
)

type Request struct {
	Records []struct {
		EventBridge struct {
			Event string `json:"event"`
		}
	}
}

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

type AmplifyBranch struct {
	DisplayName string
}

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(request Request) string {

	config.LoadConfig()

	var event Event

	// EventBridgeから転送されたイベントをマッピングします．
	err := json.Unmarshal([]byte(request.Records[0].EventBridge.Event), &event)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	client, err := NewAmplifyClient()

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	response, err := client.getBranchFromAmplify(event)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	slack := NewSlackClient()

	message := slack.buildMessage(
		event,
		AmplifyBranch{DisplayName: aws.ToString(response.Branch.DisplayName)},
	)

	err = slack.postMessage(message)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	return fmt.Sprintln("Exit")
}
