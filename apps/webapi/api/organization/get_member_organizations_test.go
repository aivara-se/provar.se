package organization_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/testutils"
)

func TestMemberOrganizations(t *testing.T) {
	app := testutils.Create()

	t.Run("success - empty", func(t *testing.T) {
		_, key := testutils.CreateUser()
		req := httptest.NewRequest("GET", "/organizations", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, []*organization.Organization{})
		if len(responseBody) != 0 {
			t.Fatalf("expected the list to be empty: %v", responseBody)
		}
	})

	t.Run("success - multiple", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		o1 := testutils.CreateOrganization(usr.ID)
		o2 := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organizations", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, []*organization.Organization{})
		if len(responseBody) != 2 {
			t.Fatalf("expected the list to have 2 orgs: %v", responseBody)
		}
		if (responseBody[0].ID != o1.ID || responseBody[1].ID != o2.ID) && (responseBody[1].ID != o1.ID || responseBody[0].ID != o2.ID) {
			t.Fatalf("unexpected orgs: %v", responseBody)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/organizations", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/organizations", nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
