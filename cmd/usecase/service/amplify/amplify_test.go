package amplify

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/entity/detail"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	aws_amplify "github.com/aws/aws-sdk-go/service/amplify"
	m_amplify "github.com/hiroki-it/notify-slack-of-amplify-events/mock/amplify"
)

/**
 * ユニットテストの前処理の結果と，後処理の関数を返却します．
 */
func setup() (*m_amplify.MockedAmplifyAPI, func()) {

	// AmplifyAPIのスタブを作成します．
	mockedAPI := &m_amplify.MockedAmplifyAPI{}

	return mockedAPI, func() {
	}
}

/**
 * GetBranchFromAmplifyメソッドが成功することをテストします．
 */
func TestGetBranchFromAmplify(t *testing.T) {

	t.Helper()

	mockedAPI, teardown := setup()
	defer teardown()

	// テストケース
	cases := []struct {
		// テストケース名
		name string
		// 期待値
		expected string
		// テストデータ
		detail *detail.Detail
	}{
		{
			name:     "TestGetBranchFromAmplify_detail_ReturnDisplayName",
			expected: "feature-test",
			detail: &detail.Detail{
				AppId:      "1",
				BranchName: "test",
			},
		},
	}

	// 反復処理で全てのテストケースを検証します．
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			// スタブのメソッドに処理の内容を定義します．
			mockedAPI.On("GetBranch", mock.Anything).Return(
				&aws_amplify.GetBranchOutput{
					Branch: &aws_amplify.Branch{DisplayName: aws.String("feature-test")},
				},
				nil,
			)

			amplifyClient := NewAmplifyClient(mockedAPI)

			// 検証対象の関数を実行します．スタブを含む一連の処理が実行されます．
			getBranchOutput, err := amplifyClient.GetBranchFromAmplify(tt.detail)

			if err != nil {
				t.Fatal(err.Error())
			}

			// 関数内部でスタブのメソッドがコールされているかを検証します．
			mockedAPI.AssertExpectations(t)

			// 最終的な返却値が正しいかを検証します．
			assert.Exactly(t, aws.String(tt.expected), getBranchOutput.Branch.DisplayName)
		})
	}
}
