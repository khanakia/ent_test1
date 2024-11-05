// Code generated by ent, DO NOT EDIT.

package templ

import (
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Templ {
	return predicate.Templ(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Templ {
	return predicate.Templ(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Templ {
	return predicate.Templ(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Templ {
	return predicate.Templ(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Templ {
	return predicate.Templ(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Templ {
	return predicate.Templ(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Templ {
	return predicate.Templ(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Templ {
	return predicate.Templ(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldAppID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldName, v))
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldBody, v))
}

// Compiled applies equality check predicate on the "compiled" field. It's identical to CompiledEQ.
func Compiled(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldCompiled, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Templ {
	return predicate.Templ(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Templ {
	return predicate.Templ(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Templ {
	return predicate.Templ(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Templ {
	return predicate.Templ(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Templ {
	return predicate.Templ(sql.FieldNotNull(FieldUpdatedAt))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.Templ {
	return predicate.Templ(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.Templ {
	return predicate.Templ(sql.FieldNotNull(FieldAppID))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContainsFold(FieldAppID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Templ {
	return predicate.Templ(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Templ {
	return predicate.Templ(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContainsFold(FieldName, v))
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldBody, v))
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldBody, v))
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldIn(FieldBody, vs...))
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldNotIn(FieldBody, vs...))
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGT(FieldBody, v))
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGTE(FieldBody, v))
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLT(FieldBody, v))
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLTE(FieldBody, v))
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContains(FieldBody, v))
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasPrefix(FieldBody, v))
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasSuffix(FieldBody, v))
}

// BodyIsNil applies the IsNil predicate on the "body" field.
func BodyIsNil() predicate.Templ {
	return predicate.Templ(sql.FieldIsNull(FieldBody))
}

// BodyNotNil applies the NotNil predicate on the "body" field.
func BodyNotNil() predicate.Templ {
	return predicate.Templ(sql.FieldNotNull(FieldBody))
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEqualFold(FieldBody, v))
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContainsFold(FieldBody, v))
}

// CompiledEQ applies the EQ predicate on the "compiled" field.
func CompiledEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldCompiled, v))
}

// CompiledNEQ applies the NEQ predicate on the "compiled" field.
func CompiledNEQ(v string) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldCompiled, v))
}

// CompiledIn applies the In predicate on the "compiled" field.
func CompiledIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldIn(FieldCompiled, vs...))
}

// CompiledNotIn applies the NotIn predicate on the "compiled" field.
func CompiledNotIn(vs ...string) predicate.Templ {
	return predicate.Templ(sql.FieldNotIn(FieldCompiled, vs...))
}

// CompiledGT applies the GT predicate on the "compiled" field.
func CompiledGT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGT(FieldCompiled, v))
}

// CompiledGTE applies the GTE predicate on the "compiled" field.
func CompiledGTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldGTE(FieldCompiled, v))
}

// CompiledLT applies the LT predicate on the "compiled" field.
func CompiledLT(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLT(FieldCompiled, v))
}

// CompiledLTE applies the LTE predicate on the "compiled" field.
func CompiledLTE(v string) predicate.Templ {
	return predicate.Templ(sql.FieldLTE(FieldCompiled, v))
}

// CompiledContains applies the Contains predicate on the "compiled" field.
func CompiledContains(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContains(FieldCompiled, v))
}

// CompiledHasPrefix applies the HasPrefix predicate on the "compiled" field.
func CompiledHasPrefix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasPrefix(FieldCompiled, v))
}

// CompiledHasSuffix applies the HasSuffix predicate on the "compiled" field.
func CompiledHasSuffix(v string) predicate.Templ {
	return predicate.Templ(sql.FieldHasSuffix(FieldCompiled, v))
}

// CompiledIsNil applies the IsNil predicate on the "compiled" field.
func CompiledIsNil() predicate.Templ {
	return predicate.Templ(sql.FieldIsNull(FieldCompiled))
}

// CompiledNotNil applies the NotNil predicate on the "compiled" field.
func CompiledNotNil() predicate.Templ {
	return predicate.Templ(sql.FieldNotNull(FieldCompiled))
}

// CompiledEqualFold applies the EqualFold predicate on the "compiled" field.
func CompiledEqualFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldEqualFold(FieldCompiled, v))
}

// CompiledContainsFold applies the ContainsFold predicate on the "compiled" field.
func CompiledContainsFold(v string) predicate.Templ {
	return predicate.Templ(sql.FieldContainsFold(FieldCompiled, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Templ {
	return predicate.Templ(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Templ {
	return predicate.Templ(sql.FieldNEQ(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Templ {
	return predicate.Templ(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Templ {
	return predicate.Templ(sql.FieldNotNull(FieldStatus))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Templ) predicate.Templ {
	return predicate.Templ(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Templ) predicate.Templ {
	return predicate.Templ(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Templ) predicate.Templ {
	return predicate.Templ(sql.NotPredicates(p))
}
