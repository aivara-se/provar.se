package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/user"
)

func SetupOrganizationMembers(app *fiber.App) {
	app.Get("/organization/:id/member/list", access.AuthenticatedGuard())
	app.Get("/organization/:id/member/list", organization.Loader(router.FromPathParam("id")))
	app.Get("/organization/:id/member/list", access.MembershipGuard())

	app.Get("/organization/:id/member/list", func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		members, err := user.FindByOrganizationID(org.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(members)
	})
}
