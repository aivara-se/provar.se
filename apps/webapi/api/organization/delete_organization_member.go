package organization

import (
	"github.com/gofiber/fiber/v2"
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
		orgID := c.Params("organizationId")
		userID := c.Params("userId")
		if err := organization.RemoveMember(orgID, userID); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
