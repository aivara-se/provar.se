package credential

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/database/repository"
)

const (
	// organizationIDKey is the key used to store the organization id
	organizationIDKey = "auth:organizationId"

	// authHeaderPrefix is the expected prefix on auth header
	authHeaderPrefix       = "bearer "
	authHeaderPrefixLength = len(authHeaderPrefix)
)

// extractApiKeyFromRequest extracts the api key from the request
func extractApiKeyFromRequest(c *fiber.Ctx) string {
	auth := c.Get(fiber.HeaderAuthorization)
	if !strings.HasPrefix(auth, authHeaderPrefix) {
		return ""
	}
	return auth[authHeaderPrefixLength:]
}

// GetOrganizationID returns the organization id from the request
func GetOrganizationID(c *fiber.Ctx) string {
	return c.Locals(organizationIDKey).(string)
}

// GetMiddleware returns a fiber middleware to validate credentials
func GetMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := extractApiKeyFromRequest(c)
		if key == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		repo := repository.GetCredentialRepository()
		cred, err := repo.FindCredenial(key)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Locals(organizationIDKey, cred.OrganizationID)
		return c.Next()
	}
}
