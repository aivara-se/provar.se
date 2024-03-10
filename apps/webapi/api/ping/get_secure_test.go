package ping_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testutils"
)

func TestSecureHealthcheck(t *testing.T) {
	app := testutils.Create()

	t.Run("success - user", func(t *testing.T) {
		_, token := testutils.CreateUser()
		req := httptest.NewRequest("GET", "/ping/secure", nil)
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping/secure", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping/secure", nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with deleted user", func(t *testing.T) {
		usr, token := testutils.CreateUser()
		usr.Delete()
		req := httptest.NewRequest("GET", "/ping/secure", nil)
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
