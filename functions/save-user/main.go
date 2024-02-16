package main

import (
	"context"
	"encoding/json"
	"go-serverless-api/config"
	"go-serverless-api/internal/dao"
	middlewares "go-serverless-api/internal/middleware"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/oklog/ulid/v2"
	. "go-serverless-api/internal/types"
	. "go-serverless-api/internal/utils"
)

func Handler(ctx context.Context, request Request) (Response, error) {
	Logger.Info().Msg("Entering Handler - Save User")

	env := config.GetConfig()
	var user dao.UserInfo
	//	err := ParseAndValidate(request.Body, &user)
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		Logger.Error().Err(err).Msgf("Error in Parsing User, %s", request.Body)
		return NewJSONErrorResponse(http.StatusBadRequest)
	}
	user.ID = ulid.Make().String()

	Logger.Debug().Ctx(ctx).Msgf("User Details: %+v", user)

	db, _ := dao.GetDynamoDBClient(env.DynamoTable, env.AwsRegion)

	err = dao.PutUserInfo(ctx, db, user)
	if err != nil {
		Logger.Error().Err(err).Msg("Error in Saving User")
		return NewJSONErrorResponse(http.StatusInternalServerError)
	}

	Logger.Info().Msg("Exiting Handler")
	return NewJSONResponse(http.StatusCreated, map[string]string{"msg": "User Created"})
}

func main() {
	config.Init()
	LoggerInit()
	handler := middlewares.LoggingMiddleware(Handler)
	lambda.Start(handler)
}
