package amplify

import (
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
)

/**
 * Amplifyのクライアントを構成します．
 */
type AmplifyClient struct {
	Api amplifyiface.AmplifyAPI
}
