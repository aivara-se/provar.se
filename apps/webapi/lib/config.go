package lib

import (
	"os"
)

// Config holds configuration for the application
type Config struct {
	MongoURI  string
	JWTSecret string
	Address   string
}

// ReadConfig loads configuration from environment variables
func ReadConfig() Config {
	// Load configuration from environment
	return Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Address:   ":" + os.Getenv("PORT"),
	}
}
