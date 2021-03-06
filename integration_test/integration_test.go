package integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/entities"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/values"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/file/repositories"
	"github.com/stretchr/testify/assert"
)

// setup ユニットテストの前処理の結果と，後処理の関数を返却します．
func setup() (*http.Client, func()) {

	// クライアントを作成します．
	client := &http.Client{}

	return client, func() {
	}
}

// TestIntegration 統合テストを実行します．
func TestIntegration(t *testing.T) {

	t.Helper()
	client, teardown := setup()
	defer teardown()

	fileRepositoty := repositories.NewFileRepository()

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// 期待値
		expected string
		// テストデータ
		detail []byte
	}{
		{
			name:     "TestIntegration_SucceedStatus_ReturnOk",
			expected: string(fileRepositoty.GetFile(entities.NewFile(values.NewPath("./test_data/response/normal_failed_status.golden")))),
			detail:   fileRepositoty.GetFile(entities.NewFile(values.NewPath("./test_data/request/normal_succeed_status.json"))),
		},
		{
			name:     "TestIntegration_FailedStatus_ReturnOk",
			expected: string(fileRepositoty.GetFile(entities.NewFile(values.NewPath("./test_data/response/normal_failed_status.golden")))),
			detail:   fileRepositoty.GetFile(entities.NewFile(values.NewPath("./test_data/request/normal_failed_status.json"))),
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

			// スラッシュを削除します．
			actual, err := strconv.Unquote(string(b))

			if err != nil {
				t.Fatal(err.Error())
			}

			t.Log(actual)

			assert.JSONEq(t, tt.expected, actual)
		})
	}
}
