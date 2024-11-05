// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/postcategory"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PostCategory is the model entity for the PostCategory schema.
type PostCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID string `json:"app_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
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
	// The values are being populated by the PostCategoryQuery when eager-loading is set.
	Edges        PostCategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PostCategoryEdges holds the relations/edges for other nodes in the graph.
type PostCategoryEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedPosts map[string][]*Post
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e PostCategoryEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PostCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case postcategory.FieldID, postcategory.FieldAppID, postcategory.FieldName, postcategory.FieldSlug, postcategory.FieldStatus, postcategory.FieldExcerpt, postcategory.FieldContent, postcategory.FieldMetaTitle, postcategory.FieldMetaDescr, postcategory.FieldMetaCanonicalURL, postcategory.FieldMetaRobots:
			values[i] = new(sql.NullString)
		case postcategory.FieldCreatedAt, postcategory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PostCategory fields.
func (pc *PostCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case postcategory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pc.ID = value.String
			}
		case postcategory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pc.CreatedAt = value.Time
			}
		case postcategory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pc.UpdatedAt = value.Time
			}
		case postcategory.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				pc.AppID = value.String
			}
		case postcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pc.Name = value.String
			}
		case postcategory.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				pc.Slug = value.String
			}
		case postcategory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pc.Status = value.String
			}
		case postcategory.FieldExcerpt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field excerpt", values[i])
			} else if value.Valid {
				pc.Excerpt = value.String
			}
		case postcategory.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pc.Content = value.String
			}
		case postcategory.FieldMetaTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_title", values[i])
			} else if value.Valid {
				pc.MetaTitle = value.String
			}
		case postcategory.FieldMetaDescr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_descr", values[i])
			} else if value.Valid {
				pc.MetaDescr = value.String
			}
		case postcategory.FieldMetaCanonicalURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_canonical_url", values[i])
			} else if value.Valid {
				pc.MetaCanonicalURL = value.String
			}
		case postcategory.FieldMetaRobots:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_robots", values[i])
			} else if value.Valid {
				pc.MetaRobots = value.String
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PostCategory.
// This includes values selected through modifiers, order, etc.
func (pc *PostCategory) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// QueryPosts queries the "posts" edge of the PostCategory entity.
func (pc *PostCategory) QueryPosts() *PostQuery {
	return NewPostCategoryClient(pc.config).QueryPosts(pc)
}

// Update returns a builder for updating this PostCategory.
// Note that you need to call PostCategory.Unwrap() before calling this method if this PostCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PostCategory) Update() *PostCategoryUpdateOne {
	return NewPostCategoryClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the PostCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PostCategory) Unwrap() *PostCategory {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PostCategory is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PostCategory) String() string {
	var builder strings.Builder
	builder.WriteString("PostCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(pc.AppID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pc.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(pc.Slug)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(pc.Status)
	builder.WriteString(", ")
	builder.WriteString("excerpt=")
	builder.WriteString(pc.Excerpt)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pc.Content)
	builder.WriteString(", ")
	builder.WriteString("meta_title=")
	builder.WriteString(pc.MetaTitle)
	builder.WriteString(", ")
	builder.WriteString("meta_descr=")
	builder.WriteString(pc.MetaDescr)
	builder.WriteString(", ")
	builder.WriteString("meta_canonical_url=")
	builder.WriteString(pc.MetaCanonicalURL)
	builder.WriteString(", ")
	builder.WriteString("meta_robots=")
	builder.WriteString(pc.MetaRobots)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPosts returns the Posts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pc *PostCategory) NamedPosts(name string) ([]*Post, error) {
	if pc.Edges.namedPosts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pc.Edges.namedPosts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pc *PostCategory) appendNamedPosts(name string, edges ...*Post) {
	if pc.Edges.namedPosts == nil {
		pc.Edges.namedPosts = make(map[string][]*Post)
	}
	if len(edges) == 0 {
		pc.Edges.namedPosts[name] = []*Post{}
	} else {
		pc.Edges.namedPosts[name] = append(pc.Edges.namedPosts[name], edges...)
	}
}

// PostCategories is a parsable slice of PostCategory.
type PostCategories []*PostCategory
