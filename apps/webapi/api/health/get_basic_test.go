package health_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testapp"
)

func TestBasicHealthcheck(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health/basic", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
