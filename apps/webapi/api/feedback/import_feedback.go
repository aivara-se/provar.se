package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/feedback/importer"
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
		imp := importer.NewImporter(organizationID, body.ProjectID, body.Link)
		defer imp.Close()
		if err := imp.Setup(); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := imp.ReadAll(); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if len(imp.ImportedData) == 0 {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := repo.ImportFeedback(imp.ImportedData); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
