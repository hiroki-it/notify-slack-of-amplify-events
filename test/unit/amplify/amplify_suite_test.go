package unit

import (
	"encoding/json"
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
	"github.com/stretchr/testify/suite"
)

/**
 * Amplifyのテストスイートを構成します．
 */
type SuiteAmplify struct {
	suite.Suite
	eventDetail *eventbridge.EventDetail
}

/**
 * Amplifyのテストの前にテストデータを読み込みます．
 */
func (suite SuiteAmplify) BeforeTest() {

	log := logger.NewLogger()

	detail, err := file.ReadTestDataFile("../testdata/request/event.json")

	if err != nil {
		log.Error(err.Error())
	}

	eventDetail := eventbridge.NewEventDetail()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err = json.Unmarshal(detail, &eventDetail)

	if err != nil {
		log.Error(err.Error())
	}

	suite.eventDetail = eventDetail
}

func TestSuiteAmplify(t *testing.T) {
	suite.Run(t, &SuiteAmplify{})
}
