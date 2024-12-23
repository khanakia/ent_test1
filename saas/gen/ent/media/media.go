// Code generated by ent, DO NOT EDIT.

package media

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldDisk holds the string denoting the disk field in the database.
	FieldDisk = "disk"
	// FieldDirectory holds the string denoting the directory field in the database.
	FieldDirectory = "directory"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOriginalName holds the string denoting the original_name field in the database.
	FieldOriginalName = "original_name"
	// FieldExtension holds the string denoting the extension field in the database.
	FieldExtension = "extension"
	// FieldMimeType holds the string denoting the mime_type field in the database.
	FieldMimeType = "mime_type"
	// FieldAggregateType holds the string denoting the aggregate_type field in the database.
	FieldAggregateType = "aggregate_type"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsVariant holds the string denoting the is_variant field in the database.
	FieldIsVariant = "is_variant"
	// FieldVariantName holds the string denoting the variant_name field in the database.
	FieldVariantName = "variant_name"
	// FieldOriginalMediaID holds the string denoting the original_media_id field in the database.
	FieldOriginalMediaID = "original_media_id"
	// FieldChecksum holds the string denoting the checksum field in the database.
	FieldChecksum = "checksum"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldAlt holds the string denoting the alt field in the database.
	FieldAlt = "alt"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeMediables holds the string denoting the mediables edge name in mutations.
	EdgeMediables = "mediables"
	// Table holds the table name of the media in the database.
	Table = "media"
	// MediablesTable is the table that holds the mediables relation/edge.
	MediablesTable = "mediables"
	// MediablesInverseTable is the table name for the Mediable entity.
	// It exists in this package in order to avoid circular dependency with the "mediable" package.
	MediablesInverseTable = "mediables"
	// MediablesColumn is the table column denoting the mediables relation/edge.
	MediablesColumn = "media_id"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldDisk,
	FieldDirectory,
	FieldName,
	FieldOriginalName,
	FieldExtension,
	FieldMimeType,
	FieldAggregateType,
	FieldSize,
	FieldDescription,
	FieldIsVariant,
	FieldVariantName,
	FieldOriginalMediaID,
	FieldChecksum,
	FieldWorkspaceID,
	FieldAlt,
	FieldUID,
	FieldStatus,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Media queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAppID orders the results by the app_id field.
func ByAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppID, opts...).ToFunc()
}

// ByDisk orders the results by the disk field.
func ByDisk(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisk, opts...).ToFunc()
}

// ByDirectory orders the results by the directory field.
func ByDirectory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDirectory, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOriginalName orders the results by the original_name field.
func ByOriginalName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginalName, opts...).ToFunc()
}

// ByExtension orders the results by the extension field.
func ByExtension(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtension, opts...).ToFunc()
}

// ByMimeType orders the results by the mime_type field.
func ByMimeType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMimeType, opts...).ToFunc()
}

// ByAggregateType orders the results by the aggregate_type field.
func ByAggregateType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAggregateType, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIsVariant orders the results by the is_variant field.
func ByIsVariant(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsVariant, opts...).ToFunc()
}

// ByVariantName orders the results by the variant_name field.
func ByVariantName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVariantName, opts...).ToFunc()
}

// ByOriginalMediaID orders the results by the original_media_id field.
func ByOriginalMediaID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginalMediaID, opts...).ToFunc()
}

// ByChecksum orders the results by the checksum field.
func ByChecksum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChecksum, opts...).ToFunc()
}

// ByWorkspaceID orders the results by the workspace_id field.
func ByWorkspaceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkspaceID, opts...).ToFunc()
}

// ByAlt orders the results by the alt field.
func ByAlt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlt, opts...).ToFunc()
}

// ByUID orders the results by the uid field.
func ByUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByMediablesCount orders the results by mediables count.
func ByMediablesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMediablesStep(), opts...)
	}
}

// ByMediables orders the results by mediables terms.
func ByMediables(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediablesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMediablesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediablesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MediablesTable, MediablesColumn),
	)
}
