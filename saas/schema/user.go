package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	// incrementalEnabled := true

	return []ent.Field{
		field.String("email").Unique(),
		field.String("phone").Unique().Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("company").Optional(),
		field.Bool("status").Optional().Default(false),
		field.String("password").
			Sensitive().
			Optional(),
		field.String("secret").Sensitive().Optional(),
		field.String("api_key").Optional(),
		field.Bool("welcome_email_sent").Annotations().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", Session.Type),
		edge.From("workspaces", Workspace.Type).
			Through("workspace_users", WorkspaceUser.Type).
			Ref("users"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
