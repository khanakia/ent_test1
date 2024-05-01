package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Keyvalue holds the schema definition for the Keyvalue entity.
type Keyvalue struct {
	ent.Schema
}

// Fields of the Keyvalue.
func (Keyvalue) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),
		field.Time("created_at").
			Optional().
			Default(time.Now),
		field.Text("value").Optional(),
	}
}

// Edges of the Keyvalue.
func (Keyvalue) Edges() []ent.Edge {
	return nil
}

func (Keyvalue) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
