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
func NewLambdaController(controller *Controller, eventPostUseCase *usecases.EventPostUseCase) *LambdaController {

	return &LambdaController{
		Controller:       controller,
		EventPostUseCase: eventPostUseCase,
	}
}

// PostEvent イベントをハンドリングします．
func (c *LambdaController) PostEvent(eventBridge events.CloudWatchEvent) (string, error) {

	log := logger.NewLogger()

	v := detail.NewDetailValidator()

	log.Info(string(eventBridge.Detail))

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal(eventBridge.Detail, v)

	if err != nil {
		log.Error(err.Error())
		return c.sendErrorJson(&Error{
			Status: 400,
			Errors: []string{err.Error()},
		}), nil
	}

	err = v.Validate()

	if err != nil {
		log.Error(err.Error())
		return c.sendErrorJson(&Error{
			Status: 400,
			Errors: []string{err.Error()},
		}), nil
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
		log.Error(err.Error())
		return c.sendErrorJson(&Error{
			Status: 400,
			Errors: []string{err.Error()},
		}), nil
	}

	return c.sendJson(&Success{
		Status:  200,
		Message: "Succeed to handle request",
	}), nil
}
