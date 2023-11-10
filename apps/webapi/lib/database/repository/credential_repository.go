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

// IsValidCredenial make sure that the credential exists in the database and is valid
func (repo *CredentialRepository) IsValidCredenial(organizationID string, key string) error {
	var result CredentialModel
	filter := bson.M{"organizationId": organizationID, "key": key}
	return repo.coll.FindOne(context.TODO(), filter).Decode(&result)
}
