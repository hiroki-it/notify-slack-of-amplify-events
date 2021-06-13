module github.com/hiroki-it/notify-slack-of-amplify-events

go 1.15

require (
	github.com/aws/aws-lambda-go v1.23.0
	github.com/aws/aws-sdk-go v1.25.40
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.16.0
)

replace (
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/application/controllers => /cmd/application/controllers
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/amplify => /cmd/domain/amplify
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/eventbridge => /cmd/domain/eventbridge
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/slack => /cmd/domain/slack
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file => /cmd/infrastructure/file
	github.com/hiroki-it/notify-slack-of-amplify-events/mock/amplify => /mock/amplify
)
