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

const MinimalRateFeedback = `{
	"type": "rate",
	"data": { "rate": 0.75 }
}`

const FeedbackWithProjectID = `{
	"projectId": "test-project-id",
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

	t.Run("success - rate", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/feedback", strings.NewReader(MinimalRateFeedback))
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
