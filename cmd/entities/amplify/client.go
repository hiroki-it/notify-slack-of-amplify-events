package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"

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
 * ゲッター
 * AmplifyAPIを返却します．
 */
func (client *AmplifyClient) GetAmplifyAPI() amplifyiface.AmplifyAPI {
	return client.api
}

/**
 * GetBranchInputを作成します．
 */
func (client *AmplifyClient) CreateGetBranchInput(eventDetail *eventbridge.EventDetail) *aws_amplify.GetBranchInput {

	return &aws_amplify.GetBranchInput{
		AppId:      aws.String(eventDetail.AppId),
		BranchName: aws.String(eventDetail.BranchName),
	}
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func (client *AmplifyClient) GetBranchFromAmplify(getBranchInput *aws_amplify.GetBranchInput) (*aws_amplify.GetBranchOutput, error) {

	// ブランチ情報を構造体として取得します．
	getBranchOutput, err := client.api.GetBranch(getBranchInput)

	if err != nil {
		return nil, err
	}

	return getBranchOutput, nil
}
