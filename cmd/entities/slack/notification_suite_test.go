package slack

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

/**
 * ユニットテストのテストスイートを構成します．
 */
type NotificationSuite struct {
	suite.Suite
}

/**
 * ユニットテストの事前処理を実行します．
 */
func (suite *NotificationSuite) BeforeTest(suiteName string, testName string) {

}

/**
 * ユニットテストのテストスイートを実行します．
 */
func TestNotificationSuite(t *testing.T) {
	suite.Run(t, &NotificationSuite{})
}
