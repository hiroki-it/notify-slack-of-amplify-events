package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/application/controllers"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
