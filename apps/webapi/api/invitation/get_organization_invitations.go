package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupInvitations(app *fiber.App) {
	path := "/organization/:organizationId/invitation/list"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Get(path, access.MembershipGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		invites, err := invitation.FindPendingByOrganizationID(org.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(invites)
	})
}
