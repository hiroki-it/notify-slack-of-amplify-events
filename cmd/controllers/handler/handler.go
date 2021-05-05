package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/exception"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
)

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(event events.CloudWatchEvent) string {

	eventDetail := eventbridge.NewEventDetail()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal([]byte(event.Detail), eventDetail)

	if err != nil {
		logger.ErrorLog(exception.NewException(err, "Failed to parse json."))
		return fmt.Sprintln("Failed to handle request")
	}

	amplifyApi, ex := amplify.NewAmplifyAPI(os.Getenv("AWS_REGION"))

	if ex != nil {
		logger.ErrorLog(ex)
		return fmt.Sprintln("Failed to handle request")
	}

	amplifyClient := amplify.NewAmplifyClient(amplifyApi)

	getBranchInput := amplifyClient.CreateGetBranchInput(eventDetail)

	getBranchOutput, ex := amplifyClient.GetBranchFromAmplify(getBranchInput)

	if ex != nil {
		logger.ErrorLog(ex)
		return fmt.Sprintln("Failed to handle request")
	}

	slackClient := slack.NewSlackClient()

	message := slackClient.BuildMessage(
		eventDetail,
		getBranchOutput.Branch,
	)

	ex = slackClient.PostMessage(message)

	if ex != nil {
		logger.ErrorLog(ex)
		return fmt.Sprintln("Failed to handle request")
	}

	return fmt.Sprintln("Succeed to handle request")
}
