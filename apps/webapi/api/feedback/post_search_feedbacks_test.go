package feedback_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	feedbackAPI "provar.se/webapi/api/feedback"
	"provar.se/webapi/lib/testutils"
)

func TestSearchFeedback(t *testing.T) {
	app := testutils.Create()

	t.Run("success - empty body", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		fb := testutils.CreateFeedback(org.ID)
		requestBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedbacks", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, &feedbackAPI.SearchFeedbackResponseBody{})
		if responseBody.Total != 1 {
			t.Fatalf("unexpected total: %d", responseBody.Total)
		}
		if len(responseBody.Feedbacks) != 1 {
			t.Fatalf("unexpected feedbacks length: %d", len(responseBody.Feedbacks))
		}
		if responseBody.Feedbacks[0].ID != fb.ID {
			t.Fatalf("unexpected feedback data: %s", responseBody.Feedbacks[0].OrganizationID)
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		testutils.CreateFeedback(org.ID)
		requestBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedbacks", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		testutils.CreateFeedback(org.ID)
		requestBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedbacks", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		testutils.CreateFeedback(org.ID)
		requestBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/feedbacks", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
