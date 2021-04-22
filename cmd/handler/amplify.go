package handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
)

type Event struct {
	Version    string   `json:"version"`
	Id         string   `json:"id"`
	DetailType string   `json:"detail-type"`
	Source     string   `json:"source"`
	Account    string   `json:"account"`
	Time       string   `json:"time"`
	Region     string   `json:"region"`
	Resources  []string `json:"resources"`
	Detail     struct {
		AppId      string `json:"appId"`
		BranchName string `json:"branchName"`
		JobId      string `json:"jobId"`
		JobStatus  string `json:"jobStatus"`
	} `json:"detail"`
}

type AmplifyBranch struct {
	DisplayName string
}

type AmplifyClientInterface interface {
	getBranchFromAmplify(event Event) (*amplify.GetBranchOutput, error)
}

type AmplifyClientImpl struct {
	AmplifyClientInterface
	Svc *amplify.Client
}

/**
 * コンストラクタ
 */
func NewAmplifyClient() (*AmplifyClientImpl, error) {

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))

	if err != nil {
		return nil, err
	}

	return &AmplifyClientImpl{
		Svc: amplify.NewFromConfig(config),
	}, nil
}

/**
 * Amplifyからブランチ情報を取得します．
 */
func (client AmplifyClientImpl) getBranchFromAmplify(event Event) (*amplify.GetBranchOutput, error) {

	input := amplify.GetBranchInput{
		AppId:      aws.String(event.Detail.AppId),
		BranchName: aws.String(event.Detail.BranchName),
	}

	// ブランチ情報を構造体として取得します．
	response, err := client.Svc.GetBranch(context.TODO(), &input)

	if err != nil {
		return nil, err
	}

	return response, err
}
