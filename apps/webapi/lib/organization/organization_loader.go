package organization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/router"
)

var (
	// organizationKey is the key used to store the loaded organization
	organizationKey = "app:organization"
)

// Loader loads an organization from the database and attaches it to the context
func Loader(getID router.ParamFetcher) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		org, err := FindByID(getID(c))
		if err != nil {
			logger.Info("Failed to load organization", err)
			return fiber.NewError(fiber.StatusNotFound, "Organization not found")
		}
		c.Locals(organizationKey, org)
		return c.Next()
	}
}

// GetOrganization returns the loaded organization from the context
func GetOrganization(c *fiber.Ctx) *Organization {
	return c.Locals(organizationKey).(*Organization)
}
