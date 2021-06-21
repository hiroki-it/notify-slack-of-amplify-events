package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/detail"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/service/notification"
)

/**
 * Lambdaハンドラー関数
 */
func HandleRequest(eventBridge events.CloudWatchEvent) (string, error) {

	log := logger.NewLogger()

	Detail := detail.NewDetail()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal([]byte(eventBridge.Detail), Detail)

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

	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(Detail)

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
	}

	message := notification.NewMessage(
		Detail,
		getBranchOutput.Branch,
	)

	slackMessage := message.BuildSlackMessage()

	slackClient := notification.NewSlackClient(
		&http.Client{},
		"https://slack.com/api/chat.postMessage",
	)

	slackNotification := notification.NewSlackNotification(
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
