// Code generated by ent, DO NOT EDIT.

package posttag

import (
	"time"

	"entgo.io/ent/dialect/sql"
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
	// Table holds the table name of the posttag in the database.
	Table = "post_tags"
)

// Columns holds all SQL columns for posttag fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
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