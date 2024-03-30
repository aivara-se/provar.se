package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/router"
)

func SetupDeleteInvitation(app *fiber.App) {
	path := "/organization/:organizationId/invitation/:invitationId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, access.PermissionGuard(&access.PermissionGuardOptions{
		OrganizationID: router.FromPathParam("organizationId"),
		ResourceType:   permission.ResourceTypeOrganization,
		ResourceID:     router.FromPathParam("organizationId"),
		Permission:     permission.PermissionTypeOrganizationAdmin,
	}))

	app.Delete(path, func(c *fiber.Ctx) error {
		inviteID := c.Params("invitationId")
		err := invitation.DeleteByID(inviteID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
