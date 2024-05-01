package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MailConnection holds the schema definition for the mail connection e.g. smtp, mailgun, sendgrid.
// we use this table to get the connection and send emails currently we support only smtp later we will add sendgrid and others
type MailConnection struct {
	ent.Schema
}

func (MailConnection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("host").Optional(),
		field.String("port").Optional(),
		field.String("username").Optional(),
		field.String("password").Sensitive().Optional(),
		field.Bool("status").Optional().Default(false),
	}
}

func (MailConnection) Edges() []ent.Edge {
	return nil
}

func (MailConnection) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}
