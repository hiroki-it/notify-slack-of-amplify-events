package amplify

import (
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/stretchr/testify/suite"
)

/**
 * ユニットテストのテストスイートを構成します．
 */
type SuiteAmplify struct {
	suite.Suite
	eventDetail *eventbridge.EventDetail
}

/**
 * ユニットテストの事前処理を実行します．
 */
func (suite *SuiteAmplify) BeforeTest(suiteName string, testName string) {

	eventDetail := eventbridge.NewEventDetail()

	eventDetail.AppId = "test"
	eventDetail.BranchName = "test"

	// テストデータを準備します．
	suite.eventDetail = eventDetail
}

/**
 * ユニットテストのテストスイートを実行します．
 */
func TestSuiteAmplify(t *testing.T) {
	suite.Run(t, &SuiteAmplify{})
}
