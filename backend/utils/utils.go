package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONB is a jsonb array type for postgres with methods to save and read from db
type JSONB []interface{}

// Scan reads data from database and unmarshall it to struct field
func (b *JSONB) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		if string(src) == "null" {
			src = []byte("{}")
		}
		return json.Unmarshal(src, b)
	case string:
		return json.Unmarshal([]byte(src), b)
	case nil:
		return nil
	}
	return fmt.Errorf("cannot convert %T", src)
}

// Value returns data to save to database
func (b JSONB) Value() (driver.Value, error) {
	if b == nil {
		return json.Marshal([]interface{}{})
	}
	return json.Marshal(b)
}

// ToGenericArray is a shortcut to represent array as interface
func ToGenericArray[T any](arr []T) []interface{} {
	res := make([]interface{}, len(arr))
	for idx, v := range arr {
		res[idx] = v
	}
	return res
}
