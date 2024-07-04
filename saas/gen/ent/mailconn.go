// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/mailconn"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MailConn is the model entity for the MailConn schema.
type MailConn struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Host holds the value of the "host" field.
	Host string `json:"host,omitempty"`
	// Port holds the value of the "port" field.
	Port int `json:"port,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Encryption holds the value of the "encryption" field.
	Encryption int `json:"encryption,omitempty"`
	// FromName holds the value of the "from_name" field.
	FromName string `json:"from_name,omitempty"`
	// FromEmail holds the value of the "from_email" field.
	FromEmail string `json:"from_email,omitempty"`
	// Status holds the value of the "status" field.
	Status       bool `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MailConn) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mailconn.FieldStatus:
			values[i] = new(sql.NullBool)
		case mailconn.FieldPort, mailconn.FieldEncryption:
			values[i] = new(sql.NullInt64)
		case mailconn.FieldID, mailconn.FieldName, mailconn.FieldHost, mailconn.FieldUsername, mailconn.FieldPassword, mailconn.FieldFromName, mailconn.FieldFromEmail:
			values[i] = new(sql.NullString)
		case mailconn.FieldCreatedAt, mailconn.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MailConn fields.
func (mc *MailConn) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mailconn.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				mc.ID = value.String
			}
		case mailconn.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mc.CreatedAt = value.Time
			}
		case mailconn.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mc.UpdatedAt = value.Time
			}
		case mailconn.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mc.Name = value.String
			}
		case mailconn.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				mc.Host = value.String
			}
		case mailconn.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				mc.Port = int(value.Int64)
			}
		case mailconn.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				mc.Username = value.String
			}
		case mailconn.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				mc.Password = value.String
			}
		case mailconn.FieldEncryption:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field encryption", values[i])
			} else if value.Valid {
				mc.Encryption = int(value.Int64)
			}
		case mailconn.FieldFromName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from_name", values[i])
			} else if value.Valid {
				mc.FromName = value.String
			}
		case mailconn.FieldFromEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from_email", values[i])
			} else if value.Valid {
				mc.FromEmail = value.String
			}
		case mailconn.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mc.Status = value.Bool
			}
		default:
			mc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MailConn.
// This includes values selected through modifiers, order, etc.
func (mc *MailConn) Value(name string) (ent.Value, error) {
	return mc.selectValues.Get(name)
}

// Update returns a builder for updating this MailConn.
// Note that you need to call MailConn.Unwrap() before calling this method if this MailConn
// was returned from a transaction, and the transaction was committed or rolled back.
func (mc *MailConn) Update() *MailConnUpdateOne {
	return NewMailConnClient(mc.config).UpdateOne(mc)
}

// Unwrap unwraps the MailConn entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mc *MailConn) Unwrap() *MailConn {
	_tx, ok := mc.config.driver.(*txDriver)
	if !ok {
		panic("ent: MailConn is not a transactional entity")
	}
	mc.config.driver = _tx.drv
	return mc
}

// String implements the fmt.Stringer.
func (mc *MailConn) String() string {
	var builder strings.Builder
	builder.WriteString("MailConn(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(mc.Name)
	builder.WriteString(", ")
	builder.WriteString("host=")
	builder.WriteString(mc.Host)
	builder.WriteString(", ")
	builder.WriteString("port=")
	builder.WriteString(fmt.Sprintf("%v", mc.Port))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(mc.Username)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("encryption=")
	builder.WriteString(fmt.Sprintf("%v", mc.Encryption))
	builder.WriteString(", ")
	builder.WriteString("from_name=")
	builder.WriteString(mc.FromName)
	builder.WriteString(", ")
	builder.WriteString("from_email=")
	builder.WriteString(mc.FromEmail)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mc.Status))
	builder.WriteByte(')')
	return builder.String()
}

// MailConns is a parsable slice of MailConn.
type MailConns []*MailConn
