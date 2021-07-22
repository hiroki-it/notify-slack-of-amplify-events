package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/controllers"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/usecases"
)

func main() {

	l := logger.NewLogger()

	c := controllers.NewLambdaController(usecases.NewEventPostUseCase(), l)

	lambda.Start(c.PostEvent)
}
