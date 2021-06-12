package integration

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file"
	"github.com/stretchr/testify/assert"
)

/**
 * 統合テストを実行します．
 */
func TestIntegration(t *testing.T) {

	t.Helper()

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// 期待値
		expected int
		// テストデータ
		detail []byte
	}{
		{
			name:     "TestIntegration_Succeed_ReturnOk",
			expected: http.StatusOK,
			detail:   file.ReadDataFile("./testdata/succeed.json"),
		},
		{
			name:     "TestIntegration_Failed_ReturnOk",
			expected: http.StatusOK,
			detail:   file.ReadDataFile("./testdata/failed.json"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			// リクエストを作成する．
			req, err := http.NewRequest(
				"POST",
				fmt.Sprint("http://", os.Getenv("LAMBDA_HOST"), ":9000/2015-03-31/functions/function/invocations"),
				bytes.NewBuffer(tt.detail),
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

			assert.Exactly(t, tt.expected, res.StatusCode)
		})
	}
}
