package unit

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
	"github.com/stretchr/testify/assert"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/test/mock/amplify"
)

/**
 * GetBranchFromAmplifyメソッドをテストします．
 */
func TestGetBranchFromAmplify(t *testing.T) {

	t.Helper()

	log := logger.NewLogger()

	eventDetail, err := dataTestGetBranchFromAmplify()

	if err != nil {
		log.Error(err.Error())
	}

	// AmplifyAPIのモックを作成する．
	mockedAPI := &m_amplify.MockedAmplifyAPI{}

	amplifyClient := amplify.NewAmplifyClient(mockedAPI)

	getBranchInput := amplifyClient.CreateGetBranchInput(eventDetail)

	// スタブに引数として渡される値と，その時の返却値を定義する．
	mockedAPI.On("GetBranch", getBranchInput).Return(
		&aws_amplify.GetBranchOutput{
			Branch: &aws_amplify.Branch{
				DisplayName: aws.String("feature-test"),
			},
		},
		nil,
	)

	// 検証対象の関数を実行する．スタブを含む一連の処理が実行される．
	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(getBranchInput)

	if err != nil {
		log.Error(err.Error())
	}

	//関数内部でスタブがコールされているかを検証する．
	mockedAPI.AssertExpectations(t)

	// 最終的な返却値が正しいかを検証する．
	assert.Exactly(t, aws.String("feature-test"), getBranchOutput.Branch.DisplayName)
}

/**
 * GetBranchFromAmplifyテストデータ
 */
func dataTestGetBranchFromAmplify() (*eventbridge.EventDetail, error) {

	detail, err := file.ReadTestDataFile("../testdata/request/event.json")

	if err != nil {
		return nil, err
	}

	eventDetail := eventbridge.NewEventDetail()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	json.Unmarshal(detail, eventDetail)

	if err != nil {
		return nil, err
	}

	return eventDetail, err
}
