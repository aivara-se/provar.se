package lib

import (
	"os"
)

// Config holds configuration for the application
type Config struct {
	MongoURI string
	Geolite2 string
	Address  string
}

// ReadConfig loads configuration from environment variables
func ReadConfig() Config {
	// Load configuration from environment
	return Config{
		MongoURI: os.Getenv("MONGO_URI"),
		Geolite2: os.Getenv("GEOLITE2_DB"),
		Address:  ":" + os.Getenv("PORT"),
	}
}
