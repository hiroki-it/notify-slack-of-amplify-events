package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	aws_amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
)

/**
 * コンストラクタ
 * AmplifyAPIを作成します．
 */
func NewAmplifyAPI() (*AmplifyAPI, error) {

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))

	if err != nil {
		return nil, err
	}

	return &AmplifyAPI{
		Client: aws_amplify.NewFromConfig(config),
	}, nil
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func GetBranchFromAmplify(api *AmplifyAPI, event eventbridge.Event) (*aws_amplify.GetBranchOutput, error) {

	input := aws_amplify.GetBranchInput{
		AppId:      aws.String(event.Detail.AppId),
		BranchName: aws.String(event.Detail.BranchName),
	}

	// ブランチ情報を構造体として取得します．
	response, err := api.Client.GetBranch(context.TODO(), &input)

	if err != nil {
		return nil, err
	}

	return response, err
}
