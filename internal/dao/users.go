package dao

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	. "go-serverless-api/internal/utils"
)

func GetUserDBItem(ctx context.Context, user UserInfo) UserDBItem {
	primaryKey := primaryKey{
		Pk: fmt.Sprintf("userId#%s", user.ID),
		Sk: fmt.Sprintf("userId#%s", user.ID),
	}
	userInfoDBItem := UserDBItem{
		primaryKey: primaryKey,
		UserInfo:   user,
	}
	return userInfoDBItem
}

func PutUserInfo(ctx context.Context, db *DB, user UserInfo) error {
	Logger.Info().Ctx(ctx).Msgf("Entering Into PutUserInfo")
	userDBItem := GetUserDBItem(ctx, user)

	Logger.Debug().Any("User DB Item", userDBItem).Send()

	_, err := putItemWrapper(db, userDBItem)
	if err != nil {
		Logger.Error().Ctx(ctx).Err(err).Msg("DBError.DBPutError")
		return err
	}
	Logger.Info().Ctx(ctx).Msgf("Exiting Into PutUserInfo")
	return nil
}

func GetUserInfo(ctx context.Context, db *DB, userID string) (*UserInfo, error) {
	Logger.Info().Ctx(ctx).Msgf("Entering Into GetMatchInfo")
	metadata := UserInfo{
		ID: userID,
	}
	key := GetUserDBItem(ctx, metadata).primaryKey

	Logger.Debug().Ctx(ctx).Msgf("PrimaryKey: %s", key)

	response, err := getItemWrapper(db, key)
	if err != nil {
		return nil, err
	}

	Logger.Debug().Ctx(ctx).Any("DB Response", response).Send()

	userDBItem := UserDBItem{}
	err = attributevalue.UnmarshalMap(response.Item, &userDBItem)
	if err != nil {
		Logger.Error().Ctx(ctx).Err(err).Msg("DBError.DBUnmarshalError")
		return nil, err
	}
	Logger.Info().Ctx(ctx).Msgf("Exiting Into GetMatchInfo")
	return &userDBItem.UserInfo, nil
}
