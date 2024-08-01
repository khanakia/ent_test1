// Code generated by ent, DO NOT EDIT.

package poststatus

import (
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldName, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldSlug, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldStatus, v))
}

// PostTypeID applies equality check predicate on the "post_type_id" field. It's identical to PostTypeIDEQ.
func PostTypeID(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldPostTypeID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotNull(FieldUpdatedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldContainsFold(FieldName, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugIsNil applies the IsNil predicate on the "slug" field.
func SlugIsNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIsNull(FieldSlug))
}

// SlugNotNil applies the NotNil predicate on the "slug" field.
func SlugNotNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotNull(FieldSlug))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldContainsFold(FieldSlug, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNEQ(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotNull(FieldStatus))
}

// PostTypeIDEQ applies the EQ predicate on the "post_type_id" field.
func PostTypeIDEQ(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEQ(FieldPostTypeID, v))
}

// PostTypeIDNEQ applies the NEQ predicate on the "post_type_id" field.
func PostTypeIDNEQ(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNEQ(FieldPostTypeID, v))
}

// PostTypeIDIn applies the In predicate on the "post_type_id" field.
func PostTypeIDIn(vs ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIn(FieldPostTypeID, vs...))
}

// PostTypeIDNotIn applies the NotIn predicate on the "post_type_id" field.
func PostTypeIDNotIn(vs ...string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotIn(FieldPostTypeID, vs...))
}

// PostTypeIDGT applies the GT predicate on the "post_type_id" field.
func PostTypeIDGT(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGT(FieldPostTypeID, v))
}

// PostTypeIDGTE applies the GTE predicate on the "post_type_id" field.
func PostTypeIDGTE(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldGTE(FieldPostTypeID, v))
}

// PostTypeIDLT applies the LT predicate on the "post_type_id" field.
func PostTypeIDLT(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLT(FieldPostTypeID, v))
}

// PostTypeIDLTE applies the LTE predicate on the "post_type_id" field.
func PostTypeIDLTE(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldLTE(FieldPostTypeID, v))
}

// PostTypeIDContains applies the Contains predicate on the "post_type_id" field.
func PostTypeIDContains(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldContains(FieldPostTypeID, v))
}

// PostTypeIDHasPrefix applies the HasPrefix predicate on the "post_type_id" field.
func PostTypeIDHasPrefix(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldHasPrefix(FieldPostTypeID, v))
}

// PostTypeIDHasSuffix applies the HasSuffix predicate on the "post_type_id" field.
func PostTypeIDHasSuffix(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldHasSuffix(FieldPostTypeID, v))
}

// PostTypeIDIsNil applies the IsNil predicate on the "post_type_id" field.
func PostTypeIDIsNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldIsNull(FieldPostTypeID))
}

// PostTypeIDNotNil applies the NotNil predicate on the "post_type_id" field.
func PostTypeIDNotNil() predicate.PostStatus {
	return predicate.PostStatus(sql.FieldNotNull(FieldPostTypeID))
}

// PostTypeIDEqualFold applies the EqualFold predicate on the "post_type_id" field.
func PostTypeIDEqualFold(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldEqualFold(FieldPostTypeID, v))
}

// PostTypeIDContainsFold applies the ContainsFold predicate on the "post_type_id" field.
func PostTypeIDContainsFold(v string) predicate.PostStatus {
	return predicate.PostStatus(sql.FieldContainsFold(FieldPostTypeID, v))
}

// HasPostType applies the HasEdge predicate on the "post_type" edge.
func HasPostType() predicate.PostStatus {
	return predicate.PostStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTypeTable, PostTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostTypeWith applies the HasEdge predicate on the "post_type" edge with a given conditions (other predicates).
func HasPostTypeWith(preds ...predicate.PostType) predicate.PostStatus {
	return predicate.PostStatus(func(s *sql.Selector) {
		step := newPostTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.PostStatus {
	return predicate.PostStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.PostStatus {
	return predicate.PostStatus(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PostStatus) predicate.PostStatus {
	return predicate.PostStatus(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PostStatus) predicate.PostStatus {
	return predicate.PostStatus(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PostStatus) predicate.PostStatus {
	return predicate.PostStatus(sql.NotPredicates(p))
}
