package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/config"
	"github.com/hiroki-it/notify-slack-of-amplify-events/mock"
	"github.com/stretchr/testify/assert"
)

func SlackResponse(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "200")
}

/**
 * ハンドラ関数をテストします．
 */
func TestLambdaHandler(t *testing.T) {

	input := aws_amplify.GetBranchInput{
		AppId:      aws.String("123456789"),
		BranchName: aws.String("feature/test"),
	}

	api, _ := mock.NewMockedAmplifyAPI()

	// スタブに引数として渡される値と，その時の返却値を定義する．
	api.Client.On("GetBranch", context.TODO(), &input).Return(Branch{DisplayName: aws.String("feature-test")}, nil)

	var event eventbridge.Event

	response, _ := amplify.GetBranchFromAmplify(api, event)

	slack := slack.NewSlackClient()

	message := slack.BuildMessage(
		event,
		amplify.AmplifyBranch{DisplayName: aws.ToString(response.Branch.DisplayName)},
	)

	json, _ := json.Marshal(message)

	request := httptest.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage",
		bytes.NewBuffer(json),
	)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SLACK_API_TOKEN")))

	// HTTPリクエストを送信する．
	writer := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, writer)
}
