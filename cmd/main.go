package main

import (
	"github.com/hiroki-it/notify_slack_of_amplify_events/cmd/slack"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(slack.Handler)
}
