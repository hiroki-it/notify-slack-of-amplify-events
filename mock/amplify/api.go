package amplify

import (
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/stretchr/testify/mock"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * モックを構成します．
 */
type MockedAmplifyAPI struct {
	amplifyiface.AmplifyAPI
	mock.Mock
}

func (m *MockedAmplifyAPI) GetBranch(input *aws_amplify.GetBranchInput) (*aws_amplify.GetBranchOutput, error) {
	arguments := m.Called(input)
	return arguments.Get(0).(*aws_amplify.GetBranchOutput), arguments.Error(1)
}
