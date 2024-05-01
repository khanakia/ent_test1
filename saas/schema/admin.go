package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.Bool("status").Optional().Default(false),
		field.String("password").Sensitive().Optional(),
		field.String("secret").Sensitive().Optional(),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Admin) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
