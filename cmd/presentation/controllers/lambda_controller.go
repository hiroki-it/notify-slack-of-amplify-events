package controllers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/validators/detail"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/inputs"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/usecases"
)

type LambdaController struct {
	*Controller
	*usecases.EventPostUseCase
}

// NewLambdaController コンストラクタ
func NewLambdaController(eventPostUseCase *usecases.EventPostUseCase, logger *logger.Logger) *LambdaController {

	return &LambdaController{
		Controller:       &Controller{logger},
		EventPostUseCase: eventPostUseCase,
	}
}

// PostEvent イベントをハンドリングします．
func (c *LambdaController) PostEvent(eventBridge events.CloudWatchEvent) (string, error) {

	c.logger.Log.Info(string(eventBridge.Detail))

	v := detail.NewDetailValidator()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal(eventBridge.Detail, v)

	if err != nil {
		c.logger.Log.Error(err.Error())
		return "", c.sendErrorJson(err)
	}

	errorMessage := v.Validate()

	if errorMessage != nil {
		c.logger.Log.Error(errorMessage.Error())
		return "", c.sendErrorJson(err)
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
		c.logger.Log.Error(err.Error())
		return "", c.sendErrorJson(err)
	}

	return c.sendJson(&Success{
		Status:  200,
		Message: "Succeed to handle request",
	}), nil
}
