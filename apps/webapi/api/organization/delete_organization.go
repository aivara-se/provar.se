package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/router"
)

func SetupDeleteOrganization(app *fiber.App) {
	path := "/organization/:organizationId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Delete(path, access.PermissionGuard(&access.PermissionGuardOptions{
		OrganizationID: router.FromPathParam("organizationId"),
		ResourceType:   permission.ResourceTypeOrganization,
		ResourceID:     router.FromPathParam("organizationId"),
		Permission:     permission.PermissionTypeOrganizationAdmin,
	}))

	app.Delete(path, func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		if err := org.Delete(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
