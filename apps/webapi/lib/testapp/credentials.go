package testapp

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/database"
)

const (
	// credentialsCollectionName is the name of the credentials collection
	credentialsCollectionName = "credentials"

	// TestCredentialName is the name of the test credential
	TestCredentialName = "Test Credential"

	// TestCredentialKey is the api key of the test credential
	TestCredentialKey = "test-api-key"

	// TestOrganizationID is the ID of the test organization
	TestOrganizationID = "657d5b9553420cf487c856e5"
)

// createTestCredentials creates a test credential in the database if missing
func createTestCredentials() error {
	coll := database.GetDatabase().Collection(credentialsCollectionName)
	organizationID, err := primitive.ObjectIDFromHex(TestOrganizationID)
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(context.TODO(), credential.CredentialDocument{
		ID:             primitive.NewObjectID(),
		CreatedAt:      time.Now(),
		LastUsedAt:     time.Now(),
		Name:           TestCredentialName,
		OrganizationID: organizationID,
		Key:            TestCredentialKey,
	})
	return err
}
