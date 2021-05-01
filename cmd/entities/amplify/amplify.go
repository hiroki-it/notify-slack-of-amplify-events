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
	getBranchOutput, err := client.Api.GetBranch(getBranchInput)

	if err != nil {
		return nil, err
	}

	return getBranchOutput, err
}
