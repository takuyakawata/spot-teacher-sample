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
	"github.com/takuyakawta/spot-teacher-sample/db/ent/educationcategory"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplan"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplaneducationcategory"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// EducationCategoryUpdate is the builder for updating EducationCategory entities.
type EducationCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *EducationCategoryMutation
}

// Where appends a list predicates to the EducationCategoryUpdate builder.
func (ecu *EducationCategoryUpdate) Where(ps ...predicate.EducationCategory) *EducationCategoryUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetUpdatedAt sets the "updated_at" field.
func (ecu *EducationCategoryUpdate) SetUpdatedAt(t time.Time) *EducationCategoryUpdate {
	ecu.mutation.SetUpdatedAt(t)
	return ecu
}

// SetName sets the "name" field.
func (ecu *EducationCategoryUpdate) SetName(s string) *EducationCategoryUpdate {
	ecu.mutation.SetName(s)
	return ecu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ecu *EducationCategoryUpdate) SetNillableName(s *string) *EducationCategoryUpdate {
	if s != nil {
		ecu.SetName(*s)
	}
	return ecu
}

// SetCode sets the "code" field.
func (ecu *EducationCategoryUpdate) SetCode(s string) *EducationCategoryUpdate {
	ecu.mutation.SetCode(s)
	return ecu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ecu *EducationCategoryUpdate) SetNillableCode(s *string) *EducationCategoryUpdate {
	if s != nil {
		ecu.SetCode(*s)
	}
	return ecu
}

// AddLessonPlanIDs adds the "lesson_plans" edge to the LessonPlan entity by IDs.
func (ecu *EducationCategoryUpdate) AddLessonPlanIDs(ids ...int64) *EducationCategoryUpdate {
	ecu.mutation.AddLessonPlanIDs(ids...)
	return ecu
}

// AddLessonPlans adds the "lesson_plans" edges to the LessonPlan entity.
func (ecu *EducationCategoryUpdate) AddLessonPlans(l ...*LessonPlan) *EducationCategoryUpdate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecu.AddLessonPlanIDs(ids...)
}

// AddLessonPlanEducationCategoryIDs adds the "lesson_plan_education_categories" edge to the LessonPlanEducationCategory entity by IDs.
func (ecu *EducationCategoryUpdate) AddLessonPlanEducationCategoryIDs(ids ...int64) *EducationCategoryUpdate {
	ecu.mutation.AddLessonPlanEducationCategoryIDs(ids...)
	return ecu
}

// AddLessonPlanEducationCategories adds the "lesson_plan_education_categories" edges to the LessonPlanEducationCategory entity.
func (ecu *EducationCategoryUpdate) AddLessonPlanEducationCategories(l ...*LessonPlanEducationCategory) *EducationCategoryUpdate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecu.AddLessonPlanEducationCategoryIDs(ids...)
}

// Mutation returns the EducationCategoryMutation object of the builder.
func (ecu *EducationCategoryUpdate) Mutation() *EducationCategoryMutation {
	return ecu.mutation
}

// ClearLessonPlans clears all "lesson_plans" edges to the LessonPlan entity.
func (ecu *EducationCategoryUpdate) ClearLessonPlans() *EducationCategoryUpdate {
	ecu.mutation.ClearLessonPlans()
	return ecu
}

// RemoveLessonPlanIDs removes the "lesson_plans" edge to LessonPlan entities by IDs.
func (ecu *EducationCategoryUpdate) RemoveLessonPlanIDs(ids ...int64) *EducationCategoryUpdate {
	ecu.mutation.RemoveLessonPlanIDs(ids...)
	return ecu
}

// RemoveLessonPlans removes "lesson_plans" edges to LessonPlan entities.
func (ecu *EducationCategoryUpdate) RemoveLessonPlans(l ...*LessonPlan) *EducationCategoryUpdate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecu.RemoveLessonPlanIDs(ids...)
}

// ClearLessonPlanEducationCategories clears all "lesson_plan_education_categories" edges to the LessonPlanEducationCategory entity.
func (ecu *EducationCategoryUpdate) ClearLessonPlanEducationCategories() *EducationCategoryUpdate {
	ecu.mutation.ClearLessonPlanEducationCategories()
	return ecu
}

// RemoveLessonPlanEducationCategoryIDs removes the "lesson_plan_education_categories" edge to LessonPlanEducationCategory entities by IDs.
func (ecu *EducationCategoryUpdate) RemoveLessonPlanEducationCategoryIDs(ids ...int64) *EducationCategoryUpdate {
	ecu.mutation.RemoveLessonPlanEducationCategoryIDs(ids...)
	return ecu
}

// RemoveLessonPlanEducationCategories removes "lesson_plan_education_categories" edges to LessonPlanEducationCategory entities.
func (ecu *EducationCategoryUpdate) RemoveLessonPlanEducationCategories(l ...*LessonPlanEducationCategory) *EducationCategoryUpdate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecu.RemoveLessonPlanEducationCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EducationCategoryUpdate) Save(ctx context.Context) (int, error) {
	ecu.defaults()
	return withHooks(ctx, ecu.sqlSave, ecu.mutation, ecu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EducationCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EducationCategoryUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EducationCategoryUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecu *EducationCategoryUpdate) defaults() {
	if _, ok := ecu.mutation.UpdatedAt(); !ok {
		v := educationcategory.UpdateDefaultUpdatedAt()
		ecu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecu *EducationCategoryUpdate) check() error {
	if v, ok := ecu.mutation.Name(); ok {
		if err := educationcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EducationCategory.name": %w`, err)}
		}
	}
	if v, ok := ecu.mutation.Code(); ok {
		if err := educationcategory.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "EducationCategory.code": %w`, err)}
		}
	}
	return nil
}

func (ecu *EducationCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ecu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(educationcategory.Table, educationcategory.Columns, sqlgraph.NewFieldSpec(educationcategory.FieldID, field.TypeInt64))
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.UpdatedAt(); ok {
		_spec.SetField(educationcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ecu.mutation.Name(); ok {
		_spec.SetField(educationcategory.FieldName, field.TypeString, value)
	}
	if value, ok := ecu.mutation.Code(); ok {
		_spec.SetField(educationcategory.FieldCode, field.TypeString, value)
	}
	if ecu.mutation.LessonPlansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   educationcategory.LessonPlansTable,
			Columns: educationcategory.LessonPlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		createE := &LessonPlanEducationCategoryCreate{config: ecu.config, mutation: newLessonPlanEducationCategoryMutation(ecu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.RemovedLessonPlansIDs(); len(nodes) > 0 && !ecu.mutation.LessonPlansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   educationcategory.LessonPlansTable,
			Columns: educationcategory.LessonPlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &LessonPlanEducationCategoryCreate{config: ecu.config, mutation: newLessonPlanEducationCategoryMutation(ecu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.LessonPlansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   educationcategory.LessonPlansTable,
			Columns: educationcategory.LessonPlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &LessonPlanEducationCategoryCreate{config: ecu.config, mutation: newLessonPlanEducationCategoryMutation(ecu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ecu.mutation.LessonPlanEducationCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   educationcategory.LessonPlanEducationCategoriesTable,
			Columns: []string{educationcategory.LessonPlanEducationCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplaneducationcategory.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.RemovedLessonPlanEducationCategoriesIDs(); len(nodes) > 0 && !ecu.mutation.LessonPlanEducationCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   educationcategory.LessonPlanEducationCategoriesTable,
			Columns: []string{educationcategory.LessonPlanEducationCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplaneducationcategory.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.LessonPlanEducationCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   educationcategory.LessonPlanEducationCategoriesTable,
			Columns: []string{educationcategory.LessonPlanEducationCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplaneducationcategory.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{educationcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecu.mutation.done = true
	return n, nil
}

// EducationCategoryUpdateOne is the builder for updating a single EducationCategory entity.
type EducationCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EducationCategoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ecuo *EducationCategoryUpdateOne) SetUpdatedAt(t time.Time) *EducationCategoryUpdateOne {
	ecuo.mutation.SetUpdatedAt(t)
	return ecuo
}

// SetName sets the "name" field.
func (ecuo *EducationCategoryUpdateOne) SetName(s string) *EducationCategoryUpdateOne {
	ecuo.mutation.SetName(s)
	return ecuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ecuo *EducationCategoryUpdateOne) SetNillableName(s *string) *EducationCategoryUpdateOne {
	if s != nil {
		ecuo.SetName(*s)
	}
	return ecuo
}

// SetCode sets the "code" field.
func (ecuo *EducationCategoryUpdateOne) SetCode(s string) *EducationCategoryUpdateOne {
	ecuo.mutation.SetCode(s)
	return ecuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ecuo *EducationCategoryUpdateOne) SetNillableCode(s *string) *EducationCategoryUpdateOne {
	if s != nil {
		ecuo.SetCode(*s)
	}
	return ecuo
}

// AddLessonPlanIDs adds the "lesson_plans" edge to the LessonPlan entity by IDs.
func (ecuo *EducationCategoryUpdateOne) AddLessonPlanIDs(ids ...int64) *EducationCategoryUpdateOne {
	ecuo.mutation.AddLessonPlanIDs(ids...)
	return ecuo
}

// AddLessonPlans adds the "lesson_plans" edges to the LessonPlan entity.
func (ecuo *EducationCategoryUpdateOne) AddLessonPlans(l ...*LessonPlan) *EducationCategoryUpdateOne {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecuo.AddLessonPlanIDs(ids...)
}

// AddLessonPlanEducationCategoryIDs adds the "lesson_plan_education_categories" edge to the LessonPlanEducationCategory entity by IDs.
func (ecuo *EducationCategoryUpdateOne) AddLessonPlanEducationCategoryIDs(ids ...int64) *EducationCategoryUpdateOne {
	ecuo.mutation.AddLessonPlanEducationCategoryIDs(ids...)
	return ecuo
}

// AddLessonPlanEducationCategories adds the "lesson_plan_education_categories" edges to the LessonPlanEducationCategory entity.
func (ecuo *EducationCategoryUpdateOne) AddLessonPlanEducationCategories(l ...*LessonPlanEducationCategory) *EducationCategoryUpdateOne {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecuo.AddLessonPlanEducationCategoryIDs(ids...)
}

// Mutation returns the EducationCategoryMutation object of the builder.
func (ecuo *EducationCategoryUpdateOne) Mutation() *EducationCategoryMutation {
	return ecuo.mutation
}

// ClearLessonPlans clears all "lesson_plans" edges to the LessonPlan entity.
func (ecuo *EducationCategoryUpdateOne) ClearLessonPlans() *EducationCategoryUpdateOne {
	ecuo.mutation.ClearLessonPlans()
	return ecuo
}

// RemoveLessonPlanIDs removes the "lesson_plans" edge to LessonPlan entities by IDs.
func (ecuo *EducationCategoryUpdateOne) RemoveLessonPlanIDs(ids ...int64) *EducationCategoryUpdateOne {
	ecuo.mutation.RemoveLessonPlanIDs(ids...)
	return ecuo
}

// RemoveLessonPlans removes "lesson_plans" edges to LessonPlan entities.
func (ecuo *EducationCategoryUpdateOne) RemoveLessonPlans(l ...*LessonPlan) *EducationCategoryUpdateOne {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecuo.RemoveLessonPlanIDs(ids...)
}

// ClearLessonPlanEducationCategories clears all "lesson_plan_education_categories" edges to the LessonPlanEducationCategory entity.
func (ecuo *EducationCategoryUpdateOne) ClearLessonPlanEducationCategories() *EducationCategoryUpdateOne {
	ecuo.mutation.ClearLessonPlanEducationCategories()
	return ecuo
}

// RemoveLessonPlanEducationCategoryIDs removes the "lesson_plan_education_categories" edge to LessonPlanEducationCategory entities by IDs.
func (ecuo *EducationCategoryUpdateOne) RemoveLessonPlanEducationCategoryIDs(ids ...int64) *EducationCategoryUpdateOne {
	ecuo.mutation.RemoveLessonPlanEducationCategoryIDs(ids...)
	return ecuo
}

// RemoveLessonPlanEducationCategories removes "lesson_plan_education_categories" edges to LessonPlanEducationCategory entities.
func (ecuo *EducationCategoryUpdateOne) RemoveLessonPlanEducationCategories(l ...*LessonPlanEducationCategory) *EducationCategoryUpdateOne {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ecuo.RemoveLessonPlanEducationCategoryIDs(ids...)
}

// Where appends a list predicates to the EducationCategoryUpdate builder.
func (ecuo *EducationCategoryUpdateOne) Where(ps ...predicate.EducationCategory) *EducationCategoryUpdateOne {
	ecuo.mutation.Where(ps...)
	return ecuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EducationCategoryUpdateOne) Select(field string, fields ...string) *EducationCategoryUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EducationCategory entity.
func (ecuo *EducationCategoryUpdateOne) Save(ctx context.Context) (*EducationCategory, error) {
	ecuo.defaults()
	return withHooks(ctx, ecuo.sqlSave, ecuo.mutation, ecuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EducationCategoryUpdateOne) SaveX(ctx context.Context) *EducationCategory {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EducationCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EducationCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecuo *EducationCategoryUpdateOne) defaults() {
	if _, ok := ecuo.mutation.UpdatedAt(); !ok {
		v := educationcategory.UpdateDefaultUpdatedAt()
		ecuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecuo *EducationCategoryUpdateOne) check() error {
	if v, ok := ecuo.mutation.Name(); ok {
		if err := educationcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "EducationCategory.name": %w`, err)}
		}
	}
	if v, ok := ecuo.mutation.Code(); ok {
		if err := educationcategory.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "EducationCategory.code": %w`, err)}
		}
	}
	return nil
}

func (ecuo *EducationCategoryUpdateOne) sqlSave(ctx context.Context) (_node *EducationCategory, err error) {
	if err := ecuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(educationcategory.Table, educationcategory.Columns, sqlgraph.NewFieldSpec(educationcategory.FieldID, field.TypeInt64))
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EducationCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, educationcategory.FieldID)
		for _, f := range fields {
			if !educationcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != educationcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.UpdatedAt(); ok {
		_spec.SetField(educationcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ecuo.mutation.Name(); ok {
		_spec.SetField(educationcategory.FieldName, field.TypeString, value)
	}
	if value, ok := ecuo.mutation.Code(); ok {
		_spec.SetField(educationcategory.FieldCode, field.TypeString, value)
	}
	if ecuo.mutation.LessonPlansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   educationcategory.LessonPlansTable,
			Columns: educationcategory.LessonPlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		createE := &LessonPlanEducationCategoryCreate{config: ecuo.config, mutation: newLessonPlanEducationCategoryMutation(ecuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.RemovedLessonPlansIDs(); len(nodes) > 0 && !ecuo.mutation.LessonPlansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   educationcategory.LessonPlansTable,
			Columns: educationcategory.LessonPlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &LessonPlanEducationCategoryCreate{config: ecuo.config, mutation: newLessonPlanEducationCategoryMutation(ecuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.LessonPlansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   educationcategory.LessonPlansTable,
			Columns: educationcategory.LessonPlansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplan.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &LessonPlanEducationCategoryCreate{config: ecuo.config, mutation: newLessonPlanEducationCategoryMutation(ecuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ecuo.mutation.LessonPlanEducationCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   educationcategory.LessonPlanEducationCategoriesTable,
			Columns: []string{educationcategory.LessonPlanEducationCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplaneducationcategory.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.RemovedLessonPlanEducationCategoriesIDs(); len(nodes) > 0 && !ecuo.mutation.LessonPlanEducationCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   educationcategory.LessonPlanEducationCategoriesTable,
			Columns: []string{educationcategory.LessonPlanEducationCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplaneducationcategory.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.LessonPlanEducationCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   educationcategory.LessonPlanEducationCategoriesTable,
			Columns: []string{educationcategory.LessonPlanEducationCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lessonplaneducationcategory.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EducationCategory{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{educationcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecuo.mutation.done = true
	return _node, nil
}
