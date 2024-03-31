package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupOrganizationPublic(app *fiber.App) {
	path := "/organization/:organizationId/public"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, organization.Loader(router.FromPathParam("organizationId")))

	app.Get(path, func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		return c.JSON(org.Public())
	})
}
