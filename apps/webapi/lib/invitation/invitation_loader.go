package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/router"
)

var (
	// invitationKey is the key used to store the loaded invitation
	invitationKey = "app:invitation"
)

// Loader loads an invitation from the database and attaches it to the context
func Loader(getID router.ParamFetcher) fiber.Handler {
	return func(c *fiber.Ctx) error {
		invite, err := FindByID(getID(c))
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		c.Locals(invitationKey, invite)
		return c.Next()
	}
}

// GetInvitation returns the loaded invitation from the context
func GetInvitation(c *fiber.Ctx) *Invitation {
	return c.Locals(invitationKey).(*Invitation)
}
