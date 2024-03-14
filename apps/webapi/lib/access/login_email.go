package access

import (
	"database/sql"

	"provar.se/webapi/lib/emails"
	"provar.se/webapi/lib/emails/templates"
	"provar.se/webapi/lib/magiclink"
	"provar.se/webapi/lib/user"
)

// PrepareLoginWithEmail sends a magic link to the user's email address.
func PrepareLoginWithEmail(email string) {
	usr, err := user.FindByEmail(email)
	if err == sql.ErrNoRows {
		usr, err = user.Create("", email, false)
	}
	if err != nil {
		return
	}
	magicLink, err := magiclink.Create(usr.ID)
	if err != nil {
		return
	}
	tmpl := &templates.LoginEmailTemplate{}
	data := map[string]string{"link": magicLink.Link()}
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

// ConfirmLoginWithEmail confirms the magic link token and logs the user in.
// It returns an access token if the token is valid. The magic link will be
// invalid after this function is called.
func ConfirmLoginWithEmail(token string) (string, error) {
	magicLink, err := magiclink.Confirm(token)
	if err != nil {
		return "", err
	}
	return CreateAccessToken(magicLink.UserID)
}
