package amplify

import (
	aws_amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
)

/**
 * EventBrdigeのイベントを構成します．
 */
type AmplifyBranch struct {
	DisplayName string
}

/**
 * APIのインターフェースを構成します．
 */
type AmplifyAPIInterface interface {
	getBranchFromAmplify(api *AmplifyAPIImpl, event eventbridge.Event) (*aws_amplify.GetBranchOutput, error)
}

/**
 * Clientをラッピングした構造体を構成します．
 */
type AmplifyAPIImpl struct {
	AmplifyAPIInterface
	Client *aws_amplify.Client
}
