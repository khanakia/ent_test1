// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/app"
	"saas/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AppUpdate is the builder for updating App entities.
type AppUpdate struct {
	config
	hooks     []Hook
	mutation  *AppMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppUpdate builder.
func (au *AppUpdate) Where(ps ...predicate.App) *AppUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AppUpdate) SetName(s string) *AppUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AppUpdate) SetNillableName(s *string) *AppUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// ClearName clears the value of the "name" field.
func (au *AppUpdate) ClearName() *AppUpdate {
	au.mutation.ClearName()
	return au
}

// SetCopyright sets the "copyright" field.
func (au *AppUpdate) SetCopyright(s string) *AppUpdate {
	au.mutation.SetCopyright(s)
	return au
}

// SetNillableCopyright sets the "copyright" field if the given value is not nil.
func (au *AppUpdate) SetNillableCopyright(s *string) *AppUpdate {
	if s != nil {
		au.SetCopyright(*s)
	}
	return au
}

// ClearCopyright clears the value of the "copyright" field.
func (au *AppUpdate) ClearCopyright() *AppUpdate {
	au.mutation.ClearCopyright()
	return au
}

// SetEmail sets the "email" field.
func (au *AppUpdate) SetEmail(s string) *AppUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (au *AppUpdate) SetNillableEmail(s *string) *AppUpdate {
	if s != nil {
		au.SetEmail(*s)
	}
	return au
}

// ClearEmail clears the value of the "email" field.
func (au *AppUpdate) ClearEmail() *AppUpdate {
	au.mutation.ClearEmail()
	return au
}

// SetAddress sets the "address" field.
func (au *AppUpdate) SetAddress(s string) *AppUpdate {
	au.mutation.SetAddress(s)
	return au
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (au *AppUpdate) SetNillableAddress(s *string) *AppUpdate {
	if s != nil {
		au.SetAddress(*s)
	}
	return au
}

// ClearAddress clears the value of the "address" field.
func (au *AppUpdate) ClearAddress() *AppUpdate {
	au.mutation.ClearAddress()
	return au
}

// SetSocialTw sets the "social_tw" field.
func (au *AppUpdate) SetSocialTw(s string) *AppUpdate {
	au.mutation.SetSocialTw(s)
	return au
}

// SetNillableSocialTw sets the "social_tw" field if the given value is not nil.
func (au *AppUpdate) SetNillableSocialTw(s *string) *AppUpdate {
	if s != nil {
		au.SetSocialTw(*s)
	}
	return au
}

// ClearSocialTw clears the value of the "social_tw" field.
func (au *AppUpdate) ClearSocialTw() *AppUpdate {
	au.mutation.ClearSocialTw()
	return au
}

// SetSocialFb sets the "social_fb" field.
func (au *AppUpdate) SetSocialFb(s string) *AppUpdate {
	au.mutation.SetSocialFb(s)
	return au
}

// SetNillableSocialFb sets the "social_fb" field if the given value is not nil.
func (au *AppUpdate) SetNillableSocialFb(s *string) *AppUpdate {
	if s != nil {
		au.SetSocialFb(*s)
	}
	return au
}

// ClearSocialFb clears the value of the "social_fb" field.
func (au *AppUpdate) ClearSocialFb() *AppUpdate {
	au.mutation.ClearSocialFb()
	return au
}

// SetSocialIn sets the "social_in" field.
func (au *AppUpdate) SetSocialIn(s string) *AppUpdate {
	au.mutation.SetSocialIn(s)
	return au
}

// SetNillableSocialIn sets the "social_in" field if the given value is not nil.
func (au *AppUpdate) SetNillableSocialIn(s *string) *AppUpdate {
	if s != nil {
		au.SetSocialIn(*s)
	}
	return au
}

// ClearSocialIn clears the value of the "social_in" field.
func (au *AppUpdate) ClearSocialIn() *AppUpdate {
	au.mutation.ClearSocialIn()
	return au
}

// SetLogoURL sets the "logo_url" field.
func (au *AppUpdate) SetLogoURL(s string) *AppUpdate {
	au.mutation.SetLogoURL(s)
	return au
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (au *AppUpdate) SetNillableLogoURL(s *string) *AppUpdate {
	if s != nil {
		au.SetLogoURL(*s)
	}
	return au
}

// ClearLogoURL clears the value of the "logo_url" field.
func (au *AppUpdate) ClearLogoURL() *AppUpdate {
	au.mutation.ClearLogoURL()
	return au
}

// SetSiteURL sets the "site_url" field.
func (au *AppUpdate) SetSiteURL(s string) *AppUpdate {
	au.mutation.SetSiteURL(s)
	return au
}

// SetNillableSiteURL sets the "site_url" field if the given value is not nil.
func (au *AppUpdate) SetNillableSiteURL(s *string) *AppUpdate {
	if s != nil {
		au.SetSiteURL(*s)
	}
	return au
}

// ClearSiteURL clears the value of the "site_url" field.
func (au *AppUpdate) ClearSiteURL() *AppUpdate {
	au.mutation.ClearSiteURL()
	return au
}

// SetDefaultMailConnID sets the "default_mail_conn_id" field.
func (au *AppUpdate) SetDefaultMailConnID(s string) *AppUpdate {
	au.mutation.SetDefaultMailConnID(s)
	return au
}

// SetNillableDefaultMailConnID sets the "default_mail_conn_id" field if the given value is not nil.
func (au *AppUpdate) SetNillableDefaultMailConnID(s *string) *AppUpdate {
	if s != nil {
		au.SetDefaultMailConnID(*s)
	}
	return au
}

// ClearDefaultMailConnID clears the value of the "default_mail_conn_id" field.
func (au *AppUpdate) ClearDefaultMailConnID() *AppUpdate {
	au.mutation.ClearDefaultMailConnID()
	return au
}

// SetMailLayoutTemplID sets the "mail_layout_templ_id" field.
func (au *AppUpdate) SetMailLayoutTemplID(s string) *AppUpdate {
	au.mutation.SetMailLayoutTemplID(s)
	return au
}

// SetNillableMailLayoutTemplID sets the "mail_layout_templ_id" field if the given value is not nil.
func (au *AppUpdate) SetNillableMailLayoutTemplID(s *string) *AppUpdate {
	if s != nil {
		au.SetMailLayoutTemplID(*s)
	}
	return au
}

// ClearMailLayoutTemplID clears the value of the "mail_layout_templ_id" field.
func (au *AppUpdate) ClearMailLayoutTemplID() *AppUpdate {
	au.mutation.ClearMailLayoutTemplID()
	return au
}

// SetWsapceInviteTemplID sets the "wsapce_invite_templ_id" field.
func (au *AppUpdate) SetWsapceInviteTemplID(s string) *AppUpdate {
	au.mutation.SetWsapceInviteTemplID(s)
	return au
}

// SetNillableWsapceInviteTemplID sets the "wsapce_invite_templ_id" field if the given value is not nil.
func (au *AppUpdate) SetNillableWsapceInviteTemplID(s *string) *AppUpdate {
	if s != nil {
		au.SetWsapceInviteTemplID(*s)
	}
	return au
}

// ClearWsapceInviteTemplID clears the value of the "wsapce_invite_templ_id" field.
func (au *AppUpdate) ClearWsapceInviteTemplID() *AppUpdate {
	au.mutation.ClearWsapceInviteTemplID()
	return au
}

// SetWsapceSuccessTemplID sets the "wsapce_success_templ_id" field.
func (au *AppUpdate) SetWsapceSuccessTemplID(s string) *AppUpdate {
	au.mutation.SetWsapceSuccessTemplID(s)
	return au
}

// SetNillableWsapceSuccessTemplID sets the "wsapce_success_templ_id" field if the given value is not nil.
func (au *AppUpdate) SetNillableWsapceSuccessTemplID(s *string) *AppUpdate {
	if s != nil {
		au.SetWsapceSuccessTemplID(*s)
	}
	return au
}

// ClearWsapceSuccessTemplID clears the value of the "wsapce_success_templ_id" field.
func (au *AppUpdate) ClearWsapceSuccessTemplID() *AppUpdate {
	au.mutation.ClearWsapceSuccessTemplID()
	return au
}

// SetAuthFpTemplID sets the "auth_fp_templ_id" field.
func (au *AppUpdate) SetAuthFpTemplID(s string) *AppUpdate {
	au.mutation.SetAuthFpTemplID(s)
	return au
}

// SetNillableAuthFpTemplID sets the "auth_fp_templ_id" field if the given value is not nil.
func (au *AppUpdate) SetNillableAuthFpTemplID(s *string) *AppUpdate {
	if s != nil {
		au.SetAuthFpTemplID(*s)
	}
	return au
}

// ClearAuthFpTemplID clears the value of the "auth_fp_templ_id" field.
func (au *AppUpdate) ClearAuthFpTemplID() *AppUpdate {
	au.mutation.ClearAuthFpTemplID()
	return au
}

// SetAuthWelcomeEmailTemplID sets the "auth_welcome_email_templ_id" field.
func (au *AppUpdate) SetAuthWelcomeEmailTemplID(s string) *AppUpdate {
	au.mutation.SetAuthWelcomeEmailTemplID(s)
	return au
}

// SetNillableAuthWelcomeEmailTemplID sets the "auth_welcome_email_templ_id" field if the given value is not nil.
func (au *AppUpdate) SetNillableAuthWelcomeEmailTemplID(s *string) *AppUpdate {
	if s != nil {
		au.SetAuthWelcomeEmailTemplID(*s)
	}
	return au
}

// ClearAuthWelcomeEmailTemplID clears the value of the "auth_welcome_email_templ_id" field.
func (au *AppUpdate) ClearAuthWelcomeEmailTemplID() *AppUpdate {
	au.mutation.ClearAuthWelcomeEmailTemplID()
	return au
}

// SetAuthVerificationTemplID sets the "auth_verification_templ_id" field.
func (au *AppUpdate) SetAuthVerificationTemplID(s string) *AppUpdate {
	au.mutation.SetAuthVerificationTemplID(s)
	return au
}

// SetNillableAuthVerificationTemplID sets the "auth_verification_templ_id" field if the given value is not nil.
func (au *AppUpdate) SetNillableAuthVerificationTemplID(s *string) *AppUpdate {
	if s != nil {
		au.SetAuthVerificationTemplID(*s)
	}
	return au
}

// ClearAuthVerificationTemplID clears the value of the "auth_verification_templ_id" field.
func (au *AppUpdate) ClearAuthVerificationTemplID() *AppUpdate {
	au.mutation.ClearAuthVerificationTemplID()
	return au
}

// SetAuthEmailVerify sets the "auth_email_verify" field.
func (au *AppUpdate) SetAuthEmailVerify(s string) *AppUpdate {
	au.mutation.SetAuthEmailVerify(s)
	return au
}

// SetNillableAuthEmailVerify sets the "auth_email_verify" field if the given value is not nil.
func (au *AppUpdate) SetNillableAuthEmailVerify(s *string) *AppUpdate {
	if s != nil {
		au.SetAuthEmailVerify(*s)
	}
	return au
}

// ClearAuthEmailVerify clears the value of the "auth_email_verify" field.
func (au *AppUpdate) ClearAuthEmailVerify() *AppUpdate {
	au.mutation.ClearAuthEmailVerify()
	return au
}

// Mutation returns the AppMutation object of the builder.
func (au *AppUpdate) Mutation() *AppMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AppUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AppUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AppUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AppUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AppUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(app.Table, app.Columns, sqlgraph.NewFieldSpec(app.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(app.FieldName, field.TypeString, value)
	}
	if au.mutation.NameCleared() {
		_spec.ClearField(app.FieldName, field.TypeString)
	}
	if value, ok := au.mutation.Copyright(); ok {
		_spec.SetField(app.FieldCopyright, field.TypeString, value)
	}
	if au.mutation.CopyrightCleared() {
		_spec.ClearField(app.FieldCopyright, field.TypeString)
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.SetField(app.FieldEmail, field.TypeString, value)
	}
	if au.mutation.EmailCleared() {
		_spec.ClearField(app.FieldEmail, field.TypeString)
	}
	if value, ok := au.mutation.Address(); ok {
		_spec.SetField(app.FieldAddress, field.TypeString, value)
	}
	if au.mutation.AddressCleared() {
		_spec.ClearField(app.FieldAddress, field.TypeString)
	}
	if value, ok := au.mutation.SocialTw(); ok {
		_spec.SetField(app.FieldSocialTw, field.TypeString, value)
	}
	if au.mutation.SocialTwCleared() {
		_spec.ClearField(app.FieldSocialTw, field.TypeString)
	}
	if value, ok := au.mutation.SocialFb(); ok {
		_spec.SetField(app.FieldSocialFb, field.TypeString, value)
	}
	if au.mutation.SocialFbCleared() {
		_spec.ClearField(app.FieldSocialFb, field.TypeString)
	}
	if value, ok := au.mutation.SocialIn(); ok {
		_spec.SetField(app.FieldSocialIn, field.TypeString, value)
	}
	if au.mutation.SocialInCleared() {
		_spec.ClearField(app.FieldSocialIn, field.TypeString)
	}
	if value, ok := au.mutation.LogoURL(); ok {
		_spec.SetField(app.FieldLogoURL, field.TypeString, value)
	}
	if au.mutation.LogoURLCleared() {
		_spec.ClearField(app.FieldLogoURL, field.TypeString)
	}
	if value, ok := au.mutation.SiteURL(); ok {
		_spec.SetField(app.FieldSiteURL, field.TypeString, value)
	}
	if au.mutation.SiteURLCleared() {
		_spec.ClearField(app.FieldSiteURL, field.TypeString)
	}
	if value, ok := au.mutation.DefaultMailConnID(); ok {
		_spec.SetField(app.FieldDefaultMailConnID, field.TypeString, value)
	}
	if au.mutation.DefaultMailConnIDCleared() {
		_spec.ClearField(app.FieldDefaultMailConnID, field.TypeString)
	}
	if value, ok := au.mutation.MailLayoutTemplID(); ok {
		_spec.SetField(app.FieldMailLayoutTemplID, field.TypeString, value)
	}
	if au.mutation.MailLayoutTemplIDCleared() {
		_spec.ClearField(app.FieldMailLayoutTemplID, field.TypeString)
	}
	if value, ok := au.mutation.WsapceInviteTemplID(); ok {
		_spec.SetField(app.FieldWsapceInviteTemplID, field.TypeString, value)
	}
	if au.mutation.WsapceInviteTemplIDCleared() {
		_spec.ClearField(app.FieldWsapceInviteTemplID, field.TypeString)
	}
	if value, ok := au.mutation.WsapceSuccessTemplID(); ok {
		_spec.SetField(app.FieldWsapceSuccessTemplID, field.TypeString, value)
	}
	if au.mutation.WsapceSuccessTemplIDCleared() {
		_spec.ClearField(app.FieldWsapceSuccessTemplID, field.TypeString)
	}
	if value, ok := au.mutation.AuthFpTemplID(); ok {
		_spec.SetField(app.FieldAuthFpTemplID, field.TypeString, value)
	}
	if au.mutation.AuthFpTemplIDCleared() {
		_spec.ClearField(app.FieldAuthFpTemplID, field.TypeString)
	}
	if value, ok := au.mutation.AuthWelcomeEmailTemplID(); ok {
		_spec.SetField(app.FieldAuthWelcomeEmailTemplID, field.TypeString, value)
	}
	if au.mutation.AuthWelcomeEmailTemplIDCleared() {
		_spec.ClearField(app.FieldAuthWelcomeEmailTemplID, field.TypeString)
	}
	if value, ok := au.mutation.AuthVerificationTemplID(); ok {
		_spec.SetField(app.FieldAuthVerificationTemplID, field.TypeString, value)
	}
	if au.mutation.AuthVerificationTemplIDCleared() {
		_spec.ClearField(app.FieldAuthVerificationTemplID, field.TypeString)
	}
	if value, ok := au.mutation.AuthEmailVerify(); ok {
		_spec.SetField(app.FieldAuthEmailVerify, field.TypeString, value)
	}
	if au.mutation.AuthEmailVerifyCleared() {
		_spec.ClearField(app.FieldAuthEmailVerify, field.TypeString)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AppUpdateOne is the builder for updating a single App entity.
type AppUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (auo *AppUpdateOne) SetName(s string) *AppUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableName(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// ClearName clears the value of the "name" field.
func (auo *AppUpdateOne) ClearName() *AppUpdateOne {
	auo.mutation.ClearName()
	return auo
}

// SetCopyright sets the "copyright" field.
func (auo *AppUpdateOne) SetCopyright(s string) *AppUpdateOne {
	auo.mutation.SetCopyright(s)
	return auo
}

// SetNillableCopyright sets the "copyright" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableCopyright(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetCopyright(*s)
	}
	return auo
}

// ClearCopyright clears the value of the "copyright" field.
func (auo *AppUpdateOne) ClearCopyright() *AppUpdateOne {
	auo.mutation.ClearCopyright()
	return auo
}

// SetEmail sets the "email" field.
func (auo *AppUpdateOne) SetEmail(s string) *AppUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableEmail(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetEmail(*s)
	}
	return auo
}

// ClearEmail clears the value of the "email" field.
func (auo *AppUpdateOne) ClearEmail() *AppUpdateOne {
	auo.mutation.ClearEmail()
	return auo
}

// SetAddress sets the "address" field.
func (auo *AppUpdateOne) SetAddress(s string) *AppUpdateOne {
	auo.mutation.SetAddress(s)
	return auo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableAddress(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetAddress(*s)
	}
	return auo
}

// ClearAddress clears the value of the "address" field.
func (auo *AppUpdateOne) ClearAddress() *AppUpdateOne {
	auo.mutation.ClearAddress()
	return auo
}

// SetSocialTw sets the "social_tw" field.
func (auo *AppUpdateOne) SetSocialTw(s string) *AppUpdateOne {
	auo.mutation.SetSocialTw(s)
	return auo
}

// SetNillableSocialTw sets the "social_tw" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableSocialTw(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetSocialTw(*s)
	}
	return auo
}

// ClearSocialTw clears the value of the "social_tw" field.
func (auo *AppUpdateOne) ClearSocialTw() *AppUpdateOne {
	auo.mutation.ClearSocialTw()
	return auo
}

// SetSocialFb sets the "social_fb" field.
func (auo *AppUpdateOne) SetSocialFb(s string) *AppUpdateOne {
	auo.mutation.SetSocialFb(s)
	return auo
}

// SetNillableSocialFb sets the "social_fb" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableSocialFb(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetSocialFb(*s)
	}
	return auo
}

// ClearSocialFb clears the value of the "social_fb" field.
func (auo *AppUpdateOne) ClearSocialFb() *AppUpdateOne {
	auo.mutation.ClearSocialFb()
	return auo
}

// SetSocialIn sets the "social_in" field.
func (auo *AppUpdateOne) SetSocialIn(s string) *AppUpdateOne {
	auo.mutation.SetSocialIn(s)
	return auo
}

// SetNillableSocialIn sets the "social_in" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableSocialIn(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetSocialIn(*s)
	}
	return auo
}

// ClearSocialIn clears the value of the "social_in" field.
func (auo *AppUpdateOne) ClearSocialIn() *AppUpdateOne {
	auo.mutation.ClearSocialIn()
	return auo
}

// SetLogoURL sets the "logo_url" field.
func (auo *AppUpdateOne) SetLogoURL(s string) *AppUpdateOne {
	auo.mutation.SetLogoURL(s)
	return auo
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableLogoURL(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetLogoURL(*s)
	}
	return auo
}

// ClearLogoURL clears the value of the "logo_url" field.
func (auo *AppUpdateOne) ClearLogoURL() *AppUpdateOne {
	auo.mutation.ClearLogoURL()
	return auo
}

// SetSiteURL sets the "site_url" field.
func (auo *AppUpdateOne) SetSiteURL(s string) *AppUpdateOne {
	auo.mutation.SetSiteURL(s)
	return auo
}

// SetNillableSiteURL sets the "site_url" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableSiteURL(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetSiteURL(*s)
	}
	return auo
}

// ClearSiteURL clears the value of the "site_url" field.
func (auo *AppUpdateOne) ClearSiteURL() *AppUpdateOne {
	auo.mutation.ClearSiteURL()
	return auo
}

// SetDefaultMailConnID sets the "default_mail_conn_id" field.
func (auo *AppUpdateOne) SetDefaultMailConnID(s string) *AppUpdateOne {
	auo.mutation.SetDefaultMailConnID(s)
	return auo
}

// SetNillableDefaultMailConnID sets the "default_mail_conn_id" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableDefaultMailConnID(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetDefaultMailConnID(*s)
	}
	return auo
}

// ClearDefaultMailConnID clears the value of the "default_mail_conn_id" field.
func (auo *AppUpdateOne) ClearDefaultMailConnID() *AppUpdateOne {
	auo.mutation.ClearDefaultMailConnID()
	return auo
}

// SetMailLayoutTemplID sets the "mail_layout_templ_id" field.
func (auo *AppUpdateOne) SetMailLayoutTemplID(s string) *AppUpdateOne {
	auo.mutation.SetMailLayoutTemplID(s)
	return auo
}

// SetNillableMailLayoutTemplID sets the "mail_layout_templ_id" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableMailLayoutTemplID(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetMailLayoutTemplID(*s)
	}
	return auo
}

// ClearMailLayoutTemplID clears the value of the "mail_layout_templ_id" field.
func (auo *AppUpdateOne) ClearMailLayoutTemplID() *AppUpdateOne {
	auo.mutation.ClearMailLayoutTemplID()
	return auo
}

// SetWsapceInviteTemplID sets the "wsapce_invite_templ_id" field.
func (auo *AppUpdateOne) SetWsapceInviteTemplID(s string) *AppUpdateOne {
	auo.mutation.SetWsapceInviteTemplID(s)
	return auo
}

// SetNillableWsapceInviteTemplID sets the "wsapce_invite_templ_id" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableWsapceInviteTemplID(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetWsapceInviteTemplID(*s)
	}
	return auo
}

// ClearWsapceInviteTemplID clears the value of the "wsapce_invite_templ_id" field.
func (auo *AppUpdateOne) ClearWsapceInviteTemplID() *AppUpdateOne {
	auo.mutation.ClearWsapceInviteTemplID()
	return auo
}

// SetWsapceSuccessTemplID sets the "wsapce_success_templ_id" field.
func (auo *AppUpdateOne) SetWsapceSuccessTemplID(s string) *AppUpdateOne {
	auo.mutation.SetWsapceSuccessTemplID(s)
	return auo
}

// SetNillableWsapceSuccessTemplID sets the "wsapce_success_templ_id" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableWsapceSuccessTemplID(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetWsapceSuccessTemplID(*s)
	}
	return auo
}

// ClearWsapceSuccessTemplID clears the value of the "wsapce_success_templ_id" field.
func (auo *AppUpdateOne) ClearWsapceSuccessTemplID() *AppUpdateOne {
	auo.mutation.ClearWsapceSuccessTemplID()
	return auo
}

// SetAuthFpTemplID sets the "auth_fp_templ_id" field.
func (auo *AppUpdateOne) SetAuthFpTemplID(s string) *AppUpdateOne {
	auo.mutation.SetAuthFpTemplID(s)
	return auo
}

// SetNillableAuthFpTemplID sets the "auth_fp_templ_id" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableAuthFpTemplID(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetAuthFpTemplID(*s)
	}
	return auo
}

// ClearAuthFpTemplID clears the value of the "auth_fp_templ_id" field.
func (auo *AppUpdateOne) ClearAuthFpTemplID() *AppUpdateOne {
	auo.mutation.ClearAuthFpTemplID()
	return auo
}

// SetAuthWelcomeEmailTemplID sets the "auth_welcome_email_templ_id" field.
func (auo *AppUpdateOne) SetAuthWelcomeEmailTemplID(s string) *AppUpdateOne {
	auo.mutation.SetAuthWelcomeEmailTemplID(s)
	return auo
}

// SetNillableAuthWelcomeEmailTemplID sets the "auth_welcome_email_templ_id" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableAuthWelcomeEmailTemplID(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetAuthWelcomeEmailTemplID(*s)
	}
	return auo
}

// ClearAuthWelcomeEmailTemplID clears the value of the "auth_welcome_email_templ_id" field.
func (auo *AppUpdateOne) ClearAuthWelcomeEmailTemplID() *AppUpdateOne {
	auo.mutation.ClearAuthWelcomeEmailTemplID()
	return auo
}

// SetAuthVerificationTemplID sets the "auth_verification_templ_id" field.
func (auo *AppUpdateOne) SetAuthVerificationTemplID(s string) *AppUpdateOne {
	auo.mutation.SetAuthVerificationTemplID(s)
	return auo
}

// SetNillableAuthVerificationTemplID sets the "auth_verification_templ_id" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableAuthVerificationTemplID(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetAuthVerificationTemplID(*s)
	}
	return auo
}

// ClearAuthVerificationTemplID clears the value of the "auth_verification_templ_id" field.
func (auo *AppUpdateOne) ClearAuthVerificationTemplID() *AppUpdateOne {
	auo.mutation.ClearAuthVerificationTemplID()
	return auo
}

// SetAuthEmailVerify sets the "auth_email_verify" field.
func (auo *AppUpdateOne) SetAuthEmailVerify(s string) *AppUpdateOne {
	auo.mutation.SetAuthEmailVerify(s)
	return auo
}

// SetNillableAuthEmailVerify sets the "auth_email_verify" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableAuthEmailVerify(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetAuthEmailVerify(*s)
	}
	return auo
}

// ClearAuthEmailVerify clears the value of the "auth_email_verify" field.
func (auo *AppUpdateOne) ClearAuthEmailVerify() *AppUpdateOne {
	auo.mutation.ClearAuthEmailVerify()
	return auo
}

// Mutation returns the AppMutation object of the builder.
func (auo *AppUpdateOne) Mutation() *AppMutation {
	return auo.mutation
}

// Where appends a list predicates to the AppUpdate builder.
func (auo *AppUpdateOne) Where(ps ...predicate.App) *AppUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AppUpdateOne) Select(field string, fields ...string) *AppUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated App entity.
func (auo *AppUpdateOne) Save(ctx context.Context) (*App, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AppUpdateOne) SaveX(ctx context.Context) *App {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AppUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AppUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AppUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AppUpdateOne) sqlSave(ctx context.Context) (_node *App, err error) {
	_spec := sqlgraph.NewUpdateSpec(app.Table, app.Columns, sqlgraph.NewFieldSpec(app.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "App.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, app.FieldID)
		for _, f := range fields {
			if !app.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != app.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(app.FieldName, field.TypeString, value)
	}
	if auo.mutation.NameCleared() {
		_spec.ClearField(app.FieldName, field.TypeString)
	}
	if value, ok := auo.mutation.Copyright(); ok {
		_spec.SetField(app.FieldCopyright, field.TypeString, value)
	}
	if auo.mutation.CopyrightCleared() {
		_spec.ClearField(app.FieldCopyright, field.TypeString)
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.SetField(app.FieldEmail, field.TypeString, value)
	}
	if auo.mutation.EmailCleared() {
		_spec.ClearField(app.FieldEmail, field.TypeString)
	}
	if value, ok := auo.mutation.Address(); ok {
		_spec.SetField(app.FieldAddress, field.TypeString, value)
	}
	if auo.mutation.AddressCleared() {
		_spec.ClearField(app.FieldAddress, field.TypeString)
	}
	if value, ok := auo.mutation.SocialTw(); ok {
		_spec.SetField(app.FieldSocialTw, field.TypeString, value)
	}
	if auo.mutation.SocialTwCleared() {
		_spec.ClearField(app.FieldSocialTw, field.TypeString)
	}
	if value, ok := auo.mutation.SocialFb(); ok {
		_spec.SetField(app.FieldSocialFb, field.TypeString, value)
	}
	if auo.mutation.SocialFbCleared() {
		_spec.ClearField(app.FieldSocialFb, field.TypeString)
	}
	if value, ok := auo.mutation.SocialIn(); ok {
		_spec.SetField(app.FieldSocialIn, field.TypeString, value)
	}
	if auo.mutation.SocialInCleared() {
		_spec.ClearField(app.FieldSocialIn, field.TypeString)
	}
	if value, ok := auo.mutation.LogoURL(); ok {
		_spec.SetField(app.FieldLogoURL, field.TypeString, value)
	}
	if auo.mutation.LogoURLCleared() {
		_spec.ClearField(app.FieldLogoURL, field.TypeString)
	}
	if value, ok := auo.mutation.SiteURL(); ok {
		_spec.SetField(app.FieldSiteURL, field.TypeString, value)
	}
	if auo.mutation.SiteURLCleared() {
		_spec.ClearField(app.FieldSiteURL, field.TypeString)
	}
	if value, ok := auo.mutation.DefaultMailConnID(); ok {
		_spec.SetField(app.FieldDefaultMailConnID, field.TypeString, value)
	}
	if auo.mutation.DefaultMailConnIDCleared() {
		_spec.ClearField(app.FieldDefaultMailConnID, field.TypeString)
	}
	if value, ok := auo.mutation.MailLayoutTemplID(); ok {
		_spec.SetField(app.FieldMailLayoutTemplID, field.TypeString, value)
	}
	if auo.mutation.MailLayoutTemplIDCleared() {
		_spec.ClearField(app.FieldMailLayoutTemplID, field.TypeString)
	}
	if value, ok := auo.mutation.WsapceInviteTemplID(); ok {
		_spec.SetField(app.FieldWsapceInviteTemplID, field.TypeString, value)
	}
	if auo.mutation.WsapceInviteTemplIDCleared() {
		_spec.ClearField(app.FieldWsapceInviteTemplID, field.TypeString)
	}
	if value, ok := auo.mutation.WsapceSuccessTemplID(); ok {
		_spec.SetField(app.FieldWsapceSuccessTemplID, field.TypeString, value)
	}
	if auo.mutation.WsapceSuccessTemplIDCleared() {
		_spec.ClearField(app.FieldWsapceSuccessTemplID, field.TypeString)
	}
	if value, ok := auo.mutation.AuthFpTemplID(); ok {
		_spec.SetField(app.FieldAuthFpTemplID, field.TypeString, value)
	}
	if auo.mutation.AuthFpTemplIDCleared() {
		_spec.ClearField(app.FieldAuthFpTemplID, field.TypeString)
	}
	if value, ok := auo.mutation.AuthWelcomeEmailTemplID(); ok {
		_spec.SetField(app.FieldAuthWelcomeEmailTemplID, field.TypeString, value)
	}
	if auo.mutation.AuthWelcomeEmailTemplIDCleared() {
		_spec.ClearField(app.FieldAuthWelcomeEmailTemplID, field.TypeString)
	}
	if value, ok := auo.mutation.AuthVerificationTemplID(); ok {
		_spec.SetField(app.FieldAuthVerificationTemplID, field.TypeString, value)
	}
	if auo.mutation.AuthVerificationTemplIDCleared() {
		_spec.ClearField(app.FieldAuthVerificationTemplID, field.TypeString)
	}
	if value, ok := auo.mutation.AuthEmailVerify(); ok {
		_spec.SetField(app.FieldAuthEmailVerify, field.TypeString, value)
	}
	if auo.mutation.AuthEmailVerifyCleared() {
		_spec.ClearField(app.FieldAuthEmailVerify, field.TypeString)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &App{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
