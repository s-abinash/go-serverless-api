package utils

import (
	"encoding/json"
	"go-serverless-api/internal/types"
	"net/http"
)

var Headers = map[string]string{
	"Content-Type": "application/json",
}

func NewJSONResponse(statusCode int, body interface{}) (types.Response, error) {
	switch b := body.(type) {
	case error:
		return NewJSONErrorResponse(statusCode)
	case string:
		headers := make(map[string]string)
		for k, v := range Headers {
			headers[k] = v
		}
		return types.Response{
			StatusCode: statusCode,
			Body:       b,
			Headers:    headers,
		}, nil
	default:
		bodyStr, err := json.Marshal(body)
		if err != nil {
			_, err := NewJSONErrorResponse(http.StatusInternalServerError)
			if err != nil {
				return types.Response{}, err
			}
		}
		headers := make(map[string]string)
		for k, v := range Headers {
			headers[k] = v
		}

		return types.Response{
			StatusCode: statusCode,
			Body:       string(bodyStr),
			Headers:    headers,
		}, nil
	}
}

func NewJSONErrorResponse(statusCode int) (types.Response, error) {
	errBody := "Error"
	bodyStr, err := json.Marshal(errBody)
	if err != nil {
		return NewJSONErrorResponse(http.StatusInternalServerError)
	}
	return types.Response{
		StatusCode: statusCode,
		Body:       string(bodyStr),
		Headers:    Headers,
	}, nil
}
