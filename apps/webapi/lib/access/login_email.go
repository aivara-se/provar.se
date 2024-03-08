package access

import (
	"provar.se/webapi/lib/magiclink"
	"provar.se/webapi/lib/user"
)

func LoginWithEmail(email string) {
	user, err := user.FindByEmail(email)
	if err != nil {
		return
	}
	magicLink, err := magiclink.Create(user.ID)
	if err != nil {
		return
	}
	// TODO: send a magic link to the user's email
}

func ConfirmLoginWithEmail(token string) error {
	// TODO: check whether the token exists in the database
	// TODO: create a new access token (jwt) for the user
	return nil
}
