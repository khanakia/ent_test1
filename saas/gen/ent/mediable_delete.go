// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/gen/ent/mediable"
	"saas/gen/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediableDelete is the builder for deleting a Mediable entity.
type MediableDelete struct {
	config
	hooks    []Hook
	mutation *MediableMutation
}

// Where appends a list predicates to the MediableDelete builder.
func (md *MediableDelete) Where(ps ...predicate.Mediable) *MediableDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MediableDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MediableDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MediableDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(mediable.Table, sqlgraph.NewFieldSpec(mediable.FieldID, field.TypeString))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MediableDeleteOne is the builder for deleting a single Mediable entity.
type MediableDeleteOne struct {
	md *MediableDelete
}

// Where appends a list predicates to the MediableDelete builder.
func (mdo *MediableDeleteOne) Where(ps ...predicate.Mediable) *MediableDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MediableDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mediable.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MediableDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
