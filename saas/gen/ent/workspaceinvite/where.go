// Code generated by ent, DO NOT EDIT.

package workspaceinvite

import (
	"saas/gen/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldUpdatedAt, v))
}

// WorkspaceID applies equality check predicate on the "workspace_id" field. It's identical to WorkspaceIDEQ.
func WorkspaceID(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldWorkspaceID, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldEmail, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldRole, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotNull(FieldUpdatedAt))
}

// WorkspaceIDEQ applies the EQ predicate on the "workspace_id" field.
func WorkspaceIDEQ(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldWorkspaceID, v))
}

// WorkspaceIDNEQ applies the NEQ predicate on the "workspace_id" field.
func WorkspaceIDNEQ(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNEQ(FieldWorkspaceID, v))
}

// WorkspaceIDIn applies the In predicate on the "workspace_id" field.
func WorkspaceIDIn(vs ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDNotIn applies the NotIn predicate on the "workspace_id" field.
func WorkspaceIDNotIn(vs ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotIn(FieldWorkspaceID, vs...))
}

// WorkspaceIDGT applies the GT predicate on the "workspace_id" field.
func WorkspaceIDGT(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGT(FieldWorkspaceID, v))
}

// WorkspaceIDGTE applies the GTE predicate on the "workspace_id" field.
func WorkspaceIDGTE(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGTE(FieldWorkspaceID, v))
}

// WorkspaceIDLT applies the LT predicate on the "workspace_id" field.
func WorkspaceIDLT(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLT(FieldWorkspaceID, v))
}

// WorkspaceIDLTE applies the LTE predicate on the "workspace_id" field.
func WorkspaceIDLTE(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLTE(FieldWorkspaceID, v))
}

// WorkspaceIDContains applies the Contains predicate on the "workspace_id" field.
func WorkspaceIDContains(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldContains(FieldWorkspaceID, v))
}

// WorkspaceIDHasPrefix applies the HasPrefix predicate on the "workspace_id" field.
func WorkspaceIDHasPrefix(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldHasPrefix(FieldWorkspaceID, v))
}

// WorkspaceIDHasSuffix applies the HasSuffix predicate on the "workspace_id" field.
func WorkspaceIDHasSuffix(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldHasSuffix(FieldWorkspaceID, v))
}

// WorkspaceIDIsNil applies the IsNil predicate on the "workspace_id" field.
func WorkspaceIDIsNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIsNull(FieldWorkspaceID))
}

// WorkspaceIDNotNil applies the NotNil predicate on the "workspace_id" field.
func WorkspaceIDNotNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotNull(FieldWorkspaceID))
}

// WorkspaceIDEqualFold applies the EqualFold predicate on the "workspace_id" field.
func WorkspaceIDEqualFold(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEqualFold(FieldWorkspaceID, v))
}

// WorkspaceIDContainsFold applies the ContainsFold predicate on the "workspace_id" field.
func WorkspaceIDContainsFold(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldContainsFold(FieldWorkspaceID, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldContainsFold(FieldEmail, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldHasSuffix(FieldRole, v))
}

// RoleIsNil applies the IsNil predicate on the "role" field.
func RoleIsNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldIsNull(FieldRole))
}

// RoleNotNil applies the NotNil predicate on the "role" field.
func RoleNotNil() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldNotNull(FieldRole))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.FieldContainsFold(FieldRole, v))
}

// HasWorkspace applies the HasEdge predicate on the "workspace" edge.
func HasWorkspace() predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkspaceTable, WorkspaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspaceWith applies the HasEdge predicate on the "workspace" edge with a given conditions (other predicates).
func HasWorkspaceWith(preds ...predicate.Workspace) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(func(s *sql.Selector) {
		step := newWorkspaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkspaceInvite) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkspaceInvite) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkspaceInvite) predicate.WorkspaceInvite {
	return predicate.WorkspaceInvite(sql.NotPredicates(p))
}