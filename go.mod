module github.com/hiroki-it/notify-slack-of-amplify-events

go 1.15

require (
	github.com/aws/aws-lambda-go v1.23.0
	github.com/aws/aws-sdk-go v1.25.40
	github.com/aws/aws-sdk-go-v2 v1.3.4
	github.com/aws/aws-sdk-go-v2/service/amplify v1.1.5
	github.com/stretchr/testify v1.7.0
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)

replace (
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/controllers/handler => /cmd/controllers/handler
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/amplify => /cmd/entities/amplify
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/eventbridge => /cmd/entities/eventbridge
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/exception => /cmd/entities/exception
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/slack => /cmd/entities/slack
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/file => /cmd/usecases/file
	github.com/hiroki-it/notify-slack-of-amplify-events/test/mock/amplify => /test/mock/amplify
)
