package account

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"provar.se/webapi/lib/config"
)

// Setup sets the secret key used to sign the JWT tokens
func Setup() error {
	goth.UseProviders(
		setupGoogleProvider(),
		setupGithubProvider(),
	)
	return nil
}

// setupGoogleProvider sets up the Google provider for Goth
func setupGoogleProvider() *google.Provider {
	cfg := config.Get()
	clientID := cfg.Auth.GoogleID
	clientSecret := cfg.Auth.GoogleSecret
	callbackURL := cfg.Provar.AppURL + "/auth/login/oauth2/google/callback"
	provider := google.New(clientID, clientSecret, callbackURL)
	return provider
}

// setupGithubProvider sets up the Github provider for Goth
func setupGithubProvider() *github.Provider {
	cfg := config.Get()
	clientID := cfg.Auth.GithubID
	clientSecret := cfg.Auth.GithubSecret
	callbackURL := cfg.Provar.AppURL + "/auth/login/oauth2/github/callback"
	provider := github.New(clientID, clientSecret, callbackURL)
	return provider
}
