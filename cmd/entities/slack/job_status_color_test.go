package slack

import (
	"testing"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge"
	"github.com/stretchr/testify/assert"
)

/**
 * 特定のステータスの場合に，PrintStatusWordメソッドが対応する文言を返却することテストします．
 */
func TestPrintStatusWord(t *testing.T) {

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// 期待値
		expected string
		// テストデータ
		eventDetail *eventbridge.EventDetail
	}{
		{
			name:     "TestPrintStatusWord_Succeed_Return成功",
			expected: "成功",
			eventDetail: &eventbridge.EventDetail{
				JobStatus: "SUCCEED",
			},
		},
		{
			name:     "TestPrintStatusWord_Failed_Return失敗",
			expected: "失敗",
			eventDetail: &eventbridge.EventDetail{
				JobStatus: "FAILED",
			},
		},
	}

	// 反復処理で全てのテストケースを検証します．
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			jobStatusColor := NewJobStatusColor(tt.eventDetail.JobStatus)

			assert.Exactly(t, tt.expected, jobStatusColor.PrintStatusWord())
		})
	}
}

/**
 * 特定のステータスの場合に，PrintStatusColorメソッドが対応するカラーコードを返却することテストします．
 */
func TestPrintStatusColor(t *testing.T) {

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// 期待値
		expected string
		// テストデータ
		eventDetail *eventbridge.EventDetail
	}{
		{
			name:     "TestPrintStatusColor_Succeed_Return#00FF00",
			expected: "#00FF00",
			eventDetail: &eventbridge.EventDetail{
				JobStatus: "SUCCEED",
			},
		},
		{
			name:     "TestPrintStatusColor_Failed_Return#ff0000",
			expected: "#ff0000",
			eventDetail: &eventbridge.EventDetail{
				JobStatus: "FAILED",
			},
		},
	}

	// 反復処理で全てのテストケースを検証します．
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			jobStatusColor := NewJobStatusColor(tt.eventDetail.JobStatus)

			assert.Exactly(t, tt.expected, jobStatusColor.PrintStatusColor())
		})
	}
}
