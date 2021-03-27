module github.com/Hiroki-IT/notify_slack_of_amplify_events

go 1.15

require (
    github.com/aws/aws-lambda-go v1.23.0
    github.com/aws/aws-sdk-go v1.38.7
    github.com/stretchr/testify v1.7.0
    )

replace (
    github.com/Hiroki-IT/notify_slack_of_amplify_events/cmd/slack => ./slack
    )
