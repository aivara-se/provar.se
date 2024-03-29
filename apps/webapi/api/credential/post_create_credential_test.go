package credential_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/testutils"
)

func TestCreateCredential(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, token := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"name": "Test.` + usr.ID + `"
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/credential", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, &credential.Credential{})
		if responseBody.Name != "Test."+usr.ID {
			t.Fatalf("unexpected response body: %+v", responseBody)
		}
		storedCred, err := credential.FindBySecret(responseBody.Secret)
		if err != nil || storedCred.ID != responseBody.ID {
			t.Fatalf("expected credential to be in the database")
		}
		if storedCred.CreatedBy != usr.ID {
			t.Fatalf("expected credential to be created by test user")
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"name": "Test.` + usr.ID + `"
		}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/credential", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		usr, token := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`not a json`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/credential", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing name", func(t *testing.T) {
		usr, token := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{}`)
		req := httptest.NewRequest("POST", "/organization/"+org.ID+"/credential", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
