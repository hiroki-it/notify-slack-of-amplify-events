package amplify

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/entities"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

type AmplifyClientInterface interface {
	GetBranchFromAmplify(*entities.Detail) (*aws_amplify.GetBranchOutput, error)
}

type AmplifyClient struct {
	api amplifyiface.AmplifyAPI
}

var _ AmplifyClientInterface = &AmplifyClient{}

// NewAmplifyClient コンストラクタ
func NewAmplifyClient(config *aws.Config) (*AmplifyClient, error) {

	s, err := session.NewSession(config)

	if err != nil {
		return nil, err
	}

	return &AmplifyClient{
		api: aws_amplify.New(s),
	}, nil
}

// GetBranchFromAmplify Amplifyからブランチ情報を取得します．
func (cl *AmplifyClient) GetBranchFromAmplify(detail *entities.Detail) (*aws_amplify.GetBranchOutput, error) {

	gbi := &aws_amplify.GetBranchInput{
		AppId:      aws.String(detail.AppId().Id()),
		BranchName: aws.String(detail.BranchName().Name()),
	}

	// ブランチ情報を構造体として取得します．
	gbo, err := cl.api.GetBranch(gbi)

	if err != nil {
		return nil, err
	}

	return gbo, nil
}
