package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/detail"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * AmplifyClientインターフェースを構成します．
 */
type AmplifyClientInterface interface {
	CreateGetBranchInput(*detail.Detail) *aws_amplify.GetBranchInput
	GetBranchFromAmplify(*detail.Detail) (*aws_amplify.GetBranchOutput, error)
}

/**
 * AmplifyClientインターフェースの実装を構成します．
 */
type AmplifyClient struct {
	AmplifyClientInterface
	api amplifyiface.AmplifyAPI
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
 * Amplifyからブランチ情報を取得します．
 */
func (cl *AmplifyClient) GetBranchFromAmplify(detail *detail.Detail) (*aws_amplify.GetBranchOutput, error) {

	getBranchInput := &aws_amplify.GetBranchInput{
		AppId:      aws.String(detail.AppId.Value()),
		BranchName: aws.String(detail.BranchName.Value()),
	}

	// ブランチ情報を構造体として取得します．
	getBranchOutput, err := cl.api.GetBranch(getBranchInput)

	if err != nil {
		return nil, err
	}

	return getBranchOutput, nil
}
