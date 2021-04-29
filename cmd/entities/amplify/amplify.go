package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
)

/**
 * コンストラクタ
 * AmplifyAPIを作成します．
 */
func NewAmplifyAPI(region string) (amplifyiface.AmplifyAPI, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		return nil, err
	}

	return aws_amplify.New(sess), err
}

/**
 * コンストラクタ
 * AmplifyClientを作成します．
 */
func NewAmplifyClient(api amplifyiface.AmplifyAPI) *AmplifyClient {

	return &AmplifyClient{
		Api: api,
	}
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func (client *AmplifyClient) GetBranchFromAmplify(event eventbridge.Event) (*aws_amplify.GetBranchOutput, error) {

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
