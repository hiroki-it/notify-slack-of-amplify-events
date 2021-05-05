package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
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
		return nil, logger.Error(err)
	}

	return getBranchOutput, err
}
