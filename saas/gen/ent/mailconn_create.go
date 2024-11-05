// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/mailconn"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MailConnCreate is the builder for creating a MailConn entity.
type MailConnCreate struct {
	config
	mutation *MailConnMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mcc *MailConnCreate) SetCreatedAt(t time.Time) *MailConnCreate {
	mcc.mutation.SetCreatedAt(t)
	return mcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableCreatedAt(t *time.Time) *MailConnCreate {
	if t != nil {
		mcc.SetCreatedAt(*t)
	}
	return mcc
}

// SetUpdatedAt sets the "updated_at" field.
func (mcc *MailConnCreate) SetUpdatedAt(t time.Time) *MailConnCreate {
	mcc.mutation.SetUpdatedAt(t)
	return mcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableUpdatedAt(t *time.Time) *MailConnCreate {
	if t != nil {
		mcc.SetUpdatedAt(*t)
	}
	return mcc
}

// SetAppID sets the "app_id" field.
func (mcc *MailConnCreate) SetAppID(s string) *MailConnCreate {
	mcc.mutation.SetAppID(s)
	return mcc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableAppID(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetAppID(*s)
	}
	return mcc
}

// SetName sets the "name" field.
func (mcc *MailConnCreate) SetName(s string) *MailConnCreate {
	mcc.mutation.SetName(s)
	return mcc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableName(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetName(*s)
	}
	return mcc
}

// SetHost sets the "host" field.
func (mcc *MailConnCreate) SetHost(s string) *MailConnCreate {
	mcc.mutation.SetHost(s)
	return mcc
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableHost(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetHost(*s)
	}
	return mcc
}

// SetPort sets the "port" field.
func (mcc *MailConnCreate) SetPort(i int) *MailConnCreate {
	mcc.mutation.SetPort(i)
	return mcc
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillablePort(i *int) *MailConnCreate {
	if i != nil {
		mcc.SetPort(*i)
	}
	return mcc
}

// SetUsername sets the "username" field.
func (mcc *MailConnCreate) SetUsername(s string) *MailConnCreate {
	mcc.mutation.SetUsername(s)
	return mcc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableUsername(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetUsername(*s)
	}
	return mcc
}

// SetPassword sets the "password" field.
func (mcc *MailConnCreate) SetPassword(s string) *MailConnCreate {
	mcc.mutation.SetPassword(s)
	return mcc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillablePassword(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetPassword(*s)
	}
	return mcc
}

// SetEncryption sets the "encryption" field.
func (mcc *MailConnCreate) SetEncryption(i int) *MailConnCreate {
	mcc.mutation.SetEncryption(i)
	return mcc
}

// SetNillableEncryption sets the "encryption" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableEncryption(i *int) *MailConnCreate {
	if i != nil {
		mcc.SetEncryption(*i)
	}
	return mcc
}

// SetFromName sets the "from_name" field.
func (mcc *MailConnCreate) SetFromName(s string) *MailConnCreate {
	mcc.mutation.SetFromName(s)
	return mcc
}

// SetNillableFromName sets the "from_name" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableFromName(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetFromName(*s)
	}
	return mcc
}

// SetFromEmail sets the "from_email" field.
func (mcc *MailConnCreate) SetFromEmail(s string) *MailConnCreate {
	mcc.mutation.SetFromEmail(s)
	return mcc
}

// SetNillableFromEmail sets the "from_email" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableFromEmail(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetFromEmail(*s)
	}
	return mcc
}

// SetStatus sets the "status" field.
func (mcc *MailConnCreate) SetStatus(b bool) *MailConnCreate {
	mcc.mutation.SetStatus(b)
	return mcc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableStatus(b *bool) *MailConnCreate {
	if b != nil {
		mcc.SetStatus(*b)
	}
	return mcc
}

// SetID sets the "id" field.
func (mcc *MailConnCreate) SetID(s string) *MailConnCreate {
	mcc.mutation.SetID(s)
	return mcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mcc *MailConnCreate) SetNillableID(s *string) *MailConnCreate {
	if s != nil {
		mcc.SetID(*s)
	}
	return mcc
}

// Mutation returns the MailConnMutation object of the builder.
func (mcc *MailConnCreate) Mutation() *MailConnMutation {
	return mcc.mutation
}

// Save creates the MailConn in the database.
func (mcc *MailConnCreate) Save(ctx context.Context) (*MailConn, error) {
	mcc.defaults()
	return withHooks(ctx, mcc.sqlSave, mcc.mutation, mcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mcc *MailConnCreate) SaveX(ctx context.Context) *MailConn {
	v, err := mcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcc *MailConnCreate) Exec(ctx context.Context) error {
	_, err := mcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcc *MailConnCreate) ExecX(ctx context.Context) {
	if err := mcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcc *MailConnCreate) defaults() {
	if _, ok := mcc.mutation.CreatedAt(); !ok {
		v := mailconn.DefaultCreatedAt()
		mcc.mutation.SetCreatedAt(v)
	}
	if _, ok := mcc.mutation.UpdatedAt(); !ok {
		v := mailconn.DefaultUpdatedAt()
		mcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mcc.mutation.Status(); !ok {
		v := mailconn.DefaultStatus
		mcc.mutation.SetStatus(v)
	}
	if _, ok := mcc.mutation.ID(); !ok {
		v := mailconn.DefaultID()
		mcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mcc *MailConnCreate) check() error {
	return nil
}

func (mcc *MailConnCreate) sqlSave(ctx context.Context) (*MailConn, error) {
	if err := mcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected MailConn.ID type: %T", _spec.ID.Value)
		}
	}
	mcc.mutation.id = &_node.ID
	mcc.mutation.done = true
	return _node, nil
}

func (mcc *MailConnCreate) createSpec() (*MailConn, *sqlgraph.CreateSpec) {
	var (
		_node = &MailConn{config: mcc.config}
		_spec = sqlgraph.NewCreateSpec(mailconn.Table, sqlgraph.NewFieldSpec(mailconn.FieldID, field.TypeString))
	)
	_spec.OnConflict = mcc.conflict
	if id, ok := mcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mcc.mutation.CreatedAt(); ok {
		_spec.SetField(mailconn.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mcc.mutation.UpdatedAt(); ok {
		_spec.SetField(mailconn.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mcc.mutation.AppID(); ok {
		_spec.SetField(mailconn.FieldAppID, field.TypeString, value)
		_node.AppID = value
	}
	if value, ok := mcc.mutation.Name(); ok {
		_spec.SetField(mailconn.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mcc.mutation.Host(); ok {
		_spec.SetField(mailconn.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := mcc.mutation.Port(); ok {
		_spec.SetField(mailconn.FieldPort, field.TypeInt, value)
		_node.Port = value
	}
	if value, ok := mcc.mutation.Username(); ok {
		_spec.SetField(mailconn.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := mcc.mutation.Password(); ok {
		_spec.SetField(mailconn.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := mcc.mutation.Encryption(); ok {
		_spec.SetField(mailconn.FieldEncryption, field.TypeInt, value)
		_node.Encryption = value
	}
	if value, ok := mcc.mutation.FromName(); ok {
		_spec.SetField(mailconn.FieldFromName, field.TypeString, value)
		_node.FromName = value
	}
	if value, ok := mcc.mutation.FromEmail(); ok {
		_spec.SetField(mailconn.FieldFromEmail, field.TypeString, value)
		_node.FromEmail = value
	}
	if value, ok := mcc.mutation.Status(); ok {
		_spec.SetField(mailconn.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MailConn.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MailConnUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcc *MailConnCreate) OnConflict(opts ...sql.ConflictOption) *MailConnUpsertOne {
	mcc.conflict = opts
	return &MailConnUpsertOne{
		create: mcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MailConn.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcc *MailConnCreate) OnConflictColumns(columns ...string) *MailConnUpsertOne {
	mcc.conflict = append(mcc.conflict, sql.ConflictColumns(columns...))
	return &MailConnUpsertOne{
		create: mcc,
	}
}

type (
	// MailConnUpsertOne is the builder for "upsert"-ing
	//  one MailConn node.
	MailConnUpsertOne struct {
		create *MailConnCreate
	}

	// MailConnUpsert is the "OnConflict" setter.
	MailConnUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *MailConnUpsert) SetUpdatedAt(v time.Time) *MailConnUpsert {
	u.Set(mailconn.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateUpdatedAt() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MailConnUpsert) ClearUpdatedAt() *MailConnUpsert {
	u.SetNull(mailconn.FieldUpdatedAt)
	return u
}

// SetAppID sets the "app_id" field.
func (u *MailConnUpsert) SetAppID(v string) *MailConnUpsert {
	u.Set(mailconn.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateAppID() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *MailConnUpsert) ClearAppID() *MailConnUpsert {
	u.SetNull(mailconn.FieldAppID)
	return u
}

// SetName sets the "name" field.
func (u *MailConnUpsert) SetName(v string) *MailConnUpsert {
	u.Set(mailconn.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateName() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *MailConnUpsert) ClearName() *MailConnUpsert {
	u.SetNull(mailconn.FieldName)
	return u
}

// SetHost sets the "host" field.
func (u *MailConnUpsert) SetHost(v string) *MailConnUpsert {
	u.Set(mailconn.FieldHost, v)
	return u
}

// UpdateHost sets the "host" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateHost() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldHost)
	return u
}

// ClearHost clears the value of the "host" field.
func (u *MailConnUpsert) ClearHost() *MailConnUpsert {
	u.SetNull(mailconn.FieldHost)
	return u
}

// SetPort sets the "port" field.
func (u *MailConnUpsert) SetPort(v int) *MailConnUpsert {
	u.Set(mailconn.FieldPort, v)
	return u
}

// UpdatePort sets the "port" field to the value that was provided on create.
func (u *MailConnUpsert) UpdatePort() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldPort)
	return u
}

// AddPort adds v to the "port" field.
func (u *MailConnUpsert) AddPort(v int) *MailConnUpsert {
	u.Add(mailconn.FieldPort, v)
	return u
}

// ClearPort clears the value of the "port" field.
func (u *MailConnUpsert) ClearPort() *MailConnUpsert {
	u.SetNull(mailconn.FieldPort)
	return u
}

// SetUsername sets the "username" field.
func (u *MailConnUpsert) SetUsername(v string) *MailConnUpsert {
	u.Set(mailconn.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateUsername() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldUsername)
	return u
}

// ClearUsername clears the value of the "username" field.
func (u *MailConnUpsert) ClearUsername() *MailConnUpsert {
	u.SetNull(mailconn.FieldUsername)
	return u
}

// SetPassword sets the "password" field.
func (u *MailConnUpsert) SetPassword(v string) *MailConnUpsert {
	u.Set(mailconn.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *MailConnUpsert) UpdatePassword() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldPassword)
	return u
}

// ClearPassword clears the value of the "password" field.
func (u *MailConnUpsert) ClearPassword() *MailConnUpsert {
	u.SetNull(mailconn.FieldPassword)
	return u
}

// SetEncryption sets the "encryption" field.
func (u *MailConnUpsert) SetEncryption(v int) *MailConnUpsert {
	u.Set(mailconn.FieldEncryption, v)
	return u
}

// UpdateEncryption sets the "encryption" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateEncryption() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldEncryption)
	return u
}

// AddEncryption adds v to the "encryption" field.
func (u *MailConnUpsert) AddEncryption(v int) *MailConnUpsert {
	u.Add(mailconn.FieldEncryption, v)
	return u
}

// ClearEncryption clears the value of the "encryption" field.
func (u *MailConnUpsert) ClearEncryption() *MailConnUpsert {
	u.SetNull(mailconn.FieldEncryption)
	return u
}

// SetFromName sets the "from_name" field.
func (u *MailConnUpsert) SetFromName(v string) *MailConnUpsert {
	u.Set(mailconn.FieldFromName, v)
	return u
}

// UpdateFromName sets the "from_name" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateFromName() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldFromName)
	return u
}

// ClearFromName clears the value of the "from_name" field.
func (u *MailConnUpsert) ClearFromName() *MailConnUpsert {
	u.SetNull(mailconn.FieldFromName)
	return u
}

// SetFromEmail sets the "from_email" field.
func (u *MailConnUpsert) SetFromEmail(v string) *MailConnUpsert {
	u.Set(mailconn.FieldFromEmail, v)
	return u
}

// UpdateFromEmail sets the "from_email" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateFromEmail() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldFromEmail)
	return u
}

// ClearFromEmail clears the value of the "from_email" field.
func (u *MailConnUpsert) ClearFromEmail() *MailConnUpsert {
	u.SetNull(mailconn.FieldFromEmail)
	return u
}

// SetStatus sets the "status" field.
func (u *MailConnUpsert) SetStatus(v bool) *MailConnUpsert {
	u.Set(mailconn.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MailConnUpsert) UpdateStatus() *MailConnUpsert {
	u.SetExcluded(mailconn.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *MailConnUpsert) ClearStatus() *MailConnUpsert {
	u.SetNull(mailconn.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MailConn.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mailconn.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MailConnUpsertOne) UpdateNewValues() *MailConnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mailconn.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(mailconn.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MailConn.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MailConnUpsertOne) Ignore() *MailConnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MailConnUpsertOne) DoNothing() *MailConnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MailConnCreate.OnConflict
// documentation for more info.
func (u *MailConnUpsertOne) Update(set func(*MailConnUpsert)) *MailConnUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MailConnUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MailConnUpsertOne) SetUpdatedAt(v time.Time) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateUpdatedAt() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MailConnUpsertOne) ClearUpdatedAt() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *MailConnUpsertOne) SetAppID(v string) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateAppID() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MailConnUpsertOne) ClearAppID() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearAppID()
	})
}

// SetName sets the "name" field.
func (u *MailConnUpsertOne) SetName(v string) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateName() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *MailConnUpsertOne) ClearName() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearName()
	})
}

// SetHost sets the "host" field.
func (u *MailConnUpsertOne) SetHost(v string) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetHost(v)
	})
}

// UpdateHost sets the "host" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateHost() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateHost()
	})
}

// ClearHost clears the value of the "host" field.
func (u *MailConnUpsertOne) ClearHost() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearHost()
	})
}

// SetPort sets the "port" field.
func (u *MailConnUpsertOne) SetPort(v int) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetPort(v)
	})
}

// AddPort adds v to the "port" field.
func (u *MailConnUpsertOne) AddPort(v int) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.AddPort(v)
	})
}

// UpdatePort sets the "port" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdatePort() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdatePort()
	})
}

// ClearPort clears the value of the "port" field.
func (u *MailConnUpsertOne) ClearPort() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearPort()
	})
}

// SetUsername sets the "username" field.
func (u *MailConnUpsertOne) SetUsername(v string) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateUsername() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *MailConnUpsertOne) ClearUsername() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearUsername()
	})
}

// SetPassword sets the "password" field.
func (u *MailConnUpsertOne) SetPassword(v string) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdatePassword() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *MailConnUpsertOne) ClearPassword() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearPassword()
	})
}

// SetEncryption sets the "encryption" field.
func (u *MailConnUpsertOne) SetEncryption(v int) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetEncryption(v)
	})
}

// AddEncryption adds v to the "encryption" field.
func (u *MailConnUpsertOne) AddEncryption(v int) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.AddEncryption(v)
	})
}

// UpdateEncryption sets the "encryption" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateEncryption() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateEncryption()
	})
}

// ClearEncryption clears the value of the "encryption" field.
func (u *MailConnUpsertOne) ClearEncryption() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearEncryption()
	})
}

// SetFromName sets the "from_name" field.
func (u *MailConnUpsertOne) SetFromName(v string) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetFromName(v)
	})
}

// UpdateFromName sets the "from_name" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateFromName() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateFromName()
	})
}

// ClearFromName clears the value of the "from_name" field.
func (u *MailConnUpsertOne) ClearFromName() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearFromName()
	})
}

// SetFromEmail sets the "from_email" field.
func (u *MailConnUpsertOne) SetFromEmail(v string) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetFromEmail(v)
	})
}

// UpdateFromEmail sets the "from_email" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateFromEmail() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateFromEmail()
	})
}

// ClearFromEmail clears the value of the "from_email" field.
func (u *MailConnUpsertOne) ClearFromEmail() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearFromEmail()
	})
}

// SetStatus sets the "status" field.
func (u *MailConnUpsertOne) SetStatus(v bool) *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MailConnUpsertOne) UpdateStatus() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *MailConnUpsertOne) ClearStatus() *MailConnUpsertOne {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *MailConnUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MailConnCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MailConnUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MailConnUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MailConnUpsertOne.ID is not supported by MySQL driver. Use MailConnUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MailConnUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MailConnCreateBulk is the builder for creating many MailConn entities in bulk.
type MailConnCreateBulk struct {
	config
	err      error
	builders []*MailConnCreate
	conflict []sql.ConflictOption
}

// Save creates the MailConn entities in the database.
func (mccb *MailConnCreateBulk) Save(ctx context.Context) ([]*MailConn, error) {
	if mccb.err != nil {
		return nil, mccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mccb.builders))
	nodes := make([]*MailConn, len(mccb.builders))
	mutators := make([]Mutator, len(mccb.builders))
	for i := range mccb.builders {
		func(i int, root context.Context) {
			builder := mccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MailConnMutation)
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
					_, err = mutators[i+1].Mutate(root, mccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mccb *MailConnCreateBulk) SaveX(ctx context.Context) []*MailConn {
	v, err := mccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mccb *MailConnCreateBulk) Exec(ctx context.Context) error {
	_, err := mccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mccb *MailConnCreateBulk) ExecX(ctx context.Context) {
	if err := mccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MailConn.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MailConnUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mccb *MailConnCreateBulk) OnConflict(opts ...sql.ConflictOption) *MailConnUpsertBulk {
	mccb.conflict = opts
	return &MailConnUpsertBulk{
		create: mccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MailConn.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mccb *MailConnCreateBulk) OnConflictColumns(columns ...string) *MailConnUpsertBulk {
	mccb.conflict = append(mccb.conflict, sql.ConflictColumns(columns...))
	return &MailConnUpsertBulk{
		create: mccb,
	}
}

// MailConnUpsertBulk is the builder for "upsert"-ing
// a bulk of MailConn nodes.
type MailConnUpsertBulk struct {
	create *MailConnCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MailConn.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mailconn.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MailConnUpsertBulk) UpdateNewValues() *MailConnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mailconn.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(mailconn.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MailConn.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MailConnUpsertBulk) Ignore() *MailConnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MailConnUpsertBulk) DoNothing() *MailConnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MailConnCreateBulk.OnConflict
// documentation for more info.
func (u *MailConnUpsertBulk) Update(set func(*MailConnUpsert)) *MailConnUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MailConnUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MailConnUpsertBulk) SetUpdatedAt(v time.Time) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateUpdatedAt() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MailConnUpsertBulk) ClearUpdatedAt() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *MailConnUpsertBulk) SetAppID(v string) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateAppID() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MailConnUpsertBulk) ClearAppID() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearAppID()
	})
}

// SetName sets the "name" field.
func (u *MailConnUpsertBulk) SetName(v string) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateName() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *MailConnUpsertBulk) ClearName() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearName()
	})
}

// SetHost sets the "host" field.
func (u *MailConnUpsertBulk) SetHost(v string) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetHost(v)
	})
}

// UpdateHost sets the "host" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateHost() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateHost()
	})
}

// ClearHost clears the value of the "host" field.
func (u *MailConnUpsertBulk) ClearHost() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearHost()
	})
}

// SetPort sets the "port" field.
func (u *MailConnUpsertBulk) SetPort(v int) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetPort(v)
	})
}

// AddPort adds v to the "port" field.
func (u *MailConnUpsertBulk) AddPort(v int) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.AddPort(v)
	})
}

// UpdatePort sets the "port" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdatePort() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdatePort()
	})
}

// ClearPort clears the value of the "port" field.
func (u *MailConnUpsertBulk) ClearPort() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearPort()
	})
}

// SetUsername sets the "username" field.
func (u *MailConnUpsertBulk) SetUsername(v string) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateUsername() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *MailConnUpsertBulk) ClearUsername() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearUsername()
	})
}

// SetPassword sets the "password" field.
func (u *MailConnUpsertBulk) SetPassword(v string) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdatePassword() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *MailConnUpsertBulk) ClearPassword() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearPassword()
	})
}

// SetEncryption sets the "encryption" field.
func (u *MailConnUpsertBulk) SetEncryption(v int) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetEncryption(v)
	})
}

// AddEncryption adds v to the "encryption" field.
func (u *MailConnUpsertBulk) AddEncryption(v int) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.AddEncryption(v)
	})
}

// UpdateEncryption sets the "encryption" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateEncryption() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateEncryption()
	})
}

// ClearEncryption clears the value of the "encryption" field.
func (u *MailConnUpsertBulk) ClearEncryption() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearEncryption()
	})
}

// SetFromName sets the "from_name" field.
func (u *MailConnUpsertBulk) SetFromName(v string) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetFromName(v)
	})
}

// UpdateFromName sets the "from_name" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateFromName() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateFromName()
	})
}

// ClearFromName clears the value of the "from_name" field.
func (u *MailConnUpsertBulk) ClearFromName() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearFromName()
	})
}

// SetFromEmail sets the "from_email" field.
func (u *MailConnUpsertBulk) SetFromEmail(v string) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetFromEmail(v)
	})
}

// UpdateFromEmail sets the "from_email" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateFromEmail() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateFromEmail()
	})
}

// ClearFromEmail clears the value of the "from_email" field.
func (u *MailConnUpsertBulk) ClearFromEmail() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearFromEmail()
	})
}

// SetStatus sets the "status" field.
func (u *MailConnUpsertBulk) SetStatus(v bool) *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MailConnUpsertBulk) UpdateStatus() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *MailConnUpsertBulk) ClearStatus() *MailConnUpsertBulk {
	return u.Update(func(s *MailConnUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *MailConnUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MailConnCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MailConnCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MailConnUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
