package organization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
)

func SetupMemberOrganizations(app *fiber.App) {
	path := "/organizations"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, access.OnlyAllowUsersGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		principal := access.GetPrincipal(c)
		orgs, err := organization.FindMemberOrganizations(principal.User.ID)
		if err != nil {
			logger.Error("Failed to get user organizations", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get user organizations")
		}
		return c.JSON(orgs)
	})
}
