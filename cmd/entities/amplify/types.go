package amplify

import (
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * AmplifyClientインターフェースを構成します．
 */
type AmplifyClientInterface interface {
	CreateGetBranchInput(*eventbridge.EventDetail) *aws_amplify.GetBranchInput
	GetBranchFromAmplify(*eventbridge.EventDetail) (*aws_amplify.GetBranchOutput, error)
}

/**
 * AmplifyClientインターフェースの実装を構成します．
 */
type AmplifyClient struct {
	AmplifyClientInterface
	api amplifyiface.AmplifyAPI
}
