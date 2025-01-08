package invitation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/router"
)

var (
	// invitationKey is the key used to store the loaded invitation
	invitationKey = "app:invitation"
)

// Loader loads an invitation from the database and attaches it to the context
func Loader(getID router.ParamFetcher) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		invite, err := FindByID(getID(c))
		if err != nil {
			logger.Info("Failed to load invitation", err)
			return fiber.NewError(fiber.StatusNotFound, "Invitation not found")
		}
		c.Locals(invitationKey, invite)
		return c.Next()
	}
}

// GetInvitation returns the loaded invitation from the context
func GetInvitation(c *fiber.Ctx) *Invitation {
	return c.Locals(invitationKey).(*Invitation)
}
