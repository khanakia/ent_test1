package schema

import (
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AdminUser holds the schema definition for the AdminUser entity.
type AdminUser struct {
	ent.Schema
}

// Fields of the AdminUser.
func (AdminUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Annotations(
			entgql.OrderField("EMAIL"),
		).Unique(),
		field.String("phone").Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("company").Optional(),
		field.String("locale").Optional(), // en
		field.String("role_id").Optional().Default(constants.UserRoleSa),
		field.Bool("status").Optional().Default(false),
		field.String("password").
			Sensitive().
			Optional(),
		field.String("secret").Sensitive().Optional(),
		field.String("api_key").Optional(),
		field.Bool("welcome_email_sent").Annotations().Optional().Annotations(
			entgql.OrderField("WELCOME_EMAIL_SENT"),
		),
	}
}

// Edges of the AdminUser.
func (AdminUser) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (AdminUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
func (AdminUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
