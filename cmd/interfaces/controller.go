package interfaces

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

type Controller struct {
	Logger *logger.Logger
}

// SendJson 正常系レスポンスをJSONで返却します．
func (c *Controller) SendJson(presenter Presenter) string {

	response, _ := json.Marshal(presenter)

	return string(response)
}

// SendErrorJson 異常系レスポンスをJSONで返却します．
func (c *Controller) SendErrorJson(error error) error {
	return error
}
