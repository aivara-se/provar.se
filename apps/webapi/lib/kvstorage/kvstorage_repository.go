package kvstorage

import (
	"time"

	"github.com/goccy/go-json"
	"provar.se/webapi/lib/database"
)

const (
	// CachedValueTTL is the time the value is stored in the cache
	CachedValueTTL = time.Hour * 24
)

// KVStorage represents the kvstorage table in the database
type KVStorage struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
	Payload   string    `db:"payload"`
}

// Get returns the cached value with the key from the database
func Get(key string, v interface{}) error {
	store := &KVStorage{}
	query := `SELECT * FROM private.kvstorage WHERE id = $1`
	err := database.DB().Get(store, query, key)
	if err != nil {
		return err
	}
	buff := []byte(store.Payload)
	return json.Unmarshal(buff, v)
}

// Set sets the value with the key in the database. This value will be stored for
// at least 24 hours. If a value already exists, it will be updated.
func Set(key string, val interface{}) error {
	buff, err := json.Marshal(val)
	if err != nil {
		return err
	}
	store := &KVStorage{
		ID:        key,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(CachedValueTTL),
		Payload:   string(buff),
	}
	query := `
		INSERT INTO private.kvstorage (id, created_at, expires_at, payload)
		VALUES (:id, :created_at, :expires_at, :payload)
		ON CONFLICT (id) DO UPDATE
		SET
			created_at = EXCLUDED.created_at,
			expires_at = EXCLUDED.expires_at,
			payload = EXCLUDED.payload
	`
	_, err = database.DB().NamedExec(query, store)
	if err != nil {
		return err
	}
	return nil
}
