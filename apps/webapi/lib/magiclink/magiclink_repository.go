package magiclink

import (
	"time"

	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/random"
)

// MagicLink represents the magiclink table in the database
type MagicLink struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	Token     string    `db:"token"`
}

// NewMagicLink creates a new magiclink in the database
func Create(userID string) (*MagicLink, error) {
	magicLink := &MagicLink{
		ID:        database.NewID(),
		UserID:    userID,
		CreatedAt: time.Now(),
		Token:     random.String(64),
	}
	query := `
		INSERT INTO public.magiclink (id, user_id, created_at, token)
		VALUES (:id, :user_id, :created_at, :token)
	`
	_, err := database.DB().NamedExec(query, magicLink)
	return magicLink, err
}
