package access

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/router"
)

// PermissionGuardOptions is used to configure the permission guard middleware
type PermissionGuardOptions struct {
	OrganizationID router.ParamFetcher
	ResourceType   permission.ResourceType
	ResourceID     router.ParamFetcher
	Permission     permission.PermissionType
}

// PermissionGuard checks if the principal has the given permission for the given
// resource. If the principal does not have the permission, it returns 403.
func PermissionGuard(opts *PermissionGuardOptions) fiber.Handler {
	return func(c *fiber.Ctx) error {
		principal := GetPrincipal(c)
		isAllowed := permission.HasPermission(&permission.PermissionQuery{
			OrganizationID: opts.OrganizationID(c),
			PrincipalType:  principal.Type,
			PrincipalID:    principal.ID(),
			ResourceType:   opts.ResourceType,
			ResourceID:     opts.ResourceID(c),
			Permission:     opts.Permission,
		})
		if !isAllowed {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}
