// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/grade"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplan"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplangrade"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// LessonPlanGradeUpdate is the builder for updating LessonPlanGrade entities.
type LessonPlanGradeUpdate struct {
	config
	hooks    []Hook
	mutation *LessonPlanGradeMutation
}

// Where appends a list predicates to the LessonPlanGradeUpdate builder.
func (lpgu *LessonPlanGradeUpdate) Where(ps ...predicate.LessonPlanGrade) *LessonPlanGradeUpdate {
	lpgu.mutation.Where(ps...)
	return lpgu
}

// SetUpdatedAt sets the "updated_at" field.
func (lpgu *LessonPlanGradeUpdate) SetUpdatedAt(t time.Time) *LessonPlanGradeUpdate {
	lpgu.mutation.SetUpdatedAt(t)
	return lpgu
}

// SetLessonPlanID sets the "lesson_plan_id" field.
func (lpgu *LessonPlanGradeUpdate) SetLessonPlanID(i int64) *LessonPlanGradeUpdate {
	lpgu.mutation.SetLessonPlanID(i)
	return lpgu
}

// SetNillableLessonPlanID sets the "lesson_plan_id" field if the given value is not nil.
func (lpgu *LessonPlanGradeUpdate) SetNillableLessonPlanID(i *int64) *LessonPlanGradeUpdate {
	if i != nil {
		lpgu.SetLessonPlanID(*i)
	}
	return lpgu
}

// SetGradeID sets the "grade_id" field.
func (lpgu *LessonPlanGradeUpdate) SetGradeID(i int64) *LessonPlanGradeUpdate {
	lpgu.mutation.SetGradeID(i)
	return lpgu
}

// SetNillableGradeID sets the "grade_id" field if the given value is not nil.
func (lpgu *LessonPlanGradeUpdate) SetNillableGradeID(i *int64) *LessonPlanGradeUpdate {
	if i != nil {
		lpgu.SetGradeID(*i)
	}
	return lpgu
}

// SetLessonPlan sets the "lesson_plan" edge to the LessonPlan entity.
func (lpgu *LessonPlanGradeUpdate) SetLessonPlan(l *LessonPlan) *LessonPlanGradeUpdate {
	return lpgu.SetLessonPlanID(l.ID)
}

// SetGrade sets the "grade" edge to the Grade entity.
func (lpgu *LessonPlanGradeUpdate) SetGrade(g *Grade) *LessonPlanGradeUpdate {
	return lpgu.SetGradeID(g.ID)
}

// Mutation returns the LessonPlanGradeMutation object of the builder.
func (lpgu *LessonPlanGradeUpdate) Mutation() *LessonPlanGradeMutation {
	return lpgu.mutation
}

// ClearLessonPlan clears the "lesson_plan" edge to the LessonPlan entity.
func (lpgu *LessonPlanGradeUpdate) ClearLessonPlan() *LessonPlanGradeUpdate {
	lpgu.mutation.ClearLessonPlan()
	return lpgu
}

// ClearGrade clears the "grade" edge to the Grade entity.
func (lpgu *LessonPlanGradeUpdate) ClearGrade() *LessonPlanGradeUpdate {
	lpgu.mutation.ClearGrade()
	return lpgu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lpgu *LessonPlanGradeUpdate) Save(ctx context.Context) (int, error) {
	lpgu.defaults()
	return withHooks(ctx, lpgu.sqlSave, lpgu.mutation, lpgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lpgu *LessonPlanGradeUpdate) SaveX(ctx context.Context) int {
	affected, err := lpgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lpgu *LessonPlanGradeUpdate) Exec(ctx context.Context) error {
	_, err := lpgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpgu *LessonPlanGradeUpdate) ExecX(ctx context.Context) {
	if err := lpgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lpgu *LessonPlanGradeUpdate) defaults() {
	if _, ok := lpgu.mutation.UpdatedAt(); !ok {
		v := lessonplangrade.UpdateDefaultUpdatedAt()
		lpgu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpgu *LessonPlanGradeUpdate) check() error {
	if lpgu.mutation.LessonPlanCleared() && len(lpgu.mutation.LessonPlanIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "LessonPlanGrade.lesson_plan"`)
	}
	if lpgu.mutation.GradeCleared() && len(lpgu.mutation.GradeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "LessonPlanGrade.grade"`)
	}
	return nil
}

func (lpgu *LessonPlanGradeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lpgu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lessonplangrade.Table, lessonplangrade.Columns, sqlgraph.NewFieldSpec(lessonplangrade.FieldID, field.TypeInt64))
	if ps := lpgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lpgu.mutation.UpdatedAt(); ok {
		_spec.SetField(lessonplangrade.FieldUpdatedAt, field.TypeTime, value)
	}
	if lpgu.mutation.LessonPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.LessonPlanTable,
			Columns: []string{lessonplangrade.LessonPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpgu.mutation.LessonPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.LessonPlanTable,
			Columns: []string{lessonplangrade.LessonPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpgu.mutation.GradeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.GradeTable,
			Columns: []string{lessonplangrade.GradeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grade.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpgu.mutation.GradeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.GradeTable,
			Columns: []string{lessonplangrade.GradeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grade.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lpgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lessonplangrade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lpgu.mutation.done = true
	return n, nil
}

// LessonPlanGradeUpdateOne is the builder for updating a single LessonPlanGrade entity.
type LessonPlanGradeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LessonPlanGradeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (lpguo *LessonPlanGradeUpdateOne) SetUpdatedAt(t time.Time) *LessonPlanGradeUpdateOne {
	lpguo.mutation.SetUpdatedAt(t)
	return lpguo
}

// SetLessonPlanID sets the "lesson_plan_id" field.
func (lpguo *LessonPlanGradeUpdateOne) SetLessonPlanID(i int64) *LessonPlanGradeUpdateOne {
	lpguo.mutation.SetLessonPlanID(i)
	return lpguo
}

// SetNillableLessonPlanID sets the "lesson_plan_id" field if the given value is not nil.
func (lpguo *LessonPlanGradeUpdateOne) SetNillableLessonPlanID(i *int64) *LessonPlanGradeUpdateOne {
	if i != nil {
		lpguo.SetLessonPlanID(*i)
	}
	return lpguo
}

// SetGradeID sets the "grade_id" field.
func (lpguo *LessonPlanGradeUpdateOne) SetGradeID(i int64) *LessonPlanGradeUpdateOne {
	lpguo.mutation.SetGradeID(i)
	return lpguo
}

// SetNillableGradeID sets the "grade_id" field if the given value is not nil.
func (lpguo *LessonPlanGradeUpdateOne) SetNillableGradeID(i *int64) *LessonPlanGradeUpdateOne {
	if i != nil {
		lpguo.SetGradeID(*i)
	}
	return lpguo
}

// SetLessonPlan sets the "lesson_plan" edge to the LessonPlan entity.
func (lpguo *LessonPlanGradeUpdateOne) SetLessonPlan(l *LessonPlan) *LessonPlanGradeUpdateOne {
	return lpguo.SetLessonPlanID(l.ID)
}

// SetGrade sets the "grade" edge to the Grade entity.
func (lpguo *LessonPlanGradeUpdateOne) SetGrade(g *Grade) *LessonPlanGradeUpdateOne {
	return lpguo.SetGradeID(g.ID)
}

// Mutation returns the LessonPlanGradeMutation object of the builder.
func (lpguo *LessonPlanGradeUpdateOne) Mutation() *LessonPlanGradeMutation {
	return lpguo.mutation
}

// ClearLessonPlan clears the "lesson_plan" edge to the LessonPlan entity.
func (lpguo *LessonPlanGradeUpdateOne) ClearLessonPlan() *LessonPlanGradeUpdateOne {
	lpguo.mutation.ClearLessonPlan()
	return lpguo
}

// ClearGrade clears the "grade" edge to the Grade entity.
func (lpguo *LessonPlanGradeUpdateOne) ClearGrade() *LessonPlanGradeUpdateOne {
	lpguo.mutation.ClearGrade()
	return lpguo
}

// Where appends a list predicates to the LessonPlanGradeUpdate builder.
func (lpguo *LessonPlanGradeUpdateOne) Where(ps ...predicate.LessonPlanGrade) *LessonPlanGradeUpdateOne {
	lpguo.mutation.Where(ps...)
	return lpguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lpguo *LessonPlanGradeUpdateOne) Select(field string, fields ...string) *LessonPlanGradeUpdateOne {
	lpguo.fields = append([]string{field}, fields...)
	return lpguo
}

// Save executes the query and returns the updated LessonPlanGrade entity.
func (lpguo *LessonPlanGradeUpdateOne) Save(ctx context.Context) (*LessonPlanGrade, error) {
	lpguo.defaults()
	return withHooks(ctx, lpguo.sqlSave, lpguo.mutation, lpguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lpguo *LessonPlanGradeUpdateOne) SaveX(ctx context.Context) *LessonPlanGrade {
	node, err := lpguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lpguo *LessonPlanGradeUpdateOne) Exec(ctx context.Context) error {
	_, err := lpguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lpguo *LessonPlanGradeUpdateOne) ExecX(ctx context.Context) {
	if err := lpguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lpguo *LessonPlanGradeUpdateOne) defaults() {
	if _, ok := lpguo.mutation.UpdatedAt(); !ok {
		v := lessonplangrade.UpdateDefaultUpdatedAt()
		lpguo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lpguo *LessonPlanGradeUpdateOne) check() error {
	if lpguo.mutation.LessonPlanCleared() && len(lpguo.mutation.LessonPlanIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "LessonPlanGrade.lesson_plan"`)
	}
	if lpguo.mutation.GradeCleared() && len(lpguo.mutation.GradeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "LessonPlanGrade.grade"`)
	}
	return nil
}

func (lpguo *LessonPlanGradeUpdateOne) sqlSave(ctx context.Context) (_node *LessonPlanGrade, err error) {
	if err := lpguo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lessonplangrade.Table, lessonplangrade.Columns, sqlgraph.NewFieldSpec(lessonplangrade.FieldID, field.TypeInt64))
	id, ok := lpguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LessonPlanGrade.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lpguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lessonplangrade.FieldID)
		for _, f := range fields {
			if !lessonplangrade.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lessonplangrade.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lpguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lpguo.mutation.UpdatedAt(); ok {
		_spec.SetField(lessonplangrade.FieldUpdatedAt, field.TypeTime, value)
	}
	if lpguo.mutation.LessonPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.LessonPlanTable,
			Columns: []string{lessonplangrade.LessonPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpguo.mutation.LessonPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.LessonPlanTable,
			Columns: []string{lessonplangrade.LessonPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lpguo.mutation.GradeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.GradeTable,
			Columns: []string{lessonplangrade.GradeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grade.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lpguo.mutation.GradeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lessonplangrade.GradeTable,
			Columns: []string{lessonplangrade.GradeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grade.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LessonPlanGrade{config: lpguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lpguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lessonplangrade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lpguo.mutation.done = true
	return _node, nil
}
