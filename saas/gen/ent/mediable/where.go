// Code generated by ent, DO NOT EDIT.

package mediable

import (
	"saas/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContainsFold(FieldID, id))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldAppID, v))
}

// MediaID applies equality check predicate on the "media_id" field. It's identical to MediaIDEQ.
func MediaID(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldMediaID, v))
}

// MediableID applies equality check predicate on the "mediable_id" field. It's identical to MediableIDEQ.
func MediableID(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldMediableID, v))
}

// MediableType applies equality check predicate on the "mediable_type" field. It's identical to MediableTypeEQ.
func MediableType(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldMediableType, v))
}

// Tag applies equality check predicate on the "tag" field. It's identical to TagEQ.
func Tag(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldTag, v))
}

// SortOrder applies equality check predicate on the "sort_order" field. It's identical to SortOrderEQ.
func SortOrder(v int) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldSortOrder, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldNotNull(FieldAppID))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContainsFold(FieldAppID, v))
}

// MediaIDEQ applies the EQ predicate on the "media_id" field.
func MediaIDEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldMediaID, v))
}

// MediaIDNEQ applies the NEQ predicate on the "media_id" field.
func MediaIDNEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNEQ(FieldMediaID, v))
}

// MediaIDIn applies the In predicate on the "media_id" field.
func MediaIDIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldIn(FieldMediaID, vs...))
}

// MediaIDNotIn applies the NotIn predicate on the "media_id" field.
func MediaIDNotIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNotIn(FieldMediaID, vs...))
}

// MediaIDGT applies the GT predicate on the "media_id" field.
func MediaIDGT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGT(FieldMediaID, v))
}

// MediaIDGTE applies the GTE predicate on the "media_id" field.
func MediaIDGTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGTE(FieldMediaID, v))
}

// MediaIDLT applies the LT predicate on the "media_id" field.
func MediaIDLT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLT(FieldMediaID, v))
}

// MediaIDLTE applies the LTE predicate on the "media_id" field.
func MediaIDLTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLTE(FieldMediaID, v))
}

// MediaIDContains applies the Contains predicate on the "media_id" field.
func MediaIDContains(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContains(FieldMediaID, v))
}

// MediaIDHasPrefix applies the HasPrefix predicate on the "media_id" field.
func MediaIDHasPrefix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasPrefix(FieldMediaID, v))
}

// MediaIDHasSuffix applies the HasSuffix predicate on the "media_id" field.
func MediaIDHasSuffix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasSuffix(FieldMediaID, v))
}

// MediaIDIsNil applies the IsNil predicate on the "media_id" field.
func MediaIDIsNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldIsNull(FieldMediaID))
}

// MediaIDNotNil applies the NotNil predicate on the "media_id" field.
func MediaIDNotNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldNotNull(FieldMediaID))
}

// MediaIDEqualFold applies the EqualFold predicate on the "media_id" field.
func MediaIDEqualFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEqualFold(FieldMediaID, v))
}

// MediaIDContainsFold applies the ContainsFold predicate on the "media_id" field.
func MediaIDContainsFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContainsFold(FieldMediaID, v))
}

// MediableIDEQ applies the EQ predicate on the "mediable_id" field.
func MediableIDEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldMediableID, v))
}

// MediableIDNEQ applies the NEQ predicate on the "mediable_id" field.
func MediableIDNEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNEQ(FieldMediableID, v))
}

// MediableIDIn applies the In predicate on the "mediable_id" field.
func MediableIDIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldIn(FieldMediableID, vs...))
}

// MediableIDNotIn applies the NotIn predicate on the "mediable_id" field.
func MediableIDNotIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNotIn(FieldMediableID, vs...))
}

// MediableIDGT applies the GT predicate on the "mediable_id" field.
func MediableIDGT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGT(FieldMediableID, v))
}

// MediableIDGTE applies the GTE predicate on the "mediable_id" field.
func MediableIDGTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGTE(FieldMediableID, v))
}

// MediableIDLT applies the LT predicate on the "mediable_id" field.
func MediableIDLT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLT(FieldMediableID, v))
}

// MediableIDLTE applies the LTE predicate on the "mediable_id" field.
func MediableIDLTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLTE(FieldMediableID, v))
}

// MediableIDContains applies the Contains predicate on the "mediable_id" field.
func MediableIDContains(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContains(FieldMediableID, v))
}

// MediableIDHasPrefix applies the HasPrefix predicate on the "mediable_id" field.
func MediableIDHasPrefix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasPrefix(FieldMediableID, v))
}

// MediableIDHasSuffix applies the HasSuffix predicate on the "mediable_id" field.
func MediableIDHasSuffix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasSuffix(FieldMediableID, v))
}

// MediableIDIsNil applies the IsNil predicate on the "mediable_id" field.
func MediableIDIsNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldIsNull(FieldMediableID))
}

// MediableIDNotNil applies the NotNil predicate on the "mediable_id" field.
func MediableIDNotNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldNotNull(FieldMediableID))
}

// MediableIDEqualFold applies the EqualFold predicate on the "mediable_id" field.
func MediableIDEqualFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEqualFold(FieldMediableID, v))
}

// MediableIDContainsFold applies the ContainsFold predicate on the "mediable_id" field.
func MediableIDContainsFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContainsFold(FieldMediableID, v))
}

// MediableTypeEQ applies the EQ predicate on the "mediable_type" field.
func MediableTypeEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldMediableType, v))
}

// MediableTypeNEQ applies the NEQ predicate on the "mediable_type" field.
func MediableTypeNEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNEQ(FieldMediableType, v))
}

// MediableTypeIn applies the In predicate on the "mediable_type" field.
func MediableTypeIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldIn(FieldMediableType, vs...))
}

// MediableTypeNotIn applies the NotIn predicate on the "mediable_type" field.
func MediableTypeNotIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNotIn(FieldMediableType, vs...))
}

// MediableTypeGT applies the GT predicate on the "mediable_type" field.
func MediableTypeGT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGT(FieldMediableType, v))
}

// MediableTypeGTE applies the GTE predicate on the "mediable_type" field.
func MediableTypeGTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGTE(FieldMediableType, v))
}

// MediableTypeLT applies the LT predicate on the "mediable_type" field.
func MediableTypeLT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLT(FieldMediableType, v))
}

// MediableTypeLTE applies the LTE predicate on the "mediable_type" field.
func MediableTypeLTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLTE(FieldMediableType, v))
}

// MediableTypeContains applies the Contains predicate on the "mediable_type" field.
func MediableTypeContains(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContains(FieldMediableType, v))
}

// MediableTypeHasPrefix applies the HasPrefix predicate on the "mediable_type" field.
func MediableTypeHasPrefix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasPrefix(FieldMediableType, v))
}

// MediableTypeHasSuffix applies the HasSuffix predicate on the "mediable_type" field.
func MediableTypeHasSuffix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasSuffix(FieldMediableType, v))
}

// MediableTypeIsNil applies the IsNil predicate on the "mediable_type" field.
func MediableTypeIsNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldIsNull(FieldMediableType))
}

// MediableTypeNotNil applies the NotNil predicate on the "mediable_type" field.
func MediableTypeNotNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldNotNull(FieldMediableType))
}

// MediableTypeEqualFold applies the EqualFold predicate on the "mediable_type" field.
func MediableTypeEqualFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEqualFold(FieldMediableType, v))
}

// MediableTypeContainsFold applies the ContainsFold predicate on the "mediable_type" field.
func MediableTypeContainsFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContainsFold(FieldMediableType, v))
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldTag, v))
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNEQ(FieldTag, v))
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldIn(FieldTag, vs...))
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...string) predicate.Mediable {
	return predicate.Mediable(sql.FieldNotIn(FieldTag, vs...))
}

// TagGT applies the GT predicate on the "tag" field.
func TagGT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGT(FieldTag, v))
}

// TagGTE applies the GTE predicate on the "tag" field.
func TagGTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldGTE(FieldTag, v))
}

// TagLT applies the LT predicate on the "tag" field.
func TagLT(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLT(FieldTag, v))
}

// TagLTE applies the LTE predicate on the "tag" field.
func TagLTE(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldLTE(FieldTag, v))
}

// TagContains applies the Contains predicate on the "tag" field.
func TagContains(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContains(FieldTag, v))
}

// TagHasPrefix applies the HasPrefix predicate on the "tag" field.
func TagHasPrefix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasPrefix(FieldTag, v))
}

// TagHasSuffix applies the HasSuffix predicate on the "tag" field.
func TagHasSuffix(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldHasSuffix(FieldTag, v))
}

// TagIsNil applies the IsNil predicate on the "tag" field.
func TagIsNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldIsNull(FieldTag))
}

// TagNotNil applies the NotNil predicate on the "tag" field.
func TagNotNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldNotNull(FieldTag))
}

// TagEqualFold applies the EqualFold predicate on the "tag" field.
func TagEqualFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldEqualFold(FieldTag, v))
}

// TagContainsFold applies the ContainsFold predicate on the "tag" field.
func TagContainsFold(v string) predicate.Mediable {
	return predicate.Mediable(sql.FieldContainsFold(FieldTag, v))
}

// SortOrderEQ applies the EQ predicate on the "sort_order" field.
func SortOrderEQ(v int) predicate.Mediable {
	return predicate.Mediable(sql.FieldEQ(FieldSortOrder, v))
}

// SortOrderNEQ applies the NEQ predicate on the "sort_order" field.
func SortOrderNEQ(v int) predicate.Mediable {
	return predicate.Mediable(sql.FieldNEQ(FieldSortOrder, v))
}

// SortOrderIn applies the In predicate on the "sort_order" field.
func SortOrderIn(vs ...int) predicate.Mediable {
	return predicate.Mediable(sql.FieldIn(FieldSortOrder, vs...))
}

// SortOrderNotIn applies the NotIn predicate on the "sort_order" field.
func SortOrderNotIn(vs ...int) predicate.Mediable {
	return predicate.Mediable(sql.FieldNotIn(FieldSortOrder, vs...))
}

// SortOrderGT applies the GT predicate on the "sort_order" field.
func SortOrderGT(v int) predicate.Mediable {
	return predicate.Mediable(sql.FieldGT(FieldSortOrder, v))
}

// SortOrderGTE applies the GTE predicate on the "sort_order" field.
func SortOrderGTE(v int) predicate.Mediable {
	return predicate.Mediable(sql.FieldGTE(FieldSortOrder, v))
}

// SortOrderLT applies the LT predicate on the "sort_order" field.
func SortOrderLT(v int) predicate.Mediable {
	return predicate.Mediable(sql.FieldLT(FieldSortOrder, v))
}

// SortOrderLTE applies the LTE predicate on the "sort_order" field.
func SortOrderLTE(v int) predicate.Mediable {
	return predicate.Mediable(sql.FieldLTE(FieldSortOrder, v))
}

// SortOrderIsNil applies the IsNil predicate on the "sort_order" field.
func SortOrderIsNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldIsNull(FieldSortOrder))
}

// SortOrderNotNil applies the NotNil predicate on the "sort_order" field.
func SortOrderNotNil() predicate.Mediable {
	return predicate.Mediable(sql.FieldNotNull(FieldSortOrder))
}

// HasMedia applies the HasEdge predicate on the "media" edge.
func HasMedia() predicate.Mediable {
	return predicate.Mediable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MediaTable, MediaColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMediaWith applies the HasEdge predicate on the "media" edge with a given conditions (other predicates).
func HasMediaWith(preds ...predicate.Media) predicate.Mediable {
	return predicate.Mediable(func(s *sql.Selector) {
		step := newMediaStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Mediable) predicate.Mediable {
	return predicate.Mediable(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Mediable) predicate.Mediable {
	return predicate.Mediable(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Mediable) predicate.Mediable {
	return predicate.Mediable(sql.NotPredicates(p))
}
