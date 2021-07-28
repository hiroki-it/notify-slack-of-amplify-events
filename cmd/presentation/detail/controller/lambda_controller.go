package controller

import (
	"encoding/json"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/input"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/detail/validator"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/usecase"
)

type LambdaController struct {
	*presentation.Controller
	*usecase.EventPostUseCase
}

// NewLambdaController コンストラクタ
func NewLambdaController(eventPostUseCase *usecase.EventPostUseCase, logger *logger.Logger) *LambdaController {

	return &LambdaController{
		Controller:       &presentation.Controller{Logger: logger},
		EventPostUseCase: eventPostUseCase,
	}
}

// PostEvent イベントをハンドリングします．
func (c *LambdaController) PostEvent(eventBridge events.CloudWatchEvent) (string, error) {

	c.Logger.Log.Info(string(eventBridge.Detail))

	v := validator.NewDetailValidator()

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

	i := input.NewEventPostInput(
		v.AppId,
		v.BranchName,
		v.JobId,
		v.JobStatusType,
	)

	uc := usecase.NewEventPostUseCase()

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
