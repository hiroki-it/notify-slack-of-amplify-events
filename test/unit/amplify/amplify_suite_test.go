package unit

import (
	"encoding/json"
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	"github.com/stretchr/testify/suite"
)

/**
 * Amplifyの単体テストのテストスイートを構成します．
 */
type SuiteAmplify struct {
	suite.Suite
	eventDetail *eventbridge.EventDetail
}

/**
 * Amplifyの単体テストの事前処理を実行します．
 */
func (suite *SuiteAmplify) BeforeTest(suiteName string, testName string) {

	detail, err := file.ReadTestDataFile("../../testdata/request/event.json")

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	eventDetail := eventbridge.NewEventDetail()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err = json.Unmarshal(detail, &eventDetail)

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	suite.eventDetail = eventDetail
}

func TestSuiteAmplify(t *testing.T) {
	suite.Run(t, &SuiteAmplify{})
}
