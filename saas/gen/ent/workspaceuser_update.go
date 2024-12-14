// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/predicate"
	"saas/gen/ent/user"
	"saas/gen/ent/workspace"
	"saas/gen/ent/workspaceuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceUserUpdate is the builder for updating WorkspaceUser entities.
type WorkspaceUserUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkspaceUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the WorkspaceUserUpdate builder.
func (wuu *WorkspaceUserUpdate) Where(ps ...predicate.WorkspaceUser) *WorkspaceUserUpdate {
	wuu.mutation.Where(ps...)
	return wuu
}

// SetUpdatedAt sets the "updated_at" field.
func (wuu *WorkspaceUserUpdate) SetUpdatedAt(t time.Time) *WorkspaceUserUpdate {
	wuu.mutation.SetUpdatedAt(t)
	return wuu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (wuu *WorkspaceUserUpdate) ClearUpdatedAt() *WorkspaceUserUpdate {
	wuu.mutation.ClearUpdatedAt()
	return wuu
}

// SetAppID sets the "app_id" field.
func (wuu *WorkspaceUserUpdate) SetAppID(s string) *WorkspaceUserUpdate {
	wuu.mutation.SetAppID(s)
	return wuu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (wuu *WorkspaceUserUpdate) SetNillableAppID(s *string) *WorkspaceUserUpdate {
	if s != nil {
		wuu.SetAppID(*s)
	}
	return wuu
}

// ClearAppID clears the value of the "app_id" field.
func (wuu *WorkspaceUserUpdate) ClearAppID() *WorkspaceUserUpdate {
	wuu.mutation.ClearAppID()
	return wuu
}

// SetWorkspaceID sets the "workspace_id" field.
func (wuu *WorkspaceUserUpdate) SetWorkspaceID(s string) *WorkspaceUserUpdate {
	wuu.mutation.SetWorkspaceID(s)
	return wuu
}

// SetNillableWorkspaceID sets the "workspace_id" field if the given value is not nil.
func (wuu *WorkspaceUserUpdate) SetNillableWorkspaceID(s *string) *WorkspaceUserUpdate {
	if s != nil {
		wuu.SetWorkspaceID(*s)
	}
	return wuu
}

// SetUserID sets the "user_id" field.
func (wuu *WorkspaceUserUpdate) SetUserID(s string) *WorkspaceUserUpdate {
	wuu.mutation.SetUserID(s)
	return wuu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wuu *WorkspaceUserUpdate) SetNillableUserID(s *string) *WorkspaceUserUpdate {
	if s != nil {
		wuu.SetUserID(*s)
	}
	return wuu
}

// SetRole sets the "role" field.
func (wuu *WorkspaceUserUpdate) SetRole(s string) *WorkspaceUserUpdate {
	wuu.mutation.SetRole(s)
	return wuu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (wuu *WorkspaceUserUpdate) SetNillableRole(s *string) *WorkspaceUserUpdate {
	if s != nil {
		wuu.SetRole(*s)
	}
	return wuu
}

// ClearRole clears the value of the "role" field.
func (wuu *WorkspaceUserUpdate) ClearRole() *WorkspaceUserUpdate {
	wuu.mutation.ClearRole()
	return wuu
}

// SetUser sets the "user" edge to the User entity.
func (wuu *WorkspaceUserUpdate) SetUser(u *User) *WorkspaceUserUpdate {
	return wuu.SetUserID(u.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (wuu *WorkspaceUserUpdate) SetWorkspace(w *Workspace) *WorkspaceUserUpdate {
	return wuu.SetWorkspaceID(w.ID)
}

// Mutation returns the WorkspaceUserMutation object of the builder.
func (wuu *WorkspaceUserUpdate) Mutation() *WorkspaceUserMutation {
	return wuu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wuu *WorkspaceUserUpdate) ClearUser() *WorkspaceUserUpdate {
	wuu.mutation.ClearUser()
	return wuu
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (wuu *WorkspaceUserUpdate) ClearWorkspace() *WorkspaceUserUpdate {
	wuu.mutation.ClearWorkspace()
	return wuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wuu *WorkspaceUserUpdate) Save(ctx context.Context) (int, error) {
	wuu.defaults()
	return withHooks(ctx, wuu.sqlSave, wuu.mutation, wuu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuu *WorkspaceUserUpdate) SaveX(ctx context.Context) int {
	affected, err := wuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wuu *WorkspaceUserUpdate) Exec(ctx context.Context) error {
	_, err := wuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuu *WorkspaceUserUpdate) ExecX(ctx context.Context) {
	if err := wuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuu *WorkspaceUserUpdate) defaults() {
	if _, ok := wuu.mutation.UpdatedAt(); !ok && !wuu.mutation.UpdatedAtCleared() {
		v := workspaceuser.UpdateDefaultUpdatedAt()
		wuu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuu *WorkspaceUserUpdate) check() error {
	if wuu.mutation.UserCleared() && len(wuu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "WorkspaceUser.user"`)
	}
	if wuu.mutation.WorkspaceCleared() && len(wuu.mutation.WorkspaceIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "WorkspaceUser.workspace"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wuu *WorkspaceUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkspaceUserUpdate {
	wuu.modifiers = append(wuu.modifiers, modifiers...)
	return wuu
}

func (wuu *WorkspaceUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wuu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workspaceuser.Table, workspaceuser.Columns, sqlgraph.NewFieldSpec(workspaceuser.FieldID, field.TypeString))
	if ps := wuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wuu.mutation.CreatedAtCleared() {
		_spec.ClearField(workspaceuser.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := wuu.mutation.UpdatedAt(); ok {
		_spec.SetField(workspaceuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if wuu.mutation.UpdatedAtCleared() {
		_spec.ClearField(workspaceuser.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := wuu.mutation.AppID(); ok {
		_spec.SetField(workspaceuser.FieldAppID, field.TypeString, value)
	}
	if wuu.mutation.AppIDCleared() {
		_spec.ClearField(workspaceuser.FieldAppID, field.TypeString)
	}
	if value, ok := wuu.mutation.Role(); ok {
		_spec.SetField(workspaceuser.FieldRole, field.TypeString, value)
	}
	if wuu.mutation.RoleCleared() {
		_spec.ClearField(workspaceuser.FieldRole, field.TypeString)
	}
	if wuu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.UserTable,
			Columns: []string{workspaceuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.UserTable,
			Columns: []string{workspaceuser.UserColumn},
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
	if wuu.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.WorkspaceTable,
			Columns: []string{workspaceuser.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuu.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.WorkspaceTable,
			Columns: []string{workspaceuser.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wuu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wuu.mutation.done = true
	return n, nil
}

// WorkspaceUserUpdateOne is the builder for updating a single WorkspaceUser entity.
type WorkspaceUserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkspaceUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (wuuo *WorkspaceUserUpdateOne) SetUpdatedAt(t time.Time) *WorkspaceUserUpdateOne {
	wuuo.mutation.SetUpdatedAt(t)
	return wuuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (wuuo *WorkspaceUserUpdateOne) ClearUpdatedAt() *WorkspaceUserUpdateOne {
	wuuo.mutation.ClearUpdatedAt()
	return wuuo
}

// SetAppID sets the "app_id" field.
func (wuuo *WorkspaceUserUpdateOne) SetAppID(s string) *WorkspaceUserUpdateOne {
	wuuo.mutation.SetAppID(s)
	return wuuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (wuuo *WorkspaceUserUpdateOne) SetNillableAppID(s *string) *WorkspaceUserUpdateOne {
	if s != nil {
		wuuo.SetAppID(*s)
	}
	return wuuo
}

// ClearAppID clears the value of the "app_id" field.
func (wuuo *WorkspaceUserUpdateOne) ClearAppID() *WorkspaceUserUpdateOne {
	wuuo.mutation.ClearAppID()
	return wuuo
}

// SetWorkspaceID sets the "workspace_id" field.
func (wuuo *WorkspaceUserUpdateOne) SetWorkspaceID(s string) *WorkspaceUserUpdateOne {
	wuuo.mutation.SetWorkspaceID(s)
	return wuuo
}

// SetNillableWorkspaceID sets the "workspace_id" field if the given value is not nil.
func (wuuo *WorkspaceUserUpdateOne) SetNillableWorkspaceID(s *string) *WorkspaceUserUpdateOne {
	if s != nil {
		wuuo.SetWorkspaceID(*s)
	}
	return wuuo
}

// SetUserID sets the "user_id" field.
func (wuuo *WorkspaceUserUpdateOne) SetUserID(s string) *WorkspaceUserUpdateOne {
	wuuo.mutation.SetUserID(s)
	return wuuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wuuo *WorkspaceUserUpdateOne) SetNillableUserID(s *string) *WorkspaceUserUpdateOne {
	if s != nil {
		wuuo.SetUserID(*s)
	}
	return wuuo
}

// SetRole sets the "role" field.
func (wuuo *WorkspaceUserUpdateOne) SetRole(s string) *WorkspaceUserUpdateOne {
	wuuo.mutation.SetRole(s)
	return wuuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (wuuo *WorkspaceUserUpdateOne) SetNillableRole(s *string) *WorkspaceUserUpdateOne {
	if s != nil {
		wuuo.SetRole(*s)
	}
	return wuuo
}

// ClearRole clears the value of the "role" field.
func (wuuo *WorkspaceUserUpdateOne) ClearRole() *WorkspaceUserUpdateOne {
	wuuo.mutation.ClearRole()
	return wuuo
}

// SetUser sets the "user" edge to the User entity.
func (wuuo *WorkspaceUserUpdateOne) SetUser(u *User) *WorkspaceUserUpdateOne {
	return wuuo.SetUserID(u.ID)
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (wuuo *WorkspaceUserUpdateOne) SetWorkspace(w *Workspace) *WorkspaceUserUpdateOne {
	return wuuo.SetWorkspaceID(w.ID)
}

// Mutation returns the WorkspaceUserMutation object of the builder.
func (wuuo *WorkspaceUserUpdateOne) Mutation() *WorkspaceUserMutation {
	return wuuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wuuo *WorkspaceUserUpdateOne) ClearUser() *WorkspaceUserUpdateOne {
	wuuo.mutation.ClearUser()
	return wuuo
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (wuuo *WorkspaceUserUpdateOne) ClearWorkspace() *WorkspaceUserUpdateOne {
	wuuo.mutation.ClearWorkspace()
	return wuuo
}

// Where appends a list predicates to the WorkspaceUserUpdate builder.
func (wuuo *WorkspaceUserUpdateOne) Where(ps ...predicate.WorkspaceUser) *WorkspaceUserUpdateOne {
	wuuo.mutation.Where(ps...)
	return wuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuuo *WorkspaceUserUpdateOne) Select(field string, fields ...string) *WorkspaceUserUpdateOne {
	wuuo.fields = append([]string{field}, fields...)
	return wuuo
}

// Save executes the query and returns the updated WorkspaceUser entity.
func (wuuo *WorkspaceUserUpdateOne) Save(ctx context.Context) (*WorkspaceUser, error) {
	wuuo.defaults()
	return withHooks(ctx, wuuo.sqlSave, wuuo.mutation, wuuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuuo *WorkspaceUserUpdateOne) SaveX(ctx context.Context) *WorkspaceUser {
	node, err := wuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuuo *WorkspaceUserUpdateOne) Exec(ctx context.Context) error {
	_, err := wuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuuo *WorkspaceUserUpdateOne) ExecX(ctx context.Context) {
	if err := wuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuuo *WorkspaceUserUpdateOne) defaults() {
	if _, ok := wuuo.mutation.UpdatedAt(); !ok && !wuuo.mutation.UpdatedAtCleared() {
		v := workspaceuser.UpdateDefaultUpdatedAt()
		wuuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuuo *WorkspaceUserUpdateOne) check() error {
	if wuuo.mutation.UserCleared() && len(wuuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "WorkspaceUser.user"`)
	}
	if wuuo.mutation.WorkspaceCleared() && len(wuuo.mutation.WorkspaceIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "WorkspaceUser.workspace"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wuuo *WorkspaceUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkspaceUserUpdateOne {
	wuuo.modifiers = append(wuuo.modifiers, modifiers...)
	return wuuo
}

func (wuuo *WorkspaceUserUpdateOne) sqlSave(ctx context.Context) (_node *WorkspaceUser, err error) {
	if err := wuuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workspaceuser.Table, workspaceuser.Columns, sqlgraph.NewFieldSpec(workspaceuser.FieldID, field.TypeString))
	id, ok := wuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkspaceUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspaceuser.FieldID)
		for _, f := range fields {
			if !workspaceuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workspaceuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wuuo.mutation.CreatedAtCleared() {
		_spec.ClearField(workspaceuser.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := wuuo.mutation.UpdatedAt(); ok {
		_spec.SetField(workspaceuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if wuuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(workspaceuser.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := wuuo.mutation.AppID(); ok {
		_spec.SetField(workspaceuser.FieldAppID, field.TypeString, value)
	}
	if wuuo.mutation.AppIDCleared() {
		_spec.ClearField(workspaceuser.FieldAppID, field.TypeString)
	}
	if value, ok := wuuo.mutation.Role(); ok {
		_spec.SetField(workspaceuser.FieldRole, field.TypeString, value)
	}
	if wuuo.mutation.RoleCleared() {
		_spec.ClearField(workspaceuser.FieldRole, field.TypeString)
	}
	if wuuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.UserTable,
			Columns: []string{workspaceuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.UserTable,
			Columns: []string{workspaceuser.UserColumn},
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
	if wuuo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.WorkspaceTable,
			Columns: []string{workspaceuser.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuuo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workspaceuser.WorkspaceTable,
			Columns: []string{workspaceuser.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(wuuo.modifiers...)
	_node = &WorkspaceUser{config: wuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuuo.mutation.done = true
	return _node, nil
}
