package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify_slack_of_amplify_events/cmd/slack"
	"github.com/hiroki-it/notify_slack_of_amplify_events/config"
)

func main() {
	config.LoadConfig()
	lambda.Start(slack.LambdaHandler)
}
