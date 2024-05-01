package schema

import (
	"entgo.io/ent"
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
			Field("user_id"),
		edge.To("workspace", Workspace.Type).
			Required().
			Unique().
			Field("workspace_id"),
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
	}
}
