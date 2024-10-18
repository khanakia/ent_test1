package schema

import (
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Workspace struct {
	ent.Schema
}

func (Workspace) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Annotations(
			entgql.OrderField("NAME"),
		),
		field.Bool("is_personal").Optional(),
		field.String("user_id").Optional(),
	}
}

func (Workspace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Through("workspace_users", WorkspaceUser.Type),
		edge.To("workspace_invites", WorkspaceInvite.Type),
	}
}

func (Workspace) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.Skip(entgql.SkipAll),
		entgql.RelayConnection(),
		entgql.QueryField().Directives(entgql.Directive{Name: constants.DirectiveCanAdmin}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Workspace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
