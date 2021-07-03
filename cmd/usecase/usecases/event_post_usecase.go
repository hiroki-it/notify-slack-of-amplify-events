package usecases

import (
	"net/http"
	"os"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/inputs"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/services/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/services/notification"
)

type EventPostUseCase struct {
}

// NewEventPostUseCase コンストラクタ
func NewEventPostUseCase() *EventPostUseCase {

	return &EventPostUseCase{}
}

// PostEvent イベントを通知します．
func (uc *EventPostUseCase) PostEvent(input *inputs.EventPostInput) error {

	amplifyApi, err := amplify.NewAmplifyAPI(os.Getenv("AWS_AMPLIFY_REGION"))

	if err != nil {
		return err
	}

	ac := amplify.NewAmplifyClient(amplifyApi)

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
