package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("slug").Optional(),
		field.String("post_status_id").Optional(),
		field.String("post_type_id").Optional(),
		field.String("primary_category_id").Optional(),
		field.String("headline").Optional(),
		field.String("excerpt").Optional(),
		field.Text("content").Optional(),
		field.String("meta_title").Optional(),
		field.String("meta_descr").Optional(),
		field.String("meta_canonical_url").Optional(),
		field.String("meta_robots").Optional(), // noindex
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post_status", PostStatus.Type).
			Field("post_status_id").
			Ref("posts").
			Unique(),

		edge.From("post_type", PostType.Type).
			Field("post_type_id").
			Ref("posts").
			Unique(),

		edge.From("primary_category", PostCategory.Type).
			Field("primary_category_id").
			Ref("posts").
			Unique(),
	}
}
func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("posts"),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
