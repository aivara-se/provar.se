package organization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/user"
)

func SetupOrganizationMembers(app *fiber.App) {
	path := "/organization/:organizationId/members"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Get(path, access.MembershipGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		org := organization.GetOrganization(c)
		members, err := user.FindByOrganizationID(org.ID)
		if err != nil {
			logger.Error("Failed to get organization members", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get organization members")
		}
		return c.JSON(members)
	})
}
