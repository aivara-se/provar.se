package feedback

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	FeedbackTypeCNPS FeedbackType = "cnps"
	FeedbackTypeCSAT FeedbackType = "csat"
)

// String returns the string representation of a FeedbackType
func (t FeedbackType) String() string {
	return string(t)
}

// FeedbackData is the data for a feedback stored in the database
// Which fields are available depends on the feedback type.
type FeedbackData struct {
	Text string  `bson:"text,omitempty"`
	CSAT float64 `bson:"csat,omitempty" validate:"gte=0,lte=1"`
	CNPS float64 `bson:"cnps,omitempty" validate:"gte=0,lte=1"`
}

// FeedbackTags is metadata for a feedback stored in the database
type FeedbackTags map[string]string

// FeedbackDocument is a MongoDB document for feedback
type FeedbackDocument struct {
	ID             primitive.ObjectID `bson:"_id"`
	OrganizationID primitive.ObjectID `bson:"organizationId"`
	ProjectID      primitive.ObjectID `bson:"projectId,omitempty"`
	Type           FeedbackType       `bson:"type"`
	Data           FeedbackData       `bson:"data"`
	Tags           FeedbackTags       `bson:"tags,omitempty"`
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

// CreateFeedbackData is the data required to create a feedback document
type CreateFeedbackData struct {
	OrganizationID string
	ProjectID      string
	CreatedAt      time.Time
	Type           FeedbackType
	Data           FeedbackData
	Tags           FeedbackTags
}

// Validate validates the CreateFeedbackData
// TODO: enforce a maximum length for the text field
// TODO: enforce a maximum length for keys and values of tags
func (data *CreateFeedbackData) Validate() error {
	return nil
}

// createFeedbackDocument creates a FeedbackDocument from a CreateFeedbackData
func createFeedbackDocument(data CreateFeedbackData) (*FeedbackDocument, error) {
	organizationIDAsObjectId, err := primitive.ObjectIDFromHex(data.OrganizationID)
	if err != nil {
		return nil, err
	}
	projectIDAsObjectId := primitive.NilObjectID
	if data.ProjectID != "" {
		projectIDAsObjectId, err = primitive.ObjectIDFromHex(data.ProjectID)
		if err != nil {
			return nil, err
		}
	}
	var createdTime time.Time
	if data.CreatedAt.IsZero() {
		createdTime = time.Now()
	} else {
		createdTime = data.CreatedAt
	}
	doc := &FeedbackDocument{
		ID:             primitive.NewObjectIDFromTimestamp(createdTime),
		OrganizationID: organizationIDAsObjectId,
		ProjectID:      projectIDAsObjectId,
		Type:           data.Type,
		Data:           data.Data,
		Tags:           data.Tags,
	}
	return doc, nil
}

// CreateFeedback creates a feedback document in the database
func (repo *FeedbackRepository) CreateFeedback(data CreateFeedbackData) error {
	doc, err := createFeedbackDocument(data)
	if err != nil {
		return err
	}
	_, err = repo.coll.InsertOne(context.TODO(), doc)
	return err
}

// ImportFeedback creates multiple feedback documents in the database
// TODO: limit the number of feedbacks that can be created at once
func (repo *FeedbackRepository) ImportFeedback(data []CreateFeedbackData) error {
	docs := make([]interface{}, len(data))
	for i, d := range data {
		doc, err := createFeedbackDocument(d)
		if err != nil {
			return err
		}
		docs[i] = doc
	}
	_, err := repo.coll.InsertMany(context.Background(), docs)
	return err
}
