package controllers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces/detail/validators"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/inputs"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/interactors"
)

type DetailController struct {
	*interfaces.Controller
	*interactors.DetailInteractor
}

// NewDetailController コンストラクタ
func NewDetailController(DetailInteractor *interactors.DetailInteractor, logger *logger.Logger) *DetailController {

	return &DetailController{
		Controller:       &interfaces.Controller{Logger: logger},
		DetailInteractor: DetailInteractor,
	}
}

// PostEvent イベントをハンドリングします．
func (c *DetailController) PostEvent(eventBridge events.CloudWatchEvent) (string, error) {

	c.Logger.Log.Info(string(eventBridge.Detail))

	v := validators.NewDetailValidator()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal(eventBridge.Detail, v)

	if err != nil {
		c.Logger.Log.Error(err.Error())
		return "", c.SendErrorJson(err)
	}

	errorMessage := v.Validate()

	if errorMessage != nil {
		c.Logger.Log.Error(errorMessage.Error())
		return "", c.SendErrorJson(err)
	}

	i := inputs.NewDetailInput(
		v.AppId,
		v.BranchName,
		v.JobId,
		v.JobStatusType,
	)

	uc := interactors.NewDetailInteractor()

	err = uc.PostEvent(i)

	if err != nil {
		c.Logger.Log.Error(err.Error())
		return "", c.SendErrorJson(err)
	}

	return c.SendJson(&interfaces.Success{
		Status:  200,
		Message: "Succeed to handle request",
	}), nil
}
