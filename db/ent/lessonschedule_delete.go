// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonschedule"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// LessonScheduleDelete is the builder for deleting a LessonSchedule entity.
type LessonScheduleDelete struct {
	config
	hooks    []Hook
	mutation *LessonScheduleMutation
}

// Where appends a list predicates to the LessonScheduleDelete builder.
func (lsd *LessonScheduleDelete) Where(ps ...predicate.LessonSchedule) *LessonScheduleDelete {
	lsd.mutation.Where(ps...)
	return lsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lsd *LessonScheduleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lsd.sqlExec, lsd.mutation, lsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lsd *LessonScheduleDelete) ExecX(ctx context.Context) int {
	n, err := lsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lsd *LessonScheduleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(lessonschedule.Table, sqlgraph.NewFieldSpec(lessonschedule.FieldID, field.TypeInt64))
	if ps := lsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lsd.mutation.done = true
	return affected, err
}

// LessonScheduleDeleteOne is the builder for deleting a single LessonSchedule entity.
type LessonScheduleDeleteOne struct {
	lsd *LessonScheduleDelete
}

// Where appends a list predicates to the LessonScheduleDelete builder.
func (lsdo *LessonScheduleDeleteOne) Where(ps ...predicate.LessonSchedule) *LessonScheduleDeleteOne {
	lsdo.lsd.mutation.Where(ps...)
	return lsdo
}

// Exec executes the deletion query.
func (lsdo *LessonScheduleDeleteOne) Exec(ctx context.Context) error {
	n, err := lsdo.lsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{lessonschedule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lsdo *LessonScheduleDeleteOne) ExecX(ctx context.Context) {
	if err := lsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
