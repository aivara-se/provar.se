package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/validator"
)

// CreateInvitationRequestBody is the request body for creating an invitation
type CreateInvitationRequestBody struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

// CreateCreateInvitationRequestBody returns a new CreateInvitationRequestBody
func CreateCreateInvitationRequestBody() interface{} {
	return new(CreateInvitationRequestBody)
}

func SetupCreateInvitation(app *fiber.App) {
	path := "/organization/:organizationId/invitation"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, validator.ValidateMiddleware(CreateCreateInvitationRequestBody))
	app.Post(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Post(path, access.PermissionGuard(&access.PermissionGuardOptions{
		OrganizationID: router.FromPathParam("organizationId"),
		ResourceType:   permission.ResourceTypeOrganization,
		ResourceID:     router.FromPathParam("organizationId"),
		Permission:     permission.PermissionTypeOrganizationAdmin,
	}))

	app.Post(path, func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		if principal.Type != permission.PrincipalTypeUser {
			// Note: only users can create organizations
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		org := organization.GetOrganization(c)
		body := validator.GetRequestBody(c).(*CreateInvitationRequestBody)
		err := invitation.Invite(org.ID, body.Name, body.Email, principal.User.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
