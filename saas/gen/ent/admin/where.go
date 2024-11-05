// Code generated by ent, DO NOT EDIT.

package admin

import (
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldAppID, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldEmail, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldLastName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldStatus, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldPassword, v))
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldSecret, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldUpdatedAt))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldAppID))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldAppID, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldEmail, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameIsNil applies the IsNil predicate on the "first_name" field.
func FirstNameIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldFirstName))
}

// FirstNameNotNil applies the NotNil predicate on the "first_name" field.
func FirstNameNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldFirstName))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldLastName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldStatus))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldPassword, v))
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEQ(FieldSecret, v))
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v string) predicate.Admin {
	return predicate.Admin(sql.FieldNEQ(FieldSecret, v))
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldIn(FieldSecret, vs...))
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...string) predicate.Admin {
	return predicate.Admin(sql.FieldNotIn(FieldSecret, vs...))
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGT(FieldSecret, v))
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldGTE(FieldSecret, v))
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLT(FieldSecret, v))
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v string) predicate.Admin {
	return predicate.Admin(sql.FieldLTE(FieldSecret, v))
}

// SecretContains applies the Contains predicate on the "secret" field.
func SecretContains(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContains(FieldSecret, v))
}

// SecretHasPrefix applies the HasPrefix predicate on the "secret" field.
func SecretHasPrefix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasPrefix(FieldSecret, v))
}

// SecretHasSuffix applies the HasSuffix predicate on the "secret" field.
func SecretHasSuffix(v string) predicate.Admin {
	return predicate.Admin(sql.FieldHasSuffix(FieldSecret, v))
}

// SecretIsNil applies the IsNil predicate on the "secret" field.
func SecretIsNil() predicate.Admin {
	return predicate.Admin(sql.FieldIsNull(FieldSecret))
}

// SecretNotNil applies the NotNil predicate on the "secret" field.
func SecretNotNil() predicate.Admin {
	return predicate.Admin(sql.FieldNotNull(FieldSecret))
}

// SecretEqualFold applies the EqualFold predicate on the "secret" field.
func SecretEqualFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldEqualFold(FieldSecret, v))
}

// SecretContainsFold applies the ContainsFold predicate on the "secret" field.
func SecretContainsFold(v string) predicate.Admin {
	return predicate.Admin(sql.FieldContainsFold(FieldSecret, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Admin) predicate.Admin {
	return predicate.Admin(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Admin) predicate.Admin {
	return predicate.Admin(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Admin) predicate.Admin {
	return predicate.Admin(sql.NotPredicates(p))
}
