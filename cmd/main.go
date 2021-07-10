package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/controllers"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecase/usecases"
)

func main() {

	c := controllers.NewLambdaController(
		controllers.NewController(),
		usecases.NewEventPostUseCase(),
	)

	lambda.Start(c.PostEvent)
}
