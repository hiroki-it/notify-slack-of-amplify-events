package usecase

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/input"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/service/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/service/notification"
)

type EventPostUseCase struct {
}

// NewEventPostUseCase コンストラクタ
func NewEventPostUseCase() *EventPostUseCase {

	return &EventPostUseCase{}
}

// PostEvent イベントを通知します．
func (uc *EventPostUseCase) PostEvent(input *input.EventPostInput) error {

	ac, err := amplify.NewAmplifyClient(&aws.Config{
		Region: aws.String(os.Getenv("AWS_AMPLIFY_REGION")),
	})

	if err != nil {
		return err
	}

	d := detail.NewDetail(
		detail.NewAppId(input.AppId),
		detail.NewBranchName(input.BranchName),
		detail.NewJobId(input.JobId),
		detail.NewJobStatusType(input.JobStatusType),
	)

	gbo, err := ac.GetBranchFromAmplify(d)

	if err != nil {
		return err
	}

	m := notification.NewMessage(
		d,
		gbo.Branch,
	)

	sm := m.BuildSlackMessage()

	sc := notification.NewSlackClient(
		&http.Client{},
		"https://slack.com/api/chat.postMessage",
	)

	sn := notification.NewSlackNotification(
		sc,
		sm,
	)

	err = sn.PostMessage()

	if err != nil {
		return err
	}

	return nil
}
