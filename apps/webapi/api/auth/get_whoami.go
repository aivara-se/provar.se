package auth

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/user"
)

// GetLoginDetailsResponseBody is the response body for getting login details.
type GetLoginDetailsResponseBody struct {
	Type       permission.PrincipalType `json:"type" validate:"required"`
	User       *user.User               `json:"user,omitempty"`
	Credential *credential.Credential   `json:"credential,omitempty"`
}

func SetupGetLoginDetails(app *fiber.App) {
	path := "/auth/whoami"

	app.Get(path, access.AuthenticatedGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		if principal.Type == permission.PrincipalTypeUser {
			return c.JSON(GetLoginDetailsResponseBody{
				Type: principal.Type,
				User: principal.User,
			})
		}
		if principal.Type == permission.PrincipalTypeCredential {
			return c.JSON(GetLoginDetailsResponseBody{
				Type:       principal.Type,
				Credential: principal.Cred,
			})
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	})
}
