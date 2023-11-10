package lib

import (
	"provar.se/webapi/lib/database"
)

// Setup initializes shared components
func Setup(config Config) error {
	// Connect to database using MONGO_URI
	if err := database.Connect(config.MongoURI); err != nil {
		return err
	}
	return nil
}
