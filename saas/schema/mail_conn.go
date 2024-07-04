package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MailConn holds the schema definition for the mail connection e.g. smtp, mailgun, sendgrid.
// we use this table to get the connection and send emails currently we support only smtp later we will add sendgrid and others
type MailConn struct {
	ent.Schema
}

func (MailConn) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("host").Optional(),
		field.Int("port").Optional(),
		field.String("username").Optional(),
		field.String("password").Sensitive().Optional(),
		field.Int("encryption").Optional(), // 0=none, 2=tls | 3=ssltls | 4=starttls
		field.String("from_name").Optional(),
		field.String("from_email").Optional(),
		field.Bool("status").Optional().Default(true),
	}
}

func (MailConn) Edges() []ent.Edge {
	return nil
}

func (MailConn) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
