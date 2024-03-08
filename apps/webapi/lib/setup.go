package lib

import (
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/config"
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/emails"
	"provar.se/webapi/lib/metadata"
)

// Setup initializes shared components
func Setup(cfg config.Config) error {
	// Connect to the database using the database uri
	if err := database.Setup(cfg.DatabaseURI); err != nil {
		return err
	}
	// Connect to geolite2 database using GEOLITE2_DB
	if err := metadata.Setup(cfg.Geolite2); err != nil {
		return err
	}
	// Setup the email server helper for sending emails
	if err := emails.Setup(cfg.Email); err != nil {
		return err
	}
	// Setup the access token helper with the secret key
	if err := access.Setup(cfg.Auth.Secret); err != nil {
		return err
	}
	return nil
}
