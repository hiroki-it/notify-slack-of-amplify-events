package controllers

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

type Controller struct {
	logger *logger.Logger
}

type Success struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// sendJson 正常系レスポンスをJSONで返却します．
func (c *Controller) sendJson(success *Success) string {

	response, _ := json.Marshal(success)

	return string(response)
}

// sendErrorJson 異常系レスポンスをJSONで返却します．
func (c *Controller) sendErrorJson(error error) error {
	return error
}
