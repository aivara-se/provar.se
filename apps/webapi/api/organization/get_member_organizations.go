package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
)

func SetupMemberOrganizations(app *fiber.App) {
	path := "/organizations"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, access.OnlyAllowUsersGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		orgs, err := organization.FindMemberOrganizations(principal.User.ID)
		if err != nil {
			return err
		}
		return c.JSON(orgs)
	})
}
