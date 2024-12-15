// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/app"
	"saas/gen/ent/appperm"
	"saas/gen/ent/approle"
	"saas/gen/ent/approleperm"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AppRoleCreate is the builder for creating a AppRole entity.
type AppRoleCreate struct {
	config
	mutation *AppRoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (arc *AppRoleCreate) SetCreatedAt(t time.Time) *AppRoleCreate {
	arc.mutation.SetCreatedAt(t)
	return arc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableCreatedAt(t *time.Time) *AppRoleCreate {
	if t != nil {
		arc.SetCreatedAt(*t)
	}
	return arc
}

// SetUpdatedAt sets the "updated_at" field.
func (arc *AppRoleCreate) SetUpdatedAt(t time.Time) *AppRoleCreate {
	arc.mutation.SetUpdatedAt(t)
	return arc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableUpdatedAt(t *time.Time) *AppRoleCreate {
	if t != nil {
		arc.SetUpdatedAt(*t)
	}
	return arc
}

// SetName sets the "name" field.
func (arc *AppRoleCreate) SetName(s string) *AppRoleCreate {
	arc.mutation.SetName(s)
	return arc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableName(s *string) *AppRoleCreate {
	if s != nil {
		arc.SetName(*s)
	}
	return arc
}

// SetAppID sets the "app_id" field.
func (arc *AppRoleCreate) SetAppID(s string) *AppRoleCreate {
	arc.mutation.SetAppID(s)
	return arc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableAppID(s *string) *AppRoleCreate {
	if s != nil {
		arc.SetAppID(*s)
	}
	return arc
}

// SetIsGlobal sets the "is_global" field.
func (arc *AppRoleCreate) SetIsGlobal(b bool) *AppRoleCreate {
	arc.mutation.SetIsGlobal(b)
	return arc
}

// SetNillableIsGlobal sets the "is_global" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableIsGlobal(b *bool) *AppRoleCreate {
	if b != nil {
		arc.SetIsGlobal(*b)
	}
	return arc
}

// SetID sets the "id" field.
func (arc *AppRoleCreate) SetID(s string) *AppRoleCreate {
	arc.mutation.SetID(s)
	return arc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableID(s *string) *AppRoleCreate {
	if s != nil {
		arc.SetID(*s)
	}
	return arc
}

// SetApp sets the "app" edge to the App entity.
func (arc *AppRoleCreate) SetApp(a *App) *AppRoleCreate {
	return arc.SetAppID(a.ID)
}

// AddAppPermIDs adds the "app_perms" edge to the AppPerm entity by IDs.
func (arc *AppRoleCreate) AddAppPermIDs(ids ...string) *AppRoleCreate {
	arc.mutation.AddAppPermIDs(ids...)
	return arc
}

// AddAppPerms adds the "app_perms" edges to the AppPerm entity.
func (arc *AppRoleCreate) AddAppPerms(a ...*AppPerm) *AppRoleCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddAppPermIDs(ids...)
}

// AddAppRolePermIDs adds the "app_role_perms" edge to the AppRolePerm entity by IDs.
func (arc *AppRoleCreate) AddAppRolePermIDs(ids ...string) *AppRoleCreate {
	arc.mutation.AddAppRolePermIDs(ids...)
	return arc
}

// AddAppRolePerms adds the "app_role_perms" edges to the AppRolePerm entity.
func (arc *AppRoleCreate) AddAppRolePerms(a ...*AppRolePerm) *AppRoleCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddAppRolePermIDs(ids...)
}

// Mutation returns the AppRoleMutation object of the builder.
func (arc *AppRoleCreate) Mutation() *AppRoleMutation {
	return arc.mutation
}

// Save creates the AppRole in the database.
func (arc *AppRoleCreate) Save(ctx context.Context) (*AppRole, error) {
	arc.defaults()
	return withHooks(ctx, arc.sqlSave, arc.mutation, arc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AppRoleCreate) SaveX(ctx context.Context) *AppRole {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AppRoleCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AppRoleCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AppRoleCreate) defaults() {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		v := approle.DefaultCreatedAt()
		arc.mutation.SetCreatedAt(v)
	}
	if _, ok := arc.mutation.UpdatedAt(); !ok {
		v := approle.DefaultUpdatedAt()
		arc.mutation.SetUpdatedAt(v)
	}
	if _, ok := arc.mutation.IsGlobal(); !ok {
		v := approle.DefaultIsGlobal
		arc.mutation.SetIsGlobal(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		v := approle.DefaultID()
		arc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AppRoleCreate) check() error {
	if _, ok := arc.mutation.IsGlobal(); !ok {
		return &ValidationError{Name: "is_global", err: errors.New(`ent: missing required field "AppRole.is_global"`)}
	}
	return nil
}

func (arc *AppRoleCreate) sqlSave(ctx context.Context) (*AppRole, error) {
	if err := arc.check(); err != nil {
		return nil, err
	}
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AppRole.ID type: %T", _spec.ID.Value)
		}
	}
	arc.mutation.id = &_node.ID
	arc.mutation.done = true
	return _node, nil
}

func (arc *AppRoleCreate) createSpec() (*AppRole, *sqlgraph.CreateSpec) {
	var (
		_node = &AppRole{config: arc.config}
		_spec = sqlgraph.NewCreateSpec(approle.Table, sqlgraph.NewFieldSpec(approle.FieldID, field.TypeString))
	)
	_spec.OnConflict = arc.conflict
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.SetField(approle.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := arc.mutation.UpdatedAt(); ok {
		_spec.SetField(approle.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := arc.mutation.Name(); ok {
		_spec.SetField(approle.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := arc.mutation.IsGlobal(); ok {
		_spec.SetField(approle.FieldIsGlobal, field.TypeBool, value)
		_node.IsGlobal = value
	}
	if nodes := arc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approle.AppTable,
			Columns: []string{approle.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AppID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.AppPermsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.AppPermsTable,
			Columns: approle.AppPermsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appperm.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePermCreate{config: arc.config, mutation: newAppRolePermMutation(arc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.AppRolePermsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   approle.AppRolePermsTable,
			Columns: []string{approle.AppRolePermsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approleperm.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppRole.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppRoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (arc *AppRoleCreate) OnConflict(opts ...sql.ConflictOption) *AppRoleUpsertOne {
	arc.conflict = opts
	return &AppRoleUpsertOne{
		create: arc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (arc *AppRoleCreate) OnConflictColumns(columns ...string) *AppRoleUpsertOne {
	arc.conflict = append(arc.conflict, sql.ConflictColumns(columns...))
	return &AppRoleUpsertOne{
		create: arc,
	}
}

type (
	// AppRoleUpsertOne is the builder for "upsert"-ing
	//  one AppRole node.
	AppRoleUpsertOne struct {
		create *AppRoleCreate
	}

	// AppRoleUpsert is the "OnConflict" setter.
	AppRoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUpsert) SetUpdatedAt(v time.Time) *AppRoleUpsert {
	u.Set(approle.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateUpdatedAt() *AppRoleUpsert {
	u.SetExcluded(approle.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppRoleUpsert) ClearUpdatedAt() *AppRoleUpsert {
	u.SetNull(approle.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *AppRoleUpsert) SetName(v string) *AppRoleUpsert {
	u.Set(approle.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateName() *AppRoleUpsert {
	u.SetExcluded(approle.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *AppRoleUpsert) ClearName() *AppRoleUpsert {
	u.SetNull(approle.FieldName)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUpsert) SetAppID(v string) *AppRoleUpsert {
	u.Set(approle.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateAppID() *AppRoleUpsert {
	u.SetExcluded(approle.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUpsert) ClearAppID() *AppRoleUpsert {
	u.SetNull(approle.FieldAppID)
	return u
}

// SetIsGlobal sets the "is_global" field.
func (u *AppRoleUpsert) SetIsGlobal(v bool) *AppRoleUpsert {
	u.Set(approle.FieldIsGlobal, v)
	return u
}

// UpdateIsGlobal sets the "is_global" field to the value that was provided on create.
func (u *AppRoleUpsert) UpdateIsGlobal() *AppRoleUpsert {
	u.SetExcluded(approle.FieldIsGlobal)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approle.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppRoleUpsertOne) UpdateNewValues() *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(approle.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(approle.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppRole.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppRoleUpsertOne) Ignore() *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppRoleUpsertOne) DoNothing() *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppRoleCreate.OnConflict
// documentation for more info.
func (u *AppRoleUpsertOne) Update(set func(*AppRoleUpsert)) *AppRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUpsertOne) SetUpdatedAt(v time.Time) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateUpdatedAt() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppRoleUpsertOne) ClearUpdatedAt() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AppRoleUpsertOne) SetName(v string) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateName() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AppRoleUpsertOne) ClearName() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearName()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUpsertOne) SetAppID(v string) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateAppID() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUpsertOne) ClearAppID() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearAppID()
	})
}

// SetIsGlobal sets the "is_global" field.
func (u *AppRoleUpsertOne) SetIsGlobal(v bool) *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetIsGlobal(v)
	})
}

// UpdateIsGlobal sets the "is_global" field to the value that was provided on create.
func (u *AppRoleUpsertOne) UpdateIsGlobal() *AppRoleUpsertOne {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateIsGlobal()
	})
}

// Exec executes the query.
func (u *AppRoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppRoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppRoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppRoleUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppRoleUpsertOne.ID is not supported by MySQL driver. Use AppRoleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppRoleUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppRoleCreateBulk is the builder for creating many AppRole entities in bulk.
type AppRoleCreateBulk struct {
	config
	err      error
	builders []*AppRoleCreate
	conflict []sql.ConflictOption
}

// Save creates the AppRole entities in the database.
func (arcb *AppRoleCreateBulk) Save(ctx context.Context) ([]*AppRole, error) {
	if arcb.err != nil {
		return nil, arcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AppRole, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = arcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AppRoleCreateBulk) SaveX(ctx context.Context) []*AppRole {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AppRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AppRoleCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppRole.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppRoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (arcb *AppRoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppRoleUpsertBulk {
	arcb.conflict = opts
	return &AppRoleUpsertBulk{
		create: arcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (arcb *AppRoleCreateBulk) OnConflictColumns(columns ...string) *AppRoleUpsertBulk {
	arcb.conflict = append(arcb.conflict, sql.ConflictColumns(columns...))
	return &AppRoleUpsertBulk{
		create: arcb,
	}
}

// AppRoleUpsertBulk is the builder for "upsert"-ing
// a bulk of AppRole nodes.
type AppRoleUpsertBulk struct {
	create *AppRoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approle.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppRoleUpsertBulk) UpdateNewValues() *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(approle.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(approle.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppRole.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppRoleUpsertBulk) Ignore() *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppRoleUpsertBulk) DoNothing() *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppRoleCreateBulk.OnConflict
// documentation for more info.
func (u *AppRoleUpsertBulk) Update(set func(*AppRoleUpsert)) *AppRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUpsertBulk) SetUpdatedAt(v time.Time) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateUpdatedAt() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppRoleUpsertBulk) ClearUpdatedAt() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AppRoleUpsertBulk) SetName(v string) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateName() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AppRoleUpsertBulk) ClearName() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearName()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUpsertBulk) SetAppID(v string) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateAppID() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUpsertBulk) ClearAppID() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.ClearAppID()
	})
}

// SetIsGlobal sets the "is_global" field.
func (u *AppRoleUpsertBulk) SetIsGlobal(v bool) *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.SetIsGlobal(v)
	})
}

// UpdateIsGlobal sets the "is_global" field to the value that was provided on create.
func (u *AppRoleUpsertBulk) UpdateIsGlobal() *AppRoleUpsertBulk {
	return u.Update(func(s *AppRoleUpsert) {
		s.UpdateIsGlobal()
	})
}

// Exec executes the query.
func (u *AppRoleUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppRoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppRoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppRoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}