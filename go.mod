module github.com/hiroki-it/notify-slack-of-amplify-events

go 1.15

require (
	github.com/aws/aws-lambda-go v1.23.0
	github.com/aws/aws-sdk-go-v2 v1.3.0
	github.com/aws/aws-sdk-go-v2/config v1.1.3
	github.com/aws/aws-sdk-go-v2/service/amplify v1.1.3
	github.com/joho/godotenv v1.3.0
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/controllers => /cmd/usecases/controllers
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/amplify => /cmd/usecases/amplify
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/eventbridge => /cmd/usecases/eventbridge
	github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/slack => /cmd/usecases/slack
	github.com/hiroki-it/notify-slack-of-amplify-events/config => /config
)
