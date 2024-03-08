package database

import (
	"provar.se/webapi/lib/random"
)

const (
	IDLength = 8
)

// NewID generates a random ID of length 8
func NewID() string {
	return random.String(IDLength)
}
