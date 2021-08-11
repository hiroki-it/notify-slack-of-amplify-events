package interactors

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/entities"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/ids"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/values"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/boundaries"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/requests"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/responses"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/services/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/services/notification"
)

type DetailInteractor struct {
}

var _ boundaries.DetailInputBoundary = &DetailInteractor{}

// NewDetailInteractor コンストラクタ
func NewDetailInteractor() *DetailInteractor {

	return &DetailInteractor{}
}

// NotifyEventDetail イベントの詳細を通知します．
func (di *DetailInteractor) NotifyEventDetail(input *requests.DetailRequest) (*responses.GetDetailResponse, error) {

	ac, err := amplify.NewAmplifyClient(&aws.Config{
		Region: aws.String(os.Getenv("AWS_AMPLIFY_REGION")),
	})

	if err != nil {
		return nil, err
	}

	d := entities.NewDetail(
		ids.NewAppId(input.AppId),
		values.NewBranchName(input.BranchName),
		ids.NewJobId(input.JobId),
		values.NewJobStatusType(input.JobStatusType),
	)

	gbo, err := ac.GetBranchFromAmplify(d)

	if err != nil {
		return nil, err
	}

	m := notification.NewMessage(
		d,
		gbo.Branch,
	)

	sn := notification.NewSlackNotification(
		notification.NewSlackClient(
			&http.Client{},
			"https://slack.com/api/chat.postMessage",
		),
		m.BuildSlackMessage(),
	)

	err = sn.PostMessage()

	if err != nil {
		return nil, err
	}

	return &responses.GetDetailResponse{
		Status:  200,
		Message: "Succeed to handle request",
	}, nil
}
