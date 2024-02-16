package dao

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	. "go-serverless-api/internal/utils"
)

func GetDynamoDBClient(table string, region string) (*DB, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		Logger.Error().Err(err).Msg("DBError.DBFailedToLoad")
		return nil, err
	}
	dynamodbClient := dynamodb.NewFromConfig(cfg)

	return &DB{
		tableName: table,
		client:    dynamodbClient,
	}, err
}

func getItemWrapper(db *DB, key interface{}) (*dynamodb.GetItemOutput, error) {

	keyAttribute, err := attributevalue.MarshalMap(key)
	if err != nil {
		Logger.Error().Err(err).Msg("DBError.DBMarshalError")
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(db.tableName),
		Key:       keyAttribute,
	}

	response, err := db.client.GetItem(context.TODO(), input)
	if err != nil {
		Logger.Error().Err(err).Msg("DBError.DBFetchError")
		return nil, err
	}

	return response, nil
}

func putItemWrapper(db *DB, inputItem interface{}) (*dynamodb.PutItemOutput, error) {

	item, err := attributevalue.MarshalMap(inputItem)
	if err != nil {
		Logger.Error().Err(err).Msg("DBError.DBMarshalError")
		return nil, err
	}

	Logger.Debug().Any("ItemMap", item).Send()

	response, err := db.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(db.tableName),
		Item:      item,
		// ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		Logger.Error().Err(err).Msg("DBError.DBWriteError")
		return nil, err
	}

	return response, nil
}
