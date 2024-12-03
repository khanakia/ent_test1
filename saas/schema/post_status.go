package schema

import (
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PostStatus struct {
	ent.Schema
}

func (PostStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("slug").Optional(), // let if user wants to give any custom name instead of remember the postStatusId
		field.Bool("status").Optional().Annotations(
			entgql.OrderField("STATUS"),
		),
		field.String("post_type_id").Optional().Annotations(
		// entgql.MapsTo("postTypeId"),
		),
	}
}

func (PostStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post_type", PostType.Type).
			Field("post_type_id").
			Ref("post_statuses").
			Unique().Annotations(entgql.OrderField("POST_TYPE_NAME")),

		edge.To("posts", Post.Type).Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

func (PostStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}

func (PostStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("postStatuses").Directives(entgql.Directive{Name: constants.DirectiveCanApp}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (PostStatus) Hooks() []ent.Hook {
	return []ent.Hook{
		slugMutateHook(),
	}
}
