package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
)

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(event events.CloudWatchEvent) (string, error) {

	log := logger.NewLogger()

	eventDetail := eventbridge.NewEventDetail()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal([]byte(event.Detail), eventDetail)

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	amplifyApi, err := amplify.NewAmplifyAPI(os.Getenv("AWS_AMPLIFY_REGION"))

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	amplifyClient := amplify.NewAmplifyClient(amplifyApi)

	getBranchInput := amplifyClient.CreateGetBranchInput(eventDetail)

	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(getBranchInput)

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	jobStatus := slack.NewJobStatus(eventDetail.JobStatus)

	status, color := jobStatus.PrintJobStatus()

	message := slack.NewMessage(
		eventDetail,
		getBranchOutput.Branch,
		status,
		color,
	)

	slackClient := slack.NewSlackClient()

	slackNotification := slack.NewSlackNotification(
		*slackClient,
		*message,
	)

	err = slackNotification.PostMessage()

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	return fmt.Sprint("Succeed to handle request"), nil
}
