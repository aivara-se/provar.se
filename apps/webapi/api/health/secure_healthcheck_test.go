package health_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testapp"
)

func TestSecureHealthcheck(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping/secure", nil)
		req.Header.Add("Authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping/secure", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping/secure", nil)
		req.Header.Add("Authorization", "bearer not-test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})
}
