package schema

import (
	"lace/jsonslice"
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PostTypeForm struct {
	ent.Schema
}

func (PostTypeForm) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Annotations(
			entgql.OrderField("NAME"),
		),
		field.Bool("status").Optional().Annotations(
			entgql.OrderField("STATUS"),
		),
		field.String("post_type_id").Optional().Annotations(
		// entgql.MapsTo("postTypeId"),
		),
		field.JSON("body", jsonslice.JsonSlice{}).Optional(),
	}
}

func (PostTypeForm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post_type", PostType.Type).
			Field("post_type_id").
			Ref("post_type_forms").
			Unique().Annotations(entgql.OrderField("POST_TYPE_NAME")),
	}
}

func (PostTypeForm) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}

func (PostTypeForm) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("postTypeForms").Directives(entgql.Directive{Name: constants.DirectiveCanApp}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
