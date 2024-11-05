package schema

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
	Prefix string
}

// Fields of the BaseMixin.
func (cls BaseMixin) Fields() []ent.Field {
	prefix := "u"
	if len(cls.Prefix) > 0 {
		prefix = cls.Prefix + "_"
	}
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return prefix + gonanoid.MustGenerate("0123456789abcdefghijklmnopqrstuvwxyz", 10)
			}).Unique(),

		field.Time("created_at").
			Optional().
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),

		field.Time("updated_at").
			Optional().
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("app_id").Optional().Annotations(
			entgql.Skip(),
		),
	}
}
