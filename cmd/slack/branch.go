package slack

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

type Branch struct {
	DisplayName string
}

type AmplifyClient struct {
	svc *amplify.Client
}

/**
 * コンストラクタ
 */
func NewAmplifyClient() (*AmplifyClient, error) {

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))

	if err != nil {
		return nil, err
	}

	return &AmplifyClient{
		svc: amplify.NewFromConfig(config),
	}, nil
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func (client AmplifyClient) getBranchFromAmplify(event Event) (Branch, error) {

	var branch Branch

	input := amplify.GetBranchInput{
		AppId:      aws.String(event.Detail.AppId),
		BranchName: aws.String(event.Detail.BranchName),
	}

	// ブランチ情報を構造体として取得します．
	response, err := client.svc.GetBranch(context.TODO(), &input)

	if err != nil {
		return branch, err
	}

	// あらかじめ定義した構造体に出力します．．
	branch = Branch{DisplayName: aws.ToString(response.Branch.DisplayName)}

	return branch, err
}
