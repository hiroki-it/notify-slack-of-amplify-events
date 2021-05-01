package amplify

import (
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
)

type AmplifyClientInterface interface {
	GetBranchFromAmplify(eventDetail *eventbridge.EventDetail) (*aws_amplify.GetBranchOutput, error)
}

/**
 * Amplifyのクライアントを構成します．
 */
type AmplifyClient struct {
	AmplifyClientInterface
	Api amplifyiface.AmplifyAPI
}
