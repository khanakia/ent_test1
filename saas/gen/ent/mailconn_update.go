// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/mailconn"
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MailConnUpdate is the builder for updating MailConn entities.
type MailConnUpdate struct {
	config
	hooks     []Hook
	mutation  *MailConnMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MailConnUpdate builder.
func (mcu *MailConnUpdate) Where(ps ...predicate.MailConn) *MailConnUpdate {
	mcu.mutation.Where(ps...)
	return mcu
}

// SetUpdatedAt sets the "updated_at" field.
func (mcu *MailConnUpdate) SetUpdatedAt(t time.Time) *MailConnUpdate {
	mcu.mutation.SetUpdatedAt(t)
	return mcu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mcu *MailConnUpdate) ClearUpdatedAt() *MailConnUpdate {
	mcu.mutation.ClearUpdatedAt()
	return mcu
}

// SetName sets the "name" field.
func (mcu *MailConnUpdate) SetName(s string) *MailConnUpdate {
	mcu.mutation.SetName(s)
	return mcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillableName(s *string) *MailConnUpdate {
	if s != nil {
		mcu.SetName(*s)
	}
	return mcu
}

// ClearName clears the value of the "name" field.
func (mcu *MailConnUpdate) ClearName() *MailConnUpdate {
	mcu.mutation.ClearName()
	return mcu
}

// SetHost sets the "host" field.
func (mcu *MailConnUpdate) SetHost(s string) *MailConnUpdate {
	mcu.mutation.SetHost(s)
	return mcu
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillableHost(s *string) *MailConnUpdate {
	if s != nil {
		mcu.SetHost(*s)
	}
	return mcu
}

// ClearHost clears the value of the "host" field.
func (mcu *MailConnUpdate) ClearHost() *MailConnUpdate {
	mcu.mutation.ClearHost()
	return mcu
}

// SetPort sets the "port" field.
func (mcu *MailConnUpdate) SetPort(i int) *MailConnUpdate {
	mcu.mutation.ResetPort()
	mcu.mutation.SetPort(i)
	return mcu
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillablePort(i *int) *MailConnUpdate {
	if i != nil {
		mcu.SetPort(*i)
	}
	return mcu
}

// AddPort adds i to the "port" field.
func (mcu *MailConnUpdate) AddPort(i int) *MailConnUpdate {
	mcu.mutation.AddPort(i)
	return mcu
}

// ClearPort clears the value of the "port" field.
func (mcu *MailConnUpdate) ClearPort() *MailConnUpdate {
	mcu.mutation.ClearPort()
	return mcu
}

// SetUsername sets the "username" field.
func (mcu *MailConnUpdate) SetUsername(s string) *MailConnUpdate {
	mcu.mutation.SetUsername(s)
	return mcu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillableUsername(s *string) *MailConnUpdate {
	if s != nil {
		mcu.SetUsername(*s)
	}
	return mcu
}

// ClearUsername clears the value of the "username" field.
func (mcu *MailConnUpdate) ClearUsername() *MailConnUpdate {
	mcu.mutation.ClearUsername()
	return mcu
}

// SetPassword sets the "password" field.
func (mcu *MailConnUpdate) SetPassword(s string) *MailConnUpdate {
	mcu.mutation.SetPassword(s)
	return mcu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillablePassword(s *string) *MailConnUpdate {
	if s != nil {
		mcu.SetPassword(*s)
	}
	return mcu
}

// ClearPassword clears the value of the "password" field.
func (mcu *MailConnUpdate) ClearPassword() *MailConnUpdate {
	mcu.mutation.ClearPassword()
	return mcu
}

// SetEncryption sets the "encryption" field.
func (mcu *MailConnUpdate) SetEncryption(i int) *MailConnUpdate {
	mcu.mutation.ResetEncryption()
	mcu.mutation.SetEncryption(i)
	return mcu
}

// SetNillableEncryption sets the "encryption" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillableEncryption(i *int) *MailConnUpdate {
	if i != nil {
		mcu.SetEncryption(*i)
	}
	return mcu
}

// AddEncryption adds i to the "encryption" field.
func (mcu *MailConnUpdate) AddEncryption(i int) *MailConnUpdate {
	mcu.mutation.AddEncryption(i)
	return mcu
}

// ClearEncryption clears the value of the "encryption" field.
func (mcu *MailConnUpdate) ClearEncryption() *MailConnUpdate {
	mcu.mutation.ClearEncryption()
	return mcu
}

// SetFromName sets the "from_name" field.
func (mcu *MailConnUpdate) SetFromName(s string) *MailConnUpdate {
	mcu.mutation.SetFromName(s)
	return mcu
}

// SetNillableFromName sets the "from_name" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillableFromName(s *string) *MailConnUpdate {
	if s != nil {
		mcu.SetFromName(*s)
	}
	return mcu
}

// ClearFromName clears the value of the "from_name" field.
func (mcu *MailConnUpdate) ClearFromName() *MailConnUpdate {
	mcu.mutation.ClearFromName()
	return mcu
}

// SetFromEmail sets the "from_email" field.
func (mcu *MailConnUpdate) SetFromEmail(s string) *MailConnUpdate {
	mcu.mutation.SetFromEmail(s)
	return mcu
}

// SetNillableFromEmail sets the "from_email" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillableFromEmail(s *string) *MailConnUpdate {
	if s != nil {
		mcu.SetFromEmail(*s)
	}
	return mcu
}

// ClearFromEmail clears the value of the "from_email" field.
func (mcu *MailConnUpdate) ClearFromEmail() *MailConnUpdate {
	mcu.mutation.ClearFromEmail()
	return mcu
}

// SetStatus sets the "status" field.
func (mcu *MailConnUpdate) SetStatus(b bool) *MailConnUpdate {
	mcu.mutation.SetStatus(b)
	return mcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mcu *MailConnUpdate) SetNillableStatus(b *bool) *MailConnUpdate {
	if b != nil {
		mcu.SetStatus(*b)
	}
	return mcu
}

// ClearStatus clears the value of the "status" field.
func (mcu *MailConnUpdate) ClearStatus() *MailConnUpdate {
	mcu.mutation.ClearStatus()
	return mcu
}

// Mutation returns the MailConnMutation object of the builder.
func (mcu *MailConnUpdate) Mutation() *MailConnMutation {
	return mcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mcu *MailConnUpdate) Save(ctx context.Context) (int, error) {
	mcu.defaults()
	return withHooks(ctx, mcu.sqlSave, mcu.mutation, mcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcu *MailConnUpdate) SaveX(ctx context.Context) int {
	affected, err := mcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mcu *MailConnUpdate) Exec(ctx context.Context) error {
	_, err := mcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcu *MailConnUpdate) ExecX(ctx context.Context) {
	if err := mcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcu *MailConnUpdate) defaults() {
	if _, ok := mcu.mutation.UpdatedAt(); !ok && !mcu.mutation.UpdatedAtCleared() {
		v := mailconn.UpdateDefaultUpdatedAt()
		mcu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mcu *MailConnUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MailConnUpdate {
	mcu.modifiers = append(mcu.modifiers, modifiers...)
	return mcu
}

func (mcu *MailConnUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mailconn.Table, mailconn.Columns, sqlgraph.NewFieldSpec(mailconn.FieldID, field.TypeString))
	if ps := mcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mcu.mutation.CreatedAtCleared() {
		_spec.ClearField(mailconn.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mcu.mutation.UpdatedAt(); ok {
		_spec.SetField(mailconn.FieldUpdatedAt, field.TypeTime, value)
	}
	if mcu.mutation.UpdatedAtCleared() {
		_spec.ClearField(mailconn.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mcu.mutation.Name(); ok {
		_spec.SetField(mailconn.FieldName, field.TypeString, value)
	}
	if mcu.mutation.NameCleared() {
		_spec.ClearField(mailconn.FieldName, field.TypeString)
	}
	if value, ok := mcu.mutation.Host(); ok {
		_spec.SetField(mailconn.FieldHost, field.TypeString, value)
	}
	if mcu.mutation.HostCleared() {
		_spec.ClearField(mailconn.FieldHost, field.TypeString)
	}
	if value, ok := mcu.mutation.Port(); ok {
		_spec.SetField(mailconn.FieldPort, field.TypeInt, value)
	}
	if value, ok := mcu.mutation.AddedPort(); ok {
		_spec.AddField(mailconn.FieldPort, field.TypeInt, value)
	}
	if mcu.mutation.PortCleared() {
		_spec.ClearField(mailconn.FieldPort, field.TypeInt)
	}
	if value, ok := mcu.mutation.Username(); ok {
		_spec.SetField(mailconn.FieldUsername, field.TypeString, value)
	}
	if mcu.mutation.UsernameCleared() {
		_spec.ClearField(mailconn.FieldUsername, field.TypeString)
	}
	if value, ok := mcu.mutation.Password(); ok {
		_spec.SetField(mailconn.FieldPassword, field.TypeString, value)
	}
	if mcu.mutation.PasswordCleared() {
		_spec.ClearField(mailconn.FieldPassword, field.TypeString)
	}
	if value, ok := mcu.mutation.Encryption(); ok {
		_spec.SetField(mailconn.FieldEncryption, field.TypeInt, value)
	}
	if value, ok := mcu.mutation.AddedEncryption(); ok {
		_spec.AddField(mailconn.FieldEncryption, field.TypeInt, value)
	}
	if mcu.mutation.EncryptionCleared() {
		_spec.ClearField(mailconn.FieldEncryption, field.TypeInt)
	}
	if value, ok := mcu.mutation.FromName(); ok {
		_spec.SetField(mailconn.FieldFromName, field.TypeString, value)
	}
	if mcu.mutation.FromNameCleared() {
		_spec.ClearField(mailconn.FieldFromName, field.TypeString)
	}
	if value, ok := mcu.mutation.FromEmail(); ok {
		_spec.SetField(mailconn.FieldFromEmail, field.TypeString, value)
	}
	if mcu.mutation.FromEmailCleared() {
		_spec.ClearField(mailconn.FieldFromEmail, field.TypeString)
	}
	if value, ok := mcu.mutation.Status(); ok {
		_spec.SetField(mailconn.FieldStatus, field.TypeBool, value)
	}
	if mcu.mutation.StatusCleared() {
		_spec.ClearField(mailconn.FieldStatus, field.TypeBool)
	}
	_spec.AddModifiers(mcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailconn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mcu.mutation.done = true
	return n, nil
}

// MailConnUpdateOne is the builder for updating a single MailConn entity.
type MailConnUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MailConnMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (mcuo *MailConnUpdateOne) SetUpdatedAt(t time.Time) *MailConnUpdateOne {
	mcuo.mutation.SetUpdatedAt(t)
	return mcuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mcuo *MailConnUpdateOne) ClearUpdatedAt() *MailConnUpdateOne {
	mcuo.mutation.ClearUpdatedAt()
	return mcuo
}

// SetName sets the "name" field.
func (mcuo *MailConnUpdateOne) SetName(s string) *MailConnUpdateOne {
	mcuo.mutation.SetName(s)
	return mcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillableName(s *string) *MailConnUpdateOne {
	if s != nil {
		mcuo.SetName(*s)
	}
	return mcuo
}

// ClearName clears the value of the "name" field.
func (mcuo *MailConnUpdateOne) ClearName() *MailConnUpdateOne {
	mcuo.mutation.ClearName()
	return mcuo
}

// SetHost sets the "host" field.
func (mcuo *MailConnUpdateOne) SetHost(s string) *MailConnUpdateOne {
	mcuo.mutation.SetHost(s)
	return mcuo
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillableHost(s *string) *MailConnUpdateOne {
	if s != nil {
		mcuo.SetHost(*s)
	}
	return mcuo
}

// ClearHost clears the value of the "host" field.
func (mcuo *MailConnUpdateOne) ClearHost() *MailConnUpdateOne {
	mcuo.mutation.ClearHost()
	return mcuo
}

// SetPort sets the "port" field.
func (mcuo *MailConnUpdateOne) SetPort(i int) *MailConnUpdateOne {
	mcuo.mutation.ResetPort()
	mcuo.mutation.SetPort(i)
	return mcuo
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillablePort(i *int) *MailConnUpdateOne {
	if i != nil {
		mcuo.SetPort(*i)
	}
	return mcuo
}

// AddPort adds i to the "port" field.
func (mcuo *MailConnUpdateOne) AddPort(i int) *MailConnUpdateOne {
	mcuo.mutation.AddPort(i)
	return mcuo
}

// ClearPort clears the value of the "port" field.
func (mcuo *MailConnUpdateOne) ClearPort() *MailConnUpdateOne {
	mcuo.mutation.ClearPort()
	return mcuo
}

// SetUsername sets the "username" field.
func (mcuo *MailConnUpdateOne) SetUsername(s string) *MailConnUpdateOne {
	mcuo.mutation.SetUsername(s)
	return mcuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillableUsername(s *string) *MailConnUpdateOne {
	if s != nil {
		mcuo.SetUsername(*s)
	}
	return mcuo
}

// ClearUsername clears the value of the "username" field.
func (mcuo *MailConnUpdateOne) ClearUsername() *MailConnUpdateOne {
	mcuo.mutation.ClearUsername()
	return mcuo
}

// SetPassword sets the "password" field.
func (mcuo *MailConnUpdateOne) SetPassword(s string) *MailConnUpdateOne {
	mcuo.mutation.SetPassword(s)
	return mcuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillablePassword(s *string) *MailConnUpdateOne {
	if s != nil {
		mcuo.SetPassword(*s)
	}
	return mcuo
}

// ClearPassword clears the value of the "password" field.
func (mcuo *MailConnUpdateOne) ClearPassword() *MailConnUpdateOne {
	mcuo.mutation.ClearPassword()
	return mcuo
}

// SetEncryption sets the "encryption" field.
func (mcuo *MailConnUpdateOne) SetEncryption(i int) *MailConnUpdateOne {
	mcuo.mutation.ResetEncryption()
	mcuo.mutation.SetEncryption(i)
	return mcuo
}

// SetNillableEncryption sets the "encryption" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillableEncryption(i *int) *MailConnUpdateOne {
	if i != nil {
		mcuo.SetEncryption(*i)
	}
	return mcuo
}

// AddEncryption adds i to the "encryption" field.
func (mcuo *MailConnUpdateOne) AddEncryption(i int) *MailConnUpdateOne {
	mcuo.mutation.AddEncryption(i)
	return mcuo
}

// ClearEncryption clears the value of the "encryption" field.
func (mcuo *MailConnUpdateOne) ClearEncryption() *MailConnUpdateOne {
	mcuo.mutation.ClearEncryption()
	return mcuo
}

// SetFromName sets the "from_name" field.
func (mcuo *MailConnUpdateOne) SetFromName(s string) *MailConnUpdateOne {
	mcuo.mutation.SetFromName(s)
	return mcuo
}

// SetNillableFromName sets the "from_name" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillableFromName(s *string) *MailConnUpdateOne {
	if s != nil {
		mcuo.SetFromName(*s)
	}
	return mcuo
}

// ClearFromName clears the value of the "from_name" field.
func (mcuo *MailConnUpdateOne) ClearFromName() *MailConnUpdateOne {
	mcuo.mutation.ClearFromName()
	return mcuo
}

// SetFromEmail sets the "from_email" field.
func (mcuo *MailConnUpdateOne) SetFromEmail(s string) *MailConnUpdateOne {
	mcuo.mutation.SetFromEmail(s)
	return mcuo
}

// SetNillableFromEmail sets the "from_email" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillableFromEmail(s *string) *MailConnUpdateOne {
	if s != nil {
		mcuo.SetFromEmail(*s)
	}
	return mcuo
}

// ClearFromEmail clears the value of the "from_email" field.
func (mcuo *MailConnUpdateOne) ClearFromEmail() *MailConnUpdateOne {
	mcuo.mutation.ClearFromEmail()
	return mcuo
}

// SetStatus sets the "status" field.
func (mcuo *MailConnUpdateOne) SetStatus(b bool) *MailConnUpdateOne {
	mcuo.mutation.SetStatus(b)
	return mcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mcuo *MailConnUpdateOne) SetNillableStatus(b *bool) *MailConnUpdateOne {
	if b != nil {
		mcuo.SetStatus(*b)
	}
	return mcuo
}

// ClearStatus clears the value of the "status" field.
func (mcuo *MailConnUpdateOne) ClearStatus() *MailConnUpdateOne {
	mcuo.mutation.ClearStatus()
	return mcuo
}

// Mutation returns the MailConnMutation object of the builder.
func (mcuo *MailConnUpdateOne) Mutation() *MailConnMutation {
	return mcuo.mutation
}

// Where appends a list predicates to the MailConnUpdate builder.
func (mcuo *MailConnUpdateOne) Where(ps ...predicate.MailConn) *MailConnUpdateOne {
	mcuo.mutation.Where(ps...)
	return mcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mcuo *MailConnUpdateOne) Select(field string, fields ...string) *MailConnUpdateOne {
	mcuo.fields = append([]string{field}, fields...)
	return mcuo
}

// Save executes the query and returns the updated MailConn entity.
func (mcuo *MailConnUpdateOne) Save(ctx context.Context) (*MailConn, error) {
	mcuo.defaults()
	return withHooks(ctx, mcuo.sqlSave, mcuo.mutation, mcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mcuo *MailConnUpdateOne) SaveX(ctx context.Context) *MailConn {
	node, err := mcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mcuo *MailConnUpdateOne) Exec(ctx context.Context) error {
	_, err := mcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcuo *MailConnUpdateOne) ExecX(ctx context.Context) {
	if err := mcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mcuo *MailConnUpdateOne) defaults() {
	if _, ok := mcuo.mutation.UpdatedAt(); !ok && !mcuo.mutation.UpdatedAtCleared() {
		v := mailconn.UpdateDefaultUpdatedAt()
		mcuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mcuo *MailConnUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MailConnUpdateOne {
	mcuo.modifiers = append(mcuo.modifiers, modifiers...)
	return mcuo
}

func (mcuo *MailConnUpdateOne) sqlSave(ctx context.Context) (_node *MailConn, err error) {
	_spec := sqlgraph.NewUpdateSpec(mailconn.Table, mailconn.Columns, sqlgraph.NewFieldSpec(mailconn.FieldID, field.TypeString))
	id, ok := mcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MailConn.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mailconn.FieldID)
		for _, f := range fields {
			if !mailconn.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mailconn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mcuo.mutation.CreatedAtCleared() {
		_spec.ClearField(mailconn.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(mailconn.FieldUpdatedAt, field.TypeTime, value)
	}
	if mcuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(mailconn.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mcuo.mutation.Name(); ok {
		_spec.SetField(mailconn.FieldName, field.TypeString, value)
	}
	if mcuo.mutation.NameCleared() {
		_spec.ClearField(mailconn.FieldName, field.TypeString)
	}
	if value, ok := mcuo.mutation.Host(); ok {
		_spec.SetField(mailconn.FieldHost, field.TypeString, value)
	}
	if mcuo.mutation.HostCleared() {
		_spec.ClearField(mailconn.FieldHost, field.TypeString)
	}
	if value, ok := mcuo.mutation.Port(); ok {
		_spec.SetField(mailconn.FieldPort, field.TypeInt, value)
	}
	if value, ok := mcuo.mutation.AddedPort(); ok {
		_spec.AddField(mailconn.FieldPort, field.TypeInt, value)
	}
	if mcuo.mutation.PortCleared() {
		_spec.ClearField(mailconn.FieldPort, field.TypeInt)
	}
	if value, ok := mcuo.mutation.Username(); ok {
		_spec.SetField(mailconn.FieldUsername, field.TypeString, value)
	}
	if mcuo.mutation.UsernameCleared() {
		_spec.ClearField(mailconn.FieldUsername, field.TypeString)
	}
	if value, ok := mcuo.mutation.Password(); ok {
		_spec.SetField(mailconn.FieldPassword, field.TypeString, value)
	}
	if mcuo.mutation.PasswordCleared() {
		_spec.ClearField(mailconn.FieldPassword, field.TypeString)
	}
	if value, ok := mcuo.mutation.Encryption(); ok {
		_spec.SetField(mailconn.FieldEncryption, field.TypeInt, value)
	}
	if value, ok := mcuo.mutation.AddedEncryption(); ok {
		_spec.AddField(mailconn.FieldEncryption, field.TypeInt, value)
	}
	if mcuo.mutation.EncryptionCleared() {
		_spec.ClearField(mailconn.FieldEncryption, field.TypeInt)
	}
	if value, ok := mcuo.mutation.FromName(); ok {
		_spec.SetField(mailconn.FieldFromName, field.TypeString, value)
	}
	if mcuo.mutation.FromNameCleared() {
		_spec.ClearField(mailconn.FieldFromName, field.TypeString)
	}
	if value, ok := mcuo.mutation.FromEmail(); ok {
		_spec.SetField(mailconn.FieldFromEmail, field.TypeString, value)
	}
	if mcuo.mutation.FromEmailCleared() {
		_spec.ClearField(mailconn.FieldFromEmail, field.TypeString)
	}
	if value, ok := mcuo.mutation.Status(); ok {
		_spec.SetField(mailconn.FieldStatus, field.TypeBool, value)
	}
	if mcuo.mutation.StatusCleared() {
		_spec.ClearField(mailconn.FieldStatus, field.TypeBool)
	}
	_spec.AddModifiers(mcuo.modifiers...)
	_node = &MailConn{config: mcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mailconn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mcuo.mutation.done = true
	return _node, nil
}