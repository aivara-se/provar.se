package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupOrganizationInvitations(app *fiber.App) {
	app.Get("/organization/:id/invitation/list", access.AuthenticatedGuard())
	app.Get("/organization/:id/invitation/list", organization.Loader(router.FromPathParam("id")))
	app.Get("/organization/:id/invitation/list", access.MembershipGuard())

	app.Get("/organization/:id/invitation/list", func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		invites, err := invitation.FindPendingByOrganizationID(org.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(invites)
	})
}
