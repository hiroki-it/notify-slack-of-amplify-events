package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
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

	api, err := amplify.NewAmplifyAPI(os.Getenv("AWS_REGION"))

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	client := amplify.NewAmplifyClient(api)

	response, err := amplify.GetBranchFromAmplify(client, event)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	slackClient := slack.NewSlackClient()

	message := slackClient.BuildMessage(
		event,
		slack.AmplifyBranch{DisplayName: aws.StringValue(response.Branch.DisplayName)},
	)

	err = slackClient.PostMessage(message)

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	return fmt.Sprintln("Exit")
}
