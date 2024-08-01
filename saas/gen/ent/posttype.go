// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/posttype"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PostType is the model entity for the PostType schema.
type PostType struct {
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
	// Status holds the value of the "status" field.
	Status posttype.Status `json:"status,omitempty"`
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
	// The values are being populated by the PostTypeQuery when eager-loading is set.
	Edges        PostTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PostTypeEdges holds the relations/edges for other nodes in the graph.
type PostTypeEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// PostStatuses holds the value of the post_statuses edge.
	PostStatuses []*PostStatus `json:"post_statuses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedPosts        map[string][]*Post
	namedPostStatuses map[string][]*PostStatus
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e PostTypeEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// PostStatusesOrErr returns the PostStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e PostTypeEdges) PostStatusesOrErr() ([]*PostStatus, error) {
	if e.loadedTypes[1] {
		return e.PostStatuses, nil
	}
	return nil, &NotLoadedError{edge: "post_statuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PostType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case posttype.FieldID, posttype.FieldName, posttype.FieldSlug, posttype.FieldStatus, posttype.FieldExcerpt, posttype.FieldContent, posttype.FieldMetaTitle, posttype.FieldMetaDescr, posttype.FieldMetaCanonicalURL, posttype.FieldMetaRobots:
			values[i] = new(sql.NullString)
		case posttype.FieldCreatedAt, posttype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PostType fields.
func (pt *PostType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case posttype.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pt.ID = value.String
			}
		case posttype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = value.Time
			}
		case posttype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pt.UpdatedAt = value.Time
			}
		case posttype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pt.Name = value.String
			}
		case posttype.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				pt.Slug = value.String
			}
		case posttype.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pt.Status = posttype.Status(value.String)
			}
		case posttype.FieldExcerpt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field excerpt", values[i])
			} else if value.Valid {
				pt.Excerpt = value.String
			}
		case posttype.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pt.Content = value.String
			}
		case posttype.FieldMetaTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_title", values[i])
			} else if value.Valid {
				pt.MetaTitle = value.String
			}
		case posttype.FieldMetaDescr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_descr", values[i])
			} else if value.Valid {
				pt.MetaDescr = value.String
			}
		case posttype.FieldMetaCanonicalURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_canonical_url", values[i])
			} else if value.Valid {
				pt.MetaCanonicalURL = value.String
			}
		case posttype.FieldMetaRobots:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta_robots", values[i])
			} else if value.Valid {
				pt.MetaRobots = value.String
			}
		default:
			pt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PostType.
// This includes values selected through modifiers, order, etc.
func (pt *PostType) Value(name string) (ent.Value, error) {
	return pt.selectValues.Get(name)
}

// QueryPosts queries the "posts" edge of the PostType entity.
func (pt *PostType) QueryPosts() *PostQuery {
	return NewPostTypeClient(pt.config).QueryPosts(pt)
}

// QueryPostStatuses queries the "post_statuses" edge of the PostType entity.
func (pt *PostType) QueryPostStatuses() *PostStatusQuery {
	return NewPostTypeClient(pt.config).QueryPostStatuses(pt)
}

// Update returns a builder for updating this PostType.
// Note that you need to call PostType.Unwrap() before calling this method if this PostType
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PostType) Update() *PostTypeUpdateOne {
	return NewPostTypeClient(pt.config).UpdateOne(pt)
}

// Unwrap unwraps the PostType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *PostType) Unwrap() *PostType {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PostType is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PostType) String() string {
	var builder strings.Builder
	builder.WriteString("PostType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pt.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(pt.Slug)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pt.Status))
	builder.WriteString(", ")
	builder.WriteString("excerpt=")
	builder.WriteString(pt.Excerpt)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pt.Content)
	builder.WriteString(", ")
	builder.WriteString("meta_title=")
	builder.WriteString(pt.MetaTitle)
	builder.WriteString(", ")
	builder.WriteString("meta_descr=")
	builder.WriteString(pt.MetaDescr)
	builder.WriteString(", ")
	builder.WriteString("meta_canonical_url=")
	builder.WriteString(pt.MetaCanonicalURL)
	builder.WriteString(", ")
	builder.WriteString("meta_robots=")
	builder.WriteString(pt.MetaRobots)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPosts returns the Posts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pt *PostType) NamedPosts(name string) ([]*Post, error) {
	if pt.Edges.namedPosts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pt.Edges.namedPosts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pt *PostType) appendNamedPosts(name string, edges ...*Post) {
	if pt.Edges.namedPosts == nil {
		pt.Edges.namedPosts = make(map[string][]*Post)
	}
	if len(edges) == 0 {
		pt.Edges.namedPosts[name] = []*Post{}
	} else {
		pt.Edges.namedPosts[name] = append(pt.Edges.namedPosts[name], edges...)
	}
}

// NamedPostStatuses returns the PostStatuses named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pt *PostType) NamedPostStatuses(name string) ([]*PostStatus, error) {
	if pt.Edges.namedPostStatuses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pt.Edges.namedPostStatuses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pt *PostType) appendNamedPostStatuses(name string, edges ...*PostStatus) {
	if pt.Edges.namedPostStatuses == nil {
		pt.Edges.namedPostStatuses = make(map[string][]*PostStatus)
	}
	if len(edges) == 0 {
		pt.Edges.namedPostStatuses[name] = []*PostStatus{}
	} else {
		pt.Edges.namedPostStatuses[name] = append(pt.Edges.namedPostStatuses[name], edges...)
	}
}

// PostTypes is a parsable slice of PostType.
type PostTypes []*PostType
