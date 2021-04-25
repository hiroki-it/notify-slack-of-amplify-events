package amplify

import (
	aws_amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
)

/**
 * EventBrdigeのイベントを構成します．
 */
type AmplifyBranch struct {
	DisplayName string
}

/**
 * Clientをラッピングした構造体を構成します．
 */
type AmplifyAPI struct {
	Client *aws_amplify.Client
}
