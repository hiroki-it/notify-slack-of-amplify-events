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
type MockClient struct {
	mockGetBranch MockGetBranch
}

/**
 * Mockのメソッドを構成します．
 */
type MockGetBranch func(ctx context.Context, params *aws_amplify.GetBranchInput, optFns ...func(*aws_amplify.Options)) (*aws_amplify.GetBranchOutput, error)

/**
 * MockをラッピングしたAPIを構成します．
 */
type MockAmplifyAPIImpl struct {
	AmplifyAPIInterface
	MockClient mock.Mock
}

/**
 * コンストラクタ
 * APIのMockを作成します．
 */
func NewMockAmplifyAPI() (*MockAmplifyAPIImpl, error) {
	return new(MockAmplifyAPIImpl), nil
}

/**
 * AmplifyのGetBranch関数のモックを返却します．
 */
func (mockClient MockClient) GetBranch(ctx context.Context, params *aws_amplify.GetBranchInput, optFns ...func(*aws_amplify.Options)) (*aws_amplify.GetBranchOutput, error) {
	return mockClient.mockGetBranch(ctx, params, optFns...)
}

/**
 * AmplifyのGetBranch関数のモックを返却します．
 */
func MockGetBranchFromAmplify(mockAPI *MockAmplifyAPIImpl, event eventbridge.Event) (*aws_amplify.GetBranchOutput, error) {
	arguments := mockAPI.MockClient.Called(event)
	return arguments.Get(0).(*aws_amplify.GetBranchOutput), arguments.Error(1)
}
