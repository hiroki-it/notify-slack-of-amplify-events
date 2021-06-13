package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/eventbridge"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * コンストラクタ
 * AmplifyClientを作成します．
 */
func NewAmplifyClient(amplifyApi amplifyiface.AmplifyAPI) *AmplifyClient {

	return &AmplifyClient{
		api: amplifyApi,
	}
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func (cl *AmplifyClient) GetBranchFromAmplify(eventDetail *eventbridge.EventDetail) (*aws_amplify.GetBranchOutput, error) {

	getBranchInput := &aws_amplify.GetBranchInput{
		AppId:      aws.String(eventDetail.AppId),
		BranchName: aws.String(eventDetail.BranchName),
	}

	// ブランチ情報を構造体として取得します．
	getBranchOutput, err := cl.api.GetBranch(getBranchInput)

	if err != nil {
		return nil, err
	}

	return getBranchOutput, nil
}
