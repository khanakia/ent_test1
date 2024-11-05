// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
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
	// FieldPostStatusID holds the string denoting the post_status_id field in the database.
	FieldPostStatusID = "post_status_id"
	// FieldPostTypeID holds the string denoting the post_type_id field in the database.
	FieldPostTypeID = "post_type_id"
	// FieldPrimaryCategoryID holds the string denoting the primary_category_id field in the database.
	FieldPrimaryCategoryID = "primary_category_id"
	// FieldHeadline holds the string denoting the headline field in the database.
	FieldHeadline = "headline"
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
	// EdgePostStatus holds the string denoting the post_status edge name in mutations.
	EdgePostStatus = "post_status"
	// EdgePostType holds the string denoting the post_type edge name in mutations.
	EdgePostType = "post_type"
	// EdgePrimaryCategory holds the string denoting the primary_category edge name in mutations.
	EdgePrimaryCategory = "primary_category"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// PostStatusTable is the table that holds the post_status relation/edge.
	PostStatusTable = "posts"
	// PostStatusInverseTable is the table name for the PostStatus entity.
	// It exists in this package in order to avoid circular dependency with the "poststatus" package.
	PostStatusInverseTable = "post_status"
	// PostStatusColumn is the table column denoting the post_status relation/edge.
	PostStatusColumn = "post_status_id"
	// PostTypeTable is the table that holds the post_type relation/edge.
	PostTypeTable = "posts"
	// PostTypeInverseTable is the table name for the PostType entity.
	// It exists in this package in order to avoid circular dependency with the "posttype" package.
	PostTypeInverseTable = "post_types"
	// PostTypeColumn is the table column denoting the post_type relation/edge.
	PostTypeColumn = "post_type_id"
	// PrimaryCategoryTable is the table that holds the primary_category relation/edge.
	PrimaryCategoryTable = "posts"
	// PrimaryCategoryInverseTable is the table name for the PostCategory entity.
	// It exists in this package in order to avoid circular dependency with the "postcategory" package.
	PrimaryCategoryInverseTable = "post_categories"
	// PrimaryCategoryColumn is the table column denoting the primary_category relation/edge.
	PrimaryCategoryColumn = "primary_category_id"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAppID,
	FieldName,
	FieldSlug,
	FieldPostStatusID,
	FieldPostTypeID,
	FieldPrimaryCategoryID,
	FieldHeadline,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Post queries.
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

// ByPostStatusID orders the results by the post_status_id field.
func ByPostStatusID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostStatusID, opts...).ToFunc()
}

// ByPostTypeID orders the results by the post_type_id field.
func ByPostTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostTypeID, opts...).ToFunc()
}

// ByPrimaryCategoryID orders the results by the primary_category_id field.
func ByPrimaryCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrimaryCategoryID, opts...).ToFunc()
}

// ByHeadline orders the results by the headline field.
func ByHeadline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeadline, opts...).ToFunc()
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

// ByPostStatusField orders the results by post_status field.
func ByPostStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStatusStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostTypeField orders the results by post_type field.
func ByPostTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByPrimaryCategoryField orders the results by primary_category field.
func ByPrimaryCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrimaryCategoryStep(), sql.OrderByField(field, opts...))
	}
}
func newPostStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostStatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostStatusTable, PostStatusColumn),
	)
}
func newPostTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PostTypeTable, PostTypeColumn),
	)
}
func newPrimaryCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrimaryCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PrimaryCategoryTable, PrimaryCategoryColumn),
	)
}
