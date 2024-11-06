// Code generated by ent, DO NOT EDIT.

package app

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the app type in the database.
	Label = "app"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCopyright holds the string denoting the copyright field in the database.
	FieldCopyright = "copyright"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldSocialTw holds the string denoting the social_tw field in the database.
	FieldSocialTw = "social_tw"
	// FieldSocialFb holds the string denoting the social_fb field in the database.
	FieldSocialFb = "social_fb"
	// FieldSocialIn holds the string denoting the social_in field in the database.
	FieldSocialIn = "social_in"
	// FieldLogoURL holds the string denoting the logo_url field in the database.
	FieldLogoURL = "logo_url"
	// FieldSiteURL holds the string denoting the site_url field in the database.
	FieldSiteURL = "site_url"
	// FieldDefaultMailConnID holds the string denoting the default_mail_conn_id field in the database.
	FieldDefaultMailConnID = "default_mail_conn_id"
	// FieldMailLayoutTemplID holds the string denoting the mail_layout_templ_id field in the database.
	FieldMailLayoutTemplID = "mail_layout_templ_id"
	// FieldWsapceInviteTemplID holds the string denoting the wsapce_invite_templ_id field in the database.
	FieldWsapceInviteTemplID = "wsapce_invite_templ_id"
	// FieldWsapceSuccessTemplID holds the string denoting the wsapce_success_templ_id field in the database.
	FieldWsapceSuccessTemplID = "wsapce_success_templ_id"
	// FieldAuthFpTemplID holds the string denoting the auth_fp_templ_id field in the database.
	FieldAuthFpTemplID = "auth_fp_templ_id"
	// FieldAuthWelcomeEmailTemplID holds the string denoting the auth_welcome_email_templ_id field in the database.
	FieldAuthWelcomeEmailTemplID = "auth_welcome_email_templ_id"
	// FieldAuthVerificationTemplID holds the string denoting the auth_verification_templ_id field in the database.
	FieldAuthVerificationTemplID = "auth_verification_templ_id"
	// FieldAuthEmailVerify holds the string denoting the auth_email_verify field in the database.
	FieldAuthEmailVerify = "auth_email_verify"
	// FieldAdminUserID holds the string denoting the admin_user_id field in the database.
	FieldAdminUserID = "admin_user_id"
	// Table holds the table name of the app in the database.
	Table = "apps"
)

// Columns holds all SQL columns for app fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldCopyright,
	FieldEmail,
	FieldAddress,
	FieldSocialTw,
	FieldSocialFb,
	FieldSocialIn,
	FieldLogoURL,
	FieldSiteURL,
	FieldDefaultMailConnID,
	FieldMailLayoutTemplID,
	FieldWsapceInviteTemplID,
	FieldWsapceSuccessTemplID,
	FieldAuthFpTemplID,
	FieldAuthWelcomeEmailTemplID,
	FieldAuthVerificationTemplID,
	FieldAuthEmailVerify,
	FieldAdminUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the App queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCopyright orders the results by the copyright field.
func ByCopyright(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCopyright, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// BySocialTw orders the results by the social_tw field.
func BySocialTw(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSocialTw, opts...).ToFunc()
}

// BySocialFb orders the results by the social_fb field.
func BySocialFb(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSocialFb, opts...).ToFunc()
}

// BySocialIn orders the results by the social_in field.
func BySocialIn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSocialIn, opts...).ToFunc()
}

// ByLogoURL orders the results by the logo_url field.
func ByLogoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoURL, opts...).ToFunc()
}

// BySiteURL orders the results by the site_url field.
func BySiteURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSiteURL, opts...).ToFunc()
}

// ByDefaultMailConnID orders the results by the default_mail_conn_id field.
func ByDefaultMailConnID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultMailConnID, opts...).ToFunc()
}

// ByMailLayoutTemplID orders the results by the mail_layout_templ_id field.
func ByMailLayoutTemplID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMailLayoutTemplID, opts...).ToFunc()
}

// ByWsapceInviteTemplID orders the results by the wsapce_invite_templ_id field.
func ByWsapceInviteTemplID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWsapceInviteTemplID, opts...).ToFunc()
}

// ByWsapceSuccessTemplID orders the results by the wsapce_success_templ_id field.
func ByWsapceSuccessTemplID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWsapceSuccessTemplID, opts...).ToFunc()
}

// ByAuthFpTemplID orders the results by the auth_fp_templ_id field.
func ByAuthFpTemplID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthFpTemplID, opts...).ToFunc()
}

// ByAuthWelcomeEmailTemplID orders the results by the auth_welcome_email_templ_id field.
func ByAuthWelcomeEmailTemplID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthWelcomeEmailTemplID, opts...).ToFunc()
}

// ByAuthVerificationTemplID orders the results by the auth_verification_templ_id field.
func ByAuthVerificationTemplID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthVerificationTemplID, opts...).ToFunc()
}

// ByAuthEmailVerify orders the results by the auth_email_verify field.
func ByAuthEmailVerify(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthEmailVerify, opts...).ToFunc()
}

// ByAdminUserID orders the results by the admin_user_id field.
func ByAdminUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdminUserID, opts...).ToFunc()
}
