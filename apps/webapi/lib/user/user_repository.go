package user

import (
	"database/sql"
	"strings"
	"time"

	"github.com/automattic/go-gravatar"
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/types"
)

// User struct represents the user table in the database
type User struct {
	ID              string         `db:"id" json:"id"`
	CreatedAt       time.Time      `db:"created_at" json:"createdAt"`
	ModifiedAt      time.Time      `db:"modified_at" json:"modifiedAt"`
	Avatar          string         `db:"avatar" json:"avatar"`
	Email           string         `db:"email" json:"email"`
	EmailVerifiedAt types.NullTime `db:"email_verified_at" json:"emailVerifiedAt"`
	Name            string         `db:"name" json:"name"`
}

// Create creates a new user in the database
func Create(name, email string, verified bool) (*User, error) {
	avatar := gravatar.NewGravatarFromEmail(email)
	avatar.Default = "mp"
	user := &User{
		ID:         database.NewID(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		Avatar:     avatar.GetURL(),
		Email:      email,
		Name:       name,
	}
	if name == "" {
		emailParts := strings.Split(email, "@")
		if len(emailParts) > 0 {
			user.Name = emailParts[0]
		}
	}
	if verified {
		user.EmailVerifiedAt = types.NewNullTime(time.Now())
	}
	query := `
		INSERT INTO private.user (id, created_at, modified_at, email, name, email_verified_at, avatar)
		VALUES (:id, :created_at, :modified_at, :email, :name, :email_verified_at, :avatar)
	`
	_, err := database.DB().NamedExec(query, user)
	return user, err
}

// Upsert creates a new user if it does not exist, otherwise it updates the
// existing user with the given details.
func Upsert(name, email, avatar string, verified bool) (*User, error) {
	tx, err := database.DB().Beginx()
	if err != nil {
		tx.Rollback()
		return nil, nil
	}
	user := &User{}
	query := "SELECT * FROM private.user WHERE email = $1"
	err = tx.Get(user, query, email)
	// If the user exists, update the user with the new given details
	if user.ID != "" {
		if !user.EmailVerifiedAt.Valid && verified {
			user.EmailVerifiedAt = types.NewNullTime(time.Now())
		}
		if name != "" {
			user.Name = name
		}
		if avatar != "" {
			user.Avatar = avatar
		}
		user.ModifiedAt = time.Now()
		query = `
			UPDATE private.user
			SET
				name = :name,
				avatar = :avatar,
				modified_at = :modified_at,
				email_verified_at = :email_verified_at
			WHERE id = :id
		`
		_, err = tx.NamedExec(query, user)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	// If the user does not exist, create a new user with the given details
	if err != sql.ErrNoRows {
		tx.Rollback()
		return nil, err
	}
	user = &User{
		ID:         database.NewID(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		Avatar:     avatar,
		Email:      email,
		Name:       name,
	}
	if name == "" {
		emailParts := strings.Split(email, "@")
		if len(emailParts) > 0 {
			user.Name = emailParts[0]
		}
	}
	if verified {
		user.EmailVerifiedAt = types.NewNullTime(time.Now())
	}
	query = `
		INSERT INTO private.user (id, created_at, modified_at, email, name, email_verified_at, avatar)
		VALUES (:id, :created_at, :modified_at, :email, :name, :email_verified_at, :avatar)
	`
	_, err = tx.NamedExec(query, user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return user, err
}

// FindByID returns a user with the given id
func FindByID(id string) (*User, error) {
	user := &User{}
	query := "SELECT * FROM private.user WHERE id = $1"
	err := database.DB().Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmail returns a user with the given email
func FindByEmail(email string) (*User, error) {
	user := &User{}
	query := "SELECT * FROM private.user WHERE email = $1"
	err := database.DB().Get(user, query, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByOrganizationID returns all members of an organization
func FindByOrganizationID(id string) ([]*User, error) {
	users := []*User{}
	query := `
		SELECT u.*
		FROM private.organizationmember om
		JOIN private.user u ON om.user_id = u.id
		WHERE om.organization_id = $1
		ORDER BY u.name
	`
	err := database.DB().Select(&users, query, id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// SetEmailVerified sets the user's email as verified
func SetEmailVerified(id string) error {
	query := "UPDATE private.user SET email_verified_at = $1 WHERE id = $2"
	_, err := database.DB().Exec(query, time.Now(), id)
	return err
}

// DeleteByID deletes a user with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.user WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// SetEmailVerified sets the user's email as verified
func (u *User) SetEmailVerified() error {
	return SetEmailVerified(u.ID)
}

// DeleteByID deletes a user with the given id
func (u *User) Delete() error {
	return DeleteByID(u.ID)
}
