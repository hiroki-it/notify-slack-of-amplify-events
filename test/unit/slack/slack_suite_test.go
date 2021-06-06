package unit

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/stretchr/testify/suite"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

/**
 * ユニットテストのテストスイートを構成します．
 */
type SuiteSlack struct {
	suite.Suite
	eventDetail     *eventbridge.EventDetail
	getBranchOutput *aws_amplify.GetBranchOutput
	jobStatusColor  *slack.JobStatusColor
}

/**
 * ユニットテストの事前処理を実行します．
 */
func (suite *SuiteSlack) BeforeTest(suiteName string, testName string) {

	eventDetail := eventbridge.NewEventDetail()

	eventDetail.AppId = "1"
	eventDetail.BranchName = "test"
	eventDetail.JobId = "1"

	// テストデータを準備します．
	suite.eventDetail = eventDetail
	suite.getBranchOutput = &aws_amplify.GetBranchOutput{
		Branch: &aws_amplify.Branch{DisplayName: aws.String("feature-test")},
	}
	suite.jobStatusColor = slack.NewJobStatusColor("SUCCESS")
}

/**
 * ユニットテストのテストスイートを実行します．
 */
func TestSuiteSlack(t *testing.T) {
	suite.Run(t, &SuiteSlack{})
}
