package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/router"
)

var (
	// organizationKey is the key used to store the loaded organization
	organizationKey = "app:organization"
)

// Loader loads an organization from the database and attaches it to the context
func Loader(getID router.ParamFetcher) fiber.Handler {
	return func(c *fiber.Ctx) error {
		org, err := FindByID(getID(c))
		if err != nil {
			return NewNotFoundError(err)
		}
		c.Locals(organizationKey, org)
		return c.Next()
	}
}

// GetOrganization returns the loaded organization from the context
func GetOrganization(c *fiber.Ctx) *Organization {
	org := c.Locals(organizationKey).(*Organization)
	if org == nil {
		panic("Organization is not available in the request context")
	}
	return org
}
