package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/user"
)

func SetupAcceptInvitation(app *fiber.App) {
	path := "/organization/:organizationId/invitation/:secret/accept"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, access.OnlyUsersGuard())
	app.Post(path, invitation.Loader(router.FromPathParam("secret")))

	app.Post(path, func(c *fiber.Ctx) error {
		orgID := c.Params("organizationId")
		principal := access.GetPrincipal(c)
		invite := invitation.GetInvitation(c)
		if invite.OrganizationID != orgID {
			return c.SendStatus(fiber.StatusNotFound)
		}
		invitedUser, err := user.FindByEmail(invite.Email)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if !invite.IsAcceptable() || principal.User.ID != invitedUser.ID {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err = organization.AddMember(orgID, invitedUser.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
