package credential

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/router"
)

func SetupDeleteCredential(app *fiber.App) {
	path := "/organization/:organizationId/credential/:credentialId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, access.PermissionGuard(&access.PermissionGuardOptions{
		OrganizationID: router.FromPathParam("organizationId"),
		ResourceType:   permission.ResourceTypeOrganization,
		ResourceID:     router.FromPathParam("organizationId"),
		Permission:     permission.PermissionTypeOrganizationAdmin,
	}))

	app.Delete(path, func(c *fiber.Ctx) error {
		credID := c.Params("credentialId")
		err := credential.DeleteByID(credID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.StatusNoContent)
	})
}
