package main

import (
	"github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/slack"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(slack.Handler)
}
