package access

import (
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/user"
)

type PrincipalType int

const (
	PrincipalTypeUnknown PrincipalType = iota
	PrincipalTypeCredential
	PrincipalTypeUser
	PrincipalTypeSystem
)

type Principal struct {
	Type PrincipalType
	User *user.User
	Cred *credential.Credential
}
