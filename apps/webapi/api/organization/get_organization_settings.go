package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupOrganizationSettings(app *fiber.App) {
	app.Get("/organization/:id/settings", access.AuthenticatedGuard())
	app.Get("/organization/:id/settings", organization.Loader(router.FromPathParam("id")))
	app.Get("/organization/:id/settings", access.MembershipGuard())

	app.Get("/organization/:id/settings", func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		settings, err := org.Settings()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(settings)
	})
}
