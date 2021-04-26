package amplify

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_amplify "github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge"
	"github.com/hiroki-it/notify-slack-of-amplify-events/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetBranchFromAmplify(t *testing.T) {

	input := aws_amplify.GetBranchInput{
		AppId:      aws.String("123456789"),
		BranchName: aws.String("feature/test"),
	}

	api, _ := mock.NewMockedAmplifyAPI()

	// スタブに引数として渡される値と，その時の返却値を定義する．
	api.Client.On("GetBranch", context.TODO(), &input).Return(Branch{DisplayName: aws.String("feature-test")}, nil)

	var event eventbridge.Event

	// 検証対象の関数を実行する．スタブを含む一連の処理が実行される．
	response, _ := GetBranchFromAmplify(api, event)

	//関数内部でスタブがコールされているかを検証する．
	api.MockedClient.AssertExpectations(t)

	// 最終的な返却値が正しいかを検証する．
	assert.Exactly(t, aws.String("feature-test"), response.Branch.DisplayName)
}
