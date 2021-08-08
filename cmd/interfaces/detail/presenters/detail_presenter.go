package presenters

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/outputs"
)

type DetailPresenter struct {
}

type GetDetailPresenter struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var _ interfaces.Presenter = &DetailPresenter{}

// GetDetailPresenter 作成レスポンスデータを作成します．
func (dp *DetailPresenter) GetDetailPresenter(cdo *outputs.GetDetailOutput) *GetDetailPresenter {
	return &GetDetailPresenter{
		Status:  cdo.Status,
		Message: cdo.Message,
	}
}
