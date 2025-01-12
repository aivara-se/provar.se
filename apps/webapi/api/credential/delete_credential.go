package credential

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupDeleteCredential(app *fiber.App) {
	path := "/organization/:organizationId/credential/:credentialId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Delete(path, access.MembershipGuard())

	app.Delete(path, func(c *fiber.Ctx) error {
		organizationID := c.Params("organizationId")
		credentialID := c.Params("credentialId")
		err := credential.DeleteByID(credentialID, organizationID)
		if err != nil {
			return err
		}
		return c.JSON(fiber.StatusNoContent)
	})
}
