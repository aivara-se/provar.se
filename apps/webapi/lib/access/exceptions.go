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
