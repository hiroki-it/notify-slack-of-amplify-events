package handler

import (
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	"github.com/stretchr/testify/suite"
)

/**
 * 統合テストのテストスイートを構成します．
 */
type SuiteIntegration struct {
	suite.Suite
	detail []byte
}

/**
 * 統合テストの事前処理を実行します．
 */
func (suite *SuiteIntegration) BeforeTest(suiteName string, testName string) {

	detail, err := file.ReadTestDataFile("./testdata/event.json")

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	suite.detail = detail
}

func TestSuiteIntegration(t *testing.T) {
	suite.Run(t, &SuiteIntegration{})
}
