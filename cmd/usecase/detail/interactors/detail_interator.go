package interactors

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/entities"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/ids"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/values"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces/detail/presenters"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/inputs"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/outputs"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/services/amplify"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/services/notification"
)

type DetailInteractor struct {
	detailPresenter *presenters.DetailPresenter
}

// NewDetailInteractor コンストラクタ
func NewDetailInteractor(detailPresenter *presenters.DetailPresenter) *DetailInteractor {

	return &DetailInteractor{
		detailPresenter: detailPresenter,
	}
}

// GetDetail イベントを通知します．
func (uc *DetailInteractor) GetDetail(input *inputs.DetailInput) (*presenters.GetDetailPresenter, error) {

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
		return nil, err
	}

	cdo := &outputs.GetDetailOutput{
		Status:  200,
		Message: "Succeed to handle request",
	}

	return uc.detailPresenter.GetDetailPresenter(cdo), nil
}
