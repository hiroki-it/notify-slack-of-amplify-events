package mock

import (
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/stretchr/testify/mock"
)

/**
 * Mockを構成します．
 */
type MockedAmplifyAPI struct {
	amplifyiface.AmplifyAPI
	mock.Mock
}

func (mockedAmplifyAPI *MockedAmplifyAPI) GetBranch(input *aws_amplify.GetBranchInput) (*aws_amplify.GetBranchOutput, error) {
	arguments := mockedAmplifyAPI.Called(input)
	return arguments.Get(0).(*aws_amplify.GetBranchOutput), arguments.Error(1)
}
