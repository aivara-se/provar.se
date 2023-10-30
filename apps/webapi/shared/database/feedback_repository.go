package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// feedbackCollectionName is the name of the feedback collection
	feedbackCollectionName = "feedbacks"
)

var (
	cachedFeedbackRepository *FeedbackRepository
)

// FeedbackRepository is a repository for feedback
type FeedbackRepository struct {
	coll *mongo.Collection
}

// GetFeedbackRepository returns a FeedbackRepository
func GetFeedbackRepository() *FeedbackRepository {
	if cachedFeedbackRepository != nil {
		return cachedFeedbackRepository
	}
	db := GetDatabase()
	coll := db.Collection(feedbackCollectionName)
	repo := &FeedbackRepository{coll: coll}
	cachedFeedbackRepository = repo
	return repo
}

// CreateFeedbackParams is the details of a feedback
type CreateFeedbackParams struct {
	OrganizationID string `bson:"organizationId"`
}

// CreateFeedback creates a feedback document in the database
func (repo *FeedbackRepository) CreateFeedback(params *CreateFeedbackParams) error {
	_, err := repo.coll.InsertOne(context.TODO(), params)
	return err
}
