package unit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/exception"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/test/mock/amplify"
	"github.com/stretchr/testify/assert"
)

func SlackResponse(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "200")
}

/**
 * 関数をテストします．
 */
func TestLambdaHandler(t *testing.T) {

	detail := file.ReadTestData("./testdata/event.json")

	eventDetail := new(eventbridge.EventDetail)

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal(detail, eventDetail)

	if err != nil {
		exception.Error(err)
	}

	mockedAPI := new(m_amplify.MockedAmplifyAPI)

	amplifyClient := amplify.NewAmplifyClient(mockedAPI)

	getBranchInput := amplifyClient.CreateGetBranchInput(eventDetail)

	// スタブに引数として渡される値と，その時の返却値を定義する．
	mockedAPI.On("GetBranch", getBranchInput).Return(
		aws_amplify.GetBranchOutput{
			Branch: &aws_amplify.Branch{
				DisplayName: aws.String("feature-test"),
			},
		},
		nil,
	)

	// 検証対象の関数を実行する．スタブを含む一連の処理が実行される．
	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(getBranchInput)

	slackClient := slack.NewSlackClient()

	message := slackClient.BuildMessage(
		eventDetail,
		getBranchOutput.Branch,
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
