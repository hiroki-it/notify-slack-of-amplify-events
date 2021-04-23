package handler

import (
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/stretchr/testify/mock"
)

type MockAmplifyAPIImpl struct {
	AmplifyAPIInterface
	mock.Mock
}

func NewMockAmplifyAPI() (*MockAmplifyAPIImpl, error) {
	return new(MockAmplifyAPIImpl), nil
}

func getMockBranchFromAmplify(mock *MockAmplifyAPIImpl, event Event) (*amplify.GetBranchOutput, error) {
	arguments := mock.Called(event)
	return arguments.Get(0).(*amplify.GetBranchOutput), arguments.Error(1)
}
