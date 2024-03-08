package credential

import (
	"database/sql"
	"time"
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
