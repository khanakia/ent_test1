// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/mediable"
	"saas/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediableUpdate is the builder for updating Mediable entities.
type MediableUpdate struct {
	config
	hooks     []Hook
	mutation  *MediableMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MediableUpdate builder.
func (mu *MediableUpdate) Where(ps ...predicate.Mediable) *MediableUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetAppID sets the "app_id" field.
func (mu *MediableUpdate) SetAppID(s string) *MediableUpdate {
	mu.mutation.SetAppID(s)
	return mu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (mu *MediableUpdate) SetNillableAppID(s *string) *MediableUpdate {
	if s != nil {
		mu.SetAppID(*s)
	}
	return mu
}

// ClearAppID clears the value of the "app_id" field.
func (mu *MediableUpdate) ClearAppID() *MediableUpdate {
	mu.mutation.ClearAppID()
	return mu
}

// SetMediaID sets the "media_id" field.
func (mu *MediableUpdate) SetMediaID(s string) *MediableUpdate {
	mu.mutation.SetMediaID(s)
	return mu
}

// SetNillableMediaID sets the "media_id" field if the given value is not nil.
func (mu *MediableUpdate) SetNillableMediaID(s *string) *MediableUpdate {
	if s != nil {
		mu.SetMediaID(*s)
	}
	return mu
}

// ClearMediaID clears the value of the "media_id" field.
func (mu *MediableUpdate) ClearMediaID() *MediableUpdate {
	mu.mutation.ClearMediaID()
	return mu
}

// SetMediableID sets the "mediable_id" field.
func (mu *MediableUpdate) SetMediableID(s string) *MediableUpdate {
	mu.mutation.SetMediableID(s)
	return mu
}

// SetNillableMediableID sets the "mediable_id" field if the given value is not nil.
func (mu *MediableUpdate) SetNillableMediableID(s *string) *MediableUpdate {
	if s != nil {
		mu.SetMediableID(*s)
	}
	return mu
}

// ClearMediableID clears the value of the "mediable_id" field.
func (mu *MediableUpdate) ClearMediableID() *MediableUpdate {
	mu.mutation.ClearMediableID()
	return mu
}

// SetMediableType sets the "mediable_type" field.
func (mu *MediableUpdate) SetMediableType(s string) *MediableUpdate {
	mu.mutation.SetMediableType(s)
	return mu
}

// SetNillableMediableType sets the "mediable_type" field if the given value is not nil.
func (mu *MediableUpdate) SetNillableMediableType(s *string) *MediableUpdate {
	if s != nil {
		mu.SetMediableType(*s)
	}
	return mu
}

// ClearMediableType clears the value of the "mediable_type" field.
func (mu *MediableUpdate) ClearMediableType() *MediableUpdate {
	mu.mutation.ClearMediableType()
	return mu
}

// SetTag sets the "tag" field.
func (mu *MediableUpdate) SetTag(s string) *MediableUpdate {
	mu.mutation.SetTag(s)
	return mu
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (mu *MediableUpdate) SetNillableTag(s *string) *MediableUpdate {
	if s != nil {
		mu.SetTag(*s)
	}
	return mu
}

// ClearTag clears the value of the "tag" field.
func (mu *MediableUpdate) ClearTag() *MediableUpdate {
	mu.mutation.ClearTag()
	return mu
}

// SetOrder sets the "order" field.
func (mu *MediableUpdate) SetOrder(i int) *MediableUpdate {
	mu.mutation.ResetOrder()
	mu.mutation.SetOrder(i)
	return mu
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (mu *MediableUpdate) SetNillableOrder(i *int) *MediableUpdate {
	if i != nil {
		mu.SetOrder(*i)
	}
	return mu
}

// AddOrder adds i to the "order" field.
func (mu *MediableUpdate) AddOrder(i int) *MediableUpdate {
	mu.mutation.AddOrder(i)
	return mu
}

// ClearOrder clears the value of the "order" field.
func (mu *MediableUpdate) ClearOrder() *MediableUpdate {
	mu.mutation.ClearOrder()
	return mu
}

// Mutation returns the MediableMutation object of the builder.
func (mu *MediableUpdate) Mutation() *MediableMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MediableUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MediableUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MediableUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MediableUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MediableUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MediableUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MediableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mediable.Table, mediable.Columns, sqlgraph.NewFieldSpec(mediable.FieldID, field.TypeString))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.AppID(); ok {
		_spec.SetField(mediable.FieldAppID, field.TypeString, value)
	}
	if mu.mutation.AppIDCleared() {
		_spec.ClearField(mediable.FieldAppID, field.TypeString)
	}
	if value, ok := mu.mutation.MediaID(); ok {
		_spec.SetField(mediable.FieldMediaID, field.TypeString, value)
	}
	if mu.mutation.MediaIDCleared() {
		_spec.ClearField(mediable.FieldMediaID, field.TypeString)
	}
	if value, ok := mu.mutation.MediableID(); ok {
		_spec.SetField(mediable.FieldMediableID, field.TypeString, value)
	}
	if mu.mutation.MediableIDCleared() {
		_spec.ClearField(mediable.FieldMediableID, field.TypeString)
	}
	if value, ok := mu.mutation.MediableType(); ok {
		_spec.SetField(mediable.FieldMediableType, field.TypeString, value)
	}
	if mu.mutation.MediableTypeCleared() {
		_spec.ClearField(mediable.FieldMediableType, field.TypeString)
	}
	if value, ok := mu.mutation.Tag(); ok {
		_spec.SetField(mediable.FieldTag, field.TypeString, value)
	}
	if mu.mutation.TagCleared() {
		_spec.ClearField(mediable.FieldTag, field.TypeString)
	}
	if value, ok := mu.mutation.Order(); ok {
		_spec.SetField(mediable.FieldOrder, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedOrder(); ok {
		_spec.AddField(mediable.FieldOrder, field.TypeInt, value)
	}
	if mu.mutation.OrderCleared() {
		_spec.ClearField(mediable.FieldOrder, field.TypeInt)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mediable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MediableUpdateOne is the builder for updating a single Mediable entity.
type MediableUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MediableMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetAppID sets the "app_id" field.
func (muo *MediableUpdateOne) SetAppID(s string) *MediableUpdateOne {
	muo.mutation.SetAppID(s)
	return muo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (muo *MediableUpdateOne) SetNillableAppID(s *string) *MediableUpdateOne {
	if s != nil {
		muo.SetAppID(*s)
	}
	return muo
}

// ClearAppID clears the value of the "app_id" field.
func (muo *MediableUpdateOne) ClearAppID() *MediableUpdateOne {
	muo.mutation.ClearAppID()
	return muo
}

// SetMediaID sets the "media_id" field.
func (muo *MediableUpdateOne) SetMediaID(s string) *MediableUpdateOne {
	muo.mutation.SetMediaID(s)
	return muo
}

// SetNillableMediaID sets the "media_id" field if the given value is not nil.
func (muo *MediableUpdateOne) SetNillableMediaID(s *string) *MediableUpdateOne {
	if s != nil {
		muo.SetMediaID(*s)
	}
	return muo
}

// ClearMediaID clears the value of the "media_id" field.
func (muo *MediableUpdateOne) ClearMediaID() *MediableUpdateOne {
	muo.mutation.ClearMediaID()
	return muo
}

// SetMediableID sets the "mediable_id" field.
func (muo *MediableUpdateOne) SetMediableID(s string) *MediableUpdateOne {
	muo.mutation.SetMediableID(s)
	return muo
}

// SetNillableMediableID sets the "mediable_id" field if the given value is not nil.
func (muo *MediableUpdateOne) SetNillableMediableID(s *string) *MediableUpdateOne {
	if s != nil {
		muo.SetMediableID(*s)
	}
	return muo
}

// ClearMediableID clears the value of the "mediable_id" field.
func (muo *MediableUpdateOne) ClearMediableID() *MediableUpdateOne {
	muo.mutation.ClearMediableID()
	return muo
}

// SetMediableType sets the "mediable_type" field.
func (muo *MediableUpdateOne) SetMediableType(s string) *MediableUpdateOne {
	muo.mutation.SetMediableType(s)
	return muo
}

// SetNillableMediableType sets the "mediable_type" field if the given value is not nil.
func (muo *MediableUpdateOne) SetNillableMediableType(s *string) *MediableUpdateOne {
	if s != nil {
		muo.SetMediableType(*s)
	}
	return muo
}

// ClearMediableType clears the value of the "mediable_type" field.
func (muo *MediableUpdateOne) ClearMediableType() *MediableUpdateOne {
	muo.mutation.ClearMediableType()
	return muo
}

// SetTag sets the "tag" field.
func (muo *MediableUpdateOne) SetTag(s string) *MediableUpdateOne {
	muo.mutation.SetTag(s)
	return muo
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (muo *MediableUpdateOne) SetNillableTag(s *string) *MediableUpdateOne {
	if s != nil {
		muo.SetTag(*s)
	}
	return muo
}

// ClearTag clears the value of the "tag" field.
func (muo *MediableUpdateOne) ClearTag() *MediableUpdateOne {
	muo.mutation.ClearTag()
	return muo
}

// SetOrder sets the "order" field.
func (muo *MediableUpdateOne) SetOrder(i int) *MediableUpdateOne {
	muo.mutation.ResetOrder()
	muo.mutation.SetOrder(i)
	return muo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (muo *MediableUpdateOne) SetNillableOrder(i *int) *MediableUpdateOne {
	if i != nil {
		muo.SetOrder(*i)
	}
	return muo
}

// AddOrder adds i to the "order" field.
func (muo *MediableUpdateOne) AddOrder(i int) *MediableUpdateOne {
	muo.mutation.AddOrder(i)
	return muo
}

// ClearOrder clears the value of the "order" field.
func (muo *MediableUpdateOne) ClearOrder() *MediableUpdateOne {
	muo.mutation.ClearOrder()
	return muo
}

// Mutation returns the MediableMutation object of the builder.
func (muo *MediableUpdateOne) Mutation() *MediableMutation {
	return muo.mutation
}

// Where appends a list predicates to the MediableUpdate builder.
func (muo *MediableUpdateOne) Where(ps ...predicate.Mediable) *MediableUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MediableUpdateOne) Select(field string, fields ...string) *MediableUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mediable entity.
func (muo *MediableUpdateOne) Save(ctx context.Context) (*Mediable, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MediableUpdateOne) SaveX(ctx context.Context) *Mediable {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MediableUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MediableUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MediableUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MediableUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MediableUpdateOne) sqlSave(ctx context.Context) (_node *Mediable, err error) {
	_spec := sqlgraph.NewUpdateSpec(mediable.Table, mediable.Columns, sqlgraph.NewFieldSpec(mediable.FieldID, field.TypeString))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mediable.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mediable.FieldID)
		for _, f := range fields {
			if !mediable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mediable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.AppID(); ok {
		_spec.SetField(mediable.FieldAppID, field.TypeString, value)
	}
	if muo.mutation.AppIDCleared() {
		_spec.ClearField(mediable.FieldAppID, field.TypeString)
	}
	if value, ok := muo.mutation.MediaID(); ok {
		_spec.SetField(mediable.FieldMediaID, field.TypeString, value)
	}
	if muo.mutation.MediaIDCleared() {
		_spec.ClearField(mediable.FieldMediaID, field.TypeString)
	}
	if value, ok := muo.mutation.MediableID(); ok {
		_spec.SetField(mediable.FieldMediableID, field.TypeString, value)
	}
	if muo.mutation.MediableIDCleared() {
		_spec.ClearField(mediable.FieldMediableID, field.TypeString)
	}
	if value, ok := muo.mutation.MediableType(); ok {
		_spec.SetField(mediable.FieldMediableType, field.TypeString, value)
	}
	if muo.mutation.MediableTypeCleared() {
		_spec.ClearField(mediable.FieldMediableType, field.TypeString)
	}
	if value, ok := muo.mutation.Tag(); ok {
		_spec.SetField(mediable.FieldTag, field.TypeString, value)
	}
	if muo.mutation.TagCleared() {
		_spec.ClearField(mediable.FieldTag, field.TypeString)
	}
	if value, ok := muo.mutation.Order(); ok {
		_spec.SetField(mediable.FieldOrder, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedOrder(); ok {
		_spec.AddField(mediable.FieldOrder, field.TypeInt, value)
	}
	if muo.mutation.OrderCleared() {
		_spec.ClearField(mediable.FieldOrder, field.TypeInt)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Mediable{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mediable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
