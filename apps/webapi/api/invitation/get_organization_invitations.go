package invitation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupInvitations(app *fiber.App) {
	path := "/organization/:organizationId/invitations"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Get(path, access.MembershipGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		org := organization.GetOrganization(c)
		invites, err := invitation.FindPendingByOrganizationID(org.ID)
		if err != nil {
			logger.Error("Failed to get organization invitations", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get organization invitations")
		}
		return c.JSON(invites)
	})
}
