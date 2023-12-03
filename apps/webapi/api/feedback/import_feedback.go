package feedback

import (
	"encoding/csv"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/validator"
)

// ImportFeedbackRequestBody is the request body for importing feedback events
type ImportFeedbackRequestBody struct {
	ProjectID string `json:"projectId"`
	Link      string `json:"link" validate:"required"`
}

// NewImportFeedbackRequestBody returns a new ImportFeedbackRequestBody
func NewImportFeedbackRequestBody() interface{} {
	return new(ImportFeedbackRequestBody)
}

// @Router      /feedback/import  [post]
// @Summary     Imports feedback from an uploaded CSV file.
// @Description Imports feedback from an uploaded CSV file.
// @Tags        feedback
// @Accept      json
// @Param       body  body  ImportFeedbackRequestBody  true  "Request body"
// @Success     204  "ok"
func SetupImportFeedback(app *fiber.App) {
	app.Post("/feedback/import", credential.GetMiddleware())
	app.Post("/feedback/import", validator.GetMiddleware(NewImportFeedbackRequestBody))

	app.Post("/feedback/import", func(c *fiber.Ctx) error {
		organizationID := credential.GetOrganizationID(c)
		body := validator.GetRequestBody(c).(*ImportFeedbackRequestBody)
		repo := feedback.GetFeedbackRepository()
		response, err := http.Get(body.Link)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		reader := csv.NewReader(response.Body)
		var data []feedback.CreateFeedbackData
		// Read the header row to determine column order
		header, err := reader.Read()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		// Create a map to store column indices based on header names
		columnMap := make(map[string]int)
		for i, columnName := range header {
			columnMap[columnName] = i
		}
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			// Convert CNPS and CSAT fields from string to float64
			cnpsVal := 0.0
			cnpsStr := record[columnMap["cnps"]]
			if cnpsStr != "" {
				cnpsVal, err = strconv.ParseFloat(record[columnMap["cnps"]], 64)
				if err != nil {
					// Handle conversion error
					return c.SendStatus(fiber.StatusInternalServerError)
				}
			}
			csatVal := 0.0
			csatStr := record[columnMap["csat"]]
			if csatStr != "" {
				csatVal, err = strconv.ParseFloat(record[columnMap["csat"]], 64)
				if err != nil {
					// Handle conversion error
					return c.SendStatus(fiber.StatusInternalServerError)
				}
			}
			createdTime := time.Now()
			createdTimeStr := record[columnMap["time"]]
			if createdTimeStr != "" {
				createdTime, err = time.Parse(time.RFC3339, record[columnMap["time"]])
				if err != nil {
					return c.SendStatus(fiber.StatusBadRequest)
				}
			}
			feedbackData := feedback.FeedbackData{
				Text: record[columnMap["text"]],
				CNPS: cnpsVal,
				CSAT: csatVal,
			}
			// Map fields dynamically based on column order
			feedback := feedback.CreateFeedbackData{
				OrganizationID: organizationID,
				ProjectID:      body.ProjectID,
				CreatedAt:      createdTime,
				Type:           feedback.FeedbackType(record[columnMap["type"]]),
				Data:           feedbackData,
			}
			data = append(data, feedback)
		}
		if err := repo.ImportFeedback(data); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
