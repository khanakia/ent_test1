// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/mediable"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediableCreate is the builder for creating a Mediable entity.
type MediableCreate struct {
	config
	mutation *MediableMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (mc *MediableCreate) SetAppID(s string) *MediableCreate {
	mc.mutation.SetAppID(s)
	return mc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (mc *MediableCreate) SetNillableAppID(s *string) *MediableCreate {
	if s != nil {
		mc.SetAppID(*s)
	}
	return mc
}

// SetMediaID sets the "media_id" field.
func (mc *MediableCreate) SetMediaID(s string) *MediableCreate {
	mc.mutation.SetMediaID(s)
	return mc
}

// SetNillableMediaID sets the "media_id" field if the given value is not nil.
func (mc *MediableCreate) SetNillableMediaID(s *string) *MediableCreate {
	if s != nil {
		mc.SetMediaID(*s)
	}
	return mc
}

// SetMediableID sets the "mediable_id" field.
func (mc *MediableCreate) SetMediableID(s string) *MediableCreate {
	mc.mutation.SetMediableID(s)
	return mc
}

// SetNillableMediableID sets the "mediable_id" field if the given value is not nil.
func (mc *MediableCreate) SetNillableMediableID(s *string) *MediableCreate {
	if s != nil {
		mc.SetMediableID(*s)
	}
	return mc
}

// SetMediableType sets the "mediable_type" field.
func (mc *MediableCreate) SetMediableType(s string) *MediableCreate {
	mc.mutation.SetMediableType(s)
	return mc
}

// SetNillableMediableType sets the "mediable_type" field if the given value is not nil.
func (mc *MediableCreate) SetNillableMediableType(s *string) *MediableCreate {
	if s != nil {
		mc.SetMediableType(*s)
	}
	return mc
}

// SetTag sets the "tag" field.
func (mc *MediableCreate) SetTag(s string) *MediableCreate {
	mc.mutation.SetTag(s)
	return mc
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (mc *MediableCreate) SetNillableTag(s *string) *MediableCreate {
	if s != nil {
		mc.SetTag(*s)
	}
	return mc
}

// SetOrder sets the "order" field.
func (mc *MediableCreate) SetOrder(i int) *MediableCreate {
	mc.mutation.SetOrder(i)
	return mc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (mc *MediableCreate) SetNillableOrder(i *int) *MediableCreate {
	if i != nil {
		mc.SetOrder(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MediableCreate) SetID(s string) *MediableCreate {
	mc.mutation.SetID(s)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MediableCreate) SetNillableID(s *string) *MediableCreate {
	if s != nil {
		mc.SetID(*s)
	}
	return mc
}

// Mutation returns the MediableMutation object of the builder.
func (mc *MediableCreate) Mutation() *MediableMutation {
	return mc.mutation
}

// Save creates the Mediable in the database.
func (mc *MediableCreate) Save(ctx context.Context) (*Mediable, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MediableCreate) SaveX(ctx context.Context) *Mediable {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MediableCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MediableCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MediableCreate) defaults() {
	if _, ok := mc.mutation.ID(); !ok {
		v := mediable.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MediableCreate) check() error {
	return nil
}

func (mc *MediableCreate) sqlSave(ctx context.Context) (*Mediable, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Mediable.ID type: %T", _spec.ID.Value)
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MediableCreate) createSpec() (*Mediable, *sqlgraph.CreateSpec) {
	var (
		_node = &Mediable{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(mediable.Table, sqlgraph.NewFieldSpec(mediable.FieldID, field.TypeString))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.AppID(); ok {
		_spec.SetField(mediable.FieldAppID, field.TypeString, value)
		_node.AppID = value
	}
	if value, ok := mc.mutation.MediaID(); ok {
		_spec.SetField(mediable.FieldMediaID, field.TypeString, value)
		_node.MediaID = value
	}
	if value, ok := mc.mutation.MediableID(); ok {
		_spec.SetField(mediable.FieldMediableID, field.TypeString, value)
		_node.MediableID = value
	}
	if value, ok := mc.mutation.MediableType(); ok {
		_spec.SetField(mediable.FieldMediableType, field.TypeString, value)
		_node.MediableType = value
	}
	if value, ok := mc.mutation.Tag(); ok {
		_spec.SetField(mediable.FieldTag, field.TypeString, value)
		_node.Tag = value
	}
	if value, ok := mc.mutation.Order(); ok {
		_spec.SetField(mediable.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mediable.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MediableUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
func (mc *MediableCreate) OnConflict(opts ...sql.ConflictOption) *MediableUpsertOne {
	mc.conflict = opts
	return &MediableUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mediable.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MediableCreate) OnConflictColumns(columns ...string) *MediableUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MediableUpsertOne{
		create: mc,
	}
}

type (
	// MediableUpsertOne is the builder for "upsert"-ing
	//  one Mediable node.
	MediableUpsertOne struct {
		create *MediableCreate
	}

	// MediableUpsert is the "OnConflict" setter.
	MediableUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *MediableUpsert) SetAppID(v string) *MediableUpsert {
	u.Set(mediable.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MediableUpsert) UpdateAppID() *MediableUpsert {
	u.SetExcluded(mediable.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *MediableUpsert) ClearAppID() *MediableUpsert {
	u.SetNull(mediable.FieldAppID)
	return u
}

// SetMediaID sets the "media_id" field.
func (u *MediableUpsert) SetMediaID(v string) *MediableUpsert {
	u.Set(mediable.FieldMediaID, v)
	return u
}

// UpdateMediaID sets the "media_id" field to the value that was provided on create.
func (u *MediableUpsert) UpdateMediaID() *MediableUpsert {
	u.SetExcluded(mediable.FieldMediaID)
	return u
}

// ClearMediaID clears the value of the "media_id" field.
func (u *MediableUpsert) ClearMediaID() *MediableUpsert {
	u.SetNull(mediable.FieldMediaID)
	return u
}

// SetMediableID sets the "mediable_id" field.
func (u *MediableUpsert) SetMediableID(v string) *MediableUpsert {
	u.Set(mediable.FieldMediableID, v)
	return u
}

// UpdateMediableID sets the "mediable_id" field to the value that was provided on create.
func (u *MediableUpsert) UpdateMediableID() *MediableUpsert {
	u.SetExcluded(mediable.FieldMediableID)
	return u
}

// ClearMediableID clears the value of the "mediable_id" field.
func (u *MediableUpsert) ClearMediableID() *MediableUpsert {
	u.SetNull(mediable.FieldMediableID)
	return u
}

// SetMediableType sets the "mediable_type" field.
func (u *MediableUpsert) SetMediableType(v string) *MediableUpsert {
	u.Set(mediable.FieldMediableType, v)
	return u
}

// UpdateMediableType sets the "mediable_type" field to the value that was provided on create.
func (u *MediableUpsert) UpdateMediableType() *MediableUpsert {
	u.SetExcluded(mediable.FieldMediableType)
	return u
}

// ClearMediableType clears the value of the "mediable_type" field.
func (u *MediableUpsert) ClearMediableType() *MediableUpsert {
	u.SetNull(mediable.FieldMediableType)
	return u
}

// SetTag sets the "tag" field.
func (u *MediableUpsert) SetTag(v string) *MediableUpsert {
	u.Set(mediable.FieldTag, v)
	return u
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *MediableUpsert) UpdateTag() *MediableUpsert {
	u.SetExcluded(mediable.FieldTag)
	return u
}

// ClearTag clears the value of the "tag" field.
func (u *MediableUpsert) ClearTag() *MediableUpsert {
	u.SetNull(mediable.FieldTag)
	return u
}

// SetOrder sets the "order" field.
func (u *MediableUpsert) SetOrder(v int) *MediableUpsert {
	u.Set(mediable.FieldOrder, v)
	return u
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *MediableUpsert) UpdateOrder() *MediableUpsert {
	u.SetExcluded(mediable.FieldOrder)
	return u
}

// AddOrder adds v to the "order" field.
func (u *MediableUpsert) AddOrder(v int) *MediableUpsert {
	u.Add(mediable.FieldOrder, v)
	return u
}

// ClearOrder clears the value of the "order" field.
func (u *MediableUpsert) ClearOrder() *MediableUpsert {
	u.SetNull(mediable.FieldOrder)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Mediable.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mediable.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MediableUpsertOne) UpdateNewValues() *MediableUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mediable.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Mediable.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MediableUpsertOne) Ignore() *MediableUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MediableUpsertOne) DoNothing() *MediableUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MediableCreate.OnConflict
// documentation for more info.
func (u *MediableUpsertOne) Update(set func(*MediableUpsert)) *MediableUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MediableUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *MediableUpsertOne) SetAppID(v string) *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MediableUpsertOne) UpdateAppID() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MediableUpsertOne) ClearAppID() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.ClearAppID()
	})
}

// SetMediaID sets the "media_id" field.
func (u *MediableUpsertOne) SetMediaID(v string) *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.SetMediaID(v)
	})
}

// UpdateMediaID sets the "media_id" field to the value that was provided on create.
func (u *MediableUpsertOne) UpdateMediaID() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateMediaID()
	})
}

// ClearMediaID clears the value of the "media_id" field.
func (u *MediableUpsertOne) ClearMediaID() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.ClearMediaID()
	})
}

// SetMediableID sets the "mediable_id" field.
func (u *MediableUpsertOne) SetMediableID(v string) *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.SetMediableID(v)
	})
}

// UpdateMediableID sets the "mediable_id" field to the value that was provided on create.
func (u *MediableUpsertOne) UpdateMediableID() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateMediableID()
	})
}

// ClearMediableID clears the value of the "mediable_id" field.
func (u *MediableUpsertOne) ClearMediableID() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.ClearMediableID()
	})
}

// SetMediableType sets the "mediable_type" field.
func (u *MediableUpsertOne) SetMediableType(v string) *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.SetMediableType(v)
	})
}

// UpdateMediableType sets the "mediable_type" field to the value that was provided on create.
func (u *MediableUpsertOne) UpdateMediableType() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateMediableType()
	})
}

// ClearMediableType clears the value of the "mediable_type" field.
func (u *MediableUpsertOne) ClearMediableType() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.ClearMediableType()
	})
}

// SetTag sets the "tag" field.
func (u *MediableUpsertOne) SetTag(v string) *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *MediableUpsertOne) UpdateTag() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateTag()
	})
}

// ClearTag clears the value of the "tag" field.
func (u *MediableUpsertOne) ClearTag() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.ClearTag()
	})
}

// SetOrder sets the "order" field.
func (u *MediableUpsertOne) SetOrder(v int) *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *MediableUpsertOne) AddOrder(v int) *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *MediableUpsertOne) UpdateOrder() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateOrder()
	})
}

// ClearOrder clears the value of the "order" field.
func (u *MediableUpsertOne) ClearOrder() *MediableUpsertOne {
	return u.Update(func(s *MediableUpsert) {
		s.ClearOrder()
	})
}

// Exec executes the query.
func (u *MediableUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MediableCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MediableUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MediableUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MediableUpsertOne.ID is not supported by MySQL driver. Use MediableUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MediableUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MediableCreateBulk is the builder for creating many Mediable entities in bulk.
type MediableCreateBulk struct {
	config
	err      error
	builders []*MediableCreate
	conflict []sql.ConflictOption
}

// Save creates the Mediable entities in the database.
func (mcb *MediableCreateBulk) Save(ctx context.Context) ([]*Mediable, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mediable, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediableMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MediableCreateBulk) SaveX(ctx context.Context) []*Mediable {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MediableCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MediableCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Mediable.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MediableUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
func (mcb *MediableCreateBulk) OnConflict(opts ...sql.ConflictOption) *MediableUpsertBulk {
	mcb.conflict = opts
	return &MediableUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Mediable.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MediableCreateBulk) OnConflictColumns(columns ...string) *MediableUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MediableUpsertBulk{
		create: mcb,
	}
}

// MediableUpsertBulk is the builder for "upsert"-ing
// a bulk of Mediable nodes.
type MediableUpsertBulk struct {
	create *MediableCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Mediable.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mediable.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MediableUpsertBulk) UpdateNewValues() *MediableUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mediable.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Mediable.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MediableUpsertBulk) Ignore() *MediableUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MediableUpsertBulk) DoNothing() *MediableUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MediableCreateBulk.OnConflict
// documentation for more info.
func (u *MediableUpsertBulk) Update(set func(*MediableUpsert)) *MediableUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MediableUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *MediableUpsertBulk) SetAppID(v string) *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MediableUpsertBulk) UpdateAppID() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MediableUpsertBulk) ClearAppID() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.ClearAppID()
	})
}

// SetMediaID sets the "media_id" field.
func (u *MediableUpsertBulk) SetMediaID(v string) *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.SetMediaID(v)
	})
}

// UpdateMediaID sets the "media_id" field to the value that was provided on create.
func (u *MediableUpsertBulk) UpdateMediaID() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateMediaID()
	})
}

// ClearMediaID clears the value of the "media_id" field.
func (u *MediableUpsertBulk) ClearMediaID() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.ClearMediaID()
	})
}

// SetMediableID sets the "mediable_id" field.
func (u *MediableUpsertBulk) SetMediableID(v string) *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.SetMediableID(v)
	})
}

// UpdateMediableID sets the "mediable_id" field to the value that was provided on create.
func (u *MediableUpsertBulk) UpdateMediableID() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateMediableID()
	})
}

// ClearMediableID clears the value of the "mediable_id" field.
func (u *MediableUpsertBulk) ClearMediableID() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.ClearMediableID()
	})
}

// SetMediableType sets the "mediable_type" field.
func (u *MediableUpsertBulk) SetMediableType(v string) *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.SetMediableType(v)
	})
}

// UpdateMediableType sets the "mediable_type" field to the value that was provided on create.
func (u *MediableUpsertBulk) UpdateMediableType() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateMediableType()
	})
}

// ClearMediableType clears the value of the "mediable_type" field.
func (u *MediableUpsertBulk) ClearMediableType() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.ClearMediableType()
	})
}

// SetTag sets the "tag" field.
func (u *MediableUpsertBulk) SetTag(v string) *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.SetTag(v)
	})
}

// UpdateTag sets the "tag" field to the value that was provided on create.
func (u *MediableUpsertBulk) UpdateTag() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateTag()
	})
}

// ClearTag clears the value of the "tag" field.
func (u *MediableUpsertBulk) ClearTag() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.ClearTag()
	})
}

// SetOrder sets the "order" field.
func (u *MediableUpsertBulk) SetOrder(v int) *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *MediableUpsertBulk) AddOrder(v int) *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *MediableUpsertBulk) UpdateOrder() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.UpdateOrder()
	})
}

// ClearOrder clears the value of the "order" field.
func (u *MediableUpsertBulk) ClearOrder() *MediableUpsertBulk {
	return u.Update(func(s *MediableUpsert) {
		s.ClearOrder()
	})
}

// Exec executes the query.
func (u *MediableUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MediableCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MediableCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MediableUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}