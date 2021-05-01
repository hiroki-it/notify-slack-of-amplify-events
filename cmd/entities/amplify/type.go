package amplify

import (
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
)

/**
 * AmplifyClientインターフェースを構成します．
 */
type AmplifyClientInterface interface {
	CreateGetBranchInput(eventDetail *eventbridge.EventDetail) *aws_amplify.GetBranchInput
	GetBranchFromAmplify(eventDetail *eventbridge.EventDetail) (*aws_amplify.GetBranchOutput, error)
}

/**
 * AmplifyClientインターフェースの実装を構成します．
 */
type AmplifyClient struct {
	AmplifyClientInterface
	Api amplifyiface.AmplifyAPI
}
