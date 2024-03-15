package magiclink

import (
	"time"

	"provar.se/webapi/lib/config"
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/random"
)

// MagicLink represents the magiclink table in the database
type MagicLink struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UserID    string    `db:"user_id"`
	Token     string    `db:"token"`
}

// Create creates a new magiclink in the database
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

// FindByUserID returns magic links for a user
func FindByUserID(id string) ([]*MagicLink, error) {
	magicLinks := []*MagicLink{}
	query := "SELECT * FROM public.magiclink WHERE user_id = $1"
	err := database.DB().Select(&magicLinks, query, id)
	if err != nil {
		return nil, err
	}
	return magicLinks, nil
}

// Confirm confirms the magiclink token and returns the magiclink if it exists.
// The magiclink will be invalid after this function is called.
func Confirm(token string) (*MagicLink, error) {
	magicLink := &MagicLink{}
	query := `
		DELETE FROM private.magiclink WHERE token = $1
		RETURNING *
	`
	err := database.DB().Get(magicLink, query, token)
	return magicLink, err
}

// DeleteByID deletes a magiclink with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.magiclink WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// Link returns the link to the application page where it is possible to login
// with the magiclink. The link is absolute and includes the magiclink token.
func (m *MagicLink) Link() string {
	cfg := config.Get()
	return cfg.Provar.AppURL + "/auth/login/email/verify/" + m.Token
}

// DeleteByID deletes a magiclink with the given id
func (m *MagicLink) Delete() error {
	return DeleteByID(m.ID)
}
