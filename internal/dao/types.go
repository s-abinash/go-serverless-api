package dao

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DB struct {
	client    *dynamodb.Client
	tableName string
}

type primaryKey struct {
	Pk string `dynamodbav:"pk" json:"pk"`
	Sk string `dynamodbav:"sk" json:"sk"`
}

type UserInfo struct {
	ID       string `dynamodbav:"id" json:"id"`
	Name     string `dynamodbav:"name" json:"name"`
	Email    string `dynamodbav:"email" json:"email"`
	Age      int    `dynamodbav:"age" json:"age"`
	IsActive bool   `dynamodbav:"isActive" json:"isActive"`
}

type UserDBItem struct {
	primaryKey
	UserInfo
}
