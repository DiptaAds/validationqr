package infrastructure

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func ConnectM(ctx context.Context) (*mongo.Database, error) {
	clientpOptions := options.Client()
	clientpOptions.ApplyURI("mongodb://192.168.18.115:27017")
	client, err := mongo.NewClient(clientpOptions)
	if err != nil {
		return nil, err

	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("validationku"), nil
}
