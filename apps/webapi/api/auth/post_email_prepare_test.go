package auth_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/magiclink"
	"provar.se/webapi/lib/testutils"
)

func TestLoginWithEmail(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		requestBody := strings.NewReader(`{"email":"` + usr.Email + `"}`)
		req := httptest.NewRequest("POST", "/auth/email/prepare", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		emails := testutils.GetSentEmails(usr.Email)
		if len(emails) != 1 {
			t.Fatalf("expected an email to be sent to user")
		}
		storedLinks, err := magiclink.FindByUserID(usr.ID)
		if err != nil || len(storedLinks) != 1 {
			t.Fatalf("expected a magiclink in the database")
		}
		htmlEmailHasLink := strings.Contains(emails[0].HTML, storedLinks[0].Link())
		textEmailHasLink := strings.Contains(emails[0].Text, storedLinks[0].Link())
		if !htmlEmailHasLink || !textEmailHasLink {
			t.Fatalf("expected email to contain magiclink")
		}
	})

	t.Run("success - new user", func(t *testing.T) {
		email := testutils.RandomEmailAddress()
		requestBody := strings.NewReader(`{"email":"` + email + `"}`)
		req := httptest.NewRequest("POST", "/auth/email/prepare", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		emails := testutils.GetSentEmails(email)
		if len(emails) != 1 {
			t.Fatalf("expected an email to be sent to user")
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		requestBody := strings.NewReader(`not a json`)
		req := httptest.NewRequest("POST", "/auth/email/prepare", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing email", func(t *testing.T) {
		requestBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/auth/email/prepare", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid email", func(t *testing.T) {
		requestBody := strings.NewReader(`{"email":"invalid"}`)
		req := httptest.NewRequest("POST", "/auth/email/prepare", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
