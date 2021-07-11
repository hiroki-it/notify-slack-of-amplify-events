package controllers

import (
	"encoding/json"
)

type Controller struct {
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
