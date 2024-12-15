// Code generated by ent, DO NOT EDIT.

package approleperm

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the approleperm type in the database.
	Label = "app_role_perm"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAppRoleID holds the string denoting the app_role_id field in the database.
	FieldAppRoleID = "app_role_id"
	// FieldAppPermID holds the string denoting the app_perm_id field in the database.
	FieldAppPermID = "app_perm_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeAppPerm holds the string denoting the app_perm edge name in mutations.
	EdgeAppPerm = "app_perm"
	// EdgeAppRole holds the string denoting the app_role edge name in mutations.
	EdgeAppRole = "app_role"
	// Table holds the table name of the approleperm in the database.
	Table = "app_role_perms"
	// AppTable is the table that holds the app relation/edge.
	AppTable = "app_role_perms"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "apps"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "app_id"
	// AppPermTable is the table that holds the app_perm relation/edge.
	AppPermTable = "app_role_perms"
	// AppPermInverseTable is the table name for the AppPerm entity.
	// It exists in this package in order to avoid circular dependency with the "appperm" package.
	AppPermInverseTable = "app_perms"
	// AppPermColumn is the table column denoting the app_perm relation/edge.
	AppPermColumn = "app_perm_id"
	// AppRoleTable is the table that holds the app_role relation/edge.
	AppRoleTable = "app_role_perms"
	// AppRoleInverseTable is the table name for the AppRole entity.
	// It exists in this package in order to avoid circular dependency with the "approle" package.
	AppRoleInverseTable = "app_roles"
	// AppRoleColumn is the table column denoting the app_role relation/edge.
	AppRoleColumn = "app_role_id"
)

// Columns holds all SQL columns for approleperm fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAppRoleID,
	FieldAppPermID,
	FieldAppID,
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

// OrderOption defines the ordering options for the AppRolePerm queries.
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

// ByAppRoleID orders the results by the app_role_id field.
func ByAppRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppRoleID, opts...).ToFunc()
}

// ByAppPermID orders the results by the app_perm_id field.
func ByAppPermID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppPermID, opts...).ToFunc()
}

// ByAppID orders the results by the app_id field.
func ByAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppID, opts...).ToFunc()
}

// ByAppField orders the results by app field.
func ByAppField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppStep(), sql.OrderByField(field, opts...))
	}
}

// ByAppPermField orders the results by app_perm field.
func ByAppPermField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppPermStep(), sql.OrderByField(field, opts...))
	}
}

// ByAppRoleField orders the results by app_role field.
func ByAppRoleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppRoleStep(), sql.OrderByField(field, opts...))
	}
}
func newAppStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AppTable, AppColumn),
	)
}
func newAppPermStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppPermInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AppPermTable, AppPermColumn),
	)
}
func newAppRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppRoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AppRoleTable, AppRoleColumn),
	)
}