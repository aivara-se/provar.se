package access

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/exceptions"
)

// NewMissingAuthTokenError creates a new error for when an authentication token
// is missing from the request.
func NewMissingAuthTokenError() error {
	err := errors.New("missing authentication token")
	return exceptions.NewError(err, fiber.StatusUnauthorized, "Missing authentication token")
}

// NewInvalidCredentialError creates a new error for when a credential is invalid.
func NewInvalidCredentialError(err error) error {
	return exceptions.NewError(err, fiber.StatusForbidden, "Invalid credential")
}

// NewOnlyUserAllowedError creates a new error for when only users are allowed
// to access a resource. Eg: when an API key is used to access a user-only API.
func NewOnlyUserAllowedError() error {
	err := errors.New("expected API to be called with a user principal type")
	return exceptions.NewError(err, fiber.StatusForbidden, "Only users are allowed")
}

// NewInvalidMagicLinkError creates a new error for when a magic link could not
// be confirmed because the magic link is not valid or they are already used once.
func NewInvalidMagicLinkError(err error) error {
	return exceptions.NewError(err, fiber.StatusForbidden, "Unable to confirm magic link")
}

// NewCredNotInOrganizationError creates a new error for when a credential does not
// belong to the organization.
func NewCredNotInOrganizationError() error {
	err := errors.New("credential does not belong to organization")
	return exceptions.NewError(err, fiber.StatusForbidden, "Credential does not belong to organization")
}

// NewUserNotInOrganizationError creates a new error for when a user is not a member
// of the organization.
func NewUserNotInOrganizationError() error {
	err := errors.New("user is not a member of organization")
	return exceptions.NewError(err, fiber.StatusForbidden, "User is not a member of organization")
}

// NewInsufficientPermissionError creates a new error for when a user does not
// have sufficient permissions to access a resource.
func NewInsufficientPermissionError() error {
	err := errors.New("insufficient permissions to access resource")
	return exceptions.NewError(err, fiber.StatusForbidden, "Insufficient permissions")
}

// NewGetGothProviderError creates a new error for when there is an issue
// getting the Goth provider.
func NewGetGothProviderError(err error) error {
	return exceptions.NewError(err, fiber.StatusInternalServerError, "Failed to get provider")
}

// NewGetSessionError creates a new error for when there is an issue
// getting a session.
func NewGetSessionError(err error) error {
	return exceptions.NewError(err, fiber.StatusInternalServerError, "Failed to create session")
}

// NewGetCreateSessionError creates a new error for when there is an issue
// creating a session or getting an already existing session.
func NewGetCreateSessionError(err error) error {
	return exceptions.NewError(err, fiber.StatusInternalServerError, "Failed to create session")
}

// NewGetSessionAuthURLError creates a new error for when there is an issue
// getting the session authentication URL.
func NewGetSessionAuthURLError(err error) error {
	return exceptions.NewError(err, fiber.StatusInternalServerError, "Failed to get auth URL")
}

// NewSessionValidationError creates a new error for when there is an issue
// validating a session.
func NewSessionValidationError(err error) error {
	return exceptions.NewError(err, fiber.StatusInternalServerError, "Failed to validate session")
}

// NewSessionUserFailedError creates a new error for when there is an issue
// with fetching session user details.
func NewSessionUserFailedError(err error) error {
	return exceptions.NewError(err, fiber.StatusInternalServerError, "Session user failed")
}

// NewSessionAuthorizeError creates a new error for when there is an issue
// authorizing a session.
func NewSessionAuthorizeError(err error) error {
	return exceptions.NewError(err, fiber.StatusInternalServerError, "Failed to authorize session")
}
