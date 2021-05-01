package unit

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/exception"
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/test/mock/amplify"
	"github.com/stretchr/testify/assert"
)

/**
 * 関数をテストします．
 */
func TestGetBranchFromAmplify(t *testing.T) {

	detail, _ := ioutil.ReadFile("/test/testdata/event.json")

	eventDetail := new(eventbridge.EventDetail)

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal([]byte(detail), eventDetail)

	if err != nil {
		exception.Error(err)
	}

	// AmplifyAPIのモックを作成する．
	mockedAPI := new(m_amplify.MockedAmplifyAPI)

	amplifyClient := amplify.NewAmplifyClient(mockedAPI)

	getBranchInput := amplifyClient.CreateGetBranchInput(eventDetail)

	// スタブに引数として渡される値と，その時の返却値を定義する．
	mockedAPI.On("GetBranch", getBranchInput).Return(Branch{DisplayName: aws.String("feature-test")}, nil)

	// 検証対象の関数を実行する．スタブを含む一連の処理が実行される．
	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(getBranchInput)

	if err != nil {
		exception.Error(err)
	}

	//関数内部でスタブがコールされているかを検証する．
	mockedAPI.AssertExpectations(t)

	// 最終的な返却値が正しいかを検証する．
	assert.Exactly(t, aws.String("feature-test"), getBranchOutput.Branch.DisplayName)
}
