package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupOrganizationCredentials(app *fiber.App) {
	app.Get("/organization/:id/credential/list", access.AuthenticatedGuard())
	app.Get("/organization/:id/credential/list", organization.Loader(router.FromPathParam("id")))
	app.Get("/organization/:id/credential/list", access.MembershipGuard())

	app.Get("/organization/:id/credential/list", func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		creds, err := credential.FindByOrganizationID(org.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(creds)
	})
}
