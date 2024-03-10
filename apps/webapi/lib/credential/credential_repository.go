package credential

import (
	"database/sql"
	"time"

	"provar.se/webapi/lib/database"
)

// Credential struct represents the credential table in the database
type Credential struct {
	ID             string       `db:"id"`
	CreatedAt      time.Time    `db:"created_at"`
	ModifiedAt     time.Time    `db:"modified_at"`
	LastUsedAt     sql.NullTime `db:"last_used_at"`
	OrganizationID string       `db:"organization_id"`
	Name           string       `db:"name"`
	Secret         string       `db:"secret"`
}

// PublicCredential represents a credential with sensitive information
// removed. This is used to send credential information to the client.
type PublicCredential struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
}

// FindBySecret returns a credential by its secret
func FindBySecret(secret string) (*Credential, error) {
	cred := &Credential{}
	query := "SELECT * FROM public.credential WHERE secret = $1"
	err := database.DB().Get(cred, query, secret)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

// DeleteByID deletes a credential with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.credential WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// Public returns a public version of the credential
func (c *Credential) Public() *PublicCredential {
	return &PublicCredential{
		ID:             c.ID,
		OrganizationID: c.OrganizationID,
		Name:           c.Name,
	}
}

// DeleteByID deletes a credential with the given id
func (c *Credential) Delete() error {
	return DeleteByID(c.ID)
}
