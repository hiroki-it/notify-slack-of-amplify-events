package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
)

/**
 * コンストラクタ
 * AmplifyAPIを作成します．
 */
func NewAmplifyClient() (*AmplifyClient, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-1")})

	if err != nil {
		return nil, err
	}

	return &AmplifyClient{
		api: aws_amplify.New(sess),
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
	response, err := client.api.GetBranch(&input)

	if err != nil {
		return nil, err
	}

	return response, err
}
