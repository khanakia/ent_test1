package schema

import (
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PostType struct {
	ent.Schema
}

func (PostType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("slug").Optional(),
		field.Enum("status").NamedValues(
			"Published", "PUBLISHED",
			"Draft", "DRAFT",
		).
			Default("PUBLISHED").
			Annotations(
				entgql.OrderField("STATUS"),
			),
		field.String("excerpt").Optional(),
		field.Text("content").Optional().MaxLen(255),
		field.String("meta_title").Optional(),
		field.String("meta_descr").Optional(),
		field.String("meta_canonical_url").Optional(),
		field.String("meta_robots").Optional(),
	}
}

func (PostType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("post_statuses", PostStatus.Type),
	}
}

func (PostType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{
			Prefix: constants.PrefixPostType,
		},
	}
}

func (PostType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
