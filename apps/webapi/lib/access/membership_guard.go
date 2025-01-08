package access

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
)

// MembershipGuard checks if the principal is a member of an organization.
// The organization MUST be loaded in the context before using this guard.
func MembershipGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		principal := GetPrincipal(c)
		org := organization.GetOrganization(c)
		if principal.Type == permission.PrincipalTypeCredential {
			if principal.Cred.OrganizationID != org.ID {
				logger.Info("Credential does not belong to organization")
				return fiber.NewError(fiber.StatusForbidden, "Credential does not belong to organization")
			}
			return c.Next()
		}
		if principal.Type == permission.PrincipalTypeUser {
			isMember, err := org.IsMember(principal.User.ID)
			if !isMember || err != nil {
				logger.Info("User is not a member of organization")
				return fiber.NewError(fiber.StatusForbidden, "User is not a member of organization")
			}
			return c.Next()
		}
		return c.SendStatus(fiber.StatusForbidden)
	}
}
