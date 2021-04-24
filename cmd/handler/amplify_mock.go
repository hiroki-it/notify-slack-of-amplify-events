package handler

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
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
type MockGetBranch func(ctx context.Context, params *amplify.GetBranchInput, optFns ...func(*amplify.Options)) (*amplify.GetBranchOutput, error)

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
func (mockClient MockClient) GetBranch(ctx context.Context, params *amplify.GetBranchInput, optFns ...func(*amplify.Options)) (*amplify.GetBranchOutput, error) {
	return mockClient.mockGetBranch(ctx, params, optFns...)
}

/**
 * AmplifyのGetBranch関数のモックを返却します．
 */
func mockGetBranchFromAmplify(mockAPI *MockAmplifyAPIImpl, event Event) (*amplify.GetBranchOutput, error) {
	arguments := mockAPI.MockClient.Called(event)
	return arguments.Get(0).(*amplify.GetBranchOutput), arguments.Error(1)
}
