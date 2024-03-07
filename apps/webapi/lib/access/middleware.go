package access

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authHeaderPrefix       = "bearer "
	authHeaderPrefixLength = len(authHeaderPrefix)
)

var (
	// principalKey is the key used to store the principal info in locals
	principalKey = "app:principal"
)

// hasPrefixCaseInsensitive checks if the string has the prefix, case insensitive
func hasPrefixCaseInsensitive(str string, prefix string) bool {
	prefixLength := len(prefix)
	if len(str) < prefixLength {
		return false
	}
	return strings.EqualFold(str[:prefixLength], prefix)
}

// extractTokenFromRequest extracts the api key from the request
func extractTokenFromRequest(c *fiber.Ctx) string {
	auth := c.Get(fiber.HeaderAuthorization)
	if !hasPrefixCaseInsensitive(auth, authHeaderPrefix) {
		return ""
	}
	return auth[authHeaderPrefixLength:]
}

// Middleware returns a fiber middleware to authenticate users
// or credentials (API keys) created on the application.
func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := extractTokenFromRequest(c)
		if key == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		// TODO: check whether the key is valid JWT, if yes, load user details
		// TODO: check whether the key is valid credential, if yes, load details
		c.Locals(principalKey, "")
		return c.Next()
	}
}

// GetAccessToken returns the access token from the fiber context
func GetAccessToken(c *fiber.Ctx) interface{} {
	return c.Locals(principalKey)
}
