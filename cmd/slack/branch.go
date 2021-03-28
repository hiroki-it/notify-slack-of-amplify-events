package slack

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * AmplifyApp構造体
 */
type AmplifyApp struct {
	Branch struct {
		ActiveJobID              string   `json:"activeJobId"`
		AssociatedResources      []string `json:"associatedResources"`
		BackendEnvironmentArn    string   `json:"backendEnvironmentArn"`
		BasicAuthCredentials     string   `json:"basicAuthCredentials"`
		BranchArn                string   `json:"branchArn"`
		BranchName               string   `json:"branchName"`
		BuildSpec                string   `json:"buildSpec"`
		CreateTime               int      `json:"createTime"`
		CustomDomains            []string `json:"customDomains"`
		Description              string   `json:"description"`
		DestinationBranch        string   `json:"destinationBranch"`
		DisplayName              string   `json:"displayName"`
		EnableAutoBuild          bool     `json:"enableAutoBuild"`
		EnableBasicAuth          bool     `json:"enableBasicAuth"`
		EnableNotification       bool     `json:"enableNotification"`
		EnablePerformanceMode    bool     `json:"enablePerformanceMode"`
		EnablePullRequestPreview bool     `json:"enablePullRequestPreview"`
		EnvironmentVariables     struct {
			String string `json:"string"`
		} `json:"environmentVariables"`
		Framework                  string `json:"framework"`
		PullRequestEnvironmentName string `json:"pullRequestEnvironmentName"`
		SourceBranch               string `json:"sourceBranch"`
		Stage                      string `json:"stage"`
		Tags                       struct {
			String string `json:"string"`
		} `json:"tags"`
		ThumbnailURL string `json:"thumbnailUrl"`
		Total1OfJobs string `json:"total1OfJobs"`
		TTL          string `json:"ttl"`
		UpdateTime   int    `json:"updateTime"`
	} `json:"branch"`
}

/**
 * Amplifyからアプリケーション情報を取得します．
 */
func getAmplifyBranch(event Event) (AmplifyApp, error) {
	sess := session.Must(session.NewSession(aws.NewConfig().WithRegion("ap-northeast-1")))

	svc := amplify.New(sess)

	option := &amplify.GetBranchInput{
		AppId:      aws.String(event.Detail.AppId),
		BranchName: aws.String(event.Detail.BranchName),
	}

	response, err := svc.GetBranch(option)

	branch := response.String()

	var amplifyApp AmplifyApp

	err = json.Unmarshal([]byte(branch), &amplifyApp)

	return amplifyApp, err
}
