// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/emailverification"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// EmailVerificationDelete is the builder for deleting a EmailVerification entity.
type EmailVerificationDelete struct {
	config
	hooks    []Hook
	mutation *EmailVerificationMutation
}

// Where appends a list predicates to the EmailVerificationDelete builder.
func (evd *EmailVerificationDelete) Where(ps ...predicate.EmailVerification) *EmailVerificationDelete {
	evd.mutation.Where(ps...)
	return evd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (evd *EmailVerificationDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, evd.sqlExec, evd.mutation, evd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (evd *EmailVerificationDelete) ExecX(ctx context.Context) int {
	n, err := evd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (evd *EmailVerificationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(emailverification.Table, sqlgraph.NewFieldSpec(emailverification.FieldID, field.TypeInt64))
	if ps := evd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, evd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	evd.mutation.done = true
	return affected, err
}

// EmailVerificationDeleteOne is the builder for deleting a single EmailVerification entity.
type EmailVerificationDeleteOne struct {
	evd *EmailVerificationDelete
}

// Where appends a list predicates to the EmailVerificationDelete builder.
func (evdo *EmailVerificationDeleteOne) Where(ps ...predicate.EmailVerification) *EmailVerificationDeleteOne {
	evdo.evd.mutation.Where(ps...)
	return evdo
}

// Exec executes the deletion query.
func (evdo *EmailVerificationDeleteOne) Exec(ctx context.Context) error {
	n, err := evdo.evd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emailverification.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (evdo *EmailVerificationDeleteOne) ExecX(ctx context.Context) {
	if err := evdo.Exec(ctx); err != nil {
		panic(err)
	}
}
