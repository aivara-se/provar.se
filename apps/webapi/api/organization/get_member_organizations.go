package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
)

func SetupMemberOrganizations(app *fiber.App) {
	app.Get("/organization/list", access.AuthenticatedGuard())
	app.Get("/organization/list", func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		if principal.Type != permission.PrincipalTypeUser {
			// Note: only users can list organizations
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		orgs, err := organization.FindMemberOrganizations(principal.User.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(orgs)
	})
}
