package amplify

import (
	"context"

	aws_amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
	"github.com/stretchr/testify/mock"
)

/**
 * Mockを構成します．
 */
type MockedClient struct {
	mockedGetBranch MockedGetBranch
}

/**
 * Mockのメソッドを構成します．
 */
type MockedGetBranch func(ctx context.Context, params *aws_amplify.GetBranchInput, optFns ...func(*aws_amplify.Options)) (*aws_amplify.GetBranchOutput, error)

/**
 * MockをラッピングしたAPIを構成します．
 */
type MockedAmplifyAPI struct {
	MockedClient mock.Mock
}

/**
 * コンストラクタ
 * APIのMockを作成します．
 */
func NewMockedAmplifyAPI() (*MockedAmplifyAPI, error) {
	return new(MockedAmplifyAPI), nil
}

/**
 * AmplifyのGetBranch関数のモックを返却します．
 */
func (mockedClient MockedClient) GetBranch(ctx context.Context, params *aws_amplify.GetBranchInput, optFns ...func(*aws_amplify.Options)) (*aws_amplify.GetBranchOutput, error) {
	return mockedClient.mockedGetBranch(ctx, params, optFns...)
}

/**
 * GetBranchFromAmplifyをモック化します．
 */
func MockedGetBranchFromAmplify(mockedAPI *MockedAmplifyAPI, event eventbridge.Event) (*aws_amplify.GetBranchOutput, error) {
	arguments := mockedAPI.MockedClient.Called(event)
	return arguments.Get(0).(*aws_amplify.GetBranchOutput), arguments.Error(1)
}
