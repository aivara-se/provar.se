package invitation

import (
	"provar.se/webapi/lib/emails"
	"provar.se/webapi/lib/emails/templates"
	"provar.se/webapi/lib/organization"
)

// Invite invites a user to an organization
func Invite(orgID, name, email, createdBy string) error {
	invite, err := Create(orgID, name, email, createdBy)
	if err != nil {
		return err
	}
	org, err := organization.FindByID(orgID)
	if err != nil {
		return err
	}
	tmpl := &templates.Invitation{}
	data := map[string]string{
		"name":             invite.Name,
		"link":             invite.Link(),
		"organizationName": org.Name,
	}
	opts := &emails.EmailOptions{
		RecvAddress: email,
		SubjectData: data,
		ContentData: data,
	}
	err = emails.Sender().Send(tmpl, opts)
	if err != nil {
		return err
	}
	return nil
}
