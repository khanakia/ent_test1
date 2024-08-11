// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/keyvalue"
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// KeyvalueUpdate is the builder for updating Keyvalue entities.
type KeyvalueUpdate struct {
	config
	hooks     []Hook
	mutation  *KeyvalueMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the KeyvalueUpdate builder.
func (ku *KeyvalueUpdate) Where(ps ...predicate.Keyvalue) *KeyvalueUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetCreatedAt sets the "created_at" field.
func (ku *KeyvalueUpdate) SetCreatedAt(t time.Time) *KeyvalueUpdate {
	ku.mutation.SetCreatedAt(t)
	return ku
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ku *KeyvalueUpdate) SetNillableCreatedAt(t *time.Time) *KeyvalueUpdate {
	if t != nil {
		ku.SetCreatedAt(*t)
	}
	return ku
}

// ClearCreatedAt clears the value of the "created_at" field.
func (ku *KeyvalueUpdate) ClearCreatedAt() *KeyvalueUpdate {
	ku.mutation.ClearCreatedAt()
	return ku
}

// SetValue sets the "value" field.
func (ku *KeyvalueUpdate) SetValue(s string) *KeyvalueUpdate {
	ku.mutation.SetValue(s)
	return ku
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (ku *KeyvalueUpdate) SetNillableValue(s *string) *KeyvalueUpdate {
	if s != nil {
		ku.SetValue(*s)
	}
	return ku
}

// ClearValue clears the value of the "value" field.
func (ku *KeyvalueUpdate) ClearValue() *KeyvalueUpdate {
	ku.mutation.ClearValue()
	return ku
}

// Mutation returns the KeyvalueMutation object of the builder.
func (ku *KeyvalueUpdate) Mutation() *KeyvalueMutation {
	return ku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KeyvalueUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ku.sqlSave, ku.mutation, ku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KeyvalueUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KeyvalueUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KeyvalueUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ku *KeyvalueUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *KeyvalueUpdate {
	ku.modifiers = append(ku.modifiers, modifiers...)
	return ku
}

func (ku *KeyvalueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(keyvalue.Table, keyvalue.Columns, sqlgraph.NewFieldSpec(keyvalue.FieldID, field.TypeString))
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.CreatedAt(); ok {
		_spec.SetField(keyvalue.FieldCreatedAt, field.TypeTime, value)
	}
	if ku.mutation.CreatedAtCleared() {
		_spec.ClearField(keyvalue.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ku.mutation.Value(); ok {
		_spec.SetField(keyvalue.FieldValue, field.TypeString, value)
	}
	if ku.mutation.ValueCleared() {
		_spec.ClearField(keyvalue.FieldValue, field.TypeString)
	}
	_spec.AddModifiers(ku.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keyvalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ku.mutation.done = true
	return n, nil
}

// KeyvalueUpdateOne is the builder for updating a single Keyvalue entity.
type KeyvalueUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *KeyvalueMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (kuo *KeyvalueUpdateOne) SetCreatedAt(t time.Time) *KeyvalueUpdateOne {
	kuo.mutation.SetCreatedAt(t)
	return kuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kuo *KeyvalueUpdateOne) SetNillableCreatedAt(t *time.Time) *KeyvalueUpdateOne {
	if t != nil {
		kuo.SetCreatedAt(*t)
	}
	return kuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (kuo *KeyvalueUpdateOne) ClearCreatedAt() *KeyvalueUpdateOne {
	kuo.mutation.ClearCreatedAt()
	return kuo
}

// SetValue sets the "value" field.
func (kuo *KeyvalueUpdateOne) SetValue(s string) *KeyvalueUpdateOne {
	kuo.mutation.SetValue(s)
	return kuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (kuo *KeyvalueUpdateOne) SetNillableValue(s *string) *KeyvalueUpdateOne {
	if s != nil {
		kuo.SetValue(*s)
	}
	return kuo
}

// ClearValue clears the value of the "value" field.
func (kuo *KeyvalueUpdateOne) ClearValue() *KeyvalueUpdateOne {
	kuo.mutation.ClearValue()
	return kuo
}

// Mutation returns the KeyvalueMutation object of the builder.
func (kuo *KeyvalueUpdateOne) Mutation() *KeyvalueMutation {
	return kuo.mutation
}

// Where appends a list predicates to the KeyvalueUpdate builder.
func (kuo *KeyvalueUpdateOne) Where(ps ...predicate.Keyvalue) *KeyvalueUpdateOne {
	kuo.mutation.Where(ps...)
	return kuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KeyvalueUpdateOne) Select(field string, fields ...string) *KeyvalueUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Keyvalue entity.
func (kuo *KeyvalueUpdateOne) Save(ctx context.Context) (*Keyvalue, error) {
	return withHooks(ctx, kuo.sqlSave, kuo.mutation, kuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KeyvalueUpdateOne) SaveX(ctx context.Context) *Keyvalue {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KeyvalueUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KeyvalueUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (kuo *KeyvalueUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *KeyvalueUpdateOne {
	kuo.modifiers = append(kuo.modifiers, modifiers...)
	return kuo
}

func (kuo *KeyvalueUpdateOne) sqlSave(ctx context.Context) (_node *Keyvalue, err error) {
	_spec := sqlgraph.NewUpdateSpec(keyvalue.Table, keyvalue.Columns, sqlgraph.NewFieldSpec(keyvalue.FieldID, field.TypeString))
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Keyvalue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keyvalue.FieldID)
		for _, f := range fields {
			if !keyvalue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != keyvalue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.CreatedAt(); ok {
		_spec.SetField(keyvalue.FieldCreatedAt, field.TypeTime, value)
	}
	if kuo.mutation.CreatedAtCleared() {
		_spec.ClearField(keyvalue.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := kuo.mutation.Value(); ok {
		_spec.SetField(keyvalue.FieldValue, field.TypeString, value)
	}
	if kuo.mutation.ValueCleared() {
		_spec.ClearField(keyvalue.FieldValue, field.TypeString)
	}
	_spec.AddModifiers(kuo.modifiers...)
	_node = &Keyvalue{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keyvalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	kuo.mutation.done = true
	return _node, nil
}