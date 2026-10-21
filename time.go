package nullable

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Time SQL type that can retrieve NULL value
type Time struct {
	Default[time.Time]
}

// NewTime creates a new nullable time, zero time is null
func NewTime(value time.Time) Time {
	n := Time{}
	n.Set(value)
	return n
}

// MarshalJSON converts current value to JSON, zero time (year 1) when invalid
func (n Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Get())
}

// UnmarshalJSON writes JSON to this type
func (n *Time) UnmarshalJSON(data []byte) error {
	dataString := string(data)
	if len(dataString) == 0 || dataString == "null" || dataString == `""` {
		n.realValue, n.isValid = time.Time{}, false
		return nil
	}

	var parsed time.Time
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	n.Set(parsed)
	return nil
}

// Scan implements scanner interface
func (n *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		n.realValue, n.isValid = time.Time{}, false
		return nil
	case time.Time:
		n.Set(v)
		return nil
	case string:
		return n.parse(v)
	case []byte:
		return n.parse(string(v))
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type nullable.Time", value)
}

// parse covers drivers that return time as text (sqlite)
func (n *Time) parse(value string) error {
	for _, layout := range []string{time.RFC3339Nano, "2006-01-02 15:04:05.999999999-07:00", "2006-01-02 15:04:05.999999999"} {
		if parsed, err := time.Parse(layout, value); err == nil {
			n.Set(parsed)
			return nil
		}
	}
	return fmt.Errorf("cannot parse %q as nullable.Time", value)
}

// Value implements the driver Valuer interface.
func (n Time) Value() (driver.Value, error) {
	if !n.isValid {
		return nil, nil
	}
	// UTC para colunas sem fuso (DATETIME do mysql, sqlite) guardarem o instante certo
	return n.realValue.UTC(), nil
}

// GormDataType gorm common data type
func (Time) GormDataType() string {
	return "timestamp_null"
}

// GormDBDataType gorm db data type
func (Time) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "DATETIME"
	case "mysql":
		// DATETIME e não TIMESTAMP: TIMESTAMP só vai até 2038-01-19.
		// O fuso fica por conta do Value(), que grava sempre em UTC.
		return "DATETIME(6) NULL"
	case "postgres":
		return "timestamptz"
	}
	return ""
}
