package creds

import (
	"provar.se/webapi/lib/database/repository"
)

// ValidateClientCredentials validates the client credentials
func ValidateClientCredentials(organizationID string, clientID string, clientSecret string) bool {
	repo := repository.GetCredentialRepository()
	credentials, err := repo.FindCredentialByClientID(organizationID, clientID)
	if err != nil {
		return false
	}
	if credentials == nil || credentials.ClientSecret != clientSecret {
		return false
	}
	return true
}
