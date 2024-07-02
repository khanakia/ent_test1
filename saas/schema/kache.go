package schema

import (
	"lace/publicid"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Kache holds the schema definition for the Kache entity.
type Kache struct {
	ent.Schema
}

// Fields of the Kache.
func (Kache) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return publicid.Must()
			}).
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable(),
		field.Time("created_at").
			Optional().
			Default(time.Now).
			Annotations(
			// entgql.Skip(entgql.SkipType),
			),
		field.String("key").Unique(),
		field.String("value").Optional(),
		field.Int("expires").Optional(),
	}
}

// Edges of the Kache.
func (Kache) Edges() []ent.Edge {
	return nil
}

func (Kache) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.Skip(entgql.SkipAll),
	}
}
