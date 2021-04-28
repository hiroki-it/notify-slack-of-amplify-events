package amplify

import (
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
)

/**
 * EventBrdigeのイベントを構成します．
 */
type AmplifyBranch struct {
	DisplayName string
}

type AmplifyClient struct {
	Api amplifyiface.AmplifyAPI
}
