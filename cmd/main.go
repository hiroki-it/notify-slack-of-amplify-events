package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/detail/controller"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/usecase"
)

func main() {

	l := logger.NewLogger()

	c := controller.NewLambdaController(usecase.NewEventPostUseCase(), l)

	lambda.Start(c.PostEvent)
}
