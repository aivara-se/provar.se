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

// FindBySecret returns a credential by its secret
func FindBySecret(secret string) (*Credential, error) {
	cred := &Credential{}
	query := "SELECT * FROM credential WHERE secret = ?"
	err := database.DB().Get(cred, query, secret)
	if err != nil {
		return nil, err
	}
	return cred, nil
}
