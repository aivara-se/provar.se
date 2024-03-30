package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/validator"
)

// CreateOrganizationRequestBody is the request body for creating an organization
type CreateOrganizationRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// CreateCreateOrganizationRequestBody returns a new CreateOrganizationRequestBody
func CreateCreateOrganizationRequestBody() interface{} {
	return new(CreateOrganizationRequestBody)
}

func SetupCreateOrganization(app *fiber.App) {
	path := "/organization"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, access.OnlyAllowUsersGuard())
	app.Post(path, validator.ValidateMiddleware(CreateCreateOrganizationRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		body := validator.GetRequestBody(c).(*CreateOrganizationRequestBody)
		organization, err := organization.Create(body.Name, body.Slug, body.Description, principal.User.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(organization)
	})
}
