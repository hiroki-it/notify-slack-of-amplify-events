package integration

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * 統合テストを実行します．
 */
func TestIntegration(t *testing.T) {

	t.Helper()

	detail, teardown := setup(t)
	defer teardown()

	// リクエストを作成する．
	req, err := http.NewRequest(
		"POST",
		fmt.Sprint("http://", os.Getenv("LAMBDA_HOST"), ":9000/2015-03-31/functions/function/invocations"),
		bytes.NewBuffer(detail),
	)

	if err != nil {
		t.Fatal(err.Error())
	}

	// クライアントを作成する．
	client := &http.Client{}

	// lambdaにリクエストを送信する．
	res, err := client.Do(req)

	if err != nil {
		t.Fatal(err.Error())
	}

	defer res.Body.Close()

	assert.Exactly(t, http.StatusOK, res.StatusCode)
}
