package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type WorkspaceInvite struct {
	ent.Schema
}

func (WorkspaceInvite) Fields() []ent.Field {
	return []ent.Field{
		field.String("workspace_id").Optional(),
		field.String("email").Optional(),
		field.String("role").Optional(),
	}
}

func (WorkspaceInvite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workspace", Workspace.Type).
			Field("workspace_id").
			Ref("workspace_invites").
			Unique(),
	}
}

func (WorkspaceInvite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("workspace_id", "email"),
	}
}

func (WorkspaceInvite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (WorkspaceInvite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
