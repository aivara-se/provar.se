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

// FeedbackType is a type of feedback. This determines what other fields
// will be available on a feedback document in the database.
type FeedbackType string

const (
	FeedbackTypeText FeedbackType = "text"
)

// String returns the string representation of a FeedbackType
func (t FeedbackType) String() string {
	return string(t)
}

// FeedbackTextData is the data stored for for a text feedback
type FeedbackTextData struct {
	Text string `bson:"text"`
}

// FeedbackData is the data for a feedback stored in the database
type FeedbackData struct {
	*FeedbackTextData `bson:"inline"`
}

// FeedbackDocument is a MongoDB document for feedback
type FeedbackDocument struct {
	OrganizationID string       `bson:"organizationId"`
	Type           FeedbackType `bson:"type"`
	Data           FeedbackData `bson:"data"`
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
func (repo *FeedbackRepository) CreateFeedback(organizationID string, feedbackType FeedbackType, feedbackData FeedbackData) error {
	doc := &FeedbackDocument{OrganizationID: organizationID}
	_, err := repo.coll.InsertOne(context.TODO(), doc)
	return err
}
