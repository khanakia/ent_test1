package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppUser struct {
	ent.Schema
}

func (AppUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_id"),
		field.String("admin_user_id"),
		field.String("app_role_id"),
	}
}

// Edges of the AppUser.
func (AppUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).
			Required().
			Unique().
			Field("app_id"),
		edge.To("adminuser", AdminUser.Type).
			Required().
			Unique().
			Field("admin_user_id"),
		edge.To("app_role", AppRole.Type).
			Required().
			Unique().
			Field("app_role_id"),
	}
}

func (AppUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
func (AppUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AppUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "admin_user_id").Unique(),
	}
}
