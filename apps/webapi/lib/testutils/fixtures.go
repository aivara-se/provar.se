package testutils

import (
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/random"
	"provar.se/webapi/lib/user"
)

// CreateUser creates a new test user in the database and returns the user and
// a valid access token for the user.
func CreateUser() (*user.User, string) {
	rnd := random.String(5)
	usr, err := user.Create("Test User "+rnd, "test.user."+rnd+"@provar.se", true)
	if err != nil {
		panic(err)
	}
	token, err := access.CreateAccessToken(usr.ID)
	if err != nil {
		panic(err)
	}
	return usr, token
}
