package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Mediable struct {
	ent.Schema
}

func (Mediable) Fields() []ent.Field {
	return []ent.Field{
		field.String("media_id").Optional(),
		field.String("mediable_id").Optional(),
		field.String("mediable_type").Optional(),
		field.String("tag").Optional(), // default | icon | featured
		field.Int("sort_order").Optional(),
	}
}

func (Mediable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("media", Media.Type).
			Ref("mediables").
			Field("media_id").
			Unique(),

		// edge.To("media", Media.Type).
		// 	Required().
		// 	Unique().
		// 	Field("media_id"),
	}
}

func (Mediable) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipWhereInput),
		// entgql.RelayConnection(),
		// entgql.QueryField("mediables").Directives(entgql.Directive{Name: constants.DirectiveCanApp}),
		// entgql.MultiOrder(),
		// entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Mediable) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}

func (Mediable) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("media_id", "mediable_id", "mediable_type", "tag").Unique(),
	}
}
