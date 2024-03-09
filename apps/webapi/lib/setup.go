package lib

import (
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/emails"
	"provar.se/webapi/lib/metadata"
)

// Setup initializes shared components
func Setup() error {
	if err := database.Setup(); err != nil {
		return err
	}
	if err := metadata.Setup(); err != nil {
		return err
	}
	if err := emails.Setup(); err != nil {
		return err
	}
	if err := access.Setup(); err != nil {
		return err
	}
	return nil
}
