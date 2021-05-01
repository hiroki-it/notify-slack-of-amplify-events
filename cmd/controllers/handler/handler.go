package handler

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
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
func HandleRequest(event events.CloudWatchEvent) {

	config.LoadConfig()

	var eventDetail *eventbridge.EventDetail

	// eventbridgeから転送されたイベントをマッピングします．
	err := json.Unmarshal([]byte(event.Detail), eventDetail)

	if err != nil {
		exception.Error(err)
	}

	amplifyApi, err := amplify.NewAmplifyAPI(os.Getenv("AWS_LAMBDA_REGION"))

	if err != nil {
		exception.Error(err)
	}

	amplifyClient := amplify.NewAmplifyClient(amplifyApi)

	response, err := amplifyClient.GetBranchFromAmplify(eventDetail)

	if err != nil {
		exception.Error(err)
	}

	slackClient := slack.NewSlackClient()

	message := slackClient.BuildMessage(
		eventDetail,
		&slack.AmplifyBranch{DisplayName: aws.StringValue(response.Branch.DisplayName)},
	)

	err = slackClient.PostMessage(message)

	if err != nil {
		exception.Error(err)
	}

	log.Println("Exit")
}
