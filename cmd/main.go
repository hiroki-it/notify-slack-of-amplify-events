package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/controllers/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
