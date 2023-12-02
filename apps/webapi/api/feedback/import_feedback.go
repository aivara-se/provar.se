package feedback

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/database/repository"
)

// @Router      /feedback/import  [post]
// @Summary     Imports feedback from an uploaded CSV file.
// @Description Imports feedback from an uploaded CSV file.
// @Tags        feedback
// @Accept      json
// @Success     204  "ok"
func SetupImportFeedback(app *fiber.App) {
	app.Post("/feedback/import", credential.GetMiddleware())

	app.Post("/feedback/import", func(c *fiber.Ctx) error {
		organizationID := credential.GetOrganizationID(c)
		repo := repository.GetFeedbackRepository()
		file, err := c.FormFile("csvfile")
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		csvfile, err := file.Open()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		defer csvfile.Close()
		reader := csv.NewReader(csvfile)
		var data []repository.CreateFeedbackData
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
			feedbackData := repository.FeedbackData{
				Text: record[columnMap["text"]],
				CNPS: cnpsVal,
				CSAT: csatVal,
			}
			// Map fields dynamically based on column order
			feedback := repository.CreateFeedbackData{
				OrganizationID: organizationID,
				Type:           repository.FeedbackType(record[columnMap["type"]]),
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
