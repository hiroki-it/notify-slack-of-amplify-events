package integration

import (
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	"github.com/stretchr/testify/suite"
)

/**
 * Lambdaの統合テストのテストスイートを構成します．
 */
type SuiteLambda struct {
	suite.Suite
	detail []byte
}

/**
 * Lambdaの統合テストの事前処理を実行します．
 */
func (suite *SuiteLambda) BeforeTest(suiteName string, testName string) {

	detail, err := file.ReadTestDataFile("./testdata/event.json")

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	suite.detail = detail
}

func TestSuiteLambda(t *testing.T) {
	suite.Run(t, &SuiteLambda{})
}
