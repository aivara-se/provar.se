package auth

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestSetupCreateAccessToken(t *testing.T) {
	type args struct {
		app *fiber.App
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupCreateAccessToken(tt.args.app)
		})
	}
}
