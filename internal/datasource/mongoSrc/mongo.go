package mongoSrc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewMongoSrc(uri string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln("Error connecting to mongoDB...", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoClient{Client: client}, nil
}

func (m *MongoClient) CloseClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := m.Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
