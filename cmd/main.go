package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify_slack_of_amplify_events/cmd/slack"
)

func main() {
	lambda.Start(slack.Handler)
}
