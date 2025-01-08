package organization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupDeleteOrganization(app *fiber.App) {
	path := "/organization/:organizationId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Delete(path, access.MembershipGuard())

	app.Delete(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		organizationID := c.Params("organizationId")
		err := organization.DeleteByID(organizationID)
		if err != nil {
			logger.Error("Failed to delete organization", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete organization")
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
