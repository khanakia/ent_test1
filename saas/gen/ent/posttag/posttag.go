// Code generated by ent, DO NOT EDIT.

package posttag

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the posttag type in the database.
	Label = "post_tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldExcerpt holds the string denoting the excerpt field in the database.
	FieldExcerpt = "excerpt"
	// FieldMetaTitle holds the string denoting the meta_title field in the database.
	FieldMetaTitle = "meta_title"
	// FieldMetaDescr holds the string denoting the meta_descr field in the database.
	FieldMetaDescr = "meta_descr"
	// FieldMetaCanonicalURL holds the string denoting the meta_canonical_url field in the database.
	FieldMetaCanonicalURL = "meta_canonical_url"
	// FieldMetaRobots holds the string denoting the meta_robots field in the database.
	FieldMetaRobots = "meta_robots"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the posttag in the database.
	Table = "post_tags"
	// PostsTable is the table that holds the posts relation/edge. The primary key declared below.
	PostsTable = "post_tag_posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
)

// Columns holds all SQL columns for posttag fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAppID,
	FieldName,
	FieldSlug,
	FieldStatus,
	FieldExcerpt,
	FieldMetaTitle,
	FieldMetaDescr,
	FieldMetaCanonicalURL,
	FieldMetaRobots,
}

var (
	// PostsPrimaryKey and PostsColumn2 are the table columns denoting the
	// primary key for the posts relation (M2M).
	PostsPrimaryKey = []string{"post_tag_id", "post_id"}
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "saas/gen/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the PostTag queries.
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

// ByAppID orders the results by the app_id field.
func ByAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByExcerpt orders the results by the excerpt field.
func ByExcerpt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExcerpt, opts...).ToFunc()
}

// ByMetaTitle orders the results by the meta_title field.
func ByMetaTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMetaTitle, opts...).ToFunc()
}

// ByMetaDescr orders the results by the meta_descr field.
func ByMetaDescr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMetaDescr, opts...).ToFunc()
}

// ByMetaCanonicalURL orders the results by the meta_canonical_url field.
func ByMetaCanonicalURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMetaCanonicalURL, opts...).ToFunc()
}

// ByMetaRobots orders the results by the meta_robots field.
func ByMetaRobots(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMetaRobots, opts...).ToFunc()
}

// ByPostsCount orders the results by posts count.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

// ByPosts orders the results by posts terms.
func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PostsTable, PostsPrimaryKey...),
	)
}
