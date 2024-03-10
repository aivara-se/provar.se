package organization

import "time"

// Organization struct represents the organization table in the database
type Organization struct {
	ID         string    `db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
	OwnerID    string    `db:"owner_id"`
	Name       string    `db:"name"`
	Slug       string    `db:"slug"`
}
