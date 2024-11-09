package schema

import (
	"saas/pkg/constants"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// App holds the settings for the whole App actually there is no team system there will be only one
// single record in the whole team
type App struct {
	ent.Schema
}

func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),      // will need on many places e.g email templates
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
		field.Bool("auth_email_verify").Optional().Default(true),

		// allow new user to registere if not found during Oauth SignIn

		field.Bool("oauth_signin_can_signup").Optional().Default(false),

		// if disabled user cannot login with password they must use oauth or other method
		field.Bool("auth_enable_password_login").Optional().Default(true),

		field.String("admin_user_id").Optional(),
	}
}

func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("default_mail_conn", MailConn.Type).
			Field("default_mail_conn_id").
			Unique(),
		edge.To("mail_layout_templ", Templ.Type).
			Field("mail_layout_templ_id").
			Unique(),
		edge.To("wsapce_invite_templ", Templ.Type).
			Field("wsapce_invite_templ_id").
			Unique(),
		edge.To("wsapce_success_templ", Templ.Type).
			Field("wsapce_success_templ_id").
			Unique(),
		edge.To("auth_fp_templ", Templ.Type).
			Field("auth_fp_templ_id").
			Unique(),
		edge.To("auth_welcome_email_templ", Templ.Type).
			Field("auth_welcome_email_templ_id").
			Unique(),
		edge.To("auth_verification_templ", Templ.Type).
			Field("auth_verification_templ_id").
			Unique(),
	}

}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (App) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.Skip(entgql.SkipAll),
		entgql.RelayConnection(),
		entgql.QueryField().Directives(entgql.Directive{Name: constants.DirectiveCanAdmin}),
		entgql.MultiOrder(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
