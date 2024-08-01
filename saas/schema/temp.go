package schema

import (
	"lace/jsontype"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Stores any temp data and for the moment let say user registration we want to store the data here and once user veriffied
// then this data will be stored to the users table to prevent ddos
type Temp struct {
	ent.Schema
}

func (Temp) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Optional(),
		field.String("type").Optional(), // user, order etc. so we can unmarshal the data accordingly

		// on user register i store the registration form here
		field.JSON("body", jsontype.JSON{}).Optional(),

		// let store i store the welcome_email_sent field here on user register more like meta data
		field.JSON("meta", jsontype.JSON{}).Optional(),
	}
}

func (Temp) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Temp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
func (Temp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
