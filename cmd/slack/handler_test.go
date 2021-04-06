package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAmplifyClientImpl struct {
	AmplifyClientInterface
	mock.Mock
}

func NewMockAmplifyClient() (*MockAmplifyClientImpl, error) {
	return new(MockAmplifyClientImpl), nil
}

func (mock *MockAmplifyClientImpl) getBranchFromAmplify(event Event) (*amplify.GetBranchOutput, error) {
	arguments := mock.Called(event)
	return arguments.Get(0).(*amplify.GetBranchOutput), arguments.Error(1)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "ping")
}

type Branch struct {
	DisplayName *string
}

/**
 * ハンドラ関数をテストします．
 */
func testLambdaHandler(t *testing.T) {

	var event Event

	// モックオブジェクトとスタブを定義します．
	client, _ := NewMockAmplifyClient()
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

	handler(writer, request)

	assert.Equal(t, http.StatusOK, writer.Code)
}
