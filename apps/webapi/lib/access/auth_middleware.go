package access

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/permission"
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

// AuthenticatedGuard returns a fiber middleware to authenticate users
// or credentials (API keys). Returns 401 if the user is not authenticated.
func AuthenticatedGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := extractTokenFromRequest(c)
		if key == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		user, err := ValidateAccessToken(key)
		if err == nil {
			c.Locals(principalKey, &permission.Principal{
				Type: permission.PrincipalTypeUser,
				User: user,
			})
			return c.Next()
		}
		cred, err := credential.FindBySecret(key)
		if err == nil {
			c.Locals(principalKey, &permission.Principal{
				Type: permission.PrincipalTypeCredential,
				Cred: cred,
			})
			return c.Next()
		}
		return c.SendStatus(fiber.StatusForbidden)
	}
}

// OnlyAllowUsersGuard returns a fiber middleware to ensure the request is
// authenticated and the principal is a user. Returns 403 if the principal
// is not a user.
func OnlyAllowUsersGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		principal := GetPrincipal(c)
		if principal.Type != permission.PrincipalTypeUser {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return c.Next()
	}
}

// GetPrincipal returns the access token from the fiber context
func GetPrincipal(c *fiber.Ctx) *permission.Principal {
	return c.Locals(principalKey).(*permission.Principal)
}
