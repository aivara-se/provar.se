package auth_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/goccy/go-json"
	"provar.se/webapi/api/auth"
	"provar.se/webapi/lib/testapp"
)

func TestCreateAccessToken(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		reqBody := strings.NewReader(`{
			"client_id": "test-client-id",
			"client_secret": "test-client-secret"
		}`)
		req := httptest.NewRequest("POST", "/auth/token", reqBody)
		req.Header.Add("x-organization", "test-organization")
		req.Header.Add("content-type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fail()
		}
		resBody := &auth.CreateAccessTokenResponseBody{}
		json.NewDecoder(res.Body).Decode(resBody)
		if resBody.Token == "" {
			t.Fail()
		}
	})

	t.Run("fail with missing organization header", func(t *testing.T) {
		reqBody := strings.NewReader(`{
			"client_id": "test-client-id",
			"client_secret": "test-client-secret"
		}`)
		req := httptest.NewRequest("POST", "/auth/token", reqBody)
		req.Header.Add("content-type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})

	t.Run("fail with missing request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/auth/token", nil)
		req.Header.Add("x-organization", "test-organization")
		req.Header.Add("content-type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})

	t.Run("fail with invalid request body", func(t *testing.T) {
		reqBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/auth/token", reqBody)
		req.Header.Add("x-organization", "test-organization")
		req.Header.Add("content-type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})

	t.Run("fail with unknown credentials", func(t *testing.T) {
		reqBody := strings.NewReader(`{
			"client_id": "not-test-client-id",
			"client_secret": "not-test-client-secret"
		}`)
		req := httptest.NewRequest("POST", "/auth/token", reqBody)
		req.Header.Add("x-organization", "test-organization")
		req.Header.Add("content-type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})
}
