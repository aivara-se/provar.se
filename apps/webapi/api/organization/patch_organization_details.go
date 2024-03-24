package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
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
	app.Patch("/organization/:id/details", access.AuthenticatedGuard())
	app.Patch("/organization/:id/details", validator.ValidateMiddleware(CreateUpdateOrganizationRequestBody))
	app.Patch("/organization/:id/details", organization.Loader(router.FromPathParam("id")))
	app.Patch("/organization/:id/details", access.PermissionGuard(&access.PermissionGuardOptions{
		OrganizationID: router.FromPathParam("id"),
		ResourceType:   permission.ResourceTypeOrganization,
		ResourceID:     router.FromPathParam("id"),
		Permission:     permission.PermissionTypeOrganizationAdmin,
	}))

	app.Patch("/organization/:id/details", func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		body := validator.GetRequestBody(c).(*UpdateOrganizationRequestBody)
		org.Name = body.Name
		org.Slug = body.Slug
		org.Description = body.Description
		if err := org.Update(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
