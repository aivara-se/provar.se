package token

import (
	"time"

	jwtw "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	// accessTokenLifetime is the lifetime of an access token
	accessTokenLifetime = 30 * time.Minute
)

var (
	// accessTokenSecret is the secret used to sign access tokens
	accessTokenSecret []byte
)

// SetSigningSecret sets the JWT secret
func SetSigningSecret(secret string) {
	accessTokenSecret = []byte(secret)
}

// GetMiddleware returns a fiber middleware to validate access tokens
func GetMiddleware() fiber.Handler {
	return jwtw.New(jwtw.Config{
		SigningKey: jwtw.SigningKey{Key: accessTokenSecret},
	})
}

// ExtractOrganizationId extracts the organization ID from a JWT token
func ExtractOrganizationId(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)["sub"].(string)
}

// CreateAccessToken creates a JWT token for an organization
func CreateAccessToken(organizationID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": organizationID,
		"exp": time.Now().Add(accessTokenLifetime).Unix(),
	})
	return token.SignedString(accessTokenSecret)
}
