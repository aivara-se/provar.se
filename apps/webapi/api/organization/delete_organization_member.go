package organization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupRemoveOrganizationMember(app *fiber.App) {
	path := "/organization/:organizationId/member/:userId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Delete(path, access.MembershipGuard())

	app.Delete(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		organizationID := c.Params("organizationId")
		userID := c.Params("userId")
		err := organization.RemoveMember(organizationID, userID)
		if err != nil {
			logger.Error("Failed to remove organization member", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to remove organization member")
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
