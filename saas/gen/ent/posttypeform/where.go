// Code generated by ent, DO NOT EDIT.

package posttypeform

import (
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldAppID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldName, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldStatus, v))
}

// PostTypeID applies equality check predicate on the "post_type_id" field. It's identical to PostTypeIDEQ.
func PostTypeID(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldPostTypeID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotNull(FieldUpdatedAt))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotNull(FieldAppID))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldContainsFold(FieldAppID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNEQ(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotNull(FieldStatus))
}

// PostTypeIDEQ applies the EQ predicate on the "post_type_id" field.
func PostTypeIDEQ(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEQ(FieldPostTypeID, v))
}

// PostTypeIDNEQ applies the NEQ predicate on the "post_type_id" field.
func PostTypeIDNEQ(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNEQ(FieldPostTypeID, v))
}

// PostTypeIDIn applies the In predicate on the "post_type_id" field.
func PostTypeIDIn(vs ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIn(FieldPostTypeID, vs...))
}

// PostTypeIDNotIn applies the NotIn predicate on the "post_type_id" field.
func PostTypeIDNotIn(vs ...string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotIn(FieldPostTypeID, vs...))
}

// PostTypeIDGT applies the GT predicate on the "post_type_id" field.
func PostTypeIDGT(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGT(FieldPostTypeID, v))
}

// PostTypeIDGTE applies the GTE predicate on the "post_type_id" field.
func PostTypeIDGTE(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldGTE(FieldPostTypeID, v))
}

// PostTypeIDLT applies the LT predicate on the "post_type_id" field.
func PostTypeIDLT(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLT(FieldPostTypeID, v))
}

// PostTypeIDLTE applies the LTE predicate on the "post_type_id" field.
func PostTypeIDLTE(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldLTE(FieldPostTypeID, v))
}

// PostTypeIDContains applies the Contains predicate on the "post_type_id" field.
func PostTypeIDContains(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldContains(FieldPostTypeID, v))
}

// PostTypeIDHasPrefix applies the HasPrefix predicate on the "post_type_id" field.
func PostTypeIDHasPrefix(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldHasPrefix(FieldPostTypeID, v))
}

// PostTypeIDHasSuffix applies the HasSuffix predicate on the "post_type_id" field.
func PostTypeIDHasSuffix(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldHasSuffix(FieldPostTypeID, v))
}

// PostTypeIDIsNil applies the IsNil predicate on the "post_type_id" field.
func PostTypeIDIsNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIsNull(FieldPostTypeID))
}

// PostTypeIDNotNil applies the NotNil predicate on the "post_type_id" field.
func PostTypeIDNotNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotNull(FieldPostTypeID))
}

// PostTypeIDEqualFold applies the EqualFold predicate on the "post_type_id" field.
func PostTypeIDEqualFold(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldEqualFold(FieldPostTypeID, v))
}

// PostTypeIDContainsFold applies the ContainsFold predicate on the "post_type_id" field.
func PostTypeIDContainsFold(v string) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldContainsFold(FieldPostTypeID, v))
}

// BodyIsNil applies the IsNil predicate on the "body" field.
func BodyIsNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldIsNull(FieldBody))
}

// BodyNotNil applies the NotNil predicate on the "body" field.
func BodyNotNil() predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.FieldNotNull(FieldBody))
}

// HasPostType applies the HasEdge predicate on the "post_type" edge.
func HasPostType() predicate.PostTypeForm {
	return predicate.PostTypeForm(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTypeTable, PostTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostTypeWith applies the HasEdge predicate on the "post_type" edge with a given conditions (other predicates).
func HasPostTypeWith(preds ...predicate.PostType) predicate.PostTypeForm {
	return predicate.PostTypeForm(func(s *sql.Selector) {
		step := newPostTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PostTypeForm) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PostTypeForm) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PostTypeForm) predicate.PostTypeForm {
	return predicate.PostTypeForm(sql.NotPredicates(p))
}