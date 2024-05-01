// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/keyvalue"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Keyvalue is the model entity for the Keyvalue schema.
type Keyvalue struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Value holds the value of the "value" field.
	Value        string `json:"value,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Keyvalue) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case keyvalue.FieldID, keyvalue.FieldValue:
			values[i] = new(sql.NullString)
		case keyvalue.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Keyvalue fields.
func (k *Keyvalue) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keyvalue.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				k.ID = value.String
			}
		case keyvalue.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				k.CreatedAt = value.Time
			}
		case keyvalue.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				k.Value = value.String
			}
		default:
			k.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Keyvalue.
// This includes values selected through modifiers, order, etc.
func (k *Keyvalue) GetValue(name string) (ent.Value, error) {
	return k.selectValues.Get(name)
}

// Update returns a builder for updating this Keyvalue.
// Note that you need to call Keyvalue.Unwrap() before calling this method if this Keyvalue
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Keyvalue) Update() *KeyvalueUpdateOne {
	return NewKeyvalueClient(k.config).UpdateOne(k)
}

// Unwrap unwraps the Keyvalue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Keyvalue) Unwrap() *Keyvalue {
	_tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("ent: Keyvalue is not a transactional entity")
	}
	k.config.driver = _tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Keyvalue) String() string {
	var builder strings.Builder
	builder.WriteString("Keyvalue(")
	builder.WriteString(fmt.Sprintf("id=%v, ", k.ID))
	builder.WriteString("created_at=")
	builder.WriteString(k.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(k.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Keyvalues is a parsable slice of Keyvalue.
type Keyvalues []*Keyvalue
