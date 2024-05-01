// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/predicate"
	"saas/gen/ent/session"
	"saas/gen/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks     []Hook
	mutation  *SessionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SessionUpdate) SetCreatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCreatedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// ClearCreatedAt clears the value of the "created_at" field.
func (su *SessionUpdate) ClearCreatedAt() *SessionUpdate {
	su.mutation.ClearCreatedAt()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SessionUpdate) SetUpdatedAt(t time.Time) *SessionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (su *SessionUpdate) ClearUpdatedAt() *SessionUpdate {
	su.mutation.ClearUpdatedAt()
	return su
}

// SetUserID sets the "user_id" field.
func (su *SessionUpdate) SetUserID(s string) *SessionUpdate {
	su.mutation.SetUserID(s)
	return su
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserID(s *string) *SessionUpdate {
	if s != nil {
		su.SetUserID(*s)
	}
	return su
}

// ClearUserID clears the value of the "user_id" field.
func (su *SessionUpdate) ClearUserID() *SessionUpdate {
	su.mutation.ClearUserID()
	return su
}

// SetIP sets the "ip" field.
func (su *SessionUpdate) SetIP(s string) *SessionUpdate {
	su.mutation.SetIP(s)
	return su
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (su *SessionUpdate) SetNillableIP(s *string) *SessionUpdate {
	if s != nil {
		su.SetIP(*s)
	}
	return su
}

// ClearIP clears the value of the "ip" field.
func (su *SessionUpdate) ClearIP() *SessionUpdate {
	su.mutation.ClearIP()
	return su
}

// SetUserAgent sets the "user_agent" field.
func (su *SessionUpdate) SetUserAgent(s string) *SessionUpdate {
	su.mutation.SetUserAgent(s)
	return su
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (su *SessionUpdate) SetNillableUserAgent(s *string) *SessionUpdate {
	if s != nil {
		su.SetUserAgent(*s)
	}
	return su
}

// ClearUserAgent clears the value of the "user_agent" field.
func (su *SessionUpdate) ClearUserAgent() *SessionUpdate {
	su.mutation.ClearUserAgent()
	return su
}

// SetPayload sets the "payload" field.
func (su *SessionUpdate) SetPayload(s string) *SessionUpdate {
	su.mutation.SetPayload(s)
	return su
}

// SetNillablePayload sets the "payload" field if the given value is not nil.
func (su *SessionUpdate) SetNillablePayload(s *string) *SessionUpdate {
	if s != nil {
		su.SetPayload(*s)
	}
	return su
}

// ClearPayload clears the value of the "payload" field.
func (su *SessionUpdate) ClearPayload() *SessionUpdate {
	su.mutation.ClearPayload()
	return su
}

// SetExpiresAt sets the "expires_at" field.
func (su *SessionUpdate) SetExpiresAt(t time.Time) *SessionUpdate {
	su.mutation.SetExpiresAt(t)
	return su
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (su *SessionUpdate) SetNillableExpiresAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetExpiresAt(*t)
	}
	return su
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (su *SessionUpdate) ClearExpiresAt() *SessionUpdate {
	su.mutation.ClearExpiresAt()
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SessionUpdate) SetUser(u *User) *SessionUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SessionUpdate) ClearUser() *SessionUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SessionUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok && !su.mutation.UpdatedAtCleared() {
		v := session.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SessionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SessionUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if su.mutation.CreatedAtCleared() {
		_spec.ClearField(session.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.UpdatedAtCleared() {
		_spec.ClearField(session.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := su.mutation.IP(); ok {
		_spec.SetField(session.FieldIP, field.TypeString, value)
	}
	if su.mutation.IPCleared() {
		_spec.ClearField(session.FieldIP, field.TypeString)
	}
	if value, ok := su.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if su.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := su.mutation.Payload(); ok {
		_spec.SetField(session.FieldPayload, field.TypeString, value)
	}
	if su.mutation.PayloadCleared() {
		_spec.ClearField(session.FieldPayload, field.TypeString)
	}
	if value, ok := su.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	if su.mutation.ExpiresAtCleared() {
		_spec.ClearField(session.FieldExpiresAt, field.TypeTime)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SessionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (suo *SessionUpdateOne) SetCreatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCreatedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (suo *SessionUpdateOne) ClearCreatedAt() *SessionUpdateOne {
	suo.mutation.ClearCreatedAt()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SessionUpdateOne) SetUpdatedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suo *SessionUpdateOne) ClearUpdatedAt() *SessionUpdateOne {
	suo.mutation.ClearUpdatedAt()
	return suo
}

// SetUserID sets the "user_id" field.
func (suo *SessionUpdateOne) SetUserID(s string) *SessionUpdateOne {
	suo.mutation.SetUserID(s)
	return suo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserID(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetUserID(*s)
	}
	return suo
}

// ClearUserID clears the value of the "user_id" field.
func (suo *SessionUpdateOne) ClearUserID() *SessionUpdateOne {
	suo.mutation.ClearUserID()
	return suo
}

// SetIP sets the "ip" field.
func (suo *SessionUpdateOne) SetIP(s string) *SessionUpdateOne {
	suo.mutation.SetIP(s)
	return suo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableIP(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetIP(*s)
	}
	return suo
}

// ClearIP clears the value of the "ip" field.
func (suo *SessionUpdateOne) ClearIP() *SessionUpdateOne {
	suo.mutation.ClearIP()
	return suo
}

// SetUserAgent sets the "user_agent" field.
func (suo *SessionUpdateOne) SetUserAgent(s string) *SessionUpdateOne {
	suo.mutation.SetUserAgent(s)
	return suo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableUserAgent(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetUserAgent(*s)
	}
	return suo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (suo *SessionUpdateOne) ClearUserAgent() *SessionUpdateOne {
	suo.mutation.ClearUserAgent()
	return suo
}

// SetPayload sets the "payload" field.
func (suo *SessionUpdateOne) SetPayload(s string) *SessionUpdateOne {
	suo.mutation.SetPayload(s)
	return suo
}

// SetNillablePayload sets the "payload" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillablePayload(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetPayload(*s)
	}
	return suo
}

// ClearPayload clears the value of the "payload" field.
func (suo *SessionUpdateOne) ClearPayload() *SessionUpdateOne {
	suo.mutation.ClearPayload()
	return suo
}

// SetExpiresAt sets the "expires_at" field.
func (suo *SessionUpdateOne) SetExpiresAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetExpiresAt(t)
	return suo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableExpiresAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetExpiresAt(*t)
	}
	return suo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (suo *SessionUpdateOne) ClearExpiresAt() *SessionUpdateOne {
	suo.mutation.ClearExpiresAt()
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SessionUpdateOne) SetUser(u *User) *SessionUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SessionUpdateOne) ClearUser() *SessionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Where appends a list predicates to the SessionUpdate builder.
func (suo *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SessionUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok && !suo.mutation.UpdatedAtCleared() {
		v := session.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SessionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SessionUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(session.FieldCreatedAt, field.TypeTime, value)
	}
	if suo.mutation.CreatedAtCleared() {
		_spec.ClearField(session.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(session.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.UpdatedAtCleared() {
		_spec.ClearField(session.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.IP(); ok {
		_spec.SetField(session.FieldIP, field.TypeString, value)
	}
	if suo.mutation.IPCleared() {
		_spec.ClearField(session.FieldIP, field.TypeString)
	}
	if value, ok := suo.mutation.UserAgent(); ok {
		_spec.SetField(session.FieldUserAgent, field.TypeString, value)
	}
	if suo.mutation.UserAgentCleared() {
		_spec.ClearField(session.FieldUserAgent, field.TypeString)
	}
	if value, ok := suo.mutation.Payload(); ok {
		_spec.SetField(session.FieldPayload, field.TypeString, value)
	}
	if suo.mutation.PayloadCleared() {
		_spec.ClearField(session.FieldPayload, field.TypeString)
	}
	if value, ok := suo.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	if suo.mutation.ExpiresAtCleared() {
		_spec.ClearField(session.FieldExpiresAt, field.TypeTime)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.UserTable,
			Columns: []string{session.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
