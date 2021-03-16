module github.com/Hiroki-IT/notify_slack_of_amplify_events

go 1.15

require (
	github.com/aws/aws-lambda-go v1.23.0
	github.com/aws/aws-sdk-go v1.35.26 // indirect
)

replace (
	github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/amplify => ./amplify
	github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/slack => ./slack
)
