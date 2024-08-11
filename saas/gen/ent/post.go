// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/post"
	"saas/gen/ent/postcategory"
	"saas/gen/ent/poststatus"
	"saas/gen/ent/posttype"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// PostStatusID holds the value of the "post_status_id" field.
	PostStatusID string `json:"post_status_id,omitempty"`
	// PostTypeID holds the value of the "post_type_id" field.
	PostTypeID string `json:"post_type_id,omitempty"`
	// PrimaryCategoryID holds the value of the "primary_category_id" field.
	PrimaryCategoryID string `json:"primary_category_id,omitempty"`
	// Headline holds the value of the "headline" field.
	Headline string `json:"headline,omitempty"`
	// Excerpt holds the value of the "excerpt" field.
	Excerpt string `json:"excerpt,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// MetaTitle holds the value of the "meta_title" field.
	MetaTitle string `json:"meta_title,omitempty"`
	// MetaDescr holds the value of the "meta_descr" field.
	MetaDescr string `json:"meta_descr,omitempty"`
	// MetaCanonicalURL holds the value of the "meta_canonical_url" field.
	MetaCanonicalURL string `json:"meta_canonical_url,omitempty"`
	// MetaRobots holds the value of the "meta_robots" field.
	MetaRobots string `json:"meta_robots,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges        PostEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// PostStatus holds the value of the post_status edge.
	PostStatus *PostStatus `json:"post_status,omitempty"`
	// PostType holds the value of the post_type edge.
	PostType *PostType `json:"post_type,omitempty"`
	// PrimaryCategory holds the value of the primary_category edge.
	PrimaryCategory *PostCategory `json:"primary_category,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// PostStatusOrErr returns the PostStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) PostStatusOrErr() (*PostStatus, error) {
	if e.loadedTypes[0] {
		if e.PostStatus == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: poststatus.Label}
		}
		return e.PostStatus, nil
	}
	return nil, &NotLoadedError{edge: "post_status"}
}

// PostTypeOrErr returns the PostType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) PostTypeOrErr() (*PostType, error) {
	if e.loadedTypes[1] {
		if e.PostType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: posttype.Label}
		}
		return e.PostType, nil
	}
	return nil, &NotLoadedError{edge: "post_type"}
}

// PrimaryCategoryOrErr returns the PrimaryCategory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) PrimaryCategoryOrErr() (*PostCategory, error) {
	if e.loadedTypes[2] {
		if e.PrimaryCategory == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: postcategory.Label}
		}
		return e.PrimaryCategory, nil
	}
	return nil, &NotLoadedError{edge: "primary_category"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldID, post.FieldName, post.FieldSlug, post.FieldPostStatusID, post.FieldPostTypeID, post.FieldPrimaryCategoryID, post.FieldHeadline, post.FieldExcerpt, post.FieldContent, post.FieldMetaTitle, post.FieldMetaDescr, post.FieldMetaCanonicalURL, post.FieldMetaRobots:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				po.ID = value.String
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case post.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				po.Name = value.String
			}
		case post.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				po.Slug = value.String
			}
		case post.FieldPostStatusID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_status_id", values[i])
			} else if value.Valid {
				po.PostStatusID = value.String
			}
		case post.FieldPostTypeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_type_id", values[i])
			} else if value.Valid {
				po.PostTypeID = value.String
			}
		case post.FieldPrimaryCategoryID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_category_id", values[i])
			} else if value.Valid {
				po.PrimaryCategoryID = value.String
			}
		case post.FieldHeadline:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field headline", values[i])
			} else if value.Valid {
				po.Headline = value.String
			}
		case post.FieldExcerpt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field excerpt", values[i])
			} else if value.Valid {
				po.Excerpt = value.String
			}
		case post.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				po.Content = value.String
			}
		case post.FieldMetaTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_title", values[i])
			} else if value.Valid {
				po.MetaTitle = value.String
			}
		case post.FieldMetaDescr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_descr", values[i])
			} else if value.Valid {
				po.MetaDescr = value.String
			}
		case post.FieldMetaCanonicalURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_canonical_url", values[i])
			} else if value.Valid {
				po.MetaCanonicalURL = value.String
			}
		case post.FieldMetaRobots:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_robots", values[i])
			} else if value.Valid {
				po.MetaRobots = value.String
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Post.
// This includes values selected through modifiers, order, etc.
func (po *Post) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryPostStatus queries the "post_status" edge of the Post entity.
func (po *Post) QueryPostStatus() *PostStatusQuery {
	return NewPostClient(po.config).QueryPostStatus(po)
}

// QueryPostType queries the "post_type" edge of the Post entity.
func (po *Post) QueryPostType() *PostTypeQuery {
	return NewPostClient(po.config).QueryPostType(po)
}

// QueryPrimaryCategory queries the "primary_category" edge of the Post entity.
func (po *Post) QueryPrimaryCategory() *PostCategoryQuery {
	return NewPostClient(po.config).QueryPrimaryCategory(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return NewPostClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(po.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(po.Slug)
	builder.WriteString(", ")
	builder.WriteString("post_status_id=")
	builder.WriteString(po.PostStatusID)
	builder.WriteString(", ")
	builder.WriteString("post_type_id=")
	builder.WriteString(po.PostTypeID)
	builder.WriteString(", ")
	builder.WriteString("primary_category_id=")
	builder.WriteString(po.PrimaryCategoryID)
	builder.WriteString(", ")
	builder.WriteString("headline=")
	builder.WriteString(po.Headline)
	builder.WriteString(", ")
	builder.WriteString("excerpt=")
	builder.WriteString(po.Excerpt)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(po.Content)
	builder.WriteString(", ")
	builder.WriteString("meta_title=")
	builder.WriteString(po.MetaTitle)
	builder.WriteString(", ")
	builder.WriteString("meta_descr=")
	builder.WriteString(po.MetaDescr)
	builder.WriteString(", ")
	builder.WriteString("meta_canonical_url=")
	builder.WriteString(po.MetaCanonicalURL)
	builder.WriteString(", ")
	builder.WriteString("meta_robots=")
	builder.WriteString(po.MetaRobots)
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post