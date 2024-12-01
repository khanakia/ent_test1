package schema

import (
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PostTag struct {
	ent.Schema
}

func (PostTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("slug").Optional(),
		field.String("status").Optional(),
		field.String("excerpt").Optional(),
		field.String("meta_title").Optional(),
		field.String("meta_descr").Optional(),
		field.String("meta_canonical_url").Optional(),
		field.String("meta_robots").Optional(),
	}
}

func (PostTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		// edge.To("posts", Post.Type).Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		// edge.To("post_statuses", PostStatus.Type).Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

func (PostTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "slug").Unique(),
	}
}

func (PostTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}

func (PostTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField().Directives(entgql.Directive{Name: constants.DirectiveCanApp}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
