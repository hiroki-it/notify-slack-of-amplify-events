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
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/hiroki-it/notify_slack_of_amplify_events/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAmplifyAPIImpl struct {
	AmplifyAPIInterface
	mock.Mock
}

func NewMockAmplifyAPI() (*MockAmplifyAPIImpl, error) {
	return new(MockAmplifyAPIImpl), nil
}

func (mock *MockAmplifyAPIImpl) getBranchFromAmplify(event Event) (*amplify.GetBranchOutput, error) {
	arguments := mock.Called(event)
	return arguments.Get(0).(*amplify.GetBranchOutput), arguments.Error(1)
}

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
	client, _ := NewMockAmplifyAPI()
	client.On("getBranchFromAmplify", event).Return(Branch{DisplayName: aws.String("feature/test")}, nil)

	response, _ := client.getBranchFromAmplify(event)

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
