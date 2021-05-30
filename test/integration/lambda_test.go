package integration

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/stretchr/testify/assert"
)

/**
 * Lambdaをテストします．．
 */
func (suite *SuiteLambda) TestLambda() {

	suite.T().Helper()

	// リクエストを作成する．
	request, err := http.NewRequest(
		"POST",
		fmt.Sprint("http://", os.Getenv("LAMBDA_HOST"), ":9000/2015-03-31/functions/function/invocations"),
		bytes.NewBuffer(suite.detail),
	)

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	// クライアントを作成する．
	client := &http.Client{}

	// lambdaにリクエストを送信する．
	response, err := client.Do(request)

	if err != nil {
		suite.T().Fatal(err.Error())
	}

	defer response.Body.Close()

	assert.Exactly(suite.T(), http.StatusOK, response.StatusCode)
}
