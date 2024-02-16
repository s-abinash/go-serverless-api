package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type EnvSpecification struct {
	Env         string `default:"local"`
	LogLevel    string `default:"0"`
	AwsRegion   string `split_words:"true" default:"us-east-1"`
	DynamoTable string `split_words:"true" required:"true"`
}

var envSpec EnvSpecification

func Init() {
	err := envconfig.Process("", &envSpec)
	if err != nil {
		fmt.Println("Error in parsing environment variables")
		log.Fatal(err.Error())
	}
	// log.Println("Config loaded successfully...")
}

func GetConfig() *EnvSpecification {
	return &envSpec
}
