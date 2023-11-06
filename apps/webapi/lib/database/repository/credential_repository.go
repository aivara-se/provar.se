package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"provar.se/webapi/lib/database"
)

const (
	// credentialsCollectionName is the name of the credentials collection
	credentialsCollectionName = "credentials"
)

var (
	// cachedCredentialRepository is a reference to a CredentialRepository
	cachedCredentialRepository *CredentialRepository
)

// CredentialModel is a MongoDB document for credentials
type CredentialModel struct {
	Name           string `bson:"name"`
	OrganizationID string `bson:"organizationId"`
	ClientID       string `bson:"clientId"`
	ClientSecret   string `bson:"clientSecret"`
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

// FindCredentialByClientID searches for a credential by organization id and the client id
func (repo *CredentialRepository) FindCredentialByClientID(organizationID string, clientID string) (*CredentialModel, error) {
	var result CredentialModel
	filter := bson.M{"organizationId": organizationID, "clientId": clientID}
	err := repo.coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
