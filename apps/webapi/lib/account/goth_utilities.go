package account

import (
	"errors"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"provar.se/webapi/lib/random"
)

const (
	// OAuth2StateParam is the name of the query parameter that stores the state
	OAuth2StateParam = "state"
)

var (
	// ErrInvalidStateParam is returned when the state parameter is invalid
	ErrInvalidStateParam = errors.New("invalid oauth2 state value")
)

// RequestParams is a struct that contains the fiber context and
// implements the goth.Params interface.
type RequestParams struct {
	*fiber.Ctx
}

// NewRequestParams returns a new RequestParams with the fiber context.
func NewRequestParams(c *fiber.Ctx) goth.Params {
	return &RequestParams{c}
}

// Get returns the value of the parameter from the fiber context.
// If the request is a POST request, it will return the form value.
// If the request is a GET request, it will return the query value.
func (p *RequestParams) Get(key string) string {
	valueFromQuery := p.Query(key)
	valueFromForm := p.FormValue(key)
	if valueFromQuery != "" {
		return valueFromQuery
	}
	if valueFromForm != "" {
		return valueFromForm
	}
	return ""
}

// MapParams is a struct that contains a map of data and
// implements the goth.Params interface.
type MapParams struct {
	data map[string]string
}

// NewMapParams returns a new MapParams with a map of data.
func NewMapParams(data map[string]string) goth.Params {
	return &MapParams{data}
}

// Get returns the value of the parameter from given map data.
func (p *MapParams) Get(key string) string {
	val, ok := p.data[key]
	if !ok {
		return ""
	}
	return val
}

// CreateNewStateValue returns the state value to be used with the auth flow.
func CreateNewStateValue() string {
	return random.String(32)
}

// ValidateStateValue validates the state parameter received from
// the OAuth2 provider. This will return an error if it is expired.
func ValidateStateValue(sess goth.Session, state string) error {
	authURLString, err := sess.GetAuthURL()
	if err != nil {
		return err
	}
	authURL, err := url.Parse(authURLString)
	if err != nil {
		return err
	}
	originalState := authURL.Query().Get(OAuth2StateParam)
	if originalState == "" {
		return ErrInvalidStateParam
	}
	if originalState != state {
		return ErrInvalidStateParam
	}
	return nil
}

// GetOrCreateSession returns the session from the database if it exists. If it
// does not exist, it will create a new session and return it along with the
// session ID to store in the cookie.
func GetOrCreateSession(provider goth.Provider, state string) (goth.Session, error) {
	sess, err := getOAuth2Session(provider, state)
	if err != nil {
		return nil, err
	}
	if sess != nil {
		return sess, nil
	}
	return newOAuth2Session(provider, state)
}

// getOAuth2Session returns the session from the database if it exists. If it does not
// exist, it will return nil.
func getOAuth2Session(provider goth.Provider, state string) (goth.Session, error) {
	acc, err := FindByState(state)
	if acc == nil {
		return nil, nil
	}
	sess, err := provider.UnmarshalSession(acc.Session)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// newOAuth2Session creates a new session and stores it in the database. It returns
// the session and the session ID to store in the cookie.
func newOAuth2Session(provider goth.Provider, state string) (goth.Session, error) {
	sess, err := provider.BeginAuth(state)
	if err != nil {
		return nil, err
	}
	acc, err := Create(provider.Name(), sess.Marshal())
	if err != nil {
		return nil, err
	}
	_, err = CreateState(acc.ID, state)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// FetchUser fetches the user from the provider. If the user is not found, it will
// retry the fetch after authorizing the session.
func FetchUser(provider goth.Provider, sess goth.Session, c *fiber.Ctx) (*goth.User, error) {
	user, err := provider.FetchUser(sess)
	if err == nil {
		return &user, err
	}
	_, err = sess.Authorize(provider, NewRequestParams(c))
	if err != nil {
		return nil, err
	}
	user, err = provider.FetchUser(sess)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
