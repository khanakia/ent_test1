package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return "u" + gonanoid.MustGenerate("0123456789abcdefghijklmnopqrstuvwxyz", 32)
			}).Unique(),

		field.Time("created_at").
			Optional().
			Default(time.Now),

		field.Time("updated_at").
			Optional().
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("user_id").Optional(),
		field.String("ip").Optional(),
		field.String("user_agent").Optional(),
		field.String("payload").Optional(),
		field.Time("expires_at").Optional(),
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Field("user_id").
			Ref("sessions").
			Unique(),
	}
}

func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// BaseMixin{},
	}
}
