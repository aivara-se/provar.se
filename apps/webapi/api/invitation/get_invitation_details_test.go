package invitation_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/testutils"
)

func TestInvitationDetails(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		inv := testutils.CreateInvitation(org.ID, usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/invitation/"+inv.Secret, nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, &invitation.Invitation{})
		if responseBody.ID != inv.ID {
			t.Fatalf("unexpected invitation: %v", responseBody)
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		inv := testutils.CreateInvitation(org.ID, u1.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/invitation/"+inv.Secret, nil)
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
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/invitation/"+inv.Secret, nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		inv := testutils.CreateInvitation(org.ID, usr.ID)
		req := httptest.NewRequest("GET", "/organization/"+org.ID+"/invitation/"+inv.Secret, nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}