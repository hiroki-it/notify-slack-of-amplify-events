package amplify

import (
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

func (m *MockedAmplifyAPI) GetBranch(input *aws_amplify.GetBranchInput) (*aws_amplify.GetBranchOutput, error) {
	arguments := m.Called(input)
	return arguments.Get(0).(*aws_amplify.GetBranchOutput), arguments.Error(1)
}
