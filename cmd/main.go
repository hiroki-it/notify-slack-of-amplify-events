package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/slack"
    )

func main() {
    lambda.Start(handler)
}
