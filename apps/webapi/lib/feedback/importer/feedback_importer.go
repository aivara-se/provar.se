package importer

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"

	"provar.se/webapi/lib/feedback"
)

const (
	// ImportBatchSize is the default number of events to import in a batch.
	ImportBatchSize = 100
)

var (
	// ErrFetchingFile is returned when the CSV file could not be fetched.
	ErrFetchingFile = errors.New("could not fetch the CSV file")

	// ErrInvalidFile is returned when the CSV file is invalid.
	ErrInvalidFile = errors.New("could not read the CSV file")
)

// Importer imports feedback events from a CSV file using the given url.
type Importer struct {
	OrganizationID string
	ProjectID      string
	CSVFileLink    string
	ImportedData   []feedback.CreateFeedbackData
	SuccessCount   int
	FailureCount   int
	RecordParser   *Parser

	responseBody io.ReadCloser
	csvReader    *csv.Reader
}

// NewImporter is used to create a new CSV Importer with given data.
func NewImporter(OrganizationID, ProjectID, CSVFileLink string) *Importer {
	return &Importer{
		OrganizationID: OrganizationID,
		ProjectID:      ProjectID,
		CSVFileLink:    CSVFileLink,
	}
}

// Setup is used to setup the importer and prepare it for importing.
func (i *Importer) Setup() error {
	response, err := http.Get(i.CSVFileLink)
	if err != nil {
		return ErrFetchingFile
	}
	i.responseBody = response.Body
	if response.StatusCode != http.StatusOK {
		return ErrFetchingFile
	}
	i.csvReader = csv.NewReader(response.Body)
	header, err := i.csvReader.Read()
	if err != nil {
		return ErrInvalidFile
	}
	columnMap := make(map[string]int)
	for idx, columnName := range header {
		columnMap[columnName] = idx
	}
	i.RecordParser = NewParser(i.OrganizationID, i.ProjectID, columnMap)
	return nil
}

// ReadAll is used to read all the rows from the CSV file into an array.
func (i *Importer) ReadAll() error {
	for {
		data, err := i.readOne()
		if err == io.EOF {
			break
		}
		if err != nil {
			i.FailureCount++
			continue
		}
		i.ImportedData = append(i.ImportedData, data)
		i.SuccessCount++
	}
	return nil
}

// readOne is used to read the next row from the CSV file and parse fields.
func (i *Importer) readOne() (data feedback.CreateFeedbackData, err error) {
	record, err := i.csvReader.Read()
	if err != nil {
		return data, err
	}
	return i.RecordParser.Parse(record)
}

// Close is used to close the importer and cleanup resources used here.
func (i *Importer) Close() error {
	if i.responseBody != nil {
		return i.responseBody.Close()
	}
	return nil
}
