package feedback

import (
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

// Create creates a new feedback in the database
func Create(organizationID string, questionType string, feedbackTime time.Time, feedbackType FeedbackType, feedbackData map[string]string, feedbackTags map[string]string, feedbackMeta map[string]string, feedbackUser map[string]string) (*Feedback, error) {
	fb := &Feedback{
		ID:             database.NewID(),
		OrganizationID: organizationID,
		CreatedAt:      time.Now(),
		QuestionType:   questionType,
		FeedbackTime:   feedbackTime,
		FeedbackType:   feedbackType,
		FeedbackData:   feedbackData,
		FeedbackTags:   feedbackTags,
		FeedbackMeta:   feedbackMeta,
		FeedbackUser:   feedbackUser,
	}
	query := `
		INSERT INTO private.feedback (id, organization_id, created_at, question_type, feedback_time, feedback_type, feedback_data, feedback_tags, feedback_meta, feedback_user)
		VALUES (:id, :organization_id, :created_at, :question_type, :feedback_time, :feedback_type, :feedback_data, :feedback_tags, :feedback_meta, :feedback_user)
	`
	_, err := database.DB().NamedExec(query, fb)
	return fb, err
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
