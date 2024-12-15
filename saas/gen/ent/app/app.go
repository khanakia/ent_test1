// Code generated by ent, DO NOT EDIT.

package app

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldOauthSigninCanSignup holds the string denoting the oauth_signin_can_signup field in the database.
	FieldOauthSigninCanSignup = "oauth_signin_can_signup"
	// FieldAuthEnablePasswordLogin holds the string denoting the auth_enable_password_login field in the database.
	FieldAuthEnablePasswordLogin = "auth_enable_password_login"
	// FieldAdminUserID holds the string denoting the admin_user_id field in the database.
	FieldAdminUserID = "admin_user_id"
	// EdgeDefaultMailConn holds the string denoting the default_mail_conn edge name in mutations.
	EdgeDefaultMailConn = "default_mail_conn"
	// EdgeMailLayoutTempl holds the string denoting the mail_layout_templ edge name in mutations.
	EdgeMailLayoutTempl = "mail_layout_templ"
	// EdgeWsapceInviteTempl holds the string denoting the wsapce_invite_templ edge name in mutations.
	EdgeWsapceInviteTempl = "wsapce_invite_templ"
	// EdgeWsapceSuccessTempl holds the string denoting the wsapce_success_templ edge name in mutations.
	EdgeWsapceSuccessTempl = "wsapce_success_templ"
	// EdgeAuthFpTempl holds the string denoting the auth_fp_templ edge name in mutations.
	EdgeAuthFpTempl = "auth_fp_templ"
	// EdgeAuthWelcomeEmailTempl holds the string denoting the auth_welcome_email_templ edge name in mutations.
	EdgeAuthWelcomeEmailTempl = "auth_welcome_email_templ"
	// EdgeAuthVerificationTempl holds the string denoting the auth_verification_templ edge name in mutations.
	EdgeAuthVerificationTempl = "auth_verification_templ"
	// EdgeAdminUser holds the string denoting the admin_user edge name in mutations.
	EdgeAdminUser = "admin_user"
	// EdgeAppUsers holds the string denoting the app_users edge name in mutations.
	EdgeAppUsers = "app_users"
	// Table holds the table name of the app in the database.
	Table = "apps"
	// DefaultMailConnTable is the table that holds the default_mail_conn relation/edge.
	DefaultMailConnTable = "apps"
	// DefaultMailConnInverseTable is the table name for the MailConn entity.
	// It exists in this package in order to avoid circular dependency with the "mailconn" package.
	DefaultMailConnInverseTable = "mail_conns"
	// DefaultMailConnColumn is the table column denoting the default_mail_conn relation/edge.
	DefaultMailConnColumn = "default_mail_conn_id"
	// MailLayoutTemplTable is the table that holds the mail_layout_templ relation/edge.
	MailLayoutTemplTable = "apps"
	// MailLayoutTemplInverseTable is the table name for the Templ entity.
	// It exists in this package in order to avoid circular dependency with the "templ" package.
	MailLayoutTemplInverseTable = "templs"
	// MailLayoutTemplColumn is the table column denoting the mail_layout_templ relation/edge.
	MailLayoutTemplColumn = "mail_layout_templ_id"
	// WsapceInviteTemplTable is the table that holds the wsapce_invite_templ relation/edge.
	WsapceInviteTemplTable = "apps"
	// WsapceInviteTemplInverseTable is the table name for the Templ entity.
	// It exists in this package in order to avoid circular dependency with the "templ" package.
	WsapceInviteTemplInverseTable = "templs"
	// WsapceInviteTemplColumn is the table column denoting the wsapce_invite_templ relation/edge.
	WsapceInviteTemplColumn = "wsapce_invite_templ_id"
	// WsapceSuccessTemplTable is the table that holds the wsapce_success_templ relation/edge.
	WsapceSuccessTemplTable = "apps"
	// WsapceSuccessTemplInverseTable is the table name for the Templ entity.
	// It exists in this package in order to avoid circular dependency with the "templ" package.
	WsapceSuccessTemplInverseTable = "templs"
	// WsapceSuccessTemplColumn is the table column denoting the wsapce_success_templ relation/edge.
	WsapceSuccessTemplColumn = "wsapce_success_templ_id"
	// AuthFpTemplTable is the table that holds the auth_fp_templ relation/edge.
	AuthFpTemplTable = "apps"
	// AuthFpTemplInverseTable is the table name for the Templ entity.
	// It exists in this package in order to avoid circular dependency with the "templ" package.
	AuthFpTemplInverseTable = "templs"
	// AuthFpTemplColumn is the table column denoting the auth_fp_templ relation/edge.
	AuthFpTemplColumn = "auth_fp_templ_id"
	// AuthWelcomeEmailTemplTable is the table that holds the auth_welcome_email_templ relation/edge.
	AuthWelcomeEmailTemplTable = "apps"
	// AuthWelcomeEmailTemplInverseTable is the table name for the Templ entity.
	// It exists in this package in order to avoid circular dependency with the "templ" package.
	AuthWelcomeEmailTemplInverseTable = "templs"
	// AuthWelcomeEmailTemplColumn is the table column denoting the auth_welcome_email_templ relation/edge.
	AuthWelcomeEmailTemplColumn = "auth_welcome_email_templ_id"
	// AuthVerificationTemplTable is the table that holds the auth_verification_templ relation/edge.
	AuthVerificationTemplTable = "apps"
	// AuthVerificationTemplInverseTable is the table name for the Templ entity.
	// It exists in this package in order to avoid circular dependency with the "templ" package.
	AuthVerificationTemplInverseTable = "templs"
	// AuthVerificationTemplColumn is the table column denoting the auth_verification_templ relation/edge.
	AuthVerificationTemplColumn = "auth_verification_templ_id"
	// AdminUserTable is the table that holds the admin_user relation/edge. The primary key declared below.
	AdminUserTable = "app_users"
	// AdminUserInverseTable is the table name for the AdminUser entity.
	// It exists in this package in order to avoid circular dependency with the "adminuser" package.
	AdminUserInverseTable = "admin_users"
	// AppUsersTable is the table that holds the app_users relation/edge.
	AppUsersTable = "app_users"
	// AppUsersInverseTable is the table name for the AppUser entity.
	// It exists in this package in order to avoid circular dependency with the "appuser" package.
	AppUsersInverseTable = "app_users"
	// AppUsersColumn is the table column denoting the app_users relation/edge.
	AppUsersColumn = "app_id"
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
	FieldOauthSigninCanSignup,
	FieldAuthEnablePasswordLogin,
	FieldAdminUserID,
}

var (
	// AdminUserPrimaryKey and AdminUserColumn2 are the table columns denoting the
	// primary key for the admin_user relation (M2M).
	AdminUserPrimaryKey = []string{"admin_user_id", "app_id"}
)

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
	// DefaultAuthEmailVerify holds the default value on creation for the "auth_email_verify" field.
	DefaultAuthEmailVerify bool
	// DefaultOauthSigninCanSignup holds the default value on creation for the "oauth_signin_can_signup" field.
	DefaultOauthSigninCanSignup bool
	// DefaultAuthEnablePasswordLogin holds the default value on creation for the "auth_enable_password_login" field.
	DefaultAuthEnablePasswordLogin bool
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

// ByOauthSigninCanSignup orders the results by the oauth_signin_can_signup field.
func ByOauthSigninCanSignup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOauthSigninCanSignup, opts...).ToFunc()
}

// ByAuthEnablePasswordLogin orders the results by the auth_enable_password_login field.
func ByAuthEnablePasswordLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthEnablePasswordLogin, opts...).ToFunc()
}

// ByAdminUserID orders the results by the admin_user_id field.
func ByAdminUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdminUserID, opts...).ToFunc()
}

// ByDefaultMailConnField orders the results by default_mail_conn field.
func ByDefaultMailConnField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDefaultMailConnStep(), sql.OrderByField(field, opts...))
	}
}

// ByMailLayoutTemplField orders the results by mail_layout_templ field.
func ByMailLayoutTemplField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMailLayoutTemplStep(), sql.OrderByField(field, opts...))
	}
}

// ByWsapceInviteTemplField orders the results by wsapce_invite_templ field.
func ByWsapceInviteTemplField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWsapceInviteTemplStep(), sql.OrderByField(field, opts...))
	}
}

// ByWsapceSuccessTemplField orders the results by wsapce_success_templ field.
func ByWsapceSuccessTemplField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWsapceSuccessTemplStep(), sql.OrderByField(field, opts...))
	}
}

// ByAuthFpTemplField orders the results by auth_fp_templ field.
func ByAuthFpTemplField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthFpTemplStep(), sql.OrderByField(field, opts...))
	}
}

// ByAuthWelcomeEmailTemplField orders the results by auth_welcome_email_templ field.
func ByAuthWelcomeEmailTemplField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthWelcomeEmailTemplStep(), sql.OrderByField(field, opts...))
	}
}

// ByAuthVerificationTemplField orders the results by auth_verification_templ field.
func ByAuthVerificationTemplField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthVerificationTemplStep(), sql.OrderByField(field, opts...))
	}
}

// ByAdminUserCount orders the results by admin_user count.
func ByAdminUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdminUserStep(), opts...)
	}
}

// ByAdminUser orders the results by admin_user terms.
func ByAdminUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdminUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAppUsersCount orders the results by app_users count.
func ByAppUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppUsersStep(), opts...)
	}
}

// ByAppUsers orders the results by app_users terms.
func ByAppUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDefaultMailConnStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DefaultMailConnInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DefaultMailConnTable, DefaultMailConnColumn),
	)
}
func newMailLayoutTemplStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MailLayoutTemplInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MailLayoutTemplTable, MailLayoutTemplColumn),
	)
}
func newWsapceInviteTemplStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WsapceInviteTemplInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, WsapceInviteTemplTable, WsapceInviteTemplColumn),
	)
}
func newWsapceSuccessTemplStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WsapceSuccessTemplInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, WsapceSuccessTemplTable, WsapceSuccessTemplColumn),
	)
}
func newAuthFpTemplStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthFpTemplInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AuthFpTemplTable, AuthFpTemplColumn),
	)
}
func newAuthWelcomeEmailTemplStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthWelcomeEmailTemplInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AuthWelcomeEmailTemplTable, AuthWelcomeEmailTemplColumn),
	)
}
func newAuthVerificationTemplStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthVerificationTemplInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AuthVerificationTemplTable, AuthVerificationTemplColumn),
	)
}
func newAdminUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdminUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AdminUserTable, AdminUserPrimaryKey...),
	)
}
func newAppUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AppUsersTable, AppUsersColumn),
	)
}
