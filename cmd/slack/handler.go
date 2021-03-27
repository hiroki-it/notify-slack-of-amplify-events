package slack

import (
	"encoding/json"
)

/**
 * Amplifyイベント構造体
 */
type Request struct {
	Records []struct {
		Amplify struct {
			Event string `json:event`
		}
	}
}

/**
 * Event構造体
 */
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
func Handler(request Request) error {

	var event Event

	// EventBridgeから転送されたイベントをマッピングします．
	err := json.Unmarshal([]byte(request.Records[0].Amplify.Event), &event)

	if err != nil {
		return err
	}

	message := buildMessage(event)

	return postMessage(message)
}
