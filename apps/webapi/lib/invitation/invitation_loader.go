package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

var (
	// invitationKey is the key used to store the loaded invitation
	invitationKey = "app:invitation"
)

// Loader loads an invitation from the database and attaches it to the context
// Note: Organization loader middleware must be used before using this loader
func Loader(getID router.ParamFetcher) fiber.Handler {
	return func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		invite, err := FindByID(getID(c), org.ID)
		if err != nil {
			return NewNotFoundError(err)
		}
		c.Locals(invitationKey, invite)
		return c.Next()
	}
}

// GetInvitation returns the loaded invitation from the context
func GetInvitation(c *fiber.Ctx) *Invitation {
	inv := c.Locals(invitationKey).(*Invitation)
	if inv == nil {
		panic("Invitation is not available in the request context")
	}
	return inv
}
