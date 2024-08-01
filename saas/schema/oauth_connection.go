package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type OauthConnection struct {
	ent.Schema
}

func (OauthConnection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("provider").Optional(), // google, facebook, github
		field.String("client_id").Optional(),
		field.String("client_secret").Optional(),
		field.String("scopes").Optional(),
		field.String("redirect_url").Optional(),
		field.String("dashboard_link").Optional(),
		field.String("note").Optional(),
		field.Bool("status").Optional(),
	}
}

func (OauthConnection) Edges() []ent.Edge {
	return nil
}

func (OauthConnection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (OauthConnection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
