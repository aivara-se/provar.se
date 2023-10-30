package database

import (
	"context"
	"errors"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cachedClient   *mongo.Client
	cachedDatabase *mongo.Database

	ErrConnectionString = errors.New("the database connection string is not valid")
)

// Connect connects to a MongoDB database and caches the client and database.
// The connection string is expected to be in the following format:
//
//	mongodb://<username>:<password>@<host>:<port>/<database>
func Connect(mongoURI string) error {
	if cachedClient != nil {
		return nil
	}
	client, err := createClient(mongoURI)
	if err != nil {
		return err
	}
	db, err := getDatabase(client, mongoURI)
	if err != nil {
		return err
	}
	err = testConnection(db)
	if err != nil {
		return err
	}
	cachedClient = client
	cachedDatabase = db
	return nil
}

// GetClient returns the connected MongoDB client.
// NOTE: The Connect function must be called before this function.
func GetClient() *mongo.Client {
	return cachedClient
}

// GetDatabase returns the connected MongoDB database.
// NOTE: The Connect function must be called before this function.
func GetDatabase() *mongo.Database {
	return cachedDatabase
}

// createClient creates a new MongoDB client with given connection string.
func createClient(mongoURI string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// getDatabase returns the database from a MongoDB client. The database name is
// extracted from the connection string. It expects the database name to be the
// "path" of the connection string.
func getDatabase(client *mongo.Client, mongoURI string) (*mongo.Database, error) {
	u, err := url.Parse(mongoURI)
	if err != nil {
		return nil, err
	}
	name := u.Path[1:]
	if name == "" {
		return nil, ErrConnectionString
	}
	db := client.Database(name)
	return db, nil
}

// testConnection tests the connection to a MongoDB database.
func testConnection(db *mongo.Database) error {
	query := bson.D{{Key: "ping", Value: 1}}
	var result bson.M
	if err := db.RunCommand(context.TODO(), query).Decode(&result); err != nil {
		return err
	}
	return nil
}
