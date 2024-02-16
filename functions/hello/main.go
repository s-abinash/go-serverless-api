package main

import (
	"context"
	"go-serverless-api/config"
	middlewares "go-serverless-api/internal/middleware"
	. "go-serverless-api/internal/utils"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	. "go-serverless-api/internal/types"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request Request) (Response, error) {
	Logger.Info().Ctx(ctx).Msg("Entering Handler")
	body := map[string]string{
		"message": "Go Serverless - Voila! Your function executed successfully!",
	}
	return NewJSONResponse(http.StatusOK, body)
}

func main() {
	config.Init()
	handler := middlewares.LoggingMiddleware(Handler)
	lambda.Start(handler)
}
