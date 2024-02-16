package middlewares

import (
	"context"
	"go-serverless-api/config"
	"io"
	"os"
	"time"

	. "go-serverless-api/internal/types"

	"github.com/rs/zerolog"
)

// LoggingMiddleware Logging middleware equivalent for Lambda
func LoggingMiddleware(handler func(context.Context, Request) (Response, error)) func(context.Context, Request) (Response, error) {
	env := config.GetConfig()
	return func(ctx context.Context, req Request) (Response, error) {
		startTime := time.Now()
		requestID := req.RequestContext.RequestID

		// Set the request ID in the context
		ctx = context.WithValue(ctx, "request_id", requestID)

		// Extract client details from the request
		clientIP := req.Headers["X-Forwarded-For"]
		method := req.HTTPMethod
		path := req.Path
		queryParams := req.QueryStringParameters
		pathParams := req.PathParameters

		var output io.Writer

		if env.Env == "local" {
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			}
		} else {
			output = os.Stdout
		}

		// Set up logger
		logger := zerolog.New(output).With().
			Timestamp().
			Str("request_id", requestID).
			Str("client_ip", clientIP).
			Str("method", method).
			Str("path", path).
			Interface("query_params", queryParams).
			Interface("path_params", pathParams).
			Logger()

		logger.Info().Msg("Request Received")

		// Execute the handler
		response, err := handler(ctx, req)
		latency := time.Since(startTime)

		// Update logger with additional fields
		logger = logger.With().
			Timestamp().
			Str("request_id", requestID).
			Str("client_ip", clientIP).
			Str("method", method).
			Str("path", path).
			Int("status_code", response.StatusCode).
			Int64("latency_ms", latency.Milliseconds()).
			Logger()

		if err != nil {
			logger.Error().Err(err).Msg("Request completed with error")
			return response, err
		}

		response.Headers["X-Request-ID"] = requestID

		logger.Info().Msg("Request completed successfully")
		return response, nil
	}
}
