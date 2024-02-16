package utils

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

func ParseAndValidate(body string, target interface{}) error {
	err := json.Unmarshal([]byte(body), target)
	if err != nil {
		return err
	}

	validate := validator.New()
	return validate.Struct(target)
}
