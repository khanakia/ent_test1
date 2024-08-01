// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/predicate"
	"saas/gen/ent/workspace"
	"saas/gen/ent/workspaceinvite"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceInviteUpdate is the builder for updating WorkspaceInvite entities.
type WorkspaceInviteUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkspaceInviteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the WorkspaceInviteUpdate builder.
func (wiu *WorkspaceInviteUpdate) Where(ps ...predicate.WorkspaceInvite) *WorkspaceInviteUpdate {
	wiu.mutation.Where(ps...)
	return wiu
}

// SetUpdatedAt sets the "updated_at" field.
func (wiu *WorkspaceInviteUpdate) SetUpdatedAt(t time.Time) *WorkspaceInviteUpdate {
	wiu.mutation.SetUpdatedAt(t)
	return wiu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (wiu *WorkspaceInviteUpdate) ClearUpdatedAt() *WorkspaceInviteUpdate {
	wiu.mutation.ClearUpdatedAt()
	return wiu
}

// SetWorkspaceID sets the "workspace_id" field.
func (wiu *WorkspaceInviteUpdate) SetWorkspaceID(s string) *WorkspaceInviteUpdate {
	wiu.mutation.SetWorkspaceID(s)
	return wiu
}

// SetNillableWorkspaceID sets the "workspace_id" field if the given value is not nil.
func (wiu *WorkspaceInviteUpdate) SetNillableWorkspaceID(s *string) *WorkspaceInviteUpdate {
	if s != nil {
		wiu.SetWorkspaceID(*s)
	}
	return wiu
}

// ClearWorkspaceID clears the value of the "workspace_id" field.
func (wiu *WorkspaceInviteUpdate) ClearWorkspaceID() *WorkspaceInviteUpdate {
	wiu.mutation.ClearWorkspaceID()
	return wiu
}

// SetEmail sets the "email" field.
func (wiu *WorkspaceInviteUpdate) SetEmail(s string) *WorkspaceInviteUpdate {
	wiu.mutation.SetEmail(s)
	return wiu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (wiu *WorkspaceInviteUpdate) SetNillableEmail(s *string) *WorkspaceInviteUpdate {
	if s != nil {
		wiu.SetEmail(*s)
	}
	return wiu
}

// ClearEmail clears the value of the "email" field.
func (wiu *WorkspaceInviteUpdate) ClearEmail() *WorkspaceInviteUpdate {
	wiu.mutation.ClearEmail()
	return wiu
}

// SetRole sets the "role" field.
func (wiu *WorkspaceInviteUpdate) SetRole(s string) *WorkspaceInviteUpdate {
	wiu.mutation.SetRole(s)
	return wiu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (wiu *WorkspaceInviteUpdate) SetNillableRole(s *string) *WorkspaceInviteUpdate {
	if s != nil {
		wiu.SetRole(*s)
	}
	return wiu
}

// ClearRole clears the value of the "role" field.
func (wiu *WorkspaceInviteUpdate) ClearRole() *WorkspaceInviteUpdate {
	wiu.mutation.ClearRole()
	return wiu
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (wiu *WorkspaceInviteUpdate) SetWorkspace(w *Workspace) *WorkspaceInviteUpdate {
	return wiu.SetWorkspaceID(w.ID)
}

// Mutation returns the WorkspaceInviteMutation object of the builder.
func (wiu *WorkspaceInviteUpdate) Mutation() *WorkspaceInviteMutation {
	return wiu.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (wiu *WorkspaceInviteUpdate) ClearWorkspace() *WorkspaceInviteUpdate {
	wiu.mutation.ClearWorkspace()
	return wiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wiu *WorkspaceInviteUpdate) Save(ctx context.Context) (int, error) {
	wiu.defaults()
	return withHooks(ctx, wiu.sqlSave, wiu.mutation, wiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wiu *WorkspaceInviteUpdate) SaveX(ctx context.Context) int {
	affected, err := wiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wiu *WorkspaceInviteUpdate) Exec(ctx context.Context) error {
	_, err := wiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wiu *WorkspaceInviteUpdate) ExecX(ctx context.Context) {
	if err := wiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wiu *WorkspaceInviteUpdate) defaults() {
	if _, ok := wiu.mutation.UpdatedAt(); !ok && !wiu.mutation.UpdatedAtCleared() {
		v := workspaceinvite.UpdateDefaultUpdatedAt()
		wiu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wiu *WorkspaceInviteUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkspaceInviteUpdate {
	wiu.modifiers = append(wiu.modifiers, modifiers...)
	return wiu
}

func (wiu *WorkspaceInviteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(workspaceinvite.Table, workspaceinvite.Columns, sqlgraph.NewFieldSpec(workspaceinvite.FieldID, field.TypeString))
	if ps := wiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wiu.mutation.CreatedAtCleared() {
		_spec.ClearField(workspaceinvite.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := wiu.mutation.UpdatedAt(); ok {
		_spec.SetField(workspaceinvite.FieldUpdatedAt, field.TypeTime, value)
	}
	if wiu.mutation.UpdatedAtCleared() {
		_spec.ClearField(workspaceinvite.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := wiu.mutation.Email(); ok {
		_spec.SetField(workspaceinvite.FieldEmail, field.TypeString, value)
	}
	if wiu.mutation.EmailCleared() {
		_spec.ClearField(workspaceinvite.FieldEmail, field.TypeString)
	}
	if value, ok := wiu.mutation.Role(); ok {
		_spec.SetField(workspaceinvite.FieldRole, field.TypeString, value)
	}
	if wiu.mutation.RoleCleared() {
		_spec.ClearField(workspaceinvite.FieldRole, field.TypeString)
	}
	if wiu.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceinvite.WorkspaceTable,
			Columns: []string{workspaceinvite.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wiu.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceinvite.WorkspaceTable,
			Columns: []string{workspaceinvite.WorkspaceColumn},
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
	_spec.AddModifiers(wiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceinvite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wiu.mutation.done = true
	return n, nil
}

// WorkspaceInviteUpdateOne is the builder for updating a single WorkspaceInvite entity.
type WorkspaceInviteUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkspaceInviteMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (wiuo *WorkspaceInviteUpdateOne) SetUpdatedAt(t time.Time) *WorkspaceInviteUpdateOne {
	wiuo.mutation.SetUpdatedAt(t)
	return wiuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (wiuo *WorkspaceInviteUpdateOne) ClearUpdatedAt() *WorkspaceInviteUpdateOne {
	wiuo.mutation.ClearUpdatedAt()
	return wiuo
}

// SetWorkspaceID sets the "workspace_id" field.
func (wiuo *WorkspaceInviteUpdateOne) SetWorkspaceID(s string) *WorkspaceInviteUpdateOne {
	wiuo.mutation.SetWorkspaceID(s)
	return wiuo
}

// SetNillableWorkspaceID sets the "workspace_id" field if the given value is not nil.
func (wiuo *WorkspaceInviteUpdateOne) SetNillableWorkspaceID(s *string) *WorkspaceInviteUpdateOne {
	if s != nil {
		wiuo.SetWorkspaceID(*s)
	}
	return wiuo
}

// ClearWorkspaceID clears the value of the "workspace_id" field.
func (wiuo *WorkspaceInviteUpdateOne) ClearWorkspaceID() *WorkspaceInviteUpdateOne {
	wiuo.mutation.ClearWorkspaceID()
	return wiuo
}

// SetEmail sets the "email" field.
func (wiuo *WorkspaceInviteUpdateOne) SetEmail(s string) *WorkspaceInviteUpdateOne {
	wiuo.mutation.SetEmail(s)
	return wiuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (wiuo *WorkspaceInviteUpdateOne) SetNillableEmail(s *string) *WorkspaceInviteUpdateOne {
	if s != nil {
		wiuo.SetEmail(*s)
	}
	return wiuo
}

// ClearEmail clears the value of the "email" field.
func (wiuo *WorkspaceInviteUpdateOne) ClearEmail() *WorkspaceInviteUpdateOne {
	wiuo.mutation.ClearEmail()
	return wiuo
}

// SetRole sets the "role" field.
func (wiuo *WorkspaceInviteUpdateOne) SetRole(s string) *WorkspaceInviteUpdateOne {
	wiuo.mutation.SetRole(s)
	return wiuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (wiuo *WorkspaceInviteUpdateOne) SetNillableRole(s *string) *WorkspaceInviteUpdateOne {
	if s != nil {
		wiuo.SetRole(*s)
	}
	return wiuo
}

// ClearRole clears the value of the "role" field.
func (wiuo *WorkspaceInviteUpdateOne) ClearRole() *WorkspaceInviteUpdateOne {
	wiuo.mutation.ClearRole()
	return wiuo
}

// SetWorkspace sets the "workspace" edge to the Workspace entity.
func (wiuo *WorkspaceInviteUpdateOne) SetWorkspace(w *Workspace) *WorkspaceInviteUpdateOne {
	return wiuo.SetWorkspaceID(w.ID)
}

// Mutation returns the WorkspaceInviteMutation object of the builder.
func (wiuo *WorkspaceInviteUpdateOne) Mutation() *WorkspaceInviteMutation {
	return wiuo.mutation
}

// ClearWorkspace clears the "workspace" edge to the Workspace entity.
func (wiuo *WorkspaceInviteUpdateOne) ClearWorkspace() *WorkspaceInviteUpdateOne {
	wiuo.mutation.ClearWorkspace()
	return wiuo
}

// Where appends a list predicates to the WorkspaceInviteUpdate builder.
func (wiuo *WorkspaceInviteUpdateOne) Where(ps ...predicate.WorkspaceInvite) *WorkspaceInviteUpdateOne {
	wiuo.mutation.Where(ps...)
	return wiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wiuo *WorkspaceInviteUpdateOne) Select(field string, fields ...string) *WorkspaceInviteUpdateOne {
	wiuo.fields = append([]string{field}, fields...)
	return wiuo
}

// Save executes the query and returns the updated WorkspaceInvite entity.
func (wiuo *WorkspaceInviteUpdateOne) Save(ctx context.Context) (*WorkspaceInvite, error) {
	wiuo.defaults()
	return withHooks(ctx, wiuo.sqlSave, wiuo.mutation, wiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wiuo *WorkspaceInviteUpdateOne) SaveX(ctx context.Context) *WorkspaceInvite {
	node, err := wiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wiuo *WorkspaceInviteUpdateOne) Exec(ctx context.Context) error {
	_, err := wiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wiuo *WorkspaceInviteUpdateOne) ExecX(ctx context.Context) {
	if err := wiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wiuo *WorkspaceInviteUpdateOne) defaults() {
	if _, ok := wiuo.mutation.UpdatedAt(); !ok && !wiuo.mutation.UpdatedAtCleared() {
		v := workspaceinvite.UpdateDefaultUpdatedAt()
		wiuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wiuo *WorkspaceInviteUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkspaceInviteUpdateOne {
	wiuo.modifiers = append(wiuo.modifiers, modifiers...)
	return wiuo
}

func (wiuo *WorkspaceInviteUpdateOne) sqlSave(ctx context.Context) (_node *WorkspaceInvite, err error) {
	_spec := sqlgraph.NewUpdateSpec(workspaceinvite.Table, workspaceinvite.Columns, sqlgraph.NewFieldSpec(workspaceinvite.FieldID, field.TypeString))
	id, ok := wiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkspaceInvite.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspaceinvite.FieldID)
		for _, f := range fields {
			if !workspaceinvite.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workspaceinvite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wiuo.mutation.CreatedAtCleared() {
		_spec.ClearField(workspaceinvite.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := wiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(workspaceinvite.FieldUpdatedAt, field.TypeTime, value)
	}
	if wiuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(workspaceinvite.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := wiuo.mutation.Email(); ok {
		_spec.SetField(workspaceinvite.FieldEmail, field.TypeString, value)
	}
	if wiuo.mutation.EmailCleared() {
		_spec.ClearField(workspaceinvite.FieldEmail, field.TypeString)
	}
	if value, ok := wiuo.mutation.Role(); ok {
		_spec.SetField(workspaceinvite.FieldRole, field.TypeString, value)
	}
	if wiuo.mutation.RoleCleared() {
		_spec.ClearField(workspaceinvite.FieldRole, field.TypeString)
	}
	if wiuo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceinvite.WorkspaceTable,
			Columns: []string{workspaceinvite.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wiuo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workspaceinvite.WorkspaceTable,
			Columns: []string{workspaceinvite.WorkspaceColumn},
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
	_spec.AddModifiers(wiuo.modifiers...)
	_node = &WorkspaceInvite{config: wiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspaceinvite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wiuo.mutation.done = true
	return _node, nil
}
