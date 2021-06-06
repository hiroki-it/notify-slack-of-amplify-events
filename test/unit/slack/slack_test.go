package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack"
	"github.com/stretchr/testify/assert"
)

/**
 * PostMessageメソッドをテストします．
 */
func (suite *SuiteSlack) TestPostMessage() {

	suite.T().Helper()

	// 外部サーバのモックを構築します．
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{}")
	}))

	defer ts.Close()

	slackMessage := slack.NewSlackMessage(
		suite.eventDetail,
		suite.getBranchOutput.Branch,
		suite.jobStatusColor,
	)

	slackClient := slack.NewSlackClient(
		&http.Client{},
		ts.URL, // モックサーバのURLに差し替えます．
	)

	slackNotification := slack.NewSlackNotification(
		slackClient,
		slackMessage,
	)

	err := slackNotification.PostMessage()

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	assert.Nil(suite.T(), err)
}
