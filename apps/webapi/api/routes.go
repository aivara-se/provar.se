package api

import (
	"os"
	"time"

	"log/slog"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	slogfiber "github.com/samber/slog-fiber"
	"provar.se/webapi/api/auth"
	"provar.se/webapi/api/credential"
	"provar.se/webapi/api/feedback"
	"provar.se/webapi/api/invitation"
	"provar.se/webapi/api/organization"
	"provar.se/webapi/api/ping"
	"provar.se/webapi/lib/config"
	"provar.se/webapi/lib/random"
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

	// Use requestid middleware
	app.Use(requestid.New(requestid.Config{
		Generator: random.WithLength(32),
	}))

	// Use cors middleware
	app.Use(cors.New())

	// Use etag middleware
	// app.Use(etag.New())

	// Use logger middleware (text or json)
	if config.Get().LogFormat == "json" {
		app.Use(slogfiber.New(slog.New(slog.NewJSONHandler(os.Stdout, nil))))
	} else {
		app.Use(slogfiber.New(slog.New(slog.NewTextHandler(os.Stdout, nil))))
	}

	// Load all app routes
	auth.SetupGetLoginDetails(app)
	auth.SetupLoginWithEmail(app)
	auth.SetupLoginWithEmailConfirm(app)
	auth.SetupOAuth2Callback(app)
	auth.SetupOAuth2Login(app)
	credential.SetupCreateCredential(app)
	credential.SetupDeleteCredential(app)
	credential.SetupOrganizationCredentials(app)
	feedback.SetupCountFeedback(app)
	feedback.SetupCreateFeedback(app)
	feedback.SetupDeleteInvitation(app)
	feedback.SetupFeedbackDetails(app)
	feedback.SetupSearchFeedback(app)
	invitation.SetupAcceptInvitation(app)
	invitation.SetupCreateInvitation(app)
	invitation.SetupDeleteInvitation(app)
	invitation.SetupInvitationDetails(app)
	invitation.SetupInvitations(app)
	organization.SetupCreateOrganization(app)
	organization.SetupDeleteOrganization(app)
	organization.SetupMemberOrganizations(app)
	organization.SetupOrganizationDetails(app)
	organization.SetupOrganizationMembers(app)
	organization.SetupOrganizationPublic(app)
	organization.SetupOrganizationSettings(app)
	organization.SetupRemoveOrganizationMember(app)
	organization.SetupUpdateOrganizationDetails(app)
	ping.SetupBasicHealthcheck(app)
	ping.SetupSecureHealthcheck(app)

	// Serve static files
	app.Static("/", "./public")

	return app
}
