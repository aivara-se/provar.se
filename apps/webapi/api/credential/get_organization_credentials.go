package credential

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupOrganizationCredentials(app *fiber.App) {
	path := "/organization/:organizationId/credentials"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Get(path, access.MembershipGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		creds, err := credential.FindByOrganizationID(org.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(creds)
	})
}
