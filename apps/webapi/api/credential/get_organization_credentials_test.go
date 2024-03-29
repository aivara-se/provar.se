package credential_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/testutils"
)

func TestOrganizationCredentials(t *testing.T) {
	app := testutils.Create()

	t.Run("success - empty", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/credential/list", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, []*credential.Credential{})
		if len(responseBody) != 0 {
			t.Fatalf("expected the list to be empty: %v", responseBody)
		}
	})

	t.Run("success - multiple", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		c1 := testutils.CreateCredential(org.ID, usr.ID)
		c2 := testutils.CreateCredential(org.ID, usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/credential/list", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, []*credential.Credential{})
		if len(responseBody) != 2 {
			t.Fatalf("expected the list to have elements: %v", responseBody)
		}
		if (responseBody[0].ID != c1.ID || responseBody[1].ID != c2.ID) && (responseBody[1].ID != c1.ID || responseBody[0].ID != c2.ID) {
			t.Fatalf("unexpected credentials: %v", responseBody)
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		o1 := testutils.CreateOrganization(u1.ID)
		req := httptest.NewRequest("GET", "/organization/"+o1.ID+"/credential/list", nil)
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/credential/list", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/credential/list", nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
