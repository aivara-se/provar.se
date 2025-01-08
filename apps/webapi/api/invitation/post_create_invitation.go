package invitation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
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
	app.Post(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Post(path, access.MembershipGuard())
	app.Post(path, validator.ValidateMiddleware(CreateCreateInvitationRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		organizationID := c.Params("organizationId")
		principal := access.GetPrincipal(c)
		body := validator.GetRequestBody(c).(*CreateInvitationRequestBody)
		err := invitation.Invite(organizationID, body.Name, body.Email, principal.ID())
		if err != nil {
			logger.Error("Failed to create invitation", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create invitation")
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
