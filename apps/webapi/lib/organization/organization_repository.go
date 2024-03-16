package organization

import (
	"time"

	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/user"
)

// Organization struct represents the organization table in the database
type Organization struct {
	ID         string    `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	CreatedBy  string    `db:"created_by" json:"created_by"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`
	Name       string    `db:"name" json:"name"`
	Slug       string    `db:"slug" json:"slug"`
}

// OrganizationMember struct represents the organizationmember table in the database
type OrganizationMember struct {
	ID             string `db:"id"`
	UserID         string `db:"user_id"`
	OrganizationID string `db:"organization_id"`
}

// Create creates a new organization in the database
func Create(name, slug, createdBy string) (*Organization, error) {
	tx, err := database.DB().Beginx()
	if err != nil {
		return nil, err
	}
	org := &Organization{
		ID:         database.NewID(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		CreatedBy:  createdBy,
		Name:       name,
		Slug:       slug,
	}
	query := `
		INSERT INTO private.organization (id, created_at, modified_at, created_by, name, slug)
		VALUES (:id, :created_at, :modified_at, :created_by, :name, :slug)
	`
	_, err = tx.NamedExec(query, org)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	member := &OrganizationMember{
		ID:             database.NewID(),
		UserID:         createdBy,
		OrganizationID: org.ID,
	}
	query = `
		INSERT INTO private.organizationmember (id, user_id, organization_id)
		VALUES (:id, :user_id, :organization_id)
	`
	_, err = tx.NamedExec(query, member)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return org, nil
}

// FindByID returns an organization with the given id
func FindByID(id string) (*Organization, error) {
	org := &Organization{}
	query := "SELECT * FROM public.organization WHERE id = $1"
	err := database.DB().Get(org, query, id)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// DeleteByID deletes an organization with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.organization WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// FindMemberOrganizations returns all organizations a user is a member of
func FindMemberOrganizations(id string) ([]*Organization, error) {
	orgs := []*Organization{}
	query := `
		SELECT o.*
		FROM public.organizationmember om
		JOIN public.organization o ON om.organization_id = o.id
		WHERE om.user_id = $1
		ORDER BY o.name
	`
	err := database.DB().Select(&orgs, query, id)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

// FindOrganizationMembers returns all members of an organization
func FindOrganizationMembers(id string) ([]*user.User, error) {
	users := []*user.User{}
	query := `
		SELECT u.*
		FROM public.organizationmember om
		JOIN public.user u ON om.user_id = u.id
		WHERE om.organization_id = $1
		ORDER BY u.name
	`
	err := database.DB().Select(&users, query, id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteByID deletes an organization with the given id
func (o *Organization) Delete() error {
	return DeleteByID(o.ID)
}

// Members returns all members of an organization
func (o *Organization) Members() ([]*user.User, error) {
	return FindOrganizationMembers(o.ID)
}
