package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("excerpt").Optional(),
		field.String("description").Optional(),
		field.Bool("status").Optional().Default(false),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{}

}

func (Plan) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (Plan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
