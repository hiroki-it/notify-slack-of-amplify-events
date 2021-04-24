package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/config"
)

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(request Request) string {

	config.LoadConfig()

	var event eventbridge.Event

	// eventbridgeから転送されたイベントをマッピングします．
	err := json.Unmarshal([]byte(request.Records[0].EventBridge.Event), &event)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	api, err := amplify.NewAmplifyAPI()

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	response, err := amplify.GetBranchFromAmplify(api, event)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	slack := slack.NewSlackClient()

	message := slack.BuildMessage(
		event,
		amplify.AmplifyBranch{DisplayName: aws.ToString(response.Branch.DisplayName)},
	)

	err = slack.PostMessage(message)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	return fmt.Sprintln("Exit")
}
