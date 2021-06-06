package amplify

import (
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"

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
func (cl *AmplifyClient) GetAmplifyAPI() amplifyiface.AmplifyAPI {
	return cl.api
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func (cl *AmplifyClient) GetBranchFromAmplify(getBranchInput *aws_amplify.GetBranchInput) (*aws_amplify.GetBranchOutput, error) {

	// ブランチ情報を構造体として取得します．
	getBranchOutput, err := cl.api.GetBranch(getBranchInput)

	if err != nil {
		return nil, err
	}

	return getBranchOutput, nil
}
