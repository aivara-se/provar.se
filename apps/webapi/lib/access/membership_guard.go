package access

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
)

// MembershipGuard checks if the principal is a member of an organization.
// The organization MUST be loaded in the context before using this guard.
func MembershipGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		principal := GetPrincipal(c)
		org := organization.GetOrganization(c)
		if principal.Type == permission.PrincipalTypeCredential {
			if principal.Cred.OrganizationID != org.ID {
				return NewCredNotInOrganizationError()
			}
			return c.Next()
		}
		if principal.Type == permission.PrincipalTypeUser {
			isMember, err := org.IsMember(principal.User.ID)
			if !isMember || err != nil {
				return NewUserNotInOrganizationError()
			}
			return c.Next()
		}
		return permission.ErrInvalidPrincipalType
	}
}
