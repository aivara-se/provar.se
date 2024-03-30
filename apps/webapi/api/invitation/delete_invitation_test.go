package invitation_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/testutils"
)

func TestDeleteInvitation(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		inv := testutils.CreateInvitation(org.ID, usr.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/invitation/"+inv.ID, nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		_, err := invitation.FindByID(inv.ID)
		if err == nil {
			t.Fatalf("expected invitation to be deleted")
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		inv := testutils.CreateInvitation(org.ID, u1.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/invitation/"+inv.ID, nil)
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		inv := testutils.CreateInvitation(org.ID, usr.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/invitation/"+inv.ID, nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
