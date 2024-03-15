// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/user"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetFirstName sets the "first_name" field.
func (uc *UserCreate) SetFirstName(s string) *UserCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableFirstName(s *string) *UserCreate {
	if s != nil {
		uc.SetFirstName(*s)
	}
	return uc
}

// SetLastName sets the "last_name" field.
func (uc *UserCreate) SetLastName(s string) *UserCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastName(s *string) *UserCreate {
	if s != nil {
		uc.SetLastName(*s)
	}
	return uc
}

// SetCompany sets the "company" field.
func (uc *UserCreate) SetCompany(s string) *UserCreate {
	uc.mutation.SetCompany(s)
	return uc
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (uc *UserCreate) SetNillableCompany(s *string) *UserCreate {
	if s != nil {
		uc.SetCompany(*s)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(b bool) *UserCreate {
	uc.mutation.SetStatus(b)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(b *bool) *UserCreate {
	if b != nil {
		uc.SetStatus(*b)
	}
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetSecret sets the "secret" field.
func (uc *UserCreate) SetSecret(s string) *UserCreate {
	uc.mutation.SetSecret(s)
	return uc
}

// SetNillableSecret sets the "secret" field if the given value is not nil.
func (uc *UserCreate) SetNillableSecret(s *string) *UserCreate {
	if s != nil {
		uc.SetSecret(*s)
	}
	return uc
}

// SetWelcomeEmailSent sets the "welcome_email_sent" field.
func (uc *UserCreate) SetWelcomeEmailSent(b bool) *UserCreate {
	uc.mutation.SetWelcomeEmailSent(b)
	return uc
}

// SetNillableWelcomeEmailSent sets the "welcome_email_sent" field if the given value is not nil.
func (uc *UserCreate) SetNillableWelcomeEmailSent(b *bool) *UserCreate {
	if b != nil {
		uc.SetWelcomeEmailSent(*b)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(s *string) *UserCreate {
	if s != nil {
		uc.SetID(*s)
	}
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := uc.mutation.Company(); ok {
		_spec.SetField(user.FieldCompany, field.TypeString, value)
		_node.Company = value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := uc.mutation.Secret(); ok {
		_spec.SetField(user.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := uc.mutation.WelcomeEmailSent(); ok {
		_spec.SetField(user.FieldWelcomeEmailSent, field.TypeBool, value)
		_node.WelcomeEmailSent = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsert) SetCreatedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateCreatedAt() *UserUpsert {
	u.SetExcluded(user.FieldCreatedAt)
	return u
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *UserUpsert) ClearCreatedAt() *UserUpsert {
	u.SetNull(user.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsert) SetUpdatedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdatedAt() *UserUpsert {
	u.SetExcluded(user.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *UserUpsert) ClearUpdatedAt() *UserUpsert {
	u.SetNull(user.FieldUpdatedAt)
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// SetFirstName sets the "first_name" field.
func (u *UserUpsert) SetFirstName(v string) *UserUpsert {
	u.Set(user.FieldFirstName, v)
	return u
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *UserUpsert) UpdateFirstName() *UserUpsert {
	u.SetExcluded(user.FieldFirstName)
	return u
}

// ClearFirstName clears the value of the "first_name" field.
func (u *UserUpsert) ClearFirstName() *UserUpsert {
	u.SetNull(user.FieldFirstName)
	return u
}

// SetLastName sets the "last_name" field.
func (u *UserUpsert) SetLastName(v string) *UserUpsert {
	u.Set(user.FieldLastName, v)
	return u
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *UserUpsert) UpdateLastName() *UserUpsert {
	u.SetExcluded(user.FieldLastName)
	return u
}

// ClearLastName clears the value of the "last_name" field.
func (u *UserUpsert) ClearLastName() *UserUpsert {
	u.SetNull(user.FieldLastName)
	return u
}

// SetCompany sets the "company" field.
func (u *UserUpsert) SetCompany(v string) *UserUpsert {
	u.Set(user.FieldCompany, v)
	return u
}

// UpdateCompany sets the "company" field to the value that was provided on create.
func (u *UserUpsert) UpdateCompany() *UserUpsert {
	u.SetExcluded(user.FieldCompany)
	return u
}

// ClearCompany clears the value of the "company" field.
func (u *UserUpsert) ClearCompany() *UserUpsert {
	u.SetNull(user.FieldCompany)
	return u
}

// SetStatus sets the "status" field.
func (u *UserUpsert) SetStatus(v bool) *UserUpsert {
	u.Set(user.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *UserUpsert) UpdateStatus() *UserUpsert {
	u.SetExcluded(user.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *UserUpsert) ClearStatus() *UserUpsert {
	u.SetNull(user.FieldStatus)
	return u
}

// SetPassword sets the "password" field.
func (u *UserUpsert) SetPassword(v string) *UserUpsert {
	u.Set(user.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsert) UpdatePassword() *UserUpsert {
	u.SetExcluded(user.FieldPassword)
	return u
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsert) ClearPassword() *UserUpsert {
	u.SetNull(user.FieldPassword)
	return u
}

// SetSecret sets the "secret" field.
func (u *UserUpsert) SetSecret(v string) *UserUpsert {
	u.Set(user.FieldSecret, v)
	return u
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *UserUpsert) UpdateSecret() *UserUpsert {
	u.SetExcluded(user.FieldSecret)
	return u
}

// ClearSecret clears the value of the "secret" field.
func (u *UserUpsert) ClearSecret() *UserUpsert {
	u.SetNull(user.FieldSecret)
	return u
}

// SetWelcomeEmailSent sets the "welcome_email_sent" field.
func (u *UserUpsert) SetWelcomeEmailSent(v bool) *UserUpsert {
	u.Set(user.FieldWelcomeEmailSent, v)
	return u
}

// UpdateWelcomeEmailSent sets the "welcome_email_sent" field to the value that was provided on create.
func (u *UserUpsert) UpdateWelcomeEmailSent() *UserUpsert {
	u.SetExcluded(user.FieldWelcomeEmailSent)
	return u
}

// ClearWelcomeEmailSent clears the value of the "welcome_email_sent" field.
func (u *UserUpsert) ClearWelcomeEmailSent() *UserUpsert {
	u.SetNull(user.FieldWelcomeEmailSent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertOne) SetCreatedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCreatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *UserUpsertOne) ClearCreatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertOne) SetUpdatedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *UserUpsertOne) ClearUpdatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetFirstName sets the "first_name" field.
func (u *UserUpsertOne) SetFirstName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateFirstName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateFirstName()
	})
}

// ClearFirstName clears the value of the "first_name" field.
func (u *UserUpsertOne) ClearFirstName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *UserUpsertOne) SetLastName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLastName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastName()
	})
}

// ClearLastName clears the value of the "last_name" field.
func (u *UserUpsertOne) ClearLastName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearLastName()
	})
}

// SetCompany sets the "company" field.
func (u *UserUpsertOne) SetCompany(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCompany(v)
	})
}

// UpdateCompany sets the "company" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCompany() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCompany()
	})
}

// ClearCompany clears the value of the "company" field.
func (u *UserUpsertOne) ClearCompany() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearCompany()
	})
}

// SetStatus sets the "status" field.
func (u *UserUpsertOne) SetStatus(v bool) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateStatus() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *UserUpsertOne) ClearStatus() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearStatus()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertOne) SetPassword(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsertOne) ClearPassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearPassword()
	})
}

// SetSecret sets the "secret" field.
func (u *UserUpsertOne) SetSecret(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateSecret() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSecret()
	})
}

// ClearSecret clears the value of the "secret" field.
func (u *UserUpsertOne) ClearSecret() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearSecret()
	})
}

// SetWelcomeEmailSent sets the "welcome_email_sent" field.
func (u *UserUpsertOne) SetWelcomeEmailSent(v bool) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetWelcomeEmailSent(v)
	})
}

// UpdateWelcomeEmailSent sets the "welcome_email_sent" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateWelcomeEmailSent() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateWelcomeEmailSent()
	})
}

// ClearWelcomeEmailSent clears the value of the "welcome_email_sent" field.
func (u *UserUpsertOne) ClearWelcomeEmailSent() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearWelcomeEmailSent()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertBulk) SetCreatedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCreatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *UserUpsertBulk) ClearCreatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertBulk) SetUpdatedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *UserUpsertBulk) ClearUpdatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetFirstName sets the "first_name" field.
func (u *UserUpsertBulk) SetFirstName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetFirstName(v)
	})
}

// UpdateFirstName sets the "first_name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateFirstName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateFirstName()
	})
}

// ClearFirstName clears the value of the "first_name" field.
func (u *UserUpsertBulk) ClearFirstName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearFirstName()
	})
}

// SetLastName sets the "last_name" field.
func (u *UserUpsertBulk) SetLastName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLastName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastName()
	})
}

// ClearLastName clears the value of the "last_name" field.
func (u *UserUpsertBulk) ClearLastName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearLastName()
	})
}

// SetCompany sets the "company" field.
func (u *UserUpsertBulk) SetCompany(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCompany(v)
	})
}

// UpdateCompany sets the "company" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCompany() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCompany()
	})
}

// ClearCompany clears the value of the "company" field.
func (u *UserUpsertBulk) ClearCompany() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearCompany()
	})
}

// SetStatus sets the "status" field.
func (u *UserUpsertBulk) SetStatus(v bool) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateStatus() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *UserUpsertBulk) ClearStatus() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearStatus()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertBulk) SetPassword(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// ClearPassword clears the value of the "password" field.
func (u *UserUpsertBulk) ClearPassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearPassword()
	})
}

// SetSecret sets the "secret" field.
func (u *UserUpsertBulk) SetSecret(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetSecret(v)
	})
}

// UpdateSecret sets the "secret" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateSecret() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSecret()
	})
}

// ClearSecret clears the value of the "secret" field.
func (u *UserUpsertBulk) ClearSecret() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearSecret()
	})
}

// SetWelcomeEmailSent sets the "welcome_email_sent" field.
func (u *UserUpsertBulk) SetWelcomeEmailSent(v bool) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetWelcomeEmailSent(v)
	})
}

// UpdateWelcomeEmailSent sets the "welcome_email_sent" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateWelcomeEmailSent() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateWelcomeEmailSent()
	})
}

// ClearWelcomeEmailSent clears the value of the "welcome_email_sent" field.
func (u *UserUpsertBulk) ClearWelcomeEmailSent() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearWelcomeEmailSent()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}