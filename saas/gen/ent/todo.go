// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/todo"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Status holds the value of the "status" field.
	Status todo.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges        TodoEdges `json:"edges"`
	todo_parent  *string
	selectValues sql.SelectValues
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// Children holds the value of the children edge.
	Children []*Todo `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Todo `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedChildren map[string][]*Todo
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) ChildrenOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) ParentOrErr() (*Todo, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: todo.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldPriority:
			values[i] = new(sql.NullInt64)
		case todo.FieldID, todo.FieldText, todo.FieldStatus:
			values[i] = new(sql.NullString)
		case todo.FieldCreatedAt, todo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case todo.ForeignKeys[0]: // todo_parent
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case todo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case todo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case todo.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		case todo.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = todo.Status(value.String)
			}
		case todo.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				t.Priority = int(value.Int64)
			}
		case todo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field todo_parent", values[i])
			} else if value.Valid {
				t.todo_parent = new(string)
				*t.todo_parent = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Todo.
// This includes values selected through modifiers, order, etc.
func (t *Todo) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryChildren queries the "children" edge of the Todo entity.
func (t *Todo) QueryChildren() *TodoQuery {
	return NewTodoClient(t.config).QueryChildren(t)
}

// QueryParent queries the "parent" edge of the Todo entity.
func (t *Todo) QueryParent() *TodoQuery {
	return NewTodoClient(t.config).QueryParent(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return NewTodoClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(t.Text)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteByte(')')
	return builder.String()
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Todo) NamedChildren(name string) ([]*Todo, error) {
	if t.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Todo) appendNamedChildren(name string, edges ...*Todo) {
	if t.Edges.namedChildren == nil {
		t.Edges.namedChildren = make(map[string][]*Todo)
	}
	if len(edges) == 0 {
		t.Edges.namedChildren[name] = []*Todo{}
	} else {
		t.Edges.namedChildren[name] = append(t.Edges.namedChildren[name], edges...)
	}
}

// Todos is a parsable slice of Todo.
type Todos []*Todo
