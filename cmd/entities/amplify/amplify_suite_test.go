package amplify

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

/**
 * ユニットテストのテストスイートを構成します．
 */
type SuiteAmplify struct {
	suite.Suite
}

/**
 * ユニットテストの事前処理を実行します．
 */
func (suite *SuiteAmplify) BeforeTest(suiteName string, testName string) {
}

/**
 * ユニットテストのテストスイートを実行します．
 */
func TestSuiteAmplify(t *testing.T) {
	suite.Run(t, &SuiteAmplify{})
}
