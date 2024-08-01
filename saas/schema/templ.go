package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Templ holds the schema definition for the template rendered
// user can create custom template and render via api and it will also compile the tailwind css and create html
type Templ struct {
	ent.Schema
}

func (Templ) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.Text("body").Optional(),
		field.Text("compiled").Optional(),
		field.Bool("status").Optional().Default(true),
	}
}

func (Templ) Edges() []ent.Edge {
	return nil
}

func (Templ) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Templ) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
