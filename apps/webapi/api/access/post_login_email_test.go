package access_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/testapp"
)

func TestLoginWithEmail(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/access/login/email", strings.NewReader(`{"email":"test@example.com"}`))
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("fail with missing email", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/access/login/email", strings.NewReader(`{}`))
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})

	t.Run("fail with invalid email", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/access/login/email", strings.NewReader(`{"email":"invalid"}`))
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/access/login/email", strings.NewReader(`not a json`))
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})
}
