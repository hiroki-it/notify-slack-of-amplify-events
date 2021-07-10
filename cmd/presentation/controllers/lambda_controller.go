package controllers

import (
	"encoding/json"
	"fmt"

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
		return fmt.Sprint("Failed to handle request"), err
	}

	err = v.Validate()

	if err != nil {
		log.Error(err.Error())
		return fmt.Sprint("Failed to handle request"), err
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
		return fmt.Sprint("Failed to handle request"), err
	}

	return fmt.Sprint("Succeed to handle request"), nil
}
