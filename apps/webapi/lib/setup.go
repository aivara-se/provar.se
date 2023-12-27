package lib

import (
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/location"
)

// Setup initializes shared components
func Setup(config Config) error {
	// Connect to the mongo database using MONGO_URI
	if err := database.Setup(config.MongoURI); err != nil {
		return err
	}
	// Connect to geolite2 database using GEOLITE2_DB
	if err := location.Setup(config.Geolite2); err != nil {
		return err
	}
	return nil
}
