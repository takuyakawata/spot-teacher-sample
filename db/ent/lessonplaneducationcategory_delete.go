// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplaneducationcategory"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// LessonPlanEducationCategoryDelete is the builder for deleting a LessonPlanEducationCategory entity.
type LessonPlanEducationCategoryDelete struct {
	config
	hooks    []Hook
	mutation *LessonPlanEducationCategoryMutation
}

// Where appends a list predicates to the LessonPlanEducationCategoryDelete builder.
func (lpecd *LessonPlanEducationCategoryDelete) Where(ps ...predicate.LessonPlanEducationCategory) *LessonPlanEducationCategoryDelete {
	lpecd.mutation.Where(ps...)
	return lpecd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lpecd *LessonPlanEducationCategoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lpecd.sqlExec, lpecd.mutation, lpecd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lpecd *LessonPlanEducationCategoryDelete) ExecX(ctx context.Context) int {
	n, err := lpecd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lpecd *LessonPlanEducationCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(lessonplaneducationcategory.Table, sqlgraph.NewFieldSpec(lessonplaneducationcategory.FieldID, field.TypeInt64))
	if ps := lpecd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lpecd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lpecd.mutation.done = true
	return affected, err
}

// LessonPlanEducationCategoryDeleteOne is the builder for deleting a single LessonPlanEducationCategory entity.
type LessonPlanEducationCategoryDeleteOne struct {
	lpecd *LessonPlanEducationCategoryDelete
}

// Where appends a list predicates to the LessonPlanEducationCategoryDelete builder.
func (lpecdo *LessonPlanEducationCategoryDeleteOne) Where(ps ...predicate.LessonPlanEducationCategory) *LessonPlanEducationCategoryDeleteOne {
	lpecdo.lpecd.mutation.Where(ps...)
	return lpecdo
}

// Exec executes the deletion query.
func (lpecdo *LessonPlanEducationCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := lpecdo.lpecd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lessonplaneducationcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lpecdo *LessonPlanEducationCategoryDeleteOne) ExecX(ctx context.Context) {
	if err := lpecdo.Exec(ctx); err != nil {
		panic(err)
	}
}
