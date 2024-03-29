package invitation_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/testutils"
)

func TestAcceptInvitation(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		u2, k2 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		inv := testutils.CreateInvitationForUser(org.ID, u2)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation/"+inv.Secret+"/accept", nil)
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		isMember, err := org.IsMember(u2.ID)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if !isMember {
			t.Fatalf("expected user to be a member")
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		_, k3 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		inv := testutils.CreateInvitationForUser(org.ID, u2)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation/"+inv.Secret+"/accept", nil)
		req.Header.Add("Authorization", "bearer "+k3)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		inv := testutils.CreateInvitationForUser(org.ID, u2)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation/"+inv.Secret+"/accept", nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		u2, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		inv := testutils.CreateInvitationForUser(org.ID, u2)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/invitation/"+inv.Secret+"/accept", nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
