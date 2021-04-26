package mock

import (
	"context"

	aws_amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/stretchr/testify/mock"
)

/**
 * Mockを構成します．
 */
type MockedClient struct {
	mock.Mock
}

type AmplifyAPI struct {
	Client MockedClient
}

func NewMockedAmplifyAPI() (*AmplifyAPI, error) {
	return new(AmplifyAPI), nil
}

func (mockedClient *MockedClient) GetBranch(ctx context.Context, params *aws_amplify.GetBranchInput, optFns ...func(*aws_amplify.Options)) (*aws_amplify.GetBranchOutput, error) {
	arguments := mockedClient.Called(ctx, params, optFns)
	return arguments.Get(0).(*aws_amplify.GetBranchOutput), arguments.Error(1)
}
