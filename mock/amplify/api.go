package amplify

import (
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/stretchr/testify/mock"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

type MockedAmplifyAPI struct {
	amplifyiface.AmplifyAPI
	mock.Mock
}

// GetBranch GetBranchを模倣します．
func (m *MockedAmplifyAPI) GetBranch(input *aws_amplify.GetBranchInput) (*aws_amplify.GetBranchOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*aws_amplify.GetBranchOutput), args.Error(1)
}
