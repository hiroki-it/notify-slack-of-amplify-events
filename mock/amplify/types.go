package amplify

import (
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
