package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AppSetting holds the settings for the whole AppSetting actually there is no team system there will be only one
// single record in the whole team
type AppSetting struct {
	ent.Schema
}

func (AppSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_name").Optional(),  // will need on many places e.g email templates
		field.String("copyright").Optional(), // need on many places e.g email templates
		field.String("email").Optional(),     // support email
		field.String("address").Optional(),   // full address 123 Medalling Jr., Suite 100, Parrot Park, CA 12345
		field.String("social_tw").Optional(),
		field.String("social_fb").Optional(),
		field.String("social_in").Optional(),
		field.String("logo_url").Optional(),

		field.String("site_url").Optional(),                  // this is the react app url we need on many places e.g forogt password
		field.String("default_mail_conn_id").Optional(),      // default mail connection used for sending emails
		field.Text("mail_layout_templ_id").Optional(),        // send eamil to invite user to join the workspace
		field.Text("wsapce_invite_templ_id").Optional(),      // send eamil to invite user to join the workspace
		field.Text("wsapce_success_templ_id").Optional(),     // send eamil to once user joined the workspace successfully
		field.Text("auth_fp_templ_id").Optional(),            // email template will be used for auth forgot password
		field.Text("auth_welcome_email_templ_id").Optional(), // email template for welcome email
		field.Text("auth_verification_templ_id").Optional(),  // email template for auth register verification

		// if enabled then user must verify their email otherwise not
		field.Text("auth_email_verify").Optional(),
	}
}

func (AppSetting) Edges() []ent.Edge {
	return nil
}

func (AppSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (AppSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}
