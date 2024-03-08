package config

import (
	"net/url"
	"os"
)

type AuthConfig struct {
	Secret       string
	GithubID     string
	GithubSecret string
	GoogleID     string
	GoogleSecret string
}

type EmailConfig struct {
	EmailFrom   string
	EmailServer *url.URL
}

// Config holds configuration for the application
type Config struct {
	DatabaseURI string
	Geolite2    string
	HostPort    string
	Auth        AuthConfig
	Email       EmailConfig
}

// FromEnv loads configuration from environment variables
func FromEnv() Config {
	emailServer, err := url.Parse(os.Getenv("EMAIL_SERVER"))
	if err != nil {
		panic(err)
	}
	return Config{
		DatabaseURI: os.Getenv("DATABASE_URI"),
		Geolite2:    os.Getenv("GEOLITE2_DB"),
		HostPort:    ":" + os.Getenv("PORT"),
		Auth: AuthConfig{
			Secret:       os.Getenv("AUTH_SECRET"),
			GithubID:     os.Getenv("AUTH_GITHUB_ID"),
			GithubSecret: os.Getenv("AUTH_GITHUB_SECRET"),
			GoogleID:     os.Getenv("AUTH_GOOGLE_ID"),
			GoogleSecret: os.Getenv("AUTH_GOOGLE_SECRET"),
		},
		Email: EmailConfig{
			EmailFrom:   os.Getenv("EMAIL_FROM"),
			EmailServer: emailServer,
		},
	}
}
