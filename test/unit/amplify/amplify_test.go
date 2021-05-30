package unit

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/stretchr/testify/assert"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/test/mock/amplify"
)

/**
 * GetBranchFromAmplifyメソッドをテストします．
 */
func (suite *SuiteAmplify) TestGetBranchFromAmplify() {

	suite.T().Helper()

	// AmplifyAPIのモックを作成する．
	mockedAPI := &m_amplify.MockedAmplifyAPI{}

	amplifyClient := amplify.NewAmplifyClient(mockedAPI)

	getBranchInput := amplifyClient.CreateGetBranchInput(suite.eventDetail)

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
		suite.T().Fatal(err.Error())
	}

	//関数内部でスタブがコールされているかを検証する．
	mockedAPI.AssertExpectations(suite.T())

	// 最終的な返却値が正しいかを検証する．
	assert.Exactly(suite.T(), aws.String("feature-test"), getBranchOutput.Branch.DisplayName)
}
