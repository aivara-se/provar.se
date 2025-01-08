package credential

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
		logger := log.WithContext(c.Context())
		credentialID := c.Params("credentialId")
		err := credential.DeleteByID(credentialID)
		if err != nil {
			logger.Error("Failed to delete credential", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete credential")
		}
		return c.JSON(fiber.StatusNoContent)
	})
}
