package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"provar.se/webapi/lib/database"
)

const (
	// credentialsCollectionName is the name of the credentials collection
	credentialsCollectionName = "credentials"
)

var (
	// ErrCredentialNotFound is returned when the api key is not found
	ErrCredentialNotFound = errors.New("credential not found")

	// cachedCredentialRepository is a reference to a CredentialRepository
	cachedCredentialRepository *CredentialRepository
)

// CredentialDocument is a MongoDB document for credentials
type CredentialDocument struct {
	Name           string `bson:"name"`
	OrganizationID string `bson:"organizationId"`
	Key            string `bson:"key"`
}

// CredentialRepository is a repository for credentials
type CredentialRepository struct {
	coll *mongo.Collection
}

// GetCredentialRepository returns a CredentialRepository
func GetCredentialRepository() *CredentialRepository {
	if cachedCredentialRepository != nil {
		return cachedCredentialRepository
	}
	db := database.GetDatabase()
	coll := db.Collection(credentialsCollectionName)
	repo := &CredentialRepository{coll: coll}
	cachedCredentialRepository = repo
	return repo
}

// FindCredenial finds a credential from in the database. Returns an error if not found.
func (repo *CredentialRepository) FindCredenial(key string) (*CredentialDocument, error) {
	var result CredentialDocument
	filter := bson.M{"key": key}
	err := repo.coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, ErrCredentialNotFound
	}
	return &result, err
}
