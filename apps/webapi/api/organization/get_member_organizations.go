package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
)

func SetupMemberOrganizations(app *fiber.App) {
	path := "/organization/list"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, access.OnlyUsersGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		orgs, err := organization.FindMemberOrganizations(principal.User.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(orgs)
	})
}
