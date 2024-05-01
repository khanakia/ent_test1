package schema

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return "u" + gonanoid.MustGenerate("0123456789abcdefghijklmnopqrstuvwxyz", 21)
			}).Unique(),

		field.Time("created_at").
			Optional().
			Default(time.Now),

		field.Time("updated_at").
			Optional().
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
