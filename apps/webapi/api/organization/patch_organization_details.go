package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/validator"
)

// UpdateOrganizationRequestBody is the request body for updating an organization
type UpdateOrganizationRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// CreateUpdateOrganizationRequestBody returns a new UpdateOrganizationRequestBody
func CreateUpdateOrganizationRequestBody() interface{} {
	return new(UpdateOrganizationRequestBody)
}

func SetupUpdateOrganizationDetails(app *fiber.App) {
	path := "/organization/:organizationId"

	app.Patch(path, access.AuthenticatedGuard())
	app.Patch(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Patch(path, access.MembershipGuard())
	app.Patch(path, validator.ValidateMiddleware(CreateUpdateOrganizationRequestBody))

	app.Patch(path, func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		body := validator.GetRequestBody(c).(*UpdateOrganizationRequestBody)
		org.Name = body.Name
		org.Slug = body.Slug
		org.Description = body.Description
		if err := org.Update(); err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
