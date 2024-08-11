// Code generated by ent, DO NOT EDIT.

package workspaceuser

import (
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// WorkspaceID applies equality check predicate on the "workspace_id" field. It's identical to WorkspaceIDEQ.
func WorkspaceID(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldWorkspaceID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldUserID, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldRole, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotNull(FieldUpdatedAt))
}

// WorkspaceIDEQ applies the EQ predicate on the "workspace_id" field.
func WorkspaceIDEQ(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldWorkspaceID, v))
}

// WorkspaceIDNEQ applies the NEQ predicate on the "workspace_id" field.
func WorkspaceIDNEQ(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNEQ(FieldWorkspaceID, v))
}

// WorkspaceIDIn applies the In predicate on the "workspace_id" field.
func WorkspaceIDIn(vs ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDNotIn applies the NotIn predicate on the "workspace_id" field.
func WorkspaceIDNotIn(vs ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDGT applies the GT predicate on the "workspace_id" field.
func WorkspaceIDGT(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGT(FieldWorkspaceID, v))
}

// WorkspaceIDGTE applies the GTE predicate on the "workspace_id" field.
func WorkspaceIDGTE(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGTE(FieldWorkspaceID, v))
}

// WorkspaceIDLT applies the LT predicate on the "workspace_id" field.
func WorkspaceIDLT(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLT(FieldWorkspaceID, v))
}

// WorkspaceIDLTE applies the LTE predicate on the "workspace_id" field.
func WorkspaceIDLTE(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLTE(FieldWorkspaceID, v))
}

// WorkspaceIDContains applies the Contains predicate on the "workspace_id" field.
func WorkspaceIDContains(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldContains(FieldWorkspaceID, v))
}

// WorkspaceIDHasPrefix applies the HasPrefix predicate on the "workspace_id" field.
func WorkspaceIDHasPrefix(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldHasPrefix(FieldWorkspaceID, v))
}

// WorkspaceIDHasSuffix applies the HasSuffix predicate on the "workspace_id" field.
func WorkspaceIDHasSuffix(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldHasSuffix(FieldWorkspaceID, v))
}

// WorkspaceIDEqualFold applies the EqualFold predicate on the "workspace_id" field.
func WorkspaceIDEqualFold(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEqualFold(FieldWorkspaceID, v))
}

// WorkspaceIDContainsFold applies the ContainsFold predicate on the "workspace_id" field.
func WorkspaceIDContainsFold(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldContainsFold(FieldWorkspaceID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldContainsFold(FieldUserID, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldHasSuffix(FieldRole, v))
}

// RoleIsNil applies the IsNil predicate on the "role" field.
func RoleIsNil() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldIsNull(FieldRole))
}

// RoleNotNil applies the NotNil predicate on the "role" field.
func RoleNotNil() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldNotNull(FieldRole))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.FieldContainsFold(FieldRole, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkspace applies the HasEdge predicate on the "workspace" edge.
func HasWorkspace() predicate.WorkspaceUser {
	return predicate.WorkspaceUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WorkspaceTable, WorkspaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspaceWith applies the HasEdge predicate on the "workspace" edge with a given conditions (other predicates).
func HasWorkspaceWith(preds ...predicate.Workspace) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(func(s *sql.Selector) {
		step := newWorkspaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkspaceUser) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkspaceUser) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkspaceUser) predicate.WorkspaceUser {
	return predicate.WorkspaceUser(sql.NotPredicates(p))
}