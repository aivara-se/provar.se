package lib

import (
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/token"
)

// Setup initializes shared components
func Setup(config Config) error {
	// Connect to database using MONGO_URI
	if err := database.Connect(config.MongoURI); err != nil {
		return err
	}
	// Configure access token signing secret
	token.SetSigningSecret(config.JWTSecret)
	return nil
}
