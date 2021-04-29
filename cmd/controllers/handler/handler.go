package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/exception"
	"github.com/hiroki-it/notify-slack-of-amplify-events/configs"
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
		return exception.Error(err)
	}

	api, err := amplify.NewAmplifyAPI(os.Getenv("AWS_REGION"))

	if err != nil {
		return exception.Error(err)
	}

	client := amplify.NewAmplifyClient(api)

	response, err := client.GetBranchFromAmplify(event)

	if err != nil {
		return exception.Error(err)
	}

	slackClient := slack.NewSlackClient()

	message := slackClient.BuildMessage(
		event,
		slack.AmplifyBranch{DisplayName: aws.StringValue(response.Branch.DisplayName)},
	)

	err = slackClient.PostMessage(message)

	if err != nil {
		return exception.Error(err)
	}

	return fmt.Sprintln("Exit")
}
