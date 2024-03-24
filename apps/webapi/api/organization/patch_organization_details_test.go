package organization_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/testutils"
)

func TestUpdateOrganizationDetails(t *testing.T) {
	app := testutils.Create()

	t.Run("success", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"name": "Updated.` + usr.ID + `",
			"slug": "updated-` + usr.ID + `",
			"description": "updated description"
		}`)
		req := httptest.NewRequest("PATCH", "/organization/"+org.ID+"/details", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 204 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
		updatedOrg, err := organization.FindByID(org.ID)
		if err != nil {
			t.Fatalf("expected organization to be in the database")
		}
		if updatedOrg.Name != "Updated."+usr.ID || updatedOrg.Slug != "updated-"+usr.ID || updatedOrg.Description != "updated description" {
			t.Fatalf("unexpected organization details in the database: %+v", updatedOrg)
		}
	})

	t.Run("fail with unknown member", func(t *testing.T) {
		u1, _ := testutils.CreateUser()
		_, k2 := testutils.CreateUser()
		org := testutils.CreateOrganization(u1.ID)
		requestBody := strings.NewReader(`{
			"name": "Updated.` + u1.ID + `",
			"slug": "updated-` + u1.ID + `",
			"description": "updated description"
		}`)
		req := httptest.NewRequest("PATCH", "/organization/"+org.ID+"/details", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+k2)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 403 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing access token", func(t *testing.T) {
		usr, _ := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"name": "Updated.` + usr.ID + `",
			"slug": "updated-` + usr.ID + `",
			"description": "updated description"
		}`)
		req := httptest.NewRequest("PATCH", "/organization/"+org.ID+"/details", requestBody)
		req.Header.Add("Content-Type", "application/json")
		res, _ := app.Test(req, -1)
		if res.StatusCode != 401 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with non-json body", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`not a json`)
		req := httptest.NewRequest("PATCH", "/organization/"+org.ID+"/details", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing name", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"slug": "updated-` + usr.ID + `",
			"description": "updated description"
		}`)
		req := httptest.NewRequest("PATCH", "/organization/"+org.ID+"/details", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})

	t.Run("fail with missing slug", func(t *testing.T) {
		usr, key := testutils.CreateUser()
		org := testutils.CreateOrganization(usr.ID)
		requestBody := strings.NewReader(`{
			"name": "Updated.` + usr.ID + `",
			"description": "updated description"
		}`)
		req := httptest.NewRequest("PATCH", "/organization/"+org.ID+"/details", requestBody)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "bearer "+key)
		res, _ := app.Test(req, -1)
		if res.StatusCode != 400 {
			t.Fatalf("unexpected status code: %d", res.StatusCode)
		}
	})
}
