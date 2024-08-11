// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/gen/ent/oauthconnection"
	"saas/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OauthConnectionDelete is the builder for deleting a OauthConnection entity.
type OauthConnectionDelete struct {
	config
	hooks    []Hook
	mutation *OauthConnectionMutation
}

// Where appends a list predicates to the OauthConnectionDelete builder.
func (ocd *OauthConnectionDelete) Where(ps ...predicate.OauthConnection) *OauthConnectionDelete {
	ocd.mutation.Where(ps...)
	return ocd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ocd *OauthConnectionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ocd.sqlExec, ocd.mutation, ocd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ocd *OauthConnectionDelete) ExecX(ctx context.Context) int {
	n, err := ocd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ocd *OauthConnectionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(oauthconnection.Table, sqlgraph.NewFieldSpec(oauthconnection.FieldID, field.TypeString))
	if ps := ocd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ocd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ocd.mutation.done = true
	return affected, err
}

// OauthConnectionDeleteOne is the builder for deleting a single OauthConnection entity.
type OauthConnectionDeleteOne struct {
	ocd *OauthConnectionDelete
}

// Where appends a list predicates to the OauthConnectionDelete builder.
func (ocdo *OauthConnectionDeleteOne) Where(ps ...predicate.OauthConnection) *OauthConnectionDeleteOne {
	ocdo.ocd.mutation.Where(ps...)
	return ocdo
}

// Exec executes the deletion query.
func (ocdo *OauthConnectionDeleteOne) Exec(ctx context.Context) error {
	n, err := ocdo.ocd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oauthconnection.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ocdo *OauthConnectionDeleteOne) ExecX(ctx context.Context) {
	if err := ocdo.Exec(ctx); err != nil {
		panic(err)
	}
}