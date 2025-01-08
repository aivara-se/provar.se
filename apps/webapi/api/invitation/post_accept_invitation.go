package invitation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/user"
	"provar.se/webapi/lib/validator"
)

// AcceptInvitationRequestBody is the request body for creating an invitation
type AcceptInvitationRequestBody struct {
	Secret string `json:"secret" validate:"required"`
}

// CreateAcceptInvitationRequestBody returns a new AcceptInvitationRequestBody
func CreateAcceptInvitationRequestBody() interface{} {
	return new(AcceptInvitationRequestBody)
}

func SetupAcceptInvitation(app *fiber.App) {
	path := "/organization/:organizationId/invitation/:invitationId/accept"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, access.OnlyAllowUsersGuard())
	app.Post(path, invitation.Loader(router.FromPathParam("invitationId")))
	app.Post(path, validator.ValidateMiddleware(CreateAcceptInvitationRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		organizationID := c.Params("organizationId")
		principal := access.GetPrincipal(c)
		invite := invitation.GetInvitation(c)
		body := validator.GetRequestBody(c).(*AcceptInvitationRequestBody)
		if invite.OrganizationID != organizationID || invite.Secret != body.Secret {
			return c.SendStatus(fiber.StatusNotFound)
		}
		invitedUser, err := user.FindByEmail(invite.Email)
		if err != nil {
			logger.Info("Failed to find invited user", err)
			return fiber.NewError(fiber.StatusBadRequest, "Failed to find invited user")
		}
		if !invite.IsAcceptable() || principal.User.ID != invitedUser.ID {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err = organization.AddMember(organizationID, invitedUser.ID)
		if err != nil {
			logger.Error("Failed to add member to organization", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to add member to organization")
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
