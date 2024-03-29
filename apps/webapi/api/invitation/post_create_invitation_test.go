package invitation_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/testutils"
)

func TestCreateInvitation(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		u1, k1 := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		requestBody := strings.NewReader(`{
			"name": "Test.` + u2.ID + `",
			"email": "test-` + u2.Email + `"
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+k1)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		invites, err := invitation.FindPendingByOrganizationID(org.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(invites) != 1 {
			t.Fatalf("expected one invite, got: %v", invites)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		requestBody := strings.NewReader(`{
			"name": "Test.` + u2.ID + `",
			"email": "test-` + u2.Email + `"
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		u1, token := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		requestBody := strings.NewReader(`not a json`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing name", func(t *testing.T) {
		u1, token := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		requestBody := strings.NewReader(`{ "email": "test-` + u2.Email + `" }`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing email", func(t *testing.T) {
		u1, token := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		requestBody := strings.NewReader(`{ "name": "test-` + u2.Name + `" }`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
