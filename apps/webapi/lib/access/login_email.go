package access

import (
	"provar.se/webapi/lib/emails"
	"provar.se/webapi/lib/emails/templates"
	"provar.se/webapi/lib/magiclink"
	"provar.se/webapi/lib/user"
)

func PrepareLoginWithEmail(email string) {
	user, err := user.FindByEmail(email)
	if err != nil {
		return
	}
	magicLink, err := magiclink.Create(user.ID)
	if err != nil {
		return
	}
	tmpl := &templates.LoginEmailTemplate{}
	data := map[string]string{"link": magicLink.Token}
	opts := &emails.EmailOptions{
		RecvAddress: email,
		SubjectData: data,
		ContentData: data,
	}
	err = emails.Sender().Send(tmpl, opts)
	if err != nil {
		return
	}
}

func ConfirmLoginWithEmail(token string) error {
	// TODO: check whether the token exists in the database
	// TODO: create a new access token (jwt) for the user
	return nil
}
