package testutils

import (
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
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

// CreateOrganization creates a new test organization for the given user and returns
// the organization.
func CreateOrganization(userID string) *organization.Organization {
	rnd := random.String(5)
	org, err := organization.Create("Test Org "+rnd, "test-org-"+rnd, "test-description", userID)
	if err != nil {
		panic(err)
	}
	return org
}

// CreateInvitationForUser creates a new test invitation for the given organization and user
func CreateInvitationForUser(orgID string, usr *user.User) *invitation.Invitation {
	inv, err := invitation.Create(orgID, usr.Name, usr.Email, usr.ID)
	if err != nil {
		panic(err)
	}
	return inv
}

// CreateInvitation creates a new test invitation for the given organization and author
func CreateInvitation(orgID, userID string) *invitation.Invitation {
	rnd := random.String(5)
	inv, err := invitation.Create(orgID, "test-invitation-"+rnd, "test-invitation-"+rnd+"@provar.se", userID)
	if err != nil {
		panic(err)
	}
	return inv
}

// CreateCredential creates a new test credential for the given organization
func CreateCredential(orgID, userID string) *credential.Credential {
	rnd := random.String(5)
	cred, err := credential.Create("Test Credential "+rnd, orgID, userID)
	if err != nil {
		panic(err)
	}
	return cred
}
