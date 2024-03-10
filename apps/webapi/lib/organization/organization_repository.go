package organization

import (
	"time"

	"provar.se/webapi/lib/database"
)

// Organization struct represents the organization table in the database
type Organization struct {
	ID         string    `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	CreatedBy  string    `db:"created_by"`
	ModifiedAt time.Time `db:"modified_at"`
	Name       string    `db:"name"`
	Slug       string    `db:"slug"`
}

// PublicOrganization represents an organization with sensitive information
// removed. This is used to send organization information to the client.
type PublicOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Create creates a new organization in the database
func Create(name, slug, createdBy string) (*Organization, error) {
	user := &Organization{
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
	_, err := database.DB().NamedExec(query, user)
	return user, err
}

// FindByID returns an organization with the given id
func FindByID(id string) (*Organization, error) {
	organization := &Organization{}
	query := "SELECT * FROM public.organization WHERE id = $1"
	err := database.DB().Get(organization, query, id)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

// DeleteByID deletes an organization with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.organization WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// Public returns a public organization from an organization
func (u *Organization) Public() *PublicOrganization {
	return &PublicOrganization{
		ID:   u.ID,
		Name: u.Name,
		Slug: u.Slug,
	}
}

// DeleteByID deletes an organization with the given id
func (u *Organization) Delete() error {
	return DeleteByID(u.ID)
}
