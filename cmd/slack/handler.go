package slack

import (
	"encoding/json"
	"fmt"
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

/**
 * ハンドラー関数
 */
func Handler(request Request) string {

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

	branch, err := client.getBranchFromAmplify(event)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	message := buildMessage(event, branch)

	err = postMessage(message)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	return fmt.Sprintln("Exit")
}
