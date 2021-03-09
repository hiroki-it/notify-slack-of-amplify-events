package main

import(
  "github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/amplify"
  "github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/slack"
)

func main() {
  slackMessage := slack_message.buildMessage()
  slack_message.postMessage(slackMessage)
}
