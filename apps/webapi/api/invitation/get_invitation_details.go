package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupInvitationDetails(app *fiber.App) {
	path := "/organization/:organizationId/invitation/:secret"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Get(path, invitation.Loader(router.FromPathParam("secret")))
	app.Get(path, access.MembershipGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		invite := invitation.GetInvitation(c)
		if invite.OrganizationID != org.ID {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(invite)
	})
}
