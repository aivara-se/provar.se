package database

import (
	"database/sql/driver"

	"github.com/goccy/go-json"
)

// StringMap is a map of strings that can be stored in the database
type StringMap map[string]string

// Value returns the JSON encoding of the StringMap
func (s StringMap) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "{}", nil
	}
	return json.Marshal(s)
}

// Scan decodes a JSON-encoded map to the StringMap
func (s *StringMap) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	}
	if len(data) == 0 {
		*s = make(StringMap)
		return nil
	}
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return nil
}
