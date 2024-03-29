package organization_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testutils"
	"provar.se/webapi/lib/user"
)

func TestOrganizationMembers(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/member/list", nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, []*user.User{})
		if len(responseBody) != 1 {
			t.Fatalf("expected the list to have the creator: %v", responseBody)
		}
		if responseBody[0].ID != usr.ID {
			t.Fatalf("unexpected organization member: %v", responseBody)
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		o1 := testutils.CreateOrganization(u1.ID)
		req := httptest.NewRequest("GET", "/organization/"+o1.ID+"/member/list", nil)
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/member/list", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/member/list", nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
