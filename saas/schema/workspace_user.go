package schema

import (
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type WorkspaceUser struct {
	ent.Schema
}

func (WorkspaceUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("workspace_id"),
		field.String("user_id"),
		field.String("role").Optional(),
	}
}

func (WorkspaceUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id").
			Annotations(entgql.OrderField("USER_EMAIL")),
		edge.To("workspace", Workspace.Type).
			Required().
			Unique().
			Field("workspace_id").
			Annotations(entgql.OrderField("WORKSPACE_NAME")),
	}
}

func (WorkspaceUser) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("workspace_id", "user_id").
		// 	Unique(),
	}
}

func (WorkspaceUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}

func (WorkspaceUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField().Directives(entgql.Directive{Name: constants.DirectiveCanApp}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
