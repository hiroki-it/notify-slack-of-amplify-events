package slack

import (
	"encoding/json"
)

/**
 * イベント構造体
 */
type Request struct {
	Records []struct {
		EventBridge struct {
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
func Handler(request Request) bool {

	var event Event

	// EventBridgeから転送されたイベントをマッピングします．
	err := json.Unmarshal([]byte(request.Records[0].EventBridge.Event), &event)

	if err != nil {
		return err
	}

	branch, err := getAmplifyBranch(event)

	if err != nil {
		return err
	}

	return result
}