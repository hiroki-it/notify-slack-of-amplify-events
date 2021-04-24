package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/config"
	"github.com/stretchr/testify/assert"
)

type Branch struct {
	DisplayName *string
}

func SlackResponse(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "200")
}

/**
 * ハンドラ関数をテストします．
 */
func TestLambdaHandler(t *testing.T) {

	config.LoadConfig()

	var event eventbridge.Event

	// モックオブジェクトとスタブを定義します．
	api, _ := amplify.NewMockAmplifyAPI()
	api.MockClient.On("mockGetBranchFromAmplify", api, event).Return(Branch{DisplayName: aws.String("feature/test")}, nil)

	response, _ := amplify.MockGetBranchFromAmplify(api, event)

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
