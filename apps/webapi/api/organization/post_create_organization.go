package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/validator"
)

// CreateOrganizationRequestBody is the request body for creating an organization
type CreateOrganizationRequestBody struct {
	Name string `json:"name" validate:"required"`
	Slug string `json:"slug" validate:"required"`
}

// CreateCreateOrganizationRequestBody returns a new CreateOrganizationRequestBody
func CreateCreateOrganizationRequestBody() interface{} {
	return new(CreateOrganizationRequestBody)
}

func SetupCreateOrganization(app *fiber.App) {
	app.Post("/organization", access.AuthenticatedGuard())
	app.Post("/organization", validator.ValidateMiddleware(CreateCreateOrganizationRequestBody))
	app.Post("/organization", func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		if principal.Type != permission.PrincipalTypeUser {
			// Note: only users can create organizations
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		body := validator.GetRequestBody(c).(*CreateOrganizationRequestBody)
		organization, err := organization.Create(body.Name, body.Slug, principal.User.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(organization.Public())
	})
}
