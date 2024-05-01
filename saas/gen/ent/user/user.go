// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldAPIKey holds the string denoting the api_key field in the database.
	FieldAPIKey = "api_key"
	// FieldWelcomeEmailSent holds the string denoting the welcome_email_sent field in the database.
	FieldWelcomeEmailSent = "welcome_email_sent"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// EdgeWorkspaces holds the string denoting the workspaces edge name in mutations.
	EdgeWorkspaces = "workspaces"
	// EdgeWorkspaceUsers holds the string denoting the workspace_users edge name in mutations.
	EdgeWorkspaceUsers = "workspace_users"
	// Table holds the table name of the user in the database.
	Table = "users"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "user_id"
	// WorkspacesTable is the table that holds the workspaces relation/edge. The primary key declared below.
	WorkspacesTable = "workspace_users"
	// WorkspacesInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspacesInverseTable = "workspaces"
	// WorkspaceUsersTable is the table that holds the workspace_users relation/edge.
	WorkspaceUsersTable = "workspace_users"
	// WorkspaceUsersInverseTable is the table name for the WorkspaceUser entity.
	// It exists in this package in order to avoid circular dependency with the "workspaceuser" package.
	WorkspaceUsersInverseTable = "workspace_users"
	// WorkspaceUsersColumn is the table column denoting the workspace_users relation/edge.
	WorkspaceUsersColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmail,
	FieldPhone,
	FieldFirstName,
	FieldLastName,
	FieldCompany,
	FieldStatus,
	FieldPassword,
	FieldSecret,
	FieldAPIKey,
	FieldWelcomeEmailSent,
}

var (
	// WorkspacesPrimaryKey and WorkspacesColumn2 are the table columns denoting the
	// primary key for the workspaces relation (M2M).
	WorkspacesPrimaryKey = []string{"workspace_id", "user_id"}
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the User queries.
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

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByCompany orders the results by the company field.
func ByCompany(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompany, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// BySecret orders the results by the secret field.
func BySecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecret, opts...).ToFunc()
}

// ByAPIKey orders the results by the api_key field.
func ByAPIKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAPIKey, opts...).ToFunc()
}

// ByWelcomeEmailSent orders the results by the welcome_email_sent field.
func ByWelcomeEmailSent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWelcomeEmailSent, opts...).ToFunc()
}

// BySessionsCount orders the results by sessions count.
func BySessionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSessionsStep(), opts...)
	}
}

// BySessions orders the results by sessions terms.
func BySessions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSessionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkspacesCount orders the results by workspaces count.
func ByWorkspacesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkspacesStep(), opts...)
	}
}

// ByWorkspaces orders the results by workspaces terms.
func ByWorkspaces(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkspacesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkspaceUsersCount orders the results by workspace_users count.
func ByWorkspaceUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkspaceUsersStep(), opts...)
	}
}

// ByWorkspaceUsers orders the results by workspace_users terms.
func ByWorkspaceUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkspaceUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSessionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SessionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
	)
}
func newWorkspacesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkspacesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WorkspacesTable, WorkspacesPrimaryKey...),
	)
}
func newWorkspaceUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkspaceUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, WorkspaceUsersTable, WorkspaceUsersColumn),
	)
}
