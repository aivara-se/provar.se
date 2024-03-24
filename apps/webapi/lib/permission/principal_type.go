package permission

import (
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/user"
)

type PrincipalType string

const (
	PrincipalTypeCredential PrincipalType = "credential"
	PrincipalTypeUser       PrincipalType = "user"
)

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
