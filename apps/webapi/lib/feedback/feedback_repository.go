package feedback

import (
	"fmt"
	"time"

	"provar.se/webapi/lib/database"
)

// FeedbackType represents the type of feedback.
type FeedbackType string

const (
	TextFeedback FeedbackType = "text"
	CNPSFeedback FeedbackType = "cnps"
	CSATFeedback FeedbackType = "csat"
)

// Feedback struct represents the feedback table in the database
type Feedback struct {
	ID             string            `db:"id" json:"id"`
	OrganizationID string            `db:"organization_id" json:"organizationId"`
	CreatedAt      time.Time         `db:"created_at" json:"createdAt"`
	QuestionType   string            `db:"question_type" json:"questionType"`
	FeedbackTime   time.Time         `db:"feedback_time" json:"feedbackTime"`
	FeedbackType   FeedbackType      `db:"feedback_type" json:"feedbackType"`
	FeedbackData   map[string]string `db:"feedback_data" json:"feedbackData"`
	FeedbackTags   map[string]string `db:"feedback_tags" json:"feedbackTags"`
	FeedbackMeta   map[string]string `db:"feedback_meta" json:"feedbackMeta"`
	FeedbackUser   map[string]string `db:"feedback_user" json:"feedbackUser"`
}

// CreateFeedbackData represents the data needed to create a new feedback
type CreateFeedbackData struct {
	QuestionType string
	FeedbackTime time.Time
	FeedbackType FeedbackType
	FeedbackData map[string]string
	FeedbackTags map[string]string
	FeedbackMeta map[string]string
	FeedbackUser map[string]string
}

// Create creates a new feedback in the database
func Create(organizationID string, data *CreateFeedbackData) (*Feedback, error) {
	fb := &Feedback{
		ID:             database.NewID(),
		OrganizationID: organizationID,
		CreatedAt:      time.Now(),
		QuestionType:   data.QuestionType,
		FeedbackTime:   data.FeedbackTime,
		FeedbackType:   data.FeedbackType,
		FeedbackData:   data.FeedbackData,
		FeedbackTags:   data.FeedbackTags,
		FeedbackMeta:   data.FeedbackMeta,
		FeedbackUser:   data.FeedbackUser,
	}
	query := `
		INSERT INTO private.feedback (id, organization_id, created_at, question_type, feedback_time, feedback_type, feedback_data, feedback_tags, feedback_meta, feedback_user)
		VALUES (:id, :organization_id, :created_at, :question_type, :feedback_time, :feedback_type, :feedback_data, :feedback_tags, :feedback_meta, :feedback_user)
	`
	_, err := database.DB().NamedExec(query, fb)
	return fb, err
}

// SearchFeedbackData represents the data needed to search for feedback
type SearchFeedbackData struct {
	BegTimestamp time.Time
	EndTimestamp time.Time
	QuestionType string
	FeedbackType FeedbackType
	FeedbackTags map[string]string
	FeedbackMeta map[string]string
	FeedbackUser map[string]string
}

// Search searches for feedback in the database with given data
func Search(organizationID string, data *SearchFeedbackData) ([]*Feedback, error) {
	feedbacks := []*Feedback{}
	query := "SELECT * FROM private.feedback WHERE organization_id = $1"
	argv := []interface{}{organizationID}
	argc := 1
	if !data.BegTimestamp.IsZero() {
		query += " AND created_at >= $" + fmt.Sprint(argc)
		argv = append(argv, data.BegTimestamp)
		argc++
	}
	if !data.EndTimestamp.IsZero() {
		query += " AND created_at <= $" + fmt.Sprint(argc)
		argv = append(argv, data.EndTimestamp)
		argc++
	}
	if data.QuestionType != "" {
		query += " AND question_type = $" + fmt.Sprint(argc)
		argv = append(argv, data.QuestionType)
		argc++
	}
	if data.FeedbackType != "" {
		query += " AND feedback_type = $" + fmt.Sprint(argc)
		argv = append(argv, data.FeedbackType)
		argc++
	}
	for k, v := range data.FeedbackTags {
		query += " AND feedback_tags->>$" + fmt.Sprint(argc) + " = $" + fmt.Sprint(argc+1)
		argv = append(argv, k, v)
		argc += 2
	}
	for k, v := range data.FeedbackMeta {
		query += " AND feedback_meta->>$" + fmt.Sprint(argc) + " = $" + fmt.Sprint(argc+1)
		argv = append(argv, k, v)
		argc += 2
	}
	for k, v := range data.FeedbackUser {
		query += " AND feedback_user->>$" + fmt.Sprint(argc) + " = $" + fmt.Sprint(argc+1)
		argv = append(argv, k, v)
		argc += 2
	}
	err := database.DB().Select(&feedbacks, query, argv...)
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

// FindByID returns a feedback with the given id
func FindByID(id string) (*Feedback, error) {
	fb := &Feedback{}
	query := "SELECT * FROM private.feedback WHERE id = $1"
	err := database.DB().Get(fb, query, id)
	if err != nil {
		return nil, err
	}
	return fb, nil
}

// DeleteByID deletes a feedback with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.feedback WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// DeleteByID deletes a feedback with the given id
func (f *Feedback) Delete() error {
	return DeleteByID(f.ID)
}
