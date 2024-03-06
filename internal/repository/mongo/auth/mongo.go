package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/MamushevArup/jwt-auth/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserInserter interface {
	InsertUser(ctx context.Context, guid, hashToken string) error
	UniqueUser(ctx context.Context, guid string) (bool, error)
	UserToken(ctx context.Context, guid string) (string, error)
	UpdateRefreshToken(ctx context.Context, guid, refreshToken string) error
}

type auth struct {
	db *mongo.Collection
	lg *logger.Logger
}

func NewAuth(lg *logger.Logger, db *mongo.Collection) UserInserter {
	return &auth{
		db: db,
		lg: lg,
	}
}

type user struct {
	GUID        string `bson:"guid"`
	HashedToken string `bson:"hashed_token"`
	// Add other user fields as needed
}

func (a *auth) UpdateRefreshToken(ctx context.Context, guid, refreshToken string) error {

	filter := bson.M{"guid": guid}

	update := bson.M{"$set": bson.M{"hashed_token": refreshToken}}

	// Execute the UpdateOne operation
	result, err := a.db.UpdateOne(ctx, filter, update)
	if err != nil {
		a.lg.Errorf("unable update document %v", err)
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (a *auth) InsertUser(ctx context.Context, guid, hashToken string) error {

	u := user{
		GUID:        guid,
		HashedToken: hashToken,
	}

	_, err := a.db.InsertOne(ctx, u)
	if err != nil {
		a.lg.Errorf("inset failed %v", err)
		return err
	}

	return nil
}

// UniqueUser look for user with guid and return user existence
func (a *auth) UniqueUser(ctx context.Context, guid string) (bool, error) {

	filter := bson.M{"guid": guid}

	result := a.db.FindOne(ctx, filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			a.lg.Errorf("no document found %v", result.Err())
			// if user doesn't exist
			return false, nil
		}
		return false, result.Err()
	}
	return true, nil
}

func (a *auth) UserToken(ctx context.Context, guid string) (string, error) {

	filter := bson.M{"guid": guid}

	projection := bson.M{"hashed_token": 1}

	result := a.db.FindOne(ctx, filter, options.FindOne().SetProjection(projection))

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return "", fmt.Errorf("user not found")
		}
		return "", result.Err() // Other error occurred
	}

	var token struct {
		Refresh string `bson:"hashed_token"`
	}

	if err := result.Decode(&token); err != nil {
		a.lg.Errorf("decode fail %v", err)
		return "", err
	}

	return token.Refresh, nil
}
