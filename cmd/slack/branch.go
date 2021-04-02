package slack

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

/**
 * AmplifyApp構造体
 */
type Branch struct {
	DisplayName string
}

/**
 * Amplifyのクライアントを作成します．
 */
func NewAmplifyClient() (*amplify.Client, error) {

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))

	if err != nil {
		return nil, err
	}

	return amplify.NewFromConfig(config), nil
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func getBranchFromAmplify(event Event) (Branch, error) {

	var branch Branch

	client, err := NewAmplifyClient()

	if err != nil {
		return branch, err
	}

	input := amplify.GetBranchInput{
		AppId:      aws.String(event.Detail.AppId),
		BranchName: aws.String(event.Detail.BranchName),
	}

	// ブランチ情報を構造体として取得します．
	response, err := client.GetBranch(context.TODO(), &input)

	if err != nil {
		return branch, err
	}

	// あらかじめ定義した構造体に出力します．．
	branch = Branch{DisplayName: aws.ToString(response.Branch.DisplayName)}

	return branch, err
}
