package config

import (
	"net/url"
	"os"
)

var (
	cachedConfig *Config
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

type ProvarConfig struct {
	AppURL string
	APIURL string
}

// Config holds configuration for the application
type Config struct {
	DatabaseURI string
	Geolite2    string
	HostPort    string
	Auth        AuthConfig
	Email       EmailConfig
	Provar      ProvarConfig
}

// SetupFromEnv loads configuration from environment variables
func SetupFromEnv() *Config {
	if cachedConfig != nil {
		return cachedConfig
	}
	emailServer, err := url.Parse(os.Getenv("EMAIL_SERVER"))
	if err != nil {
		panic(err)
	}
	cachedConfig = &Config{
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
		Provar: ProvarConfig{
			AppURL: os.Getenv("PROVAR_APP_URL"),
			APIURL: os.Getenv("PROVAR_API_URL"),
		},
	}
	return cachedConfig
}

// Get returns the cached configuration, loads from env if not cached
func Get() *Config {
	if cachedConfig == nil {
		SetupFromEnv()
	}
	return cachedConfig
}
