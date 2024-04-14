package account

import (
	"database/sql"
	"time"

	"provar.se/webapi/lib/database"
)

// Account represents the account table in the database
type Account struct {
	ID        string         `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	UserID    sql.NullString `db:"user_id"`
	Provider  string         `db:"provider"`
	Session   string         `db:"session"`
}

// AccountState represents the accountstate table in the database
type AccountState struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	AccountID string    `db:"account_id"`
	State     string    `db:"state"`
}

// Create creates a new account in the database
func Create(providerName, gothSession string) (*Account, error) {
	acc := &Account{
		ID:        database.NewID(),
		CreatedAt: time.Now(),
		Provider:  providerName,
		Session:   gothSession,
	}
	query := `
		INSERT INTO private.account (id, created_at, provider, session)
		VALUES (:id, :created_at, :provider, :session)
	`
	_, err := database.DB().NamedExec(query, acc)
	return acc, err
}

// CreateState creates a new account state in the database
func CreateState(accountID, state string) (*AccountState, error) {
	stateObj := &AccountState{
		ID:        database.NewID(),
		CreatedAt: time.Now(),
		AccountID: accountID,
		State:     state,
	}
	query := `
		INSERT INTO private.accountstate (id, created_at, account_id, state)
		VALUES (:id, :created_at, :account_id, :state)
	`
	_, err := database.DB().NamedExec(query, stateObj)
	return stateObj, err
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

// FindByState returns an account by an oauth2 state value
func FindByState(state string) (*Account, error) {
	acc := &Account{}
	query := `
		SELECT a.*
		FROM private.accountstate s
		JOIN private.account a ON s.account_id = a.id
		WHERE s.state = $1
		LIMIT 1
	`
	err := database.DB().Get(acc, query, state)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
