package integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/file"
	"github.com/stretchr/testify/assert"
)

// TestIntegration 統合テストを実行します．
func TestIntegration(t *testing.T) {

	t.Helper()

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// 期待値
		expected []byte
		// テストデータ
		detail []byte
	}{
		{
			name:     "TestIntegration_Succeed_ReturnOk",
			expected: file.ReadFile("./response/response.json.golden"),
			detail:   file.ReadFile("./request/succeed.json"),
		},
		{
			name:     "TestIntegration_Failed_ReturnOk",
			expected: file.ReadFile("./response/response.json.golden"),
			detail:   file.ReadFile("./request/failed.json"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			// リクエストを作成します．
			req, err := http.NewRequest(
				"POST",
				fmt.Sprint("http://", os.Getenv("LAMBDA_HOST"), ":9000/2015-03-31/functions/function/invocations"),
				bytes.NewBuffer(tt.detail),
			)

			if err != nil {
				t.Fatal(err.Error())
			}

			// クライアントを作成します．
			client := &http.Client{}

			// lambdaにリクエストを送信します．
			res, err := client.Do(req)

			if err != nil {
				t.Fatal(err.Error())
			}

			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)

			if err != nil {
				t.Fatal(err.Error())
			}

			assert.Exactly(t, string(tt.expected), string(b))
		})
	}
}
