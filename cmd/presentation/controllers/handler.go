package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/service/api"
)

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(event events.CloudWatchEvent) (string, error) {

	log := logger.NewLogger()

	eventDetail := &eventbridge.EventDetail{}

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal([]byte(event.Detail), eventDetail)

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	amplifyApi, err := api.NewAmplifyAPI(os.Getenv("AWS_AMPLIFY_REGION"))

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	amplifyClient := amplify.NewAmplifyClient(amplifyApi)

	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(eventDetail)

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	jobStatus := eventbridge.NewJobStatus(eventDetail.JobStatus)

	message := slack.NewMessage(
		eventDetail,
		getBranchOutput.Branch,
		jobStatus,
	)

	slackMessage := message.BuildSlackMessage()

	slackClient := slack.NewSlackClient(
		&http.Client{},
		"https://slack.com/api/chat.postMessage",
	)

	slackNotification := slack.NewSlackNotification(
		slackClient,
		slackMessage,
	)

	err = slackNotification.PostMessage()

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	return fmt.Sprint("Succeed to handle request"), nil
}
