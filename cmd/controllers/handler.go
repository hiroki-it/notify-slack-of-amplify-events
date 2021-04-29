package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
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

	sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")})

	if err != nil {
		return fmt.Sprintf("ERROR: %#v\n", err)
	}

	api := aws_amplify.New(sess)

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
