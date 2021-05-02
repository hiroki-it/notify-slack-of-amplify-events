package handler

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/exception"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
)

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(event events.CloudWatchEvent) {

	file.ReadEnvFile()

	var eventDetail *eventbridge.EventDetail

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal([]byte(event.Detail), eventDetail)

	if err != nil {
		exception.Error(err)
	}

	amplifyApi, err := amplify.NewAmplifyAPI(os.Getenv("AWS_LAMBDA_REGION"))

	if err != nil {
		exception.Error(err)
	}

	amplifyClient := amplify.NewAmplifyClient(amplifyApi)

	getBranchInput := amplifyClient.CreateGetBranchInput(eventDetail)

	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(getBranchInput)

	if err != nil {
		exception.Error(err)
	}

	slackClient := slack.NewSlackClient()

	message := slackClient.BuildMessage(
		eventDetail,
		getBranchOutput.Branch,
	)

	err = slackClient.PostMessage(message)

	if err != nil {
		exception.Error(err)
	}

	log.Println("Exit")
}
