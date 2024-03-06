package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, port, host, database, collection string) (*mongo.Collection, error) {

	mongoURL := fmt.Sprintf("mongodb://%s:%s", host, port)
	cliOp := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, cliOp)
	if err != nil {
		return nil, fmt.Errorf("error with connecting to the db error is : %v", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error with ping the db error  is : %v", err)
	}
	return client.Database(database).Collection(collection), nil
}
