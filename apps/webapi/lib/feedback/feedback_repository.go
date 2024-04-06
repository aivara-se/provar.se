package feedback

import (
	"fmt"
	"time"

	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/metadata"
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
	ID             string             `db:"id" json:"id"`
	OrganizationID string             `db:"organization_id" json:"organizationId"`
	CreatedAt      time.Time          `db:"created_at" json:"createdAt"`
	FeedbackTime   time.Time          `db:"feedback_time" json:"feedbackTime"`
	FeedbackType   FeedbackType       `db:"feedback_type" json:"feedbackType"`
	FeedbackData   database.StringMap `db:"feedback_data" json:"feedbackData"`
	FeedbackTags   database.StringMap `db:"feedback_tags" json:"feedbackTags"`
	FeedbackMeta   database.StringMap `db:"feedback_meta" json:"feedbackMeta"`
	FeedbackUser   database.StringMap `db:"feedback_user" json:"feedbackUser"`
}

// CreateFeedbackData represents the data needed to create a new feedback
type CreateFeedbackData struct {
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
		FeedbackTime:   data.FeedbackTime,
		FeedbackType:   data.FeedbackType,
		FeedbackData:   data.FeedbackData,
		FeedbackTags:   data.FeedbackTags,
		FeedbackMeta:   data.FeedbackMeta,
		FeedbackUser:   data.FeedbackUser,
	}
	query := `
		INSERT INTO private.feedback (id, organization_id, created_at, feedback_time, feedback_type, feedback_data, feedback_tags, feedback_meta, feedback_user)
		VALUES (:id, :organization_id, :created_at, :feedback_time, :feedback_type, :feedback_data, :feedback_tags, :feedback_meta, :feedback_user)
	`
	_, err := database.DB().NamedExec(query, fb)
	return fb, err
}

// Enrich enriches the feedback with additional data
func Enrich(fb *Feedback) error {
	if _, ok := fb.FeedbackMeta["request-ip"]; ok {
		location, err := metadata.GetLocation(fb.FeedbackMeta["request-ip"])
		if err == nil {
			fb.FeedbackMeta["location"] = location.Country.Names["en"]
		}
	}
	return nil
}

// SearchFeedbackData represents the data needed to search for feedback
type SearchFeedbackData struct {
	PageLimit    int
	PageOffset   int
	BegTimestamp time.Time
	EndTimestamp time.Time
	FeedbackType []FeedbackType
	FeedbackTags map[string]string
	FeedbackMeta map[string]string
	FeedbackUser map[string]string
}

// Search searches for feedback in the database with given data and returns the result.
func Search(organizationID string, data *SearchFeedbackData) ([]*Feedback, error) {
	feedbacks := []*Feedback{}
	query, argv := searchQuery(organizationID, data, false)
	err := database.DB().Select(&feedbacks, query, argv...)
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

// Count searches for feedback in the database with given data and returns the count.
func Count(organizationID string, data *SearchFeedbackData) (int, error) {
	var count int
	query, argv := searchQuery(organizationID, data, true)
	err := database.DB().Get(&count, query, argv...)
	if err != nil {
		return 0, err
	}
	return count, nil
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

// searchQuery returns the query string for searching feedback with given data
func searchQuery(organizationID string, data *SearchFeedbackData, count bool) (string, []interface{}) {
	var query string
	if count {
		query = "SELECT COUNT(*) FROM private.feedback WHERE organization_id = $1"
	} else {
		query = "SELECT * FROM private.feedback WHERE organization_id = $1"
	}
	argv := []interface{}{organizationID}
	argc := 1

	next := func() int {
		argc++
		return argc
	}

	if !data.BegTimestamp.IsZero() {
		query += " AND feedback_time >= $" + fmt.Sprint(next()) + "::timestamp"
		argv = append(argv, data.BegTimestamp)
	}
	if !data.EndTimestamp.IsZero() {
		query += " AND feedback_time <= $" + fmt.Sprint(next()) + "::timestamp"
		argv = append(argv, data.EndTimestamp)
	}
	if len(data.FeedbackType) > 0 {
		query += " AND feedback_type IN ("
		for i, ft := range data.FeedbackType {
			query += "$" + fmt.Sprint(next())
			argv = append(argv, ft)
			if i < len(data.FeedbackType)-1 {
				query += ", "
			}
		}
		query += ")"
	}
	for k, v := range data.FeedbackTags {
		query += " AND feedback_tags->>$" + fmt.Sprint(next()) + " = $" + fmt.Sprint(next())
		argv = append(argv, k, v)
	}
	for k, v := range data.FeedbackMeta {
		query += " AND feedback_meta->>$" + fmt.Sprint(next()) + " = $" + fmt.Sprint(next())
		argv = append(argv, k, v)
	}
	for k, v := range data.FeedbackUser {
		query += " AND feedback_user->>$" + fmt.Sprint(next()) + " = $" + fmt.Sprint(next())
		argv = append(argv, k, v)
	}
	if !count {
		query += " ORDER BY created_at DESC"
		if data.PageLimit > 0 {
			query += " LIMIT $" + fmt.Sprint(next())
			argv = append(argv, data.PageLimit)
		}
		if data.PageOffset > 0 {
			query += " OFFSET $" + fmt.Sprint(next())
			argv = append(argv, data.PageOffset)
		}
	}
	return query, argv
}
