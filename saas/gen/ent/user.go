// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/gen/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Company holds the value of the "company" field.
	Company string `json:"company,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale string `json:"locale,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID string `json:"role_id,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"-"`
	// APIKey holds the value of the "api_key" field.
	APIKey string `json:"api_key,omitempty"`
	// WelcomeEmailSent holds the value of the "welcome_email_sent" field.
	WelcomeEmailSent bool `json:"welcome_email_sent,omitempty"`
	// CanAdmin holds the value of the "can_admin" field.
	CanAdmin bool `json:"can_admin,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Sessions holds the value of the sessions edge.
	Sessions []*Session `json:"sessions,omitempty"`
	// Workspaces holds the value of the workspaces edge.
	Workspaces []*Workspace `json:"workspaces,omitempty"`
	// WorkspaceUsers holds the value of the workspace_users edge.
	WorkspaceUsers []*WorkspaceUser `json:"workspace_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedSessions       map[string][]*Session
	namedWorkspaces     map[string][]*Workspace
	namedWorkspaceUsers map[string][]*WorkspaceUser
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[0] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// WorkspacesOrErr returns the Workspaces value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WorkspacesOrErr() ([]*Workspace, error) {
	if e.loadedTypes[1] {
		return e.Workspaces, nil
	}
	return nil, &NotLoadedError{edge: "workspaces"}
}

// WorkspaceUsersOrErr returns the WorkspaceUsers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WorkspaceUsersOrErr() ([]*WorkspaceUser, error) {
	if e.loadedTypes[2] {
		return e.WorkspaceUsers, nil
	}
	return nil, &NotLoadedError{edge: "workspace_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldStatus, user.FieldWelcomeEmailSent, user.FieldCanAdmin:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldEmail, user.FieldPhone, user.FieldFirstName, user.FieldLastName, user.FieldCompany, user.FieldLocale, user.FieldRoleID, user.FieldPassword, user.FieldSecret, user.FieldAPIKey:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldCompany:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company", values[i])
			} else if value.Valid {
				u.Company = value.String
			}
		case user.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				u.Locale = value.String
			}
		case user.FieldRoleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				u.RoleID = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = value.Bool
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				u.Secret = value.String
			}
		case user.FieldAPIKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_key", values[i])
			} else if value.Valid {
				u.APIKey = value.String
			}
		case user.FieldWelcomeEmailSent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field welcome_email_sent", values[i])
			} else if value.Valid {
				u.WelcomeEmailSent = value.Bool
			}
		case user.FieldCanAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_admin", values[i])
			} else if value.Valid {
				u.CanAdmin = value.Bool
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySessions queries the "sessions" edge of the User entity.
func (u *User) QuerySessions() *SessionQuery {
	return NewUserClient(u.config).QuerySessions(u)
}

// QueryWorkspaces queries the "workspaces" edge of the User entity.
func (u *User) QueryWorkspaces() *WorkspaceQuery {
	return NewUserClient(u.config).QueryWorkspaces(u)
}

// QueryWorkspaceUsers queries the "workspace_users" edge of the User entity.
func (u *User) QueryWorkspaceUsers() *WorkspaceUserQuery {
	return NewUserClient(u.config).QueryWorkspaceUsers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("company=")
	builder.WriteString(u.Company)
	builder.WriteString(", ")
	builder.WriteString("locale=")
	builder.WriteString(u.Locale)
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(u.RoleID)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("secret=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("api_key=")
	builder.WriteString(u.APIKey)
	builder.WriteString(", ")
	builder.WriteString("welcome_email_sent=")
	builder.WriteString(fmt.Sprintf("%v", u.WelcomeEmailSent))
	builder.WriteString(", ")
	builder.WriteString("can_admin=")
	builder.WriteString(fmt.Sprintf("%v", u.CanAdmin))
	builder.WriteByte(')')
	return builder.String()
}

// NamedSessions returns the Sessions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedSessions(name string) ([]*Session, error) {
	if u.Edges.namedSessions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedSessions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedSessions(name string, edges ...*Session) {
	if u.Edges.namedSessions == nil {
		u.Edges.namedSessions = make(map[string][]*Session)
	}
	if len(edges) == 0 {
		u.Edges.namedSessions[name] = []*Session{}
	} else {
		u.Edges.namedSessions[name] = append(u.Edges.namedSessions[name], edges...)
	}
}

// NamedWorkspaces returns the Workspaces named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedWorkspaces(name string) ([]*Workspace, error) {
	if u.Edges.namedWorkspaces == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedWorkspaces[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedWorkspaces(name string, edges ...*Workspace) {
	if u.Edges.namedWorkspaces == nil {
		u.Edges.namedWorkspaces = make(map[string][]*Workspace)
	}
	if len(edges) == 0 {
		u.Edges.namedWorkspaces[name] = []*Workspace{}
	} else {
		u.Edges.namedWorkspaces[name] = append(u.Edges.namedWorkspaces[name], edges...)
	}
}

// NamedWorkspaceUsers returns the WorkspaceUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedWorkspaceUsers(name string) ([]*WorkspaceUser, error) {
	if u.Edges.namedWorkspaceUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedWorkspaceUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedWorkspaceUsers(name string, edges ...*WorkspaceUser) {
	if u.Edges.namedWorkspaceUsers == nil {
		u.Edges.namedWorkspaceUsers = make(map[string][]*WorkspaceUser)
	}
	if len(edges) == 0 {
		u.Edges.namedWorkspaceUsers[name] = []*WorkspaceUser{}
	} else {
		u.Edges.namedWorkspaceUsers[name] = append(u.Edges.namedWorkspaceUsers[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
