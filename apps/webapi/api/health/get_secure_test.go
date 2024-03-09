package health_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testapp"
)

func TestSecureHealthcheck(t *testing.T) {
	app := testapp.Create()

	t.Run("success - user", func(t *testing.T) {
		_, key := testapp.CreateUser()
		req := httptest.NewRequest("GET", "/health/secure", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health/secure", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health/secure", nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with deleted user", func(t *testing.T) {
		usr, key := testapp.CreateUser()
		usr.Delete()
		req := httptest.NewRequest("GET", "/health/secure", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fail()
		}
	})
}
