package credential

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	// organizationIDKey is the key used to store the organization id
	organizationIDKey = "auth:organizationId"

	// authHeaderPrefix is the expected prefix on auth header
	authHeaderPrefix       = "bearer "
	authHeaderPrefixLength = len(authHeaderPrefix)
)

// hasPrefixCaseInsensitive checks if the string has the prefix, case insensitive
func hasPrefixCaseInsensitive(str string, prefix string) bool {
	prefixLength := len(prefix)
	if len(str) < prefixLength {
		return false
	}
	return strings.EqualFold(str[:prefixLength], prefix)
}

// extractApiKeyFromRequest extracts the api key from the request
func extractApiKeyFromRequest(c *fiber.Ctx) string {
	auth := c.Get(fiber.HeaderAuthorization)
	if !hasPrefixCaseInsensitive(auth, authHeaderPrefix) {
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
		repo := GetCredentialRepository()
		cred, err := repo.FindCredenial(key)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Locals(organizationIDKey, cred.OrganizationID)
		return c.Next()
	}
}
