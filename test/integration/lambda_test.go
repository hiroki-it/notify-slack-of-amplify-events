package unit

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
	"github.com/stretchr/testify/assert"
)

/**
 * Lambdaをテストします．．
 */
func TestLambda(t *testing.T) {

	log := logger.NewLogger()

	detail, err := file.ReadTestDataFile("../testdata/request/event.json")

	if err != nil {
		log.Error(err.Error())
	}

	// リクエストを作成する．
	request, _ := http.NewRequest(
		"POST",
		"http://lambda:9000/2015-03-31/functions/function/invocations",
		bytes.NewBuffer(detail),
	)

	// クライアントを作成する．
	client := &http.Client{}

	// lambdaにリクエストを送信する．
	response, err := client.Do(request)

	if err != nil {
		log.Error(err.Error())
	}

	defer response.Body.Close()

	assert.Exactly(t, http.StatusOK, response.StatusCode)
}
