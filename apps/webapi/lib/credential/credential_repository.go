package credential

import (
	"time"

	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/random"
	"provar.se/webapi/lib/types"
)

// Credential struct represents the credential table in the database
type Credential struct {
	ID             string         `db:"id" json:"id"`
	CreatedAt      time.Time      `db:"created_at" json:"createdAt"`
	CreatedBy      string         `db:"created_by" json:"createdBy"`
	ModifiedAt     time.Time      `db:"modified_at" json:"modifiedAt"`
	LastUsedAt     types.NullTime `db:"last_used_at" json:"lastUsedAt"`
	OrganizationID string         `db:"organization_id" json:"organizationId"`
	Name           string         `db:"name" json:"name"`
	Secret         string         `db:"secret" json:"secret"`
}

// Create creates a new credential in the database
func Create(name, orgID, createdBy string) (*Credential, error) {
	cred := &Credential{
		ID:             database.NewID(),
		OrganizationID: orgID,
		CreatedAt:      time.Now(),
		ModifiedAt:     time.Now(),
		CreatedBy:      createdBy,
		Name:           name,
		Secret:         random.String(64),
	}
	query := `
		INSERT INTO private.credential (id, created_at, created_by, modified_at, organization_id, name, secret)
		VALUES (:id, :created_at, :created_by, :modified_at, :organization_id, :name, :secret)
	`
	_, err := database.DB().NamedExec(query, cred)
	return cred, err
}

// FindBySecret returns a credential by its secret
func FindBySecret(secret string) (*Credential, error) {
	cred := &Credential{}
	query := "SELECT * FROM private.credential WHERE secret = $1"
	err := database.DB().Get(cred, query, secret)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

// FindByOrganizationID returns all credentials for the given organization
func FindByOrganizationID(id string) ([]*Credential, error) {
	creds := []*Credential{}
	query := "SELECT * FROM private.credential WHERE organization_id = $1"
	err := database.DB().Select(&creds, query, id)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

// DeleteByID deletes a credential with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.credential WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// DeleteByID deletes a credential with the given id
func (c *Credential) Delete() error {
	return DeleteByID(c.ID)
}
