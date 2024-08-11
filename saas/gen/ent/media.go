// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/media"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Media is the model entity for the Media schema.
type Media struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Disk holds the value of the "disk" field.
	Disk string `json:"disk,omitempty"`
	// Directory holds the value of the "directory" field.
	Directory string `json:"directory,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// OriginalName holds the value of the "original_name" field.
	OriginalName string `json:"original_name,omitempty"`
	// Extension holds the value of the "extension" field.
	Extension string `json:"extension,omitempty"`
	// MimeType holds the value of the "mime_type" field.
	MimeType string `json:"mime_type,omitempty"`
	// AggregateType holds the value of the "aggregate_type" field.
	AggregateType string `json:"aggregate_type,omitempty"`
	// Size holds the value of the "size" field.
	Size uint `json:"size,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsVariant holds the value of the "is_variant" field.
	IsVariant bool `json:"is_variant,omitempty"`
	// VariantName holds the value of the "variant_name" field.
	VariantName string `json:"variant_name,omitempty"`
	// OriginalMediaID holds the value of the "original_media_id" field.
	OriginalMediaID string `json:"original_media_id,omitempty"`
	// Checksum holds the value of the "checksum" field.
	Checksum string `json:"checksum,omitempty"`
	// WorkspaceID holds the value of the "workspace_id" field.
	WorkspaceID string `json:"workspace_id,omitempty"`
	// Alt holds the value of the "alt" field.
	Alt string `json:"alt,omitempty"`
	// UID holds the value of the "uid" field.
	UID string `json:"uid,omitempty"`
	// Status holds the value of the "status" field.
	Status       bool `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Media) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case media.FieldIsVariant, media.FieldStatus:
			values[i] = new(sql.NullBool)
		case media.FieldSize:
			values[i] = new(sql.NullInt64)
		case media.FieldID, media.FieldDisk, media.FieldDirectory, media.FieldName, media.FieldOriginalName, media.FieldExtension, media.FieldMimeType, media.FieldAggregateType, media.FieldDescription, media.FieldVariantName, media.FieldOriginalMediaID, media.FieldChecksum, media.FieldWorkspaceID, media.FieldAlt, media.FieldUID:
			values[i] = new(sql.NullString)
		case media.FieldCreatedAt, media.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Media fields.
func (m *Media) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case media.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case media.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case media.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case media.FieldDisk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field disk", values[i])
			} else if value.Valid {
				m.Disk = value.String
			}
		case media.FieldDirectory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field directory", values[i])
			} else if value.Valid {
				m.Directory = value.String
			}
		case media.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case media.FieldOriginalName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field original_name", values[i])
			} else if value.Valid {
				m.OriginalName = value.String
			}
		case media.FieldExtension:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extension", values[i])
			} else if value.Valid {
				m.Extension = value.String
			}
		case media.FieldMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime_type", values[i])
			} else if value.Valid {
				m.MimeType = value.String
			}
		case media.FieldAggregateType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field aggregate_type", values[i])
			} else if value.Valid {
				m.AggregateType = value.String
			}
		case media.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				m.Size = uint(value.Int64)
			}
		case media.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				m.Description = value.String
			}
		case media.FieldIsVariant:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_variant", values[i])
			} else if value.Valid {
				m.IsVariant = value.Bool
			}
		case media.FieldVariantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field variant_name", values[i])
			} else if value.Valid {
				m.VariantName = value.String
			}
		case media.FieldOriginalMediaID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field original_media_id", values[i])
			} else if value.Valid {
				m.OriginalMediaID = value.String
			}
		case media.FieldChecksum:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field checksum", values[i])
			} else if value.Valid {
				m.Checksum = value.String
			}
		case media.FieldWorkspaceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field workspace_id", values[i])
			} else if value.Valid {
				m.WorkspaceID = value.String
			}
		case media.FieldAlt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alt", values[i])
			} else if value.Valid {
				m.Alt = value.String
			}
		case media.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				m.UID = value.String
			}
		case media.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = value.Bool
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Media.
// This includes values selected through modifiers, order, etc.
func (m *Media) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Media.
// Note that you need to call Media.Unwrap() before calling this method if this Media
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Media) Update() *MediaUpdateOne {
	return NewMediaClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Media entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Media) Unwrap() *Media {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Media is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Media) String() string {
	var builder strings.Builder
	builder.WriteString("Media(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("disk=")
	builder.WriteString(m.Disk)
	builder.WriteString(", ")
	builder.WriteString("directory=")
	builder.WriteString(m.Directory)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("original_name=")
	builder.WriteString(m.OriginalName)
	builder.WriteString(", ")
	builder.WriteString("extension=")
	builder.WriteString(m.Extension)
	builder.WriteString(", ")
	builder.WriteString("mime_type=")
	builder.WriteString(m.MimeType)
	builder.WriteString(", ")
	builder.WriteString("aggregate_type=")
	builder.WriteString(m.AggregateType)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", m.Size))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(m.Description)
	builder.WriteString(", ")
	builder.WriteString("is_variant=")
	builder.WriteString(fmt.Sprintf("%v", m.IsVariant))
	builder.WriteString(", ")
	builder.WriteString("variant_name=")
	builder.WriteString(m.VariantName)
	builder.WriteString(", ")
	builder.WriteString("original_media_id=")
	builder.WriteString(m.OriginalMediaID)
	builder.WriteString(", ")
	builder.WriteString("checksum=")
	builder.WriteString(m.Checksum)
	builder.WriteString(", ")
	builder.WriteString("workspace_id=")
	builder.WriteString(m.WorkspaceID)
	builder.WriteString(", ")
	builder.WriteString("alt=")
	builder.WriteString(m.Alt)
	builder.WriteString(", ")
	builder.WriteString("uid=")
	builder.WriteString(m.UID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteByte(')')
	return builder.String()
}

// MediaSlice is a parsable slice of Media.
type MediaSlice []*Media
