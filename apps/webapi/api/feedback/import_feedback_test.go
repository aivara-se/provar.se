package feedback_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/testapp"
)

const MinimalImportRequest = `{
	"link": "http://localhost:9001/importable/valid"
}`

func TestImportFeedback(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback/import", strings.NewReader(MinimalImportRequest))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback/import", strings.NewReader(MinimalImportRequest))
		req.Header.Add("content-type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback/import", strings.NewReader(MinimalImportRequest))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer invalid-token")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback/import", strings.NewReader(`{}`))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})
}
