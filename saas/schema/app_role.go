package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppRole struct {
	ent.Schema
}

func (AppRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),

		// global role will have `global` as app_id
		field.String("app_id").Optional(),

		// global roles just like aws
		field.Bool("is_global").Default(false),
	}
}

// Edges of the AppRole.
func (AppRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Unique(),

		edge.To("app_perms", AppPerm.Type).
			Through("app_role_perms", AppRolePerm.Type),

		// edge.From("app_role_perms", AppPerm.Type).
		// 	Ref("app_perms").
		// 	Through("app_role_perms", AppRolePerm.Type),
		// edge.From("app", App.Type).
		// 	Ref("approles").
		// 	Unique(),
	}
}

func (AppRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
func (AppRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AppRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "name").Unique(),
	}
}
