package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupOrganizationDetails(app *fiber.App) {
	app.Get("/organization/:id/details", access.AuthenticatedGuard())
	app.Get("/organization/:id/details", organization.Loader(router.FromPathParam("id")))
	app.Get("/organization/:id/details", access.MembershipGuard())

	app.Get("/organization/:id/details", func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		return c.JSON(org)
	})
}
