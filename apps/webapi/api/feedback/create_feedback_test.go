package feedback_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testapp"
	"provar.se/webapi/lib/token"
)

func TestCreateFeedback(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		validToken, _ := token.CreateAccessToken("test-organization")
		req := httptest.NewRequest("POST", "/feedback", nil)
		req.Header.Add("authorization", "Bearer "+validToken)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", nil)
		req.Header.Add("authorization", "Bearer invalid-token")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})
}
