package organization_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/testutils"
)

func TestOrganizationPublic(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/public", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, &organization.Organization{})
		if responseBody.ID != org.ID {
			t.Fatalf("expected the response to be valid: %v", responseBody)
		}
	})

	t.Run("success - not member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		o1 := testutils.CreateOrganization(u1.ID)
		req := httptest.NewRequest("GET", "/organization/"+o1.ID+"/public", nil)
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/public", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/public", nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
