// Code generated by ent, DO NOT EDIT.

package oauthconnection

import (
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldAppID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldName, v))
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldProvider, v))
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldClientID, v))
}

// ClientSecret applies equality check predicate on the "client_secret" field. It's identical to ClientSecretEQ.
func ClientSecret(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldClientSecret, v))
}

// Scopes applies equality check predicate on the "scopes" field. It's identical to ScopesEQ.
func Scopes(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldScopes, v))
}

// RedirectURL applies equality check predicate on the "redirect_url" field. It's identical to RedirectURLEQ.
func RedirectURL(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldRedirectURL, v))
}

// DashboardLink applies equality check predicate on the "dashboard_link" field. It's identical to DashboardLinkEQ.
func DashboardLink(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldDashboardLink, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldNote, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldUpdatedAt))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldAppID))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldAppID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldName, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldProvider, v))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldProvider, v))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldProvider, vs...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldProvider, vs...))
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldProvider, v))
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldProvider, v))
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldProvider, v))
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldProvider, v))
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldProvider, v))
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldProvider, v))
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldProvider, v))
}

// ProviderIsNil applies the IsNil predicate on the "provider" field.
func ProviderIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldProvider))
}

// ProviderNotNil applies the NotNil predicate on the "provider" field.
func ProviderNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldProvider))
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldProvider, v))
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldProvider, v))
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldClientID, v))
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldClientID, v))
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldClientID, vs...))
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldClientID, vs...))
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldClientID, v))
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldClientID, v))
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldClientID, v))
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldClientID, v))
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldClientID, v))
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldClientID, v))
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldClientID, v))
}

// ClientIDIsNil applies the IsNil predicate on the "client_id" field.
func ClientIDIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldClientID))
}

// ClientIDNotNil applies the NotNil predicate on the "client_id" field.
func ClientIDNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldClientID))
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldClientID, v))
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldClientID, v))
}

// ClientSecretEQ applies the EQ predicate on the "client_secret" field.
func ClientSecretEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldClientSecret, v))
}

// ClientSecretNEQ applies the NEQ predicate on the "client_secret" field.
func ClientSecretNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldClientSecret, v))
}

// ClientSecretIn applies the In predicate on the "client_secret" field.
func ClientSecretIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldClientSecret, vs...))
}

// ClientSecretNotIn applies the NotIn predicate on the "client_secret" field.
func ClientSecretNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldClientSecret, vs...))
}

// ClientSecretGT applies the GT predicate on the "client_secret" field.
func ClientSecretGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldClientSecret, v))
}

// ClientSecretGTE applies the GTE predicate on the "client_secret" field.
func ClientSecretGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldClientSecret, v))
}

// ClientSecretLT applies the LT predicate on the "client_secret" field.
func ClientSecretLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldClientSecret, v))
}

// ClientSecretLTE applies the LTE predicate on the "client_secret" field.
func ClientSecretLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldClientSecret, v))
}

// ClientSecretContains applies the Contains predicate on the "client_secret" field.
func ClientSecretContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldClientSecret, v))
}

// ClientSecretHasPrefix applies the HasPrefix predicate on the "client_secret" field.
func ClientSecretHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldClientSecret, v))
}

// ClientSecretHasSuffix applies the HasSuffix predicate on the "client_secret" field.
func ClientSecretHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldClientSecret, v))
}

// ClientSecretIsNil applies the IsNil predicate on the "client_secret" field.
func ClientSecretIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldClientSecret))
}

// ClientSecretNotNil applies the NotNil predicate on the "client_secret" field.
func ClientSecretNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldClientSecret))
}

// ClientSecretEqualFold applies the EqualFold predicate on the "client_secret" field.
func ClientSecretEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldClientSecret, v))
}

// ClientSecretContainsFold applies the ContainsFold predicate on the "client_secret" field.
func ClientSecretContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldClientSecret, v))
}

// ScopesEQ applies the EQ predicate on the "scopes" field.
func ScopesEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldScopes, v))
}

// ScopesNEQ applies the NEQ predicate on the "scopes" field.
func ScopesNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldScopes, v))
}

// ScopesIn applies the In predicate on the "scopes" field.
func ScopesIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldScopes, vs...))
}

// ScopesNotIn applies the NotIn predicate on the "scopes" field.
func ScopesNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldScopes, vs...))
}

// ScopesGT applies the GT predicate on the "scopes" field.
func ScopesGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldScopes, v))
}

// ScopesGTE applies the GTE predicate on the "scopes" field.
func ScopesGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldScopes, v))
}

// ScopesLT applies the LT predicate on the "scopes" field.
func ScopesLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldScopes, v))
}

// ScopesLTE applies the LTE predicate on the "scopes" field.
func ScopesLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldScopes, v))
}

// ScopesContains applies the Contains predicate on the "scopes" field.
func ScopesContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldScopes, v))
}

// ScopesHasPrefix applies the HasPrefix predicate on the "scopes" field.
func ScopesHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldScopes, v))
}

// ScopesHasSuffix applies the HasSuffix predicate on the "scopes" field.
func ScopesHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldScopes, v))
}

// ScopesIsNil applies the IsNil predicate on the "scopes" field.
func ScopesIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldScopes))
}

// ScopesNotNil applies the NotNil predicate on the "scopes" field.
func ScopesNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldScopes))
}

// ScopesEqualFold applies the EqualFold predicate on the "scopes" field.
func ScopesEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldScopes, v))
}

// ScopesContainsFold applies the ContainsFold predicate on the "scopes" field.
func ScopesContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldScopes, v))
}

// RedirectURLEQ applies the EQ predicate on the "redirect_url" field.
func RedirectURLEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldRedirectURL, v))
}

// RedirectURLNEQ applies the NEQ predicate on the "redirect_url" field.
func RedirectURLNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldRedirectURL, v))
}

// RedirectURLIn applies the In predicate on the "redirect_url" field.
func RedirectURLIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldRedirectURL, vs...))
}

// RedirectURLNotIn applies the NotIn predicate on the "redirect_url" field.
func RedirectURLNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldRedirectURL, vs...))
}

// RedirectURLGT applies the GT predicate on the "redirect_url" field.
func RedirectURLGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldRedirectURL, v))
}

// RedirectURLGTE applies the GTE predicate on the "redirect_url" field.
func RedirectURLGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldRedirectURL, v))
}

// RedirectURLLT applies the LT predicate on the "redirect_url" field.
func RedirectURLLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldRedirectURL, v))
}

// RedirectURLLTE applies the LTE predicate on the "redirect_url" field.
func RedirectURLLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldRedirectURL, v))
}

// RedirectURLContains applies the Contains predicate on the "redirect_url" field.
func RedirectURLContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldRedirectURL, v))
}

// RedirectURLHasPrefix applies the HasPrefix predicate on the "redirect_url" field.
func RedirectURLHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldRedirectURL, v))
}

// RedirectURLHasSuffix applies the HasSuffix predicate on the "redirect_url" field.
func RedirectURLHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldRedirectURL, v))
}

// RedirectURLIsNil applies the IsNil predicate on the "redirect_url" field.
func RedirectURLIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldRedirectURL))
}

// RedirectURLNotNil applies the NotNil predicate on the "redirect_url" field.
func RedirectURLNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldRedirectURL))
}

// RedirectURLEqualFold applies the EqualFold predicate on the "redirect_url" field.
func RedirectURLEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldRedirectURL, v))
}

// RedirectURLContainsFold applies the ContainsFold predicate on the "redirect_url" field.
func RedirectURLContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldRedirectURL, v))
}

// DashboardLinkEQ applies the EQ predicate on the "dashboard_link" field.
func DashboardLinkEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldDashboardLink, v))
}

// DashboardLinkNEQ applies the NEQ predicate on the "dashboard_link" field.
func DashboardLinkNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldDashboardLink, v))
}

// DashboardLinkIn applies the In predicate on the "dashboard_link" field.
func DashboardLinkIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldDashboardLink, vs...))
}

// DashboardLinkNotIn applies the NotIn predicate on the "dashboard_link" field.
func DashboardLinkNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldDashboardLink, vs...))
}

// DashboardLinkGT applies the GT predicate on the "dashboard_link" field.
func DashboardLinkGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldDashboardLink, v))
}

// DashboardLinkGTE applies the GTE predicate on the "dashboard_link" field.
func DashboardLinkGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldDashboardLink, v))
}

// DashboardLinkLT applies the LT predicate on the "dashboard_link" field.
func DashboardLinkLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldDashboardLink, v))
}

// DashboardLinkLTE applies the LTE predicate on the "dashboard_link" field.
func DashboardLinkLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldDashboardLink, v))
}

// DashboardLinkContains applies the Contains predicate on the "dashboard_link" field.
func DashboardLinkContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldDashboardLink, v))
}

// DashboardLinkHasPrefix applies the HasPrefix predicate on the "dashboard_link" field.
func DashboardLinkHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldDashboardLink, v))
}

// DashboardLinkHasSuffix applies the HasSuffix predicate on the "dashboard_link" field.
func DashboardLinkHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldDashboardLink, v))
}

// DashboardLinkIsNil applies the IsNil predicate on the "dashboard_link" field.
func DashboardLinkIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldDashboardLink))
}

// DashboardLinkNotNil applies the NotNil predicate on the "dashboard_link" field.
func DashboardLinkNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldDashboardLink))
}

// DashboardLinkEqualFold applies the EqualFold predicate on the "dashboard_link" field.
func DashboardLinkEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldDashboardLink, v))
}

// DashboardLinkContainsFold applies the ContainsFold predicate on the "dashboard_link" field.
func DashboardLinkContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldDashboardLink, v))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldContainsFold(FieldNote, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNEQ(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.OauthConnection {
	return predicate.OauthConnection(sql.FieldNotNull(FieldStatus))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OauthConnection) predicate.OauthConnection {
	return predicate.OauthConnection(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OauthConnection) predicate.OauthConnection {
	return predicate.OauthConnection(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OauthConnection) predicate.OauthConnection {
	return predicate.OauthConnection(sql.NotPredicates(p))
}
