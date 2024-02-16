package main

import (
	"context"
	"go-serverless-api/config"
	"go-serverless-api/internal/dao"
	middlewares "go-serverless-api/internal/middleware"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	. "go-serverless-api/internal/types"
	. "go-serverless-api/internal/utils"
)

func Handler(ctx context.Context, request Request) (Response, error) {
	Logger.Info().Msg("Entering Handler - Get User")
	env := config.GetConfig()

	userID := request.PathParameters["userID"]

	Logger.Debug().Ctx(ctx).Msgf("User ID: %s", userID)

	db, _ := dao.GetDynamoDBClient(env.DynamoTable, env.AwsRegion)

	user, err := dao.GetUserInfo(ctx, db, userID)
	if err != nil {
		Logger.Error().Err(err).Msg("Error in Getting User")
		return NewJSONErrorResponse(http.StatusInternalServerError)
	}

	Logger.Info().Msg("Exiting Handler")
	return NewJSONResponse(http.StatusCreated, user)
}

func main() {
	config.Init()
	LoggerInit()
	handler := middlewares.LoggingMiddleware(Handler)
	lambda.Start(handler)
}
