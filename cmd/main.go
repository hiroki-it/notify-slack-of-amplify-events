package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/detail/controllers"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/usecases"
)

func main() {

	l := logger.NewLogger()

	c := controllers.NewDetailController(usecases.NewEventPostUseCase(), l)

	lambda.Start(c.PostEvent)
}
