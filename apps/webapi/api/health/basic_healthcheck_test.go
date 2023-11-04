package health_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testapp"
)

func TestBasicHealthcheck(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/ping", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})
}
