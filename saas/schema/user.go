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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	// incrementalEnabled := true

	return []ent.Field{
		field.String("email").Annotations(
			entgql.OrderField("EMAIL"),
		),
		field.String("phone").Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("company").Optional(),
		field.String("locale").Optional(), // en
		field.String("role_id").Optional().Default(constants.UserRoleMember),
		field.Bool("status").Optional().Default(false),
		field.String("password").
			Sensitive().
			Optional(),
		field.String("secret").Sensitive().Optional(),
		field.String("api_key").Optional(),
		field.Bool("welcome_email_sent").Annotations().Optional().Annotations(
			entgql.OrderField("WELCOME_EMAIL_SENT"),
		),
		field.Bool("can_admin").Optional().Default(false).Annotations().Optional().Annotations(
			entgql.OrderField("CAN_ADMIN"),
		), // if true allow user to login to super admin
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

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "email").Unique(),
		index.Fields("app_id", "phone").Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField().Directives(entgql.Directive{Name: constants.DirectiveCanApp}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}
