package interactors

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/entities"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/ids"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/values"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/inputs"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/services/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/services/notification"
)

type EventPostInteractor struct {
}

// NewEventPostInteractor コンストラクタ
func NewEventPostInteractor() *EventPostInteractor {

	return &EventPostInteractor{}
}

// PostEvent イベントを通知します．
func (uc *EventPostInteractor) PostEvent(input *inputs.EventPostInput) error {

	ac, err := amplify.NewAmplifyClient(&aws.Config{
		Region: aws.String(os.Getenv("AWS_AMPLIFY_REGION")),
	})

	if err != nil {
		return err
	}

	d := entities.NewDetail(
		ids.NewAppId(input.AppId),
		values.NewBranchName(input.BranchName),
		ids.NewJobId(input.JobId),
		values.NewJobStatusType(input.JobStatusType),
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
