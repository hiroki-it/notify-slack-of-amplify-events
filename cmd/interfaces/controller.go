package interfaces

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

type Controller struct {
	Logger *logger.Logger
}

// JSON 正常系レスポンスをJSONで返却します．
func (c *Controller) JSON(response interface{}) string {

	byteJson, _ := json.Marshal(response)

	return string(byteJson)
}
