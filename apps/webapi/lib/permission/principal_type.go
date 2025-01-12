package permission

import (
	"errors"

	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/user"
)

var (
	// ErrInvalidPrincipalType is returned when the principal type is invalid
	ErrInvalidPrincipalType = errors.New("invalid principal type")
)

// PrincipalType is the type of the principal
type PrincipalType string

// Principal types
const (
	PrincipalTypeCredential PrincipalType = "credential"
	PrincipalTypeUser       PrincipalType = "user"
)

// Principal is the principal of the request. It can be a user or a credential
// that is making the request. This is always set in the context of the request.
type Principal struct {
	Type PrincipalType
	User *user.User
	Cred *credential.Credential
}

// ID returns the ID of the principal
func (p *Principal) ID() string {
	if p.Type == PrincipalTypeUser {
		return p.User.ID
	}
	if p.Type == PrincipalTypeCredential {
		return p.Cred.ID
	}
	return ""
}
