package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify_slack_of_amplify_events/cmd/handler"
)

func main() {
	lambda.Start(handler.LambdaHandler)
}
