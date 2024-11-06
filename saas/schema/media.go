package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Media struct {
	ent.Schema
}

func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("disk").Optional(),          // public
		field.String("directory").Optional(),     // app/1/1
		field.String("name").Optional(),          // myimage.jpg user can update this
		field.String("original_name").Optional(), // uploaded file name image.jpg

		field.String("extension").Optional(),      // webp, jpg
		field.String("mime_type").Optional(),      // image/jpeg
		field.String("aggregate_type").Optional(), // image
		field.Uint("size").Optional(),             // original file size
		field.String("description").Optional(),
		field.Bool("is_variant").Optional(),
		field.String("variant_name").Optional(),
		field.String("original_media_id").Optional(), // parent ID
		field.String("checksum").Optional(),
		field.String("workspace_id"),
		field.String("alt").Optional(),
		field.String("uid").Optional().Unique(), // just a random unique id to the media
		field.Bool("status").Optional().Default(false),
	}
}

func (Media) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Media) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}

func (Media) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseApp{},
	}
}
