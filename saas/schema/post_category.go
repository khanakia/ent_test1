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

type PostCategory struct {
	ent.Schema
}

func (PostCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("parent_id").Optional(),
		field.String("name").Optional(),
		field.String("slug").Optional(),
		field.String("status").Optional(),
		field.String("excerpt").Optional(),
		field.Text("content").Optional(),
		field.String("meta_title").Optional(),
		field.String("meta_descr").Optional(),
		field.String("meta_canonical_url").Optional(),
		field.String("meta_robots").Optional(),
	}
}

func (PostCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("children", PostCategory.Type).
			From("parent").
			Field("parent_id").
			Unique(),
	}
}

func (PostCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "slug").Unique(),
	}
}

func (PostCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}

func (PostCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("postCategories").Directives(entgql.Directive{Name: constants.DirectiveCanApp}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (PostCategory) Hooks() []ent.Hook {
	return []ent.Hook{
		slugMutateHook(),
	}
}
