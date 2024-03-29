package api

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"provar.se/webapi/api/auth"
	"provar.se/webapi/api/credential"
	"provar.se/webapi/api/invitation"
	"provar.se/webapi/api/organization"
	"provar.se/webapi/api/ping"
)

func Create() *fiber.App {
	// Create a new fiber application
	app := fiber.New(fiber.Config{
		AppName:            "Provar API",
		ReadTimeout:        time.Second,
		WriteTimeout:       time.Second,
		WriteBufferSize:    2048,
		BodyLimit:          10 * 1024,
		JSONEncoder:        json.Marshal,
		JSONDecoder:        json.Unmarshal,
		EnableIPValidation: true,
		ProxyHeader:        "X-Forwarded-For",
	})

	// Use rate limiter middleware
	app.Use(limiter.New(limiter.Config{
		Expiration: time.Second,
		Max:        100,
	}))

	// Use cors middleware
	app.Use(cors.New())

	// Use logger middleware
	app.Use(logger.New())

	// Load all app routes
	auth.SetupGetLoginDetails(app)
	auth.SetupLoginWithEmail(app)
	auth.SetupLoginWithEmailConfirm(app)
	credential.SetupCreateCredential(app)
	credential.SetupOrganizationCredentials(app)
	invitation.SetupInvitations(app)
	organization.SetupCreateOrganization(app)
	organization.SetupDeleteOrganization(app)
	organization.SetupMemberOrganizations(app)
	organization.SetupOrganizationDetails(app)
	organization.SetupOrganizationMembers(app)
	organization.SetupOrganizationSettings(app)
	organization.SetupUpdateOrganizationDetails(app)
	ping.SetupBasicHealthcheck(app)
	ping.SetupSecureHealthcheck(app)

	// Serve static files
	app.Static("/", "./public")

	return app
}
