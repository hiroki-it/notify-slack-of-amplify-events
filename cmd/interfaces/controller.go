package interfaces

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

type Controller struct {
	Logger *logger.Logger
}

type Success struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// SendJson 正常系レスポンスをJSONで返却します．
func (c *Controller) SendJson(success *Success) string {

	response, _ := json.Marshal(success)

	return string(response)
}

// SendErrorJson 異常系レスポンスをJSONで返却します．
func (c *Controller) SendErrorJson(error error) error {
	return error
}
