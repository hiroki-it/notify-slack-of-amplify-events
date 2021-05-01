package unit

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/test/mock/amplify"
	"github.com/stretchr/testify/assert"
)

/**
 * 関数をテストします．
 */
func TestGetBranchFromAmplify(t *testing.T) {

	input := aws_amplify.GetBranchInput{
		AppId:      aws.String("123456789"),
		BranchName: aws.String("feature/test"),
	}

	mockedAPI := new(m_amplify.MockedAmplifyAPI)

	// スタブに引数として渡される値と，その時の返却値を定義する．
	mockedAPI.On("GetBranch", &input).Return(Branch{DisplayName: aws.String("feature-test")}, nil)

	client := amplify.NewAmplifyClient(mockedAPI)

	eventDetail := &eventbridge.EventDetail{
		AppId:      "123456789",
		BranchName: "feature/test",
		JobId:      "123456789",
		JobStatus:  "SUCCEED",
	}

	// 検証対象の関数を実行する．スタブを含む一連の処理が実行される．
	response, err := client.GetBranchFromAmplify(eventDetail)

	if err != nil {
		exception.Error(err)
	}

	//関数内部でスタブがコールされているかを検証する．
	mockedAPI.AssertExpectations(t)

	// 最終的な返却値が正しいかを検証する．
	assert.Exactly(t, aws.String("feature-test"), response.Branch.DisplayName)
}
