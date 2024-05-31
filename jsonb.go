//go:build linux || windows || darwin
// +build linux windows darwin

package pgjsonb

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Type to support PostgreSQL jsonb data type.
type JSONB struct {
	value any
}

func (j *JSONB) unmarshal(data []byte) error {
	var v any

	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	j.value = v

	return nil
}

// Implements sql.Scanner interface
func (j *JSONB) Scan(value any) error {
	if value == nil {
		j.value = nil
		return nil
	}

	var data []byte

	switch vt := value.(type) {
	case []byte:
		data = vt
	default:
		return fmt.Errorf("driver data type not supported")
	}

	if data == nil {
		return fmt.Errorf("no data found")
	}

	return j.unmarshal(data)
}

// Implements sql.Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j.value)
}

// Implements json.Marshaler interface
func (j JSONB) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.value)
}

// Implements json.Unmarshaler interface
func (j *JSONB) UnmarshalJSON(data []byte) error {
	return j.unmarshal(data)
}

// Returns json encoded string for the underlying value
func (j JSONB) String() string {
	str, _ := j.MarshalJSON()
	return string(str)
}

// Sets underlying value, can be any valid value that
// can be encoded into valid json string
func (j *JSONB) Set(value any) {
	j.value = value
}

// Returns the underlying value. This is json decoded
// value scanned from sql.Row or any valid value set by
// the Set method.
func (j *JSONB) Get() any {
	return j.value
}

// Creates new JSONB instance and set the
// underlying value to value
func New(value any) JSONB {
	return JSONB{
		value: value,
	}
}
