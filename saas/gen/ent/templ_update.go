// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/predicate"
	"saas/gen/ent/templ"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TemplUpdate is the builder for updating Templ entities.
type TemplUpdate struct {
	config
	hooks     []Hook
	mutation  *TemplMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TemplUpdate builder.
func (tu *TemplUpdate) Where(ps ...predicate.Templ) *TemplUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TemplUpdate) SetCreatedAt(t time.Time) *TemplUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TemplUpdate) SetNillableCreatedAt(t *time.Time) *TemplUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (tu *TemplUpdate) ClearCreatedAt() *TemplUpdate {
	tu.mutation.ClearCreatedAt()
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TemplUpdate) SetUpdatedAt(t time.Time) *TemplUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tu *TemplUpdate) ClearUpdatedAt() *TemplUpdate {
	tu.mutation.ClearUpdatedAt()
	return tu
}

// SetName sets the "name" field.
func (tu *TemplUpdate) SetName(s string) *TemplUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TemplUpdate) SetNillableName(s *string) *TemplUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// ClearName clears the value of the "name" field.
func (tu *TemplUpdate) ClearName() *TemplUpdate {
	tu.mutation.ClearName()
	return tu
}

// SetBody sets the "body" field.
func (tu *TemplUpdate) SetBody(s string) *TemplUpdate {
	tu.mutation.SetBody(s)
	return tu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (tu *TemplUpdate) SetNillableBody(s *string) *TemplUpdate {
	if s != nil {
		tu.SetBody(*s)
	}
	return tu
}

// ClearBody clears the value of the "body" field.
func (tu *TemplUpdate) ClearBody() *TemplUpdate {
	tu.mutation.ClearBody()
	return tu
}

// SetCompiled sets the "compiled" field.
func (tu *TemplUpdate) SetCompiled(s string) *TemplUpdate {
	tu.mutation.SetCompiled(s)
	return tu
}

// SetNillableCompiled sets the "compiled" field if the given value is not nil.
func (tu *TemplUpdate) SetNillableCompiled(s *string) *TemplUpdate {
	if s != nil {
		tu.SetCompiled(*s)
	}
	return tu
}

// ClearCompiled clears the value of the "compiled" field.
func (tu *TemplUpdate) ClearCompiled() *TemplUpdate {
	tu.mutation.ClearCompiled()
	return tu
}

// SetStatus sets the "status" field.
func (tu *TemplUpdate) SetStatus(b bool) *TemplUpdate {
	tu.mutation.SetStatus(b)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TemplUpdate) SetNillableStatus(b *bool) *TemplUpdate {
	if b != nil {
		tu.SetStatus(*b)
	}
	return tu
}

// ClearStatus clears the value of the "status" field.
func (tu *TemplUpdate) ClearStatus() *TemplUpdate {
	tu.mutation.ClearStatus()
	return tu
}

// Mutation returns the TemplMutation object of the builder.
func (tu *TemplUpdate) Mutation() *TemplMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TemplUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TemplUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TemplUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TemplUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TemplUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok && !tu.mutation.UpdatedAtCleared() {
		v := templ.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TemplUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemplUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TemplUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(templ.Table, templ.Columns, sqlgraph.NewFieldSpec(templ.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(templ.FieldCreatedAt, field.TypeTime, value)
	}
	if tu.mutation.CreatedAtCleared() {
		_spec.ClearField(templ.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(templ.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.UpdatedAtCleared() {
		_spec.ClearField(templ.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(templ.FieldName, field.TypeString, value)
	}
	if tu.mutation.NameCleared() {
		_spec.ClearField(templ.FieldName, field.TypeString)
	}
	if value, ok := tu.mutation.Body(); ok {
		_spec.SetField(templ.FieldBody, field.TypeString, value)
	}
	if tu.mutation.BodyCleared() {
		_spec.ClearField(templ.FieldBody, field.TypeString)
	}
	if value, ok := tu.mutation.Compiled(); ok {
		_spec.SetField(templ.FieldCompiled, field.TypeString, value)
	}
	if tu.mutation.CompiledCleared() {
		_spec.ClearField(templ.FieldCompiled, field.TypeString)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(templ.FieldStatus, field.TypeBool, value)
	}
	if tu.mutation.StatusCleared() {
		_spec.ClearField(templ.FieldStatus, field.TypeBool)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templ.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TemplUpdateOne is the builder for updating a single Templ entity.
type TemplUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TemplMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TemplUpdateOne) SetCreatedAt(t time.Time) *TemplUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TemplUpdateOne) SetNillableCreatedAt(t *time.Time) *TemplUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (tuo *TemplUpdateOne) ClearCreatedAt() *TemplUpdateOne {
	tuo.mutation.ClearCreatedAt()
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TemplUpdateOne) SetUpdatedAt(t time.Time) *TemplUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tuo *TemplUpdateOne) ClearUpdatedAt() *TemplUpdateOne {
	tuo.mutation.ClearUpdatedAt()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TemplUpdateOne) SetName(s string) *TemplUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TemplUpdateOne) SetNillableName(s *string) *TemplUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// ClearName clears the value of the "name" field.
func (tuo *TemplUpdateOne) ClearName() *TemplUpdateOne {
	tuo.mutation.ClearName()
	return tuo
}

// SetBody sets the "body" field.
func (tuo *TemplUpdateOne) SetBody(s string) *TemplUpdateOne {
	tuo.mutation.SetBody(s)
	return tuo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (tuo *TemplUpdateOne) SetNillableBody(s *string) *TemplUpdateOne {
	if s != nil {
		tuo.SetBody(*s)
	}
	return tuo
}

// ClearBody clears the value of the "body" field.
func (tuo *TemplUpdateOne) ClearBody() *TemplUpdateOne {
	tuo.mutation.ClearBody()
	return tuo
}

// SetCompiled sets the "compiled" field.
func (tuo *TemplUpdateOne) SetCompiled(s string) *TemplUpdateOne {
	tuo.mutation.SetCompiled(s)
	return tuo
}

// SetNillableCompiled sets the "compiled" field if the given value is not nil.
func (tuo *TemplUpdateOne) SetNillableCompiled(s *string) *TemplUpdateOne {
	if s != nil {
		tuo.SetCompiled(*s)
	}
	return tuo
}

// ClearCompiled clears the value of the "compiled" field.
func (tuo *TemplUpdateOne) ClearCompiled() *TemplUpdateOne {
	tuo.mutation.ClearCompiled()
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TemplUpdateOne) SetStatus(b bool) *TemplUpdateOne {
	tuo.mutation.SetStatus(b)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TemplUpdateOne) SetNillableStatus(b *bool) *TemplUpdateOne {
	if b != nil {
		tuo.SetStatus(*b)
	}
	return tuo
}

// ClearStatus clears the value of the "status" field.
func (tuo *TemplUpdateOne) ClearStatus() *TemplUpdateOne {
	tuo.mutation.ClearStatus()
	return tuo
}

// Mutation returns the TemplMutation object of the builder.
func (tuo *TemplUpdateOne) Mutation() *TemplMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TemplUpdate builder.
func (tuo *TemplUpdateOne) Where(ps ...predicate.Templ) *TemplUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TemplUpdateOne) Select(field string, fields ...string) *TemplUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Templ entity.
func (tuo *TemplUpdateOne) Save(ctx context.Context) (*Templ, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TemplUpdateOne) SaveX(ctx context.Context) *Templ {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TemplUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TemplUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TemplUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok && !tuo.mutation.UpdatedAtCleared() {
		v := templ.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TemplUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TemplUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TemplUpdateOne) sqlSave(ctx context.Context) (_node *Templ, err error) {
	_spec := sqlgraph.NewUpdateSpec(templ.Table, templ.Columns, sqlgraph.NewFieldSpec(templ.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Templ.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, templ.FieldID)
		for _, f := range fields {
			if !templ.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != templ.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(templ.FieldCreatedAt, field.TypeTime, value)
	}
	if tuo.mutation.CreatedAtCleared() {
		_spec.ClearField(templ.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(templ.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(templ.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(templ.FieldName, field.TypeString, value)
	}
	if tuo.mutation.NameCleared() {
		_spec.ClearField(templ.FieldName, field.TypeString)
	}
	if value, ok := tuo.mutation.Body(); ok {
		_spec.SetField(templ.FieldBody, field.TypeString, value)
	}
	if tuo.mutation.BodyCleared() {
		_spec.ClearField(templ.FieldBody, field.TypeString)
	}
	if value, ok := tuo.mutation.Compiled(); ok {
		_spec.SetField(templ.FieldCompiled, field.TypeString, value)
	}
	if tuo.mutation.CompiledCleared() {
		_spec.ClearField(templ.FieldCompiled, field.TypeString)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(templ.FieldStatus, field.TypeBool, value)
	}
	if tuo.mutation.StatusCleared() {
		_spec.ClearField(templ.FieldStatus, field.TypeBool)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Templ{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templ.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
