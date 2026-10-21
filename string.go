package nullable

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// String SQL type that can retrieve NULL value
type String struct {
	Default[string]
}

// NewString creates a new nullable string
func NewString(value string) String {
	n := String{}
	n.Set(value)
	return n
}

// MarshalJSON converts current value to JSON
func (n String) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *String) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" {
		return nil
	}

	var parsed string
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.Set(parsed)
	return nil
}

// Scan implements scanner interface
func (n *String) Scan(value interface{}) error {
	if value == nil {
		n.realValue, n.isValid = "", false
		return nil
	}
	var scanned string
	if err := convertAssign(&scanned, value); err != nil {
		return err
	}
	n.Set(scanned)
	return nil
}

// Value implements the driver Valuer interface.
func (n String) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	return n.realValue, nil
}

// GormDataType gorm common data type
func (String) GormDataType() string {
	return "string_null"
}

// GormDBDataType gorm db data type
func (String) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite", "mysql":
		return "TEXT"
	case "postgres":
		return "text"
	}
	return ""
}
