package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/router"
)

func SetupRemoveOrganizationMember(app *fiber.App) {
	path := "/organization/:organizationId/member/:userId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, organization.Loader(router.FromPathParam("organizationId")))

	// Custom permission guard to allow the user to remove themselves
	app.Delete(path, func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		userID := c.Params("userId")
		if principal.ID() == userID {
			return c.Next()
		}
		isAllowed := permission.HasPermission(&permission.PermissionQuery{
			OrganizationID: c.Params("organizationId"),
			PrincipalType:  principal.Type,
			PrincipalID:    principal.ID(),
			ResourceType:   permission.ResourceTypeOrganization,
			ResourceID:     c.Params("organizationId"),
			Permission:     permission.PermissionTypeOrganizationAdmin,
		})
		if !isAllowed {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return c.Next()
	})

	app.Delete(path, func(c *fiber.Ctx) error {
		orgID := c.Params("organizationId")
		userID := c.Params("userId")
		if err := organization.RemoveMember(orgID, userID); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
