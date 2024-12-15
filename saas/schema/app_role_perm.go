package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppRolePerm struct {
	ent.Schema
}

func (AppRolePerm) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_role_id"),
		field.String("app_perm_id"),
		field.String("app_id"),
	}
}

func (AppRolePerm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).
			Required().
			Unique().
			Field("app_id"),
		edge.To("app_perm", AppPerm.Type).
			Required().
			Unique().
			Field("app_perm_id"),
		edge.To("app_role", AppRole.Type).
			Required().
			Unique().
			Field("app_role_id"),
		// edge.To("app", App.Type).Field("app_id").Unique(),
		// edge.From("app", App.Type).
		// 	Ref("approles").
		// 	Unique(),
	}
}

func (AppRolePerm) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
func (AppRolePerm) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AppRolePerm) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_role_id", "app_perm_id").Unique(),
	}
}
