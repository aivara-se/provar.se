package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"provar.se/webapi/lib/database"
)

const (
	// feedbackCollectionName is the name of the feedback collection
	feedbackCollectionName = "feedbacks"
)

var (
	// cachedFeedbackRepository is a reference to a FeedbackRepository
	cachedFeedbackRepository *FeedbackRepository
)

// FeedbackDocument is a MongoDB document for feedback
type FeedbackDocument struct {
	OrganizationID string `bson:"organizationId"`
}

// FeedbackRepository is a repository for feedback
type FeedbackRepository struct {
	coll *mongo.Collection
}

// GetFeedbackRepository returns a FeedbackRepository
func GetFeedbackRepository() *FeedbackRepository {
	if cachedFeedbackRepository != nil {
		return cachedFeedbackRepository
	}
	db := database.GetDatabase()
	coll := db.Collection(feedbackCollectionName)
	repo := &FeedbackRepository{coll: coll}
	cachedFeedbackRepository = repo
	return repo
}

// CreateFeedback creates a feedback document in the database
func (repo *FeedbackRepository) CreateFeedback(organizationID string) error {
	doc := &FeedbackDocument{OrganizationID: organizationID}
	_, err := repo.coll.InsertOne(context.TODO(), doc)
	return err
}
