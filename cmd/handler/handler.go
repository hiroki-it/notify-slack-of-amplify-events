package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/config"
)

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

	api, err := NewAmplifyAPI()

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	response, err := getBranchFromAmplify(api, event)

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
