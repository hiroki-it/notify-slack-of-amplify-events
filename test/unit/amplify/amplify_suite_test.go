package unit

import (
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/stretchr/testify/suite"
)

/**
 * Amplifyの単体テストのテストスイートを構成します．
 */
type SuiteAmplify struct {
	suite.Suite
	getBranchFromAmplifyData *eventbridge.EventDetail
}

/**
 * Amplifyの単体テストの事前処理を実行します．
 */
func (suite *SuiteAmplify) BeforeTest(suiteName string, testName string) {

	eventDetail := eventbridge.NewEventDetail()

	eventDetail.AppId = "test"
	eventDetail.BranchName = "test"

	// テストデータを準備します．
	suite.getBranchFromAmplifyData = eventDetail
}

func TestSuiteAmplify(t *testing.T) {
	suite.Run(t, &SuiteAmplify{})
}
