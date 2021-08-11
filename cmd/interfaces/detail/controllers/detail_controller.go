package controllers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces/detail/validators"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/boundaries"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/requests"
)

type DetailController struct {
	*interfaces.Controller
	detailInputBoundary boundaries.DetailInputBoundary
}

// NewDetailController コンストラクタ
func NewDetailController(detailInputBoundary boundaries.DetailInputBoundary, logger *logger.Logger) *DetailController {

	return &DetailController{
		Controller:          &interfaces.Controller{Logger: logger},
		detailInputBoundary: detailInputBoundary,
	}
}

// HandleEvent イベントをハンドリングします．
func (c *DetailController) HandleEvent(eventBridge events.CloudWatchEvent) (string, error) {

	c.Logger.Log.Info(string(eventBridge.Detail))

	v := validators.NewDetailValidator()

	// eventbridgeから転送されたJSONを構造体にマッピングします．
	err := json.Unmarshal(eventBridge.Detail, v)

	if err != nil {
		c.Logger.Log.Error(err.Error())
		return "", err
	}

	errorMessage := v.Validate()

	if errorMessage != nil {
		c.Logger.Log.Error(errorMessage.Error())
		return "", err
	}

	i := &requests.DetailRequest{
		AppId:         v.AppId,
		BranchName:    v.BranchName,
		JobId:         v.JobId,
		JobStatusType: v.JobStatusType,
	}

	dib, err := c.detailInputBoundary.NotifyEventDetail(i)

	if err != nil {
		c.Logger.Log.Error(err.Error())
		return "", err
	}

	return c.JSON(dib), nil
}
