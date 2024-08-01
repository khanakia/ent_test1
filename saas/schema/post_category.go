package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PostCategory struct {
	ent.Schema
}

func (PostCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("slug").Optional(),
		field.String("status").Optional(),
		field.String("excerpt").Optional(),
		field.Text("content").Optional(),
		field.String("meta_title").Optional(),
		field.String("meta_descr").Optional(),
		field.String("meta_canonical_url").Optional(),
		field.String("meta_robots").Optional(),
	}
}

func (PostCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}

func (PostCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
