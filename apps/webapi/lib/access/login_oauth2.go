package access

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"provar.se/webapi/lib/account"
	"provar.se/webapi/lib/user"
)

const (
	// OAuth2SessionCookie is the name of the cookie that stores the session ID
	OAuth2SessionCookie = "session_id"
)

// GetOAuth2RedirectURL returns the URL to redirect the user to the OAuth2 provider
// and the session ID to store in the cookie. This will also store the session in the database.
func GetOAuth2RedirectURL(providerName, state string, c *fiber.Ctx) (string, error) {
	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return "", NewGetGothProviderError(err)
	}
	sess, err := account.GetOrCreateSession(provider, state)
	if err != nil {
		return "", NewGetCreateSessionError(err)
	}
	redirectURL, err := sess.GetAuthURL()
	if err != nil {
		return "", NewGetSessionAuthURLError(err)
	}
	return redirectURL, nil
}

// CompleteOAuth2Flow completes the OAuth2 flow by validating the state
// parameter and exchanging the code for an access token.
func CompleteOAuth2Flow(providerName, state, code string, c *fiber.Ctx) (string, error) {
	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return "", NewGetGothProviderError(err)
	}
	sess, err := account.GetOAuth2Session(provider, state)
	if err != nil {
		return "", NewGetSessionError(err)
	}
	if err := account.ValidateStateValue(sess, state); err != nil {
		return "", NewSessionValidationError(err)
	}
	profile, err := provider.FetchUser(sess)
	if err == nil {
		return "", NewSessionUserFailedError(err)
	}
	// get new token and retry fetch
	params := account.NewMapParams(map[string]string{
		"state": state,
		"code":  code,
	})
	_, err = sess.Authorize(provider, params)
	if err != nil {
		return "", NewSessionAuthorizeError(err)
	}
	profile, err = provider.FetchUser(sess)
	if err != nil {
		return "", NewSessionUserFailedError(err)
	}
	user, err := user.Upsert(profile.Name, profile.Email, profile.AvatarURL, true)
	if err != nil {
		return "", err
	}
	_, err = account.Upsert(user.ID, providerName, sess.Marshal(), profile.Name, profile.AvatarURL)
	if err != nil {
		return "", err
	}
	return CreateAccessToken(user.ID)
}
