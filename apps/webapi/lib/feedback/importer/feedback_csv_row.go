package importer

import (
	"errors"
	"strconv"
	"time"

	"provar.se/webapi/lib/feedback"
)

var (
	// ErrInvalidCreatedAt is returned when the time field is invalid.
	ErrInvalidCreatedAt = errors.New("invalid 'time' field in imported CSV file")

	// ErrMissingTypeColumn is returned when the 'type' column is missing.
	ErrMissingTypeColumn = errors.New("missing 'type' column in imported CSV file")

	// ErrInvalidTypeColumn is returned when the 'type' column is invalid.
	ErrInvalidTypeColumn = errors.New("invalid 'type' column in imported CSV file")

	// ErrMissingTextColumn is returned when the 'text' column is missing.
	ErrMissingTextColumn = errors.New("missing 'text' column in imported CSV file")

	// ErrInvalidTextColumn is returned when the 'text' column is invalid.
	ErrInvalidTextColumn = errors.New("invalid 'text' column in imported CSV file")

	// ErrMissingCNPSColumn is returned when the 'cnps' column is missing.
	ErrMissingCNPSColumn = errors.New("missing 'cnps' column in imported CSV file")

	// ErrInvalidCNPSColumn is returned when the 'cnps' column is invalid.
	ErrInvalidCNPSColumn = errors.New("invalid 'cnps' column in imported CSV file")
)

// Parser is used to parse a single row read from the CSV file.
type Parser struct {
	OrganizationID string
	ProjectID      string
	ColumnMap      map[string]int
}

// NewParser is used to create a new Parser instance.
func NewParser(OrganizationID, ProjectID string, ColumnsMap map[string]int) *Parser {
	return &Parser{
		OrganizationID: OrganizationID,
		ProjectID:      ProjectID,
		ColumnMap:      ColumnsMap,
	}
}

// Parse is used to parse a single row from the imported CSV file
func (p *Parser) Parse(Record []string) (val feedback.CreateFeedbackData, err error) {
	val = feedback.CreateFeedbackData{
		OrganizationID: p.OrganizationID,
		ProjectID:      p.ProjectID,
	}
	if val.Type, err = p.parseType(Record); err != nil {
		return val, err
	}
	if val.CreatedAt, err = p.parseTime(Record); err != nil {
		return val, err
	}
	if val.Data.Text, err = p.parseText(Record); err != nil {
		return val, err
	}
	if val.Type == feedback.FeedbackTypeCNPS {
		if val.Data.CNPS, err = p.parseCNPS(Record); err != nil {
			return val, err
		}
	}
	if val.Type == feedback.FeedbackTypeCSAT {
		if val.Data.CSAT, err = p.parseCSAT(Record); err != nil {
			return val, err
		}
	}
	val.User = p.parseUser(Record)
	val.Tags = p.parseTags(Record)
	val.Meta = p.parseMeta(Record)
	return val, err
}

// parseType is used to parse the 'type' field from the CSV file.
func (p *Parser) parseType(Record []string) (feedback.FeedbackType, error) {
	idx, ok := p.ColumnMap["type"]
	if !ok {
		return "", ErrMissingTypeColumn
	}
	str := Record[idx]
	if str == "" {
		return "", ErrInvalidTypeColumn
	}
	val := feedback.FeedbackType(str)
	if val == feedback.FeedbackTypeText || val == feedback.FeedbackTypeCNPS || val == feedback.FeedbackTypeCSAT {
		return val, nil
	}
	return "", ErrInvalidTypeColumn
}

// parseTime is used to parse the 'time' field from the CSV file.
func (p *Parser) parseTime(Record []string) (time.Time, error) {
	now := time.Now()
	idx, ok := p.ColumnMap["time"]
	if !ok {
		return now, nil
	}
	str := Record[idx]
	if str == "" {
		return now, ErrInvalidCreatedAt
	}
	val, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return now, ErrInvalidCreatedAt
	}
	return val, nil
}

// parseText is used to parse the 'text' field from the CSV file.
func (p *Parser) parseText(Record []string) (string, error) {
	idx, ok := p.ColumnMap["text"]
	if !ok {
		return "", ErrMissingTextColumn
	}
	str := Record[idx]
	return str, nil
}

// parseCNPS is used to parse the 'cnps' field from the CSV file.
func (p *Parser) parseCNPS(Record []string) (float64, error) {
	idx, ok := p.ColumnMap["cnps"]
	if !ok {
		return 0, ErrMissingCNPSColumn
	}
	str := Record[idx]
	if str == "" {
		return 0, ErrInvalidCNPSColumn
	}
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, ErrInvalidCNPSColumn
	}
	return val, nil
}

// parseCSAT is used to parse the 'csat' field from the CSV file.
func (p *Parser) parseCSAT(Record []string) (float64, error) {
	idx, ok := p.ColumnMap["csat"]
	if !ok {
		return 0, ErrMissingCNPSColumn
	}
	str := Record[idx]
	if str == "" {
		return 0, ErrInvalidCNPSColumn
	}
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, ErrInvalidCNPSColumn
	}
	return val, nil
}

// parseUser is used to parse the 'user.*' fields from the CSV file.
func (p *Parser) parseUser(Record []string) feedback.FeedbackUser {
	return feedback.FeedbackUser{
		ID:    p.parseOptionalString(Record, "user.id"),
		Name:  p.parseOptionalString(Record, "user.name"),
		Email: p.parseOptionalString(Record, "user.email"),
		Image: p.parseOptionalString(Record, "user.image"),
	}
}

// parseTags is used to parse the 'tags.*' fields from the CSV file.
func (p *Parser) parseTags(Record []string) feedback.FeedbackTags {
	tags := feedback.FeedbackTags{}
	for key, idx := range p.ColumnMap {
		if len(key) >= 5 && key[:5] == "tags." {
			tags[key[5:]] = Record[idx]
		}
	}
	return tags
}

// parseMeta is used to parse the 'meta.*' fields from the CSV file.
func (p *Parser) parseMeta(Record []string) feedback.FeedbackMeta {
	meta := feedback.FeedbackMeta{}
	for key, idx := range p.ColumnMap {
		if len(key) >= 5 && key[:5] == "meta." {
			meta[key[5:]] = Record[idx]
		}
	}
	return meta
}

// parseOptionalString is used to parse a string field that is optional.
func (p *Parser) parseOptionalString(Record []string, Column string) string {
	idx, ok := p.ColumnMap[Column]
	if !ok {
		return ""
	}
	return Record[idx]
}
