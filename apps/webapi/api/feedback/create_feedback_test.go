package feedback_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/testapp"
)

const MinimalTextFeedback = `{
	"type": "text",
	"data": { "text": "This is a test feedback" }
}`

const MinimalCNPSFeedback = `{
	"type": "cnps",
	"data": { "cnps": 0.8 }
}`

const MinimalCSATFeedback = `{
	"type": "csat",
	"data": { "csat": 0.7 }
}`

const FeedbackWithProjectID = `{
	"projectId": "507f1f77bcf86cd799439011",
	"type": "text",
	"data": { "text": "This is a test feedback" },
	"tags": { "country": "LK" }
}`

func TestCreateFeedback(t *testing.T) {
	app := testapp.Create()

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(FeedbackWithProjectID))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("success - text", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(MinimalTextFeedback))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("success - cnps", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(MinimalCNPSFeedback))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("success - csat", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(MinimalCSATFeedback))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fail()
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(FeedbackWithProjectID))
		req.Header.Add("content-type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(FeedbackWithProjectID))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer invalid-token")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fail()
		}
	})

	t.Run("fail with invalid request body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(`{}`))
		req.Header.Add("content-type", "application/json")
		req.Header.Add("authorization", "bearer test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fail()
		}
	})
}
