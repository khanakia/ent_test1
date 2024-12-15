package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppPerm struct {
	ent.Schema
}

func (AppPerm) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(), // read:any
		field.String("app_id").Optional(),
	}
}

// Edges of the AppPerm.
func (AppPerm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Unique(),
		edge.From("app_roles", AppRole.Type).
			Ref("app_perms").
			Through("app_role_perms", AppRolePerm.Type),

		// edge.To("app_perms", AppPerm.Type).
		// 	Through("app_role_perms", AppRolePerm.Type),
	}
}

func (AppPerm) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
func (AppPerm) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AppPerm) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "name").Unique(),
	}
}
