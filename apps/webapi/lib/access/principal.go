package access

import (
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/user"
)

type PrincipalType int

const (
	Unknown PrincipalType = iota
	Credential
	User
)

type Principal struct {
	Type PrincipalType
	User *user.User
	Cred *credential.Credential
}
