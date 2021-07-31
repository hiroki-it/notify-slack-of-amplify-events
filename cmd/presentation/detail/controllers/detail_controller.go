package controllers

import (
	"encoding/json"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/inputs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/detail/validators"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/usecases"
)

type DetailController struct {
	*presentation.Controller
	*usecases.EventPostUseCase
}

// NewDetailController コンストラクタ
func NewDetailController(eventPostUseCase *usecases.EventPostUseCase, logger *logger.Logger) *DetailController {

	return &DetailController{
		Controller:       &presentation.Controller{Logger: logger},
		EventPostUseCase: eventPostUseCase,
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

	i := inputs.NewEventPostInput(
		v.AppId,
		v.BranchName,
		v.JobId,
		v.JobStatusType,
	)

	uc := usecases.NewEventPostUseCase()

	err = uc.PostEvent(i)

	if err != nil {
		c.Logger.Log.Error(err.Error())
		return "", c.SendErrorJson(err)
	}

	return c.SendJson(&presentation.Success{
		Status:  200,
		Message: "Succeed to handle request",
	}), nil
}
