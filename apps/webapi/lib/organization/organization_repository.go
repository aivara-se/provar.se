package organization

import "time"

// Organization struct represents the organization table in the database
type Organization struct {
	ID         string    `db:"id"`
	Slug       string    `db:"slug"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
	OwnerID    string    `db:"owner_id"`
	Name       string    `db:"name"`
}
