package feedback_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/testapp"
)

const FeedbackCSV = `type,time,cnps,csat,text
text,2023-01-21T16:34:59.512Z,,,The experience was average. There's room for enhancement.
text,2023-05-16T15:39:02.789Z,,,"The platform is reliable, but it lacks some features."
text,2023-03-06T02:31:53.609Z,,,"The service is decent, but there's room for enhancement."`

func TestImportFeedback(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		body, writer := testapp.FileUpload("csvfile", FeedbackCSV)
		req := httptest.NewRequest("POST", "/feedback/import", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		body, writer := testapp.FileUpload("csvfile", FeedbackCSV)
		req := httptest.NewRequest("POST", "/feedback/import", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		body, writer := testapp.FileUpload("csvfile", FeedbackCSV)
		req := httptest.NewRequest("POST", "/feedback/import", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
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
