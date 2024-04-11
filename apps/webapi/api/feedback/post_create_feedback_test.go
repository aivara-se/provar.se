package feedback_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/testutils"
)

func TestCreateFeedback(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"feedbacks": [
				{
					"type": "text",
					"time": "2024-04-01T22:03:39.385Z",
					"data": { "question-type": "default", "response-text": "test" }
				}
			]
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedback", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		fbs, err := feedback.Search(org.ID, &feedback.SearchFeedbackData{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		count, err := feedback.Count(org.ID, &feedback.SearchFeedbackData{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if count != 1 || len(fbs) != 1 {
			t.Fatalf("expected one feedback, got: %v", fbs)
		}
	})

	t.Run("success - cnps", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"feedbacks": [
				{
					"type": "cnps",
					"time": "2024-04-01T22:03:39.385Z",
					"data": { "question-type": "rating-11p", "response-data": "7" }
				}
			]
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedback", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		fbs, err := feedback.Search(org.ID, &feedback.SearchFeedbackData{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		count, err := feedback.Count(org.ID, &feedback.SearchFeedbackData{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if count != 1 || len(fbs) != 1 {
			t.Fatalf("expected one feedback, got: %v", fbs)
		}
	})

	t.Run("success - csat", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"feedbacks": [
				{
					"type": "csat",
					"time": "2024-04-01T22:03:39.385Z",
					"data": { "question-type": "rating-11p", "response-data": "7" }
				}
			]
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedback", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		fbs, err := feedback.Search(org.ID, &feedback.SearchFeedbackData{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		count, err := feedback.Count(org.ID, &feedback.SearchFeedbackData{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if count != 1 || len(fbs) != 1 {
			t.Fatalf("expected one feedback, got: %v", fbs)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"feedbacks": [
				{
					"type": "text",
					"time": "2024-04-01",
					"data": { "question-type": "default", "response-text": "test" }
				}
			]
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedback", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`not a json`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedback", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing question type", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
				"feedbacks": [
				{
					"type": "text",
					"time": "2024-04-01",
					"data": { "question-type": "default", "response-text": "test" }
				}
			]
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedback", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing feedback type", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"feedbacks": [
				{
					"time": "2024-04-01T22:03:39.385Z",
					"data": { "question-type": "default", "response-text": "test" }
				}
			]
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedback", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
