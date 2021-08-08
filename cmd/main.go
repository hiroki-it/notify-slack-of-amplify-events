package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/interfaces/detail/controllers"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/detail/interactors"
)

func main() {

	l := logger.NewLogger()

	c := controllers.NewDetailController(interactors.NewEventPostInteractor(), l)

	lambda.Start(c.PostEvent)
}
