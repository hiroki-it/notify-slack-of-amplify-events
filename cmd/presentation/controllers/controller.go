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

type Error struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
}

// NewController コンストラクタ
func NewController() *Controller {
	return &Controller{}
}

// sendJson 正常系レスポンスをJSONで返却します．
func (c *Controller) sendJson(success *Success) string {

	response, _ := json.Marshal(success)

	return string(response)
}

// sendErrorJson 異常系レスポンスをJSONで返却します．
func (c *Controller) sendErrorJson(error *Error) string {

	response, _ := json.Marshal(error)

	return string(response)
}
