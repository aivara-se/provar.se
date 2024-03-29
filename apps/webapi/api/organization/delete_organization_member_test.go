package organization_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/testutils"
	"provar.se/webapi/lib/user"
)

func TestDeleteOrganizationMember(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		u1, k1 := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		err := org.AddMember(u2.ID)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/member/"+u1.ID, nil)
		req.Header.Add("Authorization", "bearer "+k1)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		users, err := user.FindByOrganizationID(org.ID)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(users) != 1 {
			t.Fatalf("expected only one user in organization")
		}
	})

	t.Run("success - last member", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/member/"+usr.ID, nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		_, err := organization.FindByID(org.ID)
		if err == nil {
			t.Fatalf("expected organization to be deleted")
		}
	})

	t.Run("success - unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		u2, k2 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/member/"+u2.ID, nil)
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/member/"+usr.ID, nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
