// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/gen/ent/adminuser"
	"saas/gen/ent/app"
	"saas/gen/ent/appuser"
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminUserUpdate is the builder for updating AdminUser entities.
type AdminUserUpdate struct {
	config
	hooks     []Hook
	mutation  *AdminUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AdminUserUpdate builder.
func (auu *AdminUserUpdate) Where(ps ...predicate.AdminUser) *AdminUserUpdate {
	auu.mutation.Where(ps...)
	return auu
}

// SetUpdatedAt sets the "updated_at" field.
func (auu *AdminUserUpdate) SetUpdatedAt(t time.Time) *AdminUserUpdate {
	auu.mutation.SetUpdatedAt(t)
	return auu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auu *AdminUserUpdate) ClearUpdatedAt() *AdminUserUpdate {
	auu.mutation.ClearUpdatedAt()
	return auu
}

// SetEmail sets the "email" field.
func (auu *AdminUserUpdate) SetEmail(s string) *AdminUserUpdate {
	auu.mutation.SetEmail(s)
	return auu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableEmail(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetEmail(*s)
	}
	return auu
}

// SetPhone sets the "phone" field.
func (auu *AdminUserUpdate) SetPhone(s string) *AdminUserUpdate {
	auu.mutation.SetPhone(s)
	return auu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillablePhone(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetPhone(*s)
	}
	return auu
}

// ClearPhone clears the value of the "phone" field.
func (auu *AdminUserUpdate) ClearPhone() *AdminUserUpdate {
	auu.mutation.ClearPhone()
	return auu
}

// SetFirstName sets the "first_name" field.
func (auu *AdminUserUpdate) SetFirstName(s string) *AdminUserUpdate {
	auu.mutation.SetFirstName(s)
	return auu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableFirstName(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetFirstName(*s)
	}
	return auu
}

// ClearFirstName clears the value of the "first_name" field.
func (auu *AdminUserUpdate) ClearFirstName() *AdminUserUpdate {
	auu.mutation.ClearFirstName()
	return auu
}

// SetLastName sets the "last_name" field.
func (auu *AdminUserUpdate) SetLastName(s string) *AdminUserUpdate {
	auu.mutation.SetLastName(s)
	return auu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableLastName(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetLastName(*s)
	}
	return auu
}

// ClearLastName clears the value of the "last_name" field.
func (auu *AdminUserUpdate) ClearLastName() *AdminUserUpdate {
	auu.mutation.ClearLastName()
	return auu
}

// SetCompany sets the "company" field.
func (auu *AdminUserUpdate) SetCompany(s string) *AdminUserUpdate {
	auu.mutation.SetCompany(s)
	return auu
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableCompany(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetCompany(*s)
	}
	return auu
}

// ClearCompany clears the value of the "company" field.
func (auu *AdminUserUpdate) ClearCompany() *AdminUserUpdate {
	auu.mutation.ClearCompany()
	return auu
}

// SetLocale sets the "locale" field.
func (auu *AdminUserUpdate) SetLocale(s string) *AdminUserUpdate {
	auu.mutation.SetLocale(s)
	return auu
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableLocale(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetLocale(*s)
	}
	return auu
}

// ClearLocale clears the value of the "locale" field.
func (auu *AdminUserUpdate) ClearLocale() *AdminUserUpdate {
	auu.mutation.ClearLocale()
	return auu
}

// SetRoleID sets the "role_id" field.
func (auu *AdminUserUpdate) SetRoleID(s string) *AdminUserUpdate {
	auu.mutation.SetRoleID(s)
	return auu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableRoleID(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetRoleID(*s)
	}
	return auu
}

// ClearRoleID clears the value of the "role_id" field.
func (auu *AdminUserUpdate) ClearRoleID() *AdminUserUpdate {
	auu.mutation.ClearRoleID()
	return auu
}

// SetStatus sets the "status" field.
func (auu *AdminUserUpdate) SetStatus(b bool) *AdminUserUpdate {
	auu.mutation.SetStatus(b)
	return auu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableStatus(b *bool) *AdminUserUpdate {
	if b != nil {
		auu.SetStatus(*b)
	}
	return auu
}

// ClearStatus clears the value of the "status" field.
func (auu *AdminUserUpdate) ClearStatus() *AdminUserUpdate {
	auu.mutation.ClearStatus()
	return auu
}

// SetPassword sets the "password" field.
func (auu *AdminUserUpdate) SetPassword(s string) *AdminUserUpdate {
	auu.mutation.SetPassword(s)
	return auu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillablePassword(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetPassword(*s)
	}
	return auu
}

// ClearPassword clears the value of the "password" field.
func (auu *AdminUserUpdate) ClearPassword() *AdminUserUpdate {
	auu.mutation.ClearPassword()
	return auu
}

// SetSecret sets the "secret" field.
func (auu *AdminUserUpdate) SetSecret(s string) *AdminUserUpdate {
	auu.mutation.SetSecret(s)
	return auu
}

// SetNillableSecret sets the "secret" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableSecret(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetSecret(*s)
	}
	return auu
}

// ClearSecret clears the value of the "secret" field.
func (auu *AdminUserUpdate) ClearSecret() *AdminUserUpdate {
	auu.mutation.ClearSecret()
	return auu
}

// SetAPIKey sets the "api_key" field.
func (auu *AdminUserUpdate) SetAPIKey(s string) *AdminUserUpdate {
	auu.mutation.SetAPIKey(s)
	return auu
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableAPIKey(s *string) *AdminUserUpdate {
	if s != nil {
		auu.SetAPIKey(*s)
	}
	return auu
}

// ClearAPIKey clears the value of the "api_key" field.
func (auu *AdminUserUpdate) ClearAPIKey() *AdminUserUpdate {
	auu.mutation.ClearAPIKey()
	return auu
}

// SetWelcomeEmailSent sets the "welcome_email_sent" field.
func (auu *AdminUserUpdate) SetWelcomeEmailSent(b bool) *AdminUserUpdate {
	auu.mutation.SetWelcomeEmailSent(b)
	return auu
}

// SetNillableWelcomeEmailSent sets the "welcome_email_sent" field if the given value is not nil.
func (auu *AdminUserUpdate) SetNillableWelcomeEmailSent(b *bool) *AdminUserUpdate {
	if b != nil {
		auu.SetWelcomeEmailSent(*b)
	}
	return auu
}

// ClearWelcomeEmailSent clears the value of the "welcome_email_sent" field.
func (auu *AdminUserUpdate) ClearWelcomeEmailSent() *AdminUserUpdate {
	auu.mutation.ClearWelcomeEmailSent()
	return auu
}

// AddAppIDs adds the "apps" edge to the App entity by IDs.
func (auu *AdminUserUpdate) AddAppIDs(ids ...string) *AdminUserUpdate {
	auu.mutation.AddAppIDs(ids...)
	return auu
}

// AddApps adds the "apps" edges to the App entity.
func (auu *AdminUserUpdate) AddApps(a ...*App) *AdminUserUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auu.AddAppIDs(ids...)
}

// AddAppUserIDs adds the "app_users" edge to the AppUser entity by IDs.
func (auu *AdminUserUpdate) AddAppUserIDs(ids ...string) *AdminUserUpdate {
	auu.mutation.AddAppUserIDs(ids...)
	return auu
}

// AddAppUsers adds the "app_users" edges to the AppUser entity.
func (auu *AdminUserUpdate) AddAppUsers(a ...*AppUser) *AdminUserUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auu.AddAppUserIDs(ids...)
}

// Mutation returns the AdminUserMutation object of the builder.
func (auu *AdminUserUpdate) Mutation() *AdminUserMutation {
	return auu.mutation
}

// ClearApps clears all "apps" edges to the App entity.
func (auu *AdminUserUpdate) ClearApps() *AdminUserUpdate {
	auu.mutation.ClearApps()
	return auu
}

// RemoveAppIDs removes the "apps" edge to App entities by IDs.
func (auu *AdminUserUpdate) RemoveAppIDs(ids ...string) *AdminUserUpdate {
	auu.mutation.RemoveAppIDs(ids...)
	return auu
}

// RemoveApps removes "apps" edges to App entities.
func (auu *AdminUserUpdate) RemoveApps(a ...*App) *AdminUserUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auu.RemoveAppIDs(ids...)
}

// ClearAppUsers clears all "app_users" edges to the AppUser entity.
func (auu *AdminUserUpdate) ClearAppUsers() *AdminUserUpdate {
	auu.mutation.ClearAppUsers()
	return auu
}

// RemoveAppUserIDs removes the "app_users" edge to AppUser entities by IDs.
func (auu *AdminUserUpdate) RemoveAppUserIDs(ids ...string) *AdminUserUpdate {
	auu.mutation.RemoveAppUserIDs(ids...)
	return auu
}

// RemoveAppUsers removes "app_users" edges to AppUser entities.
func (auu *AdminUserUpdate) RemoveAppUsers(a ...*AppUser) *AdminUserUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auu.RemoveAppUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auu *AdminUserUpdate) Save(ctx context.Context) (int, error) {
	auu.defaults()
	return withHooks(ctx, auu.sqlSave, auu.mutation, auu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auu *AdminUserUpdate) SaveX(ctx context.Context) int {
	affected, err := auu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auu *AdminUserUpdate) Exec(ctx context.Context) error {
	_, err := auu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auu *AdminUserUpdate) ExecX(ctx context.Context) {
	if err := auu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auu *AdminUserUpdate) defaults() {
	if _, ok := auu.mutation.UpdatedAt(); !ok && !auu.mutation.UpdatedAtCleared() {
		v := adminuser.UpdateDefaultUpdatedAt()
		auu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auu *AdminUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdminUserUpdate {
	auu.modifiers = append(auu.modifiers, modifiers...)
	return auu
}

func (auu *AdminUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminuser.Table, adminuser.Columns, sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeString))
	if ps := auu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auu.mutation.CreatedAtCleared() {
		_spec.ClearField(adminuser.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := auu.mutation.UpdatedAt(); ok {
		_spec.SetField(adminuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if auu.mutation.UpdatedAtCleared() {
		_spec.ClearField(adminuser.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auu.mutation.Email(); ok {
		_spec.SetField(adminuser.FieldEmail, field.TypeString, value)
	}
	if value, ok := auu.mutation.Phone(); ok {
		_spec.SetField(adminuser.FieldPhone, field.TypeString, value)
	}
	if auu.mutation.PhoneCleared() {
		_spec.ClearField(adminuser.FieldPhone, field.TypeString)
	}
	if value, ok := auu.mutation.FirstName(); ok {
		_spec.SetField(adminuser.FieldFirstName, field.TypeString, value)
	}
	if auu.mutation.FirstNameCleared() {
		_spec.ClearField(adminuser.FieldFirstName, field.TypeString)
	}
	if value, ok := auu.mutation.LastName(); ok {
		_spec.SetField(adminuser.FieldLastName, field.TypeString, value)
	}
	if auu.mutation.LastNameCleared() {
		_spec.ClearField(adminuser.FieldLastName, field.TypeString)
	}
	if value, ok := auu.mutation.Company(); ok {
		_spec.SetField(adminuser.FieldCompany, field.TypeString, value)
	}
	if auu.mutation.CompanyCleared() {
		_spec.ClearField(adminuser.FieldCompany, field.TypeString)
	}
	if value, ok := auu.mutation.Locale(); ok {
		_spec.SetField(adminuser.FieldLocale, field.TypeString, value)
	}
	if auu.mutation.LocaleCleared() {
		_spec.ClearField(adminuser.FieldLocale, field.TypeString)
	}
	if value, ok := auu.mutation.RoleID(); ok {
		_spec.SetField(adminuser.FieldRoleID, field.TypeString, value)
	}
	if auu.mutation.RoleIDCleared() {
		_spec.ClearField(adminuser.FieldRoleID, field.TypeString)
	}
	if value, ok := auu.mutation.Status(); ok {
		_spec.SetField(adminuser.FieldStatus, field.TypeBool, value)
	}
	if auu.mutation.StatusCleared() {
		_spec.ClearField(adminuser.FieldStatus, field.TypeBool)
	}
	if value, ok := auu.mutation.Password(); ok {
		_spec.SetField(adminuser.FieldPassword, field.TypeString, value)
	}
	if auu.mutation.PasswordCleared() {
		_spec.ClearField(adminuser.FieldPassword, field.TypeString)
	}
	if value, ok := auu.mutation.Secret(); ok {
		_spec.SetField(adminuser.FieldSecret, field.TypeString, value)
	}
	if auu.mutation.SecretCleared() {
		_spec.ClearField(adminuser.FieldSecret, field.TypeString)
	}
	if value, ok := auu.mutation.APIKey(); ok {
		_spec.SetField(adminuser.FieldAPIKey, field.TypeString, value)
	}
	if auu.mutation.APIKeyCleared() {
		_spec.ClearField(adminuser.FieldAPIKey, field.TypeString)
	}
	if value, ok := auu.mutation.WelcomeEmailSent(); ok {
		_spec.SetField(adminuser.FieldWelcomeEmailSent, field.TypeBool, value)
	}
	if auu.mutation.WelcomeEmailSentCleared() {
		_spec.ClearField(adminuser.FieldWelcomeEmailSent, field.TypeBool)
	}
	if auu.mutation.AppsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adminuser.AppsTable,
			Columns: adminuser.AppsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeString),
			},
		}
		createE := &AppUserCreate{config: auu.config, mutation: newAppUserMutation(auu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auu.mutation.RemovedAppsIDs(); len(nodes) > 0 && !auu.mutation.AppsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adminuser.AppsTable,
			Columns: adminuser.AppsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppUserCreate{config: auu.config, mutation: newAppUserMutation(auu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auu.mutation.AppsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adminuser.AppsTable,
			Columns: adminuser.AppsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppUserCreate{config: auu.config, mutation: newAppUserMutation(auu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auu.mutation.AppUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adminuser.AppUsersTable,
			Columns: []string{adminuser.AppUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auu.mutation.RemovedAppUsersIDs(); len(nodes) > 0 && !auu.mutation.AppUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adminuser.AppUsersTable,
			Columns: []string{adminuser.AppUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auu.mutation.AppUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adminuser.AppUsersTable,
			Columns: []string{adminuser.AppUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(auu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, auu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	auu.mutation.done = true
	return n, nil
}

// AdminUserUpdateOne is the builder for updating a single AdminUser entity.
type AdminUserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AdminUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (auuo *AdminUserUpdateOne) SetUpdatedAt(t time.Time) *AdminUserUpdateOne {
	auuo.mutation.SetUpdatedAt(t)
	return auuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auuo *AdminUserUpdateOne) ClearUpdatedAt() *AdminUserUpdateOne {
	auuo.mutation.ClearUpdatedAt()
	return auuo
}

// SetEmail sets the "email" field.
func (auuo *AdminUserUpdateOne) SetEmail(s string) *AdminUserUpdateOne {
	auuo.mutation.SetEmail(s)
	return auuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableEmail(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetEmail(*s)
	}
	return auuo
}

// SetPhone sets the "phone" field.
func (auuo *AdminUserUpdateOne) SetPhone(s string) *AdminUserUpdateOne {
	auuo.mutation.SetPhone(s)
	return auuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillablePhone(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetPhone(*s)
	}
	return auuo
}

// ClearPhone clears the value of the "phone" field.
func (auuo *AdminUserUpdateOne) ClearPhone() *AdminUserUpdateOne {
	auuo.mutation.ClearPhone()
	return auuo
}

// SetFirstName sets the "first_name" field.
func (auuo *AdminUserUpdateOne) SetFirstName(s string) *AdminUserUpdateOne {
	auuo.mutation.SetFirstName(s)
	return auuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableFirstName(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetFirstName(*s)
	}
	return auuo
}

// ClearFirstName clears the value of the "first_name" field.
func (auuo *AdminUserUpdateOne) ClearFirstName() *AdminUserUpdateOne {
	auuo.mutation.ClearFirstName()
	return auuo
}

// SetLastName sets the "last_name" field.
func (auuo *AdminUserUpdateOne) SetLastName(s string) *AdminUserUpdateOne {
	auuo.mutation.SetLastName(s)
	return auuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableLastName(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetLastName(*s)
	}
	return auuo
}

// ClearLastName clears the value of the "last_name" field.
func (auuo *AdminUserUpdateOne) ClearLastName() *AdminUserUpdateOne {
	auuo.mutation.ClearLastName()
	return auuo
}

// SetCompany sets the "company" field.
func (auuo *AdminUserUpdateOne) SetCompany(s string) *AdminUserUpdateOne {
	auuo.mutation.SetCompany(s)
	return auuo
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableCompany(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetCompany(*s)
	}
	return auuo
}

// ClearCompany clears the value of the "company" field.
func (auuo *AdminUserUpdateOne) ClearCompany() *AdminUserUpdateOne {
	auuo.mutation.ClearCompany()
	return auuo
}

// SetLocale sets the "locale" field.
func (auuo *AdminUserUpdateOne) SetLocale(s string) *AdminUserUpdateOne {
	auuo.mutation.SetLocale(s)
	return auuo
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableLocale(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetLocale(*s)
	}
	return auuo
}

// ClearLocale clears the value of the "locale" field.
func (auuo *AdminUserUpdateOne) ClearLocale() *AdminUserUpdateOne {
	auuo.mutation.ClearLocale()
	return auuo
}

// SetRoleID sets the "role_id" field.
func (auuo *AdminUserUpdateOne) SetRoleID(s string) *AdminUserUpdateOne {
	auuo.mutation.SetRoleID(s)
	return auuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableRoleID(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetRoleID(*s)
	}
	return auuo
}

// ClearRoleID clears the value of the "role_id" field.
func (auuo *AdminUserUpdateOne) ClearRoleID() *AdminUserUpdateOne {
	auuo.mutation.ClearRoleID()
	return auuo
}

// SetStatus sets the "status" field.
func (auuo *AdminUserUpdateOne) SetStatus(b bool) *AdminUserUpdateOne {
	auuo.mutation.SetStatus(b)
	return auuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableStatus(b *bool) *AdminUserUpdateOne {
	if b != nil {
		auuo.SetStatus(*b)
	}
	return auuo
}

// ClearStatus clears the value of the "status" field.
func (auuo *AdminUserUpdateOne) ClearStatus() *AdminUserUpdateOne {
	auuo.mutation.ClearStatus()
	return auuo
}

// SetPassword sets the "password" field.
func (auuo *AdminUserUpdateOne) SetPassword(s string) *AdminUserUpdateOne {
	auuo.mutation.SetPassword(s)
	return auuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillablePassword(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetPassword(*s)
	}
	return auuo
}

// ClearPassword clears the value of the "password" field.
func (auuo *AdminUserUpdateOne) ClearPassword() *AdminUserUpdateOne {
	auuo.mutation.ClearPassword()
	return auuo
}

// SetSecret sets the "secret" field.
func (auuo *AdminUserUpdateOne) SetSecret(s string) *AdminUserUpdateOne {
	auuo.mutation.SetSecret(s)
	return auuo
}

// SetNillableSecret sets the "secret" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableSecret(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetSecret(*s)
	}
	return auuo
}

// ClearSecret clears the value of the "secret" field.
func (auuo *AdminUserUpdateOne) ClearSecret() *AdminUserUpdateOne {
	auuo.mutation.ClearSecret()
	return auuo
}

// SetAPIKey sets the "api_key" field.
func (auuo *AdminUserUpdateOne) SetAPIKey(s string) *AdminUserUpdateOne {
	auuo.mutation.SetAPIKey(s)
	return auuo
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableAPIKey(s *string) *AdminUserUpdateOne {
	if s != nil {
		auuo.SetAPIKey(*s)
	}
	return auuo
}

// ClearAPIKey clears the value of the "api_key" field.
func (auuo *AdminUserUpdateOne) ClearAPIKey() *AdminUserUpdateOne {
	auuo.mutation.ClearAPIKey()
	return auuo
}

// SetWelcomeEmailSent sets the "welcome_email_sent" field.
func (auuo *AdminUserUpdateOne) SetWelcomeEmailSent(b bool) *AdminUserUpdateOne {
	auuo.mutation.SetWelcomeEmailSent(b)
	return auuo
}

// SetNillableWelcomeEmailSent sets the "welcome_email_sent" field if the given value is not nil.
func (auuo *AdminUserUpdateOne) SetNillableWelcomeEmailSent(b *bool) *AdminUserUpdateOne {
	if b != nil {
		auuo.SetWelcomeEmailSent(*b)
	}
	return auuo
}

// ClearWelcomeEmailSent clears the value of the "welcome_email_sent" field.
func (auuo *AdminUserUpdateOne) ClearWelcomeEmailSent() *AdminUserUpdateOne {
	auuo.mutation.ClearWelcomeEmailSent()
	return auuo
}

// AddAppIDs adds the "apps" edge to the App entity by IDs.
func (auuo *AdminUserUpdateOne) AddAppIDs(ids ...string) *AdminUserUpdateOne {
	auuo.mutation.AddAppIDs(ids...)
	return auuo
}

// AddApps adds the "apps" edges to the App entity.
func (auuo *AdminUserUpdateOne) AddApps(a ...*App) *AdminUserUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auuo.AddAppIDs(ids...)
}

// AddAppUserIDs adds the "app_users" edge to the AppUser entity by IDs.
func (auuo *AdminUserUpdateOne) AddAppUserIDs(ids ...string) *AdminUserUpdateOne {
	auuo.mutation.AddAppUserIDs(ids...)
	return auuo
}

// AddAppUsers adds the "app_users" edges to the AppUser entity.
func (auuo *AdminUserUpdateOne) AddAppUsers(a ...*AppUser) *AdminUserUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auuo.AddAppUserIDs(ids...)
}

// Mutation returns the AdminUserMutation object of the builder.
func (auuo *AdminUserUpdateOne) Mutation() *AdminUserMutation {
	return auuo.mutation
}

// ClearApps clears all "apps" edges to the App entity.
func (auuo *AdminUserUpdateOne) ClearApps() *AdminUserUpdateOne {
	auuo.mutation.ClearApps()
	return auuo
}

// RemoveAppIDs removes the "apps" edge to App entities by IDs.
func (auuo *AdminUserUpdateOne) RemoveAppIDs(ids ...string) *AdminUserUpdateOne {
	auuo.mutation.RemoveAppIDs(ids...)
	return auuo
}

// RemoveApps removes "apps" edges to App entities.
func (auuo *AdminUserUpdateOne) RemoveApps(a ...*App) *AdminUserUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auuo.RemoveAppIDs(ids...)
}

// ClearAppUsers clears all "app_users" edges to the AppUser entity.
func (auuo *AdminUserUpdateOne) ClearAppUsers() *AdminUserUpdateOne {
	auuo.mutation.ClearAppUsers()
	return auuo
}

// RemoveAppUserIDs removes the "app_users" edge to AppUser entities by IDs.
func (auuo *AdminUserUpdateOne) RemoveAppUserIDs(ids ...string) *AdminUserUpdateOne {
	auuo.mutation.RemoveAppUserIDs(ids...)
	return auuo
}

// RemoveAppUsers removes "app_users" edges to AppUser entities.
func (auuo *AdminUserUpdateOne) RemoveAppUsers(a ...*AppUser) *AdminUserUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auuo.RemoveAppUserIDs(ids...)
}

// Where appends a list predicates to the AdminUserUpdate builder.
func (auuo *AdminUserUpdateOne) Where(ps ...predicate.AdminUser) *AdminUserUpdateOne {
	auuo.mutation.Where(ps...)
	return auuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auuo *AdminUserUpdateOne) Select(field string, fields ...string) *AdminUserUpdateOne {
	auuo.fields = append([]string{field}, fields...)
	return auuo
}

// Save executes the query and returns the updated AdminUser entity.
func (auuo *AdminUserUpdateOne) Save(ctx context.Context) (*AdminUser, error) {
	auuo.defaults()
	return withHooks(ctx, auuo.sqlSave, auuo.mutation, auuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auuo *AdminUserUpdateOne) SaveX(ctx context.Context) *AdminUser {
	node, err := auuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auuo *AdminUserUpdateOne) Exec(ctx context.Context) error {
	_, err := auuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auuo *AdminUserUpdateOne) ExecX(ctx context.Context) {
	if err := auuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auuo *AdminUserUpdateOne) defaults() {
	if _, ok := auuo.mutation.UpdatedAt(); !ok && !auuo.mutation.UpdatedAtCleared() {
		v := adminuser.UpdateDefaultUpdatedAt()
		auuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auuo *AdminUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdminUserUpdateOne {
	auuo.modifiers = append(auuo.modifiers, modifiers...)
	return auuo
}

func (auuo *AdminUserUpdateOne) sqlSave(ctx context.Context) (_node *AdminUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(adminuser.Table, adminuser.Columns, sqlgraph.NewFieldSpec(adminuser.FieldID, field.TypeString))
	id, ok := auuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdminUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminuser.FieldID)
		for _, f := range fields {
			if !adminuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adminuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auuo.mutation.CreatedAtCleared() {
		_spec.ClearField(adminuser.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := auuo.mutation.UpdatedAt(); ok {
		_spec.SetField(adminuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if auuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(adminuser.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auuo.mutation.Email(); ok {
		_spec.SetField(adminuser.FieldEmail, field.TypeString, value)
	}
	if value, ok := auuo.mutation.Phone(); ok {
		_spec.SetField(adminuser.FieldPhone, field.TypeString, value)
	}
	if auuo.mutation.PhoneCleared() {
		_spec.ClearField(adminuser.FieldPhone, field.TypeString)
	}
	if value, ok := auuo.mutation.FirstName(); ok {
		_spec.SetField(adminuser.FieldFirstName, field.TypeString, value)
	}
	if auuo.mutation.FirstNameCleared() {
		_spec.ClearField(adminuser.FieldFirstName, field.TypeString)
	}
	if value, ok := auuo.mutation.LastName(); ok {
		_spec.SetField(adminuser.FieldLastName, field.TypeString, value)
	}
	if auuo.mutation.LastNameCleared() {
		_spec.ClearField(adminuser.FieldLastName, field.TypeString)
	}
	if value, ok := auuo.mutation.Company(); ok {
		_spec.SetField(adminuser.FieldCompany, field.TypeString, value)
	}
	if auuo.mutation.CompanyCleared() {
		_spec.ClearField(adminuser.FieldCompany, field.TypeString)
	}
	if value, ok := auuo.mutation.Locale(); ok {
		_spec.SetField(adminuser.FieldLocale, field.TypeString, value)
	}
	if auuo.mutation.LocaleCleared() {
		_spec.ClearField(adminuser.FieldLocale, field.TypeString)
	}
	if value, ok := auuo.mutation.RoleID(); ok {
		_spec.SetField(adminuser.FieldRoleID, field.TypeString, value)
	}
	if auuo.mutation.RoleIDCleared() {
		_spec.ClearField(adminuser.FieldRoleID, field.TypeString)
	}
	if value, ok := auuo.mutation.Status(); ok {
		_spec.SetField(adminuser.FieldStatus, field.TypeBool, value)
	}
	if auuo.mutation.StatusCleared() {
		_spec.ClearField(adminuser.FieldStatus, field.TypeBool)
	}
	if value, ok := auuo.mutation.Password(); ok {
		_spec.SetField(adminuser.FieldPassword, field.TypeString, value)
	}
	if auuo.mutation.PasswordCleared() {
		_spec.ClearField(adminuser.FieldPassword, field.TypeString)
	}
	if value, ok := auuo.mutation.Secret(); ok {
		_spec.SetField(adminuser.FieldSecret, field.TypeString, value)
	}
	if auuo.mutation.SecretCleared() {
		_spec.ClearField(adminuser.FieldSecret, field.TypeString)
	}
	if value, ok := auuo.mutation.APIKey(); ok {
		_spec.SetField(adminuser.FieldAPIKey, field.TypeString, value)
	}
	if auuo.mutation.APIKeyCleared() {
		_spec.ClearField(adminuser.FieldAPIKey, field.TypeString)
	}
	if value, ok := auuo.mutation.WelcomeEmailSent(); ok {
		_spec.SetField(adminuser.FieldWelcomeEmailSent, field.TypeBool, value)
	}
	if auuo.mutation.WelcomeEmailSentCleared() {
		_spec.ClearField(adminuser.FieldWelcomeEmailSent, field.TypeBool)
	}
	if auuo.mutation.AppsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adminuser.AppsTable,
			Columns: adminuser.AppsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeString),
			},
		}
		createE := &AppUserCreate{config: auuo.config, mutation: newAppUserMutation(auuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auuo.mutation.RemovedAppsIDs(); len(nodes) > 0 && !auuo.mutation.AppsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adminuser.AppsTable,
			Columns: adminuser.AppsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppUserCreate{config: auuo.config, mutation: newAppUserMutation(auuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auuo.mutation.AppsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adminuser.AppsTable,
			Columns: adminuser.AppsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppUserCreate{config: auuo.config, mutation: newAppUserMutation(auuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auuo.mutation.AppUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adminuser.AppUsersTable,
			Columns: []string{adminuser.AppUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appuser.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auuo.mutation.RemovedAppUsersIDs(); len(nodes) > 0 && !auuo.mutation.AppUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adminuser.AppUsersTable,
			Columns: []string{adminuser.AppUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auuo.mutation.AppUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   adminuser.AppUsersTable,
			Columns: []string{adminuser.AppUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(auuo.modifiers...)
	_node = &AdminUser{config: auuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adminuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auuo.mutation.done = true
	return _node, nil
}
