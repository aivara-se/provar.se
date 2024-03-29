package credential

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/permission"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/validator"
)

// CreateCredentialRequestBody is the request body for creating an organization
type CreateCredentialRequestBody struct {
	Name string `json:"name" validate:"required"`
}

// CreateCreateCredentialRequestBody returns a new CreateCredentialRequestBody
func CreateCreateCredentialRequestBody() interface{} {
	return new(CreateCredentialRequestBody)
}

func SetupCreateCredential(app *fiber.App) {
	path := "/organization/:organizationId/credential"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, access.OnlyUsersGuard())
	app.Post(path, validator.ValidateMiddleware(CreateCreateCredentialRequestBody))
	app.Post(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Post(path, access.PermissionGuard(&access.PermissionGuardOptions{
		OrganizationID: router.FromPathParam("organizationId"),
		ResourceType:   permission.ResourceTypeOrganization,
		ResourceID:     router.FromPathParam("organizationId"),
		Permission:     permission.PermissionTypeOrganizationAdmin,
	}))

	app.Post(path, func(c *fiber.Ctx) error {
		principal := access.GetPrincipal(c)
		org := organization.GetOrganization(c)
		body := validator.GetRequestBody(c).(*CreateCredentialRequestBody)
		cred, err := credential.Create(body.Name, org.ID, principal.User.ID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(cred)
	})
}
