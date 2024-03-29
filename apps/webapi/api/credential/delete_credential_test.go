package credential_test

import (
	"net/http/httptest"
	"testing"

	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/testutils"
)

func TestDeleteCredential(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		c1 := testutils.CreateCredential(org.ID, usr.ID)
		c2 := testutils.CreateCredential(org.ID, usr.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/credential/"+c1.ID, nil)
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		credentials, err := credential.FindByOrganizationID(org.ID)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(credentials) != 1 {
			t.Fatalf("expected only one credential in organization")
		}
		if credentials[0].ID != c2.ID {
			t.Fatalf("unexpected credential: %v", credentials)
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		c1 := testutils.CreateCredential(org.ID, u1.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/credential/"+c1.ID, nil)
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		c1 := testutils.CreateCredential(org.ID, usr.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/credential/"+c1.ID, nil)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with invalid access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		c1 := testutils.CreateCredential(org.ID, usr.ID)
		req := httptest.NewRequest("DELETE", "/organization/"+org.ID+"/credential/"+c1.ID, nil)
		req.Header.Add("Authorization", "bear test-api-key")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
