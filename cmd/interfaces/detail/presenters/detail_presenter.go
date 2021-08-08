package presenters

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/outputs"
)

type DetailPresenter struct {
}

type CreateDetailPresenter struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var _ interfaces.Presenter = &DetailPresenter{}

// CreateDetailPresenter 作成レスポンスデータを作成します．
func (dp *DetailPresenter) CreateDetailPresenter(cdo *outputs.CreateDetailOutput) *CreateDetailPresenter {
	return &CreateDetailPresenter{
		Status:  cdo.Status,
		Message: cdo.Message,
	}
}
