package ping_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testutils"
)

func TestBasicHealthcheck(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
