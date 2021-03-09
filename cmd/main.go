package main

import (
  "github.com/cmd/amplify-cli"
  "github.com/cmd/slack-message"
)

func main() {
  slackMessage := slack_message.buildMessage()
  slack_message.postMessage(slackMessage)
}
