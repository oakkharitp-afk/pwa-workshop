package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoClient(uri string) (*mongo.Client, error) {

	ctx, cancle := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()

	clientMongo, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// context with timeout when ping to client
	ctx, cancle = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancle()

	// check client is ready
	err = clientMongo.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return clientMongo, nil

}

func ConnectSQLClient(uri, driver string) (*sql.DB, error) {
	client, err := sql.Open(driver, uri)
	if err != nil {
		return nil, err
	}

	// check client is ready
	err = client.Ping()
	if err != nil {
		return nil, err
	}

	return client, nil
}
