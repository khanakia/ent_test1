package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PostStatus struct {
	ent.Schema
}

func (PostStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("slug").Optional(), // let if user wants to give any custom name instead of remember the postStatusId
		field.Bool("status").Optional(),
		field.String("post_type_id").Optional(),
	}
}

func (PostStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post_type", PostType.Type).
			Field("post_type_id").
			Ref("post_statuses").
			Unique(),

		edge.To("posts", Post.Type),
	}
}

func (PostStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
