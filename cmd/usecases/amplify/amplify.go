package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
)

/**
 * コンストラクタ
 * AmplifyClientを作成します．
 */
func NewAmplifyClient(api amplifyiface.AmplifyAPI) (*AmplifyClient, error) {

	return &AmplifyClient{
		Api: api,
	}, nil
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func GetBranchFromAmplify(client *AmplifyClient, event eventbridge.Event) (*aws_amplify.GetBranchOutput, error) {

	input := aws_amplify.GetBranchInput{
		AppId:      aws.String(event.Detail.AppId),
		BranchName: aws.String(event.Detail.BranchName),
	}

	// ブランチ情報を構造体として取得します．
	response, err := client.Api.GetBranch(&input)

	if err != nil {
		return nil, err
	}

	return response, err
}
