package user

import "database/sql/driver"

type Level int

const (
	Unknown Level = iota
	Low
	High
)

func (p Level) String() string {
	switch p {
	case Low:
		return "LOW"
	case High:
		return "HIGH"
	default:
		return "UNKNOWN"
	}
}

// Values provides list valid values for Enum.
func (Level) Values() []string {
	return []string{Unknown.String(), Low.String(), High.String()}
}

// Value provides the DB a string from int.
func (p Level) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *Level) Scan(val any) error {
	var s string
	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}
	switch s {
	case "LOW":
		*p = Low
	case "HIGH":
		*p = High
	default:
		*p = Unknown
	}
	return nil
}
