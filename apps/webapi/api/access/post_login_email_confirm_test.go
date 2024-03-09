package access_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/magiclink"
	"provar.se/webapi/lib/testapp"
	"provar.se/webapi/lib/user"
)

func PrepareMagicLink() (*user.User, string) {
	usr, _ := testapp.CreateUser()
	access.PrepareLoginWithEmail(usr.Email)
	magicLinks, _ := magiclink.FindByUserID(usr.ID)
	return usr, magicLinks[0].Token
}

func TestLoginWithEmailConfirm(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		usr, token := PrepareMagicLink()
		requestBody := strings.NewReader(`{"token":"` + token + `"}`)
		req := httptest.NewRequest("POST", "/access/login/email/confirm", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		storedLinks, err := magiclink.FindByUserID(usr.ID)
		if err != nil || len(storedLinks) != 0 {
			t.Fatalf("expected stored magic link to be deleted")
		}
		responseBody := testapp.ReadJSON(res.Body, map[string]string{})
		if responseBody["accessToken"] == "" {
			t.Fatalf("expected access token in response")
		}
		validatedUser, err := access.ValidateAccessToken(responseBody["accessToken"])
		if err != nil || validatedUser.ID != usr.ID {
			t.Fatalf("expected access token to be valid")
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		requestBody := strings.NewReader(`not a json`)
		req := httptest.NewRequest("POST", "/access/login/email/confirm", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing token", func(t *testing.T) {
		requestBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/access/login/email/confirm", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with unknown token", func(t *testing.T) {
		requestBody := strings.NewReader(`{"token":"not-a-real-token"}`)
		req := httptest.NewRequest("POST", "/access/login/email/confirm", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
