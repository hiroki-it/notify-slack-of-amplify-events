package notification

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/entities"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/ids"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/values"
	"github.com/stretchr/testify/assert"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
)

// setup ユニットテストの前処理の結果と，後処理の関数を返却します．
func setup() (*httptest.Server, func()) {

	// 外部サーバのモックを構築します．
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{}")
	}))

	return ts, func() {
		ts.Close()
	}
}

// TestPostMessage PostMessageメソッドが成功することをテストします．
func TestPostMessage(t *testing.T) {

	t.Helper()

	ts, teardown := setup()
	defer teardown()

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// 期待値
		expected string
		// テストデータ
		detail          *entities.Detail
		getBranchOutput *aws_amplify.GetBranchOutput
	}{
		{
			name: "TestPostMessage_JobStatusTypeSucceed_ReturnNil",
			detail: entities.NewDetail(
				ids.NewAppId("1"),
				values.NewBranchName("test"),
				ids.NewJobId("1"),
				values.NewJobStatusType("SUCCEED"),
			),
			getBranchOutput: &aws_amplify.GetBranchOutput{
				Branch: &aws_amplify.Branch{DisplayName: aws.String("feature-test")},
			},
		},
		{
			name: "TestPostMessage_JobStatusTypeFailed_ReturnNil",
			detail: entities.NewDetail(
				ids.NewAppId("1"),
				values.NewBranchName("test"),
				ids.NewJobId("1"),
				values.NewJobStatusType("FAILED"),
			),
			getBranchOutput: &aws_amplify.GetBranchOutput{
				Branch: &aws_amplify.Branch{DisplayName: aws.String("feature-test")},
			},
		},
	}

	// 反復処理で全てのテストケースを検証します．
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			m := NewMessage(
				tt.detail,
				tt.getBranchOutput.Branch,
			)

			sc := NewSlackClient(
				&http.Client{},
				ts.URL, // モックサーバのURLに差し替えます．
			)

			sm := m.BuildSlackMessage()

			sn := NewSlackNotification(
				sc,
				sm,
			)

			err := sn.PostMessage()

			if err == nil {
				// 正常系テストのアサーション
				assert.Nil(t, err)
			}

			if err != nil {
				// 異常系テストのアサーション
				assert.Exactly(t, tt.expected, err)
			}
		})
	}
}
