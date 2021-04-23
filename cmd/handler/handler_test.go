package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hiroki-it/notify_slack_of_amplify_events/config"
	"github.com/stretchr/testify/assert"
)

type Branch struct {
	DisplayName *string
}

/**
 * ハンドラ関数をテストします．
 */
func TestLambdaHandler(t *testing.T) {

	config.LoadConfig()

	var event Event

	// モックオブジェクトとスタブを定義します．
	api, _ := NewMockAmplifyAPI()
	api.On("getMockBranchFromAmplify", event).Return(Branch{DisplayName: aws.String("feature/test")}, nil)

	response, _ := getMockBranchFromAmplify(api, event)

	slack := NewSlackClient()

	message := slack.buildMessage(
		event,
		AmplifyBranch{DisplayName: aws.ToString(response.Branch.DisplayName)},
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

	assert.Equal(t, http.StatusOK, writer.Code)
}
