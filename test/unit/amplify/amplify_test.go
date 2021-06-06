package unit

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/test/mock/amplify"
)

/**
 * GetBranchFromAmplifyメソッドをテストします．
 */
func (suite *SuiteAmplify) TestGetBranchFromAmplify() {

	suite.T().Helper()

	// AmplifyAPIのスタブを作成する．
	mockedAPI := &m_amplify.MockedAmplifyAPI{}

	// スタブのメソッドに処理の内容を定義する．
	mockedAPI.On("GetBranch", mock.Anything).Return(
		&aws_amplify.GetBranchOutput{
			Branch: &aws_amplify.Branch{
				DisplayName: aws.String("feature-test"),
			},
		},
		nil,
	)

	amplifyClient := amplify.NewAmplifyClient(mockedAPI)

	// 検証対象の関数を実行する．スタブを含む一連の処理が実行される．
	getBranchOutput, err := amplifyClient.GetBranchFromAmplify(suite.getBranchFromAmplifyData)

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	// 関数内部でスタブのメソッドがコールされているかを検証する．
	mockedAPI.AssertExpectations(suite.T())

	// 最終的な返却値が正しいかを検証する．
	assert.Exactly(suite.T(), aws.String("feature-test"), getBranchOutput.Branch.DisplayName)
}
