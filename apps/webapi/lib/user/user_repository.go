package user

import (
	"database/sql"
	"time"

	"provar.se/webapi/lib/database"
)

// User struct represents the user table in the database
type User struct {
	ID              string       `db:"id"`
	CreatedAt       time.Time    `db:"created_at"`
	ModifiedAt      time.Time    `db:"modified_at"`
	Email           string       `db:"email"`
	EmailVerifiedAt sql.NullTime `db:"email_verified_at"`
	Name            string       `db:"name"`
}

// Create creates a new user in the database
func Create(name, email string, verified bool) (*User, error) {
	user := &User{
		ID:         database.NewID(),
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		Email:      email,
		Name:       name,
	}
	if verified {
		user.EmailVerifiedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}
	query := `
		INSERT INTO public.user (id, created_at, modified_at, email, name, email_verified_at)
		VALUES (:id, :created_at, :modified_at, :email, :name, :email_verified_at)
	`
	_, err := database.DB().NamedExec(query, user)
	return user, err
}

// FindByID returns a user with the given id
func FindByID(id string) (*User, error) {
	user := &User{}
	query := "SELECT * FROM public.user WHERE id = $1"
	err := database.DB().Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmail returns a user with the given email
func FindByEmail(email string) (*User, error) {
	user := &User{}
	query := "SELECT * FROM public.user WHERE email = $1"
	err := database.DB().Get(user, query, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteByID deletes a user with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM public.user WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// DeleteByID deletes a user with the given id
func (u *User) Delete() error {
	return DeleteByID(u.ID)
}