package organization_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/testutils"
)

func TestCreateOrganization(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, token := testutils.CreateUser()
		requestBody := strings.NewReader(`{
			"name": "Test.` + usr.ID + `",
			"slug": "test-` + usr.ID + `"
		}`)
		req := httptest.NewRequest("POST", "/organization", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		responseBody := testutils.ReadJSON(res.Body, &organization.PublicOrganization{})
		if responseBody.Name != "Test."+usr.ID || responseBody.Slug != "test-"+usr.ID {
			t.Fatalf("unexpected response body: %+v", responseBody)
		}
		storedOrg, err := organization.FindByID(responseBody.ID)
		if err != nil || storedOrg.ID != responseBody.ID {
			t.Fatalf("expected organization to be in the database")
		}
		if storedOrg.Name != responseBody.Name || storedOrg.Slug != responseBody.Slug {
			t.Fatalf("unexpected organization details in the database: %+v", storedOrg)
		}
		if storedOrg.CreatedBy != usr.ID {
			t.Fatalf("expected organization to be created by test user")
		}
		members, err := storedOrg.Members()
		if err != nil || len(members) != 1 || members[0].ID != usr.ID {
			t.Fatalf("expected test user to be a member of the organization")
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		requestBody := strings.NewReader(`{
			"name": "Test.` + usr.ID + `",
			"slug": "test-` + usr.ID + `"
		}`)
		req := httptest.NewRequest("POST", "/organization", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		_, token := testutils.CreateUser()
		requestBody := strings.NewReader(`not a json`)
		req := httptest.NewRequest("POST", "/organization", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing name", func(t *testing.T) {
		_, token := testutils.CreateUser()
		requestBody := strings.NewReader(`{"slug": "test"}`)
		req := httptest.NewRequest("POST", "/organization", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing slug", func(t *testing.T) {
		_, token := testutils.CreateUser()
		requestBody := strings.NewReader(`{"name": "Test"}`)
		req := httptest.NewRequest("POST", "/organization", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+token)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
