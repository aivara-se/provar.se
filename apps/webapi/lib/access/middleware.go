package access

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
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
		user, err := ValidateAccessToken(key)
		if err == nil {
			c.Locals(principalKey, &Principal{
				Type: PrincipalTypeUser,
				User: user,
			})
			return c.Next()
		}
		cred, err := credential.FindBySecret(key)
		if err == nil {
			c.Locals(principalKey, &Principal{
				Type: PrincipalTypeCredential,
				Cred: cred,
			})
			return c.Next()
		}
		return c.SendStatus(fiber.StatusForbidden)
	}
}

// GetPrincipal returns the access token from the fiber context
func GetPrincipal(c *fiber.Ctx) interface{} {
	return c.Locals(principalKey)
}