// Code generated by ent, DO NOT EDIT.

package posttype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the posttype type in the database.
	Label = "post_type"
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
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
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
	// EdgePostStatuses holds the string denoting the post_statuses edge name in mutations.
	EdgePostStatuses = "post_statuses"
	// EdgePostTypeForms holds the string denoting the post_type_forms edge name in mutations.
	EdgePostTypeForms = "post_type_forms"
	// Table holds the table name of the posttype in the database.
	Table = "post_types"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "post_type_id"
	// PostStatusesTable is the table that holds the post_statuses relation/edge.
	PostStatusesTable = "post_status"
	// PostStatusesInverseTable is the table name for the PostStatus entity.
	// It exists in this package in order to avoid circular dependency with the "poststatus" package.
	PostStatusesInverseTable = "post_status"
	// PostStatusesColumn is the table column denoting the post_statuses relation/edge.
	PostStatusesColumn = "post_type_id"
	// PostTypeFormsTable is the table that holds the post_type_forms relation/edge.
	PostTypeFormsTable = "post_type_forms"
	// PostTypeFormsInverseTable is the table name for the PostTypeForm entity.
	// It exists in this package in order to avoid circular dependency with the "posttypeform" package.
	PostTypeFormsInverseTable = "post_type_forms"
	// PostTypeFormsColumn is the table column denoting the post_type_forms relation/edge.
	PostTypeFormsColumn = "post_type_id"
)

// Columns holds all SQL columns for posttype fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAppID,
	FieldName,
	FieldSlug,
	FieldStatus,
	FieldExcerpt,
	FieldContent,
	FieldMetaTitle,
	FieldMetaDescr,
	FieldMetaCanonicalURL,
	FieldMetaRobots,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the PostType queries.
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

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
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

// ByPostStatusesCount orders the results by post_statuses count.
func ByPostStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostStatusesStep(), opts...)
	}
}

// ByPostStatuses orders the results by post_statuses terms.
func ByPostStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostTypeFormsCount orders the results by post_type_forms count.
func ByPostTypeFormsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostTypeFormsStep(), opts...)
	}
}

// ByPostTypeForms orders the results by post_type_forms terms.
func ByPostTypeForms(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostTypeFormsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
	)
}
func newPostStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostStatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostStatusesTable, PostStatusesColumn),
	)
}
func newPostTypeFormsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostTypeFormsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PostTypeFormsTable, PostTypeFormsColumn),
	)
}
