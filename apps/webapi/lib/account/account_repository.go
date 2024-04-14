package account

import (
	"database/sql"
	"time"

	"provar.se/webapi/lib/database"
)

// Account represents the account table in the database
type Account struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UserID    string    `db:"user_id"`
	Provider  string    `db:"provider"`
	Session   string    `db:"session"`
	Name      string    `db:"name"`
	Avatar    string    `db:"avatar"`
}

// Create creates a new account in the database
func Create(provider, session, name, avatar string) (*Account, error) {
	acc := &Account{
		ID:        database.NewID(),
		CreatedAt: time.Now(),
		Provider:  provider,
		Session:   session,
		Name:      name,
		Avatar:    avatar,
	}
	query := `
		INSERT INTO private.account (id, created_at, provider, session, name, avatar)
		VALUES (:id, :created_at, :provider, :session, :name, :avatar)
	`
	_, err := database.DB().NamedExec(query, acc)
	return acc, err
}

// Upsert creates a new account if it does not exist, otherwise it updates the
// existing account with the given details.
func Upsert(userID, provider, session, name, avatar string) (*Account, error) {
	tx, err := database.DB().Beginx()
	if err != nil {
		tx.Rollback()
		return nil, nil
	}
	acc := &Account{}
	query := `SELECT * FROM private.account WHERE provider = $1 AND user_id = $2`
	err = tx.Get(acc, query, provider, userID)
	// If the account exists, update the account with the new given details
	if acc.ID != "" {
		acc.Provider = provider
		acc.Session = session
		acc.Name = name
		acc.Avatar = avatar
		query = `
			UPDATE private.account
			SET
				provider = :provider,
				session = :session,
				name = :name,
				avatar = :avatar
			WHERE provider = :provider AND user_id = :user_id
		`
		_, err = tx.NamedExec(query, acc)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
		return acc, nil
	}
	// If the account does not exist, create a new account with the given details
	if err != sql.ErrNoRows {
		tx.Rollback()
		return nil, err
	}
	acc = &Account{
		ID:        database.NewID(),
		CreatedAt: time.Now(),
		Provider:  provider,
		Session:   session,
		Name:      name,
		Avatar:    avatar,
	}
	query = `
		INSERT INTO private.account (id, created_at, provider, session, name, avatar)
		VALUES (:id, :created_at, :provider, :session, :name, :avatar)
	`
	_, err = database.DB().NamedExec(query, acc)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// UpdateByID updates an account with the given details
func UpdateByID(id, providerName, gothSession, userID string) error {
	query := `
		UPDATE private.account
		SET provider = $2, session = $3, user_id = $4
		WHERE id = $1
	`
	_, err := database.DB().Exec(query, id, providerName, gothSession, userID)
	return err
}

// FindByID returns an account by its id
func FindByID(id string) (*Account, error) {
	acc := &Account{}
	query := `SELECT * FROM private.account WHERE id = $1`
	err := database.DB().Get(acc, query, id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
