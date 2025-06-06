// Code generated by ent, DO NOT EDIT.

package lessonplangrade

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldUpdatedAt, v))
}

// LessonPlanID applies equality check predicate on the "lesson_plan_id" field. It's identical to LessonPlanIDEQ.
func LessonPlanID(v int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldLessonPlanID, v))
}

// GradeID applies equality check predicate on the "grade_id" field. It's identical to GradeIDEQ.
func GradeID(v int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldGradeID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldLTE(FieldUpdatedAt, v))
}

// LessonPlanIDEQ applies the EQ predicate on the "lesson_plan_id" field.
func LessonPlanIDEQ(v int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldLessonPlanID, v))
}

// LessonPlanIDNEQ applies the NEQ predicate on the "lesson_plan_id" field.
func LessonPlanIDNEQ(v int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNEQ(FieldLessonPlanID, v))
}

// LessonPlanIDIn applies the In predicate on the "lesson_plan_id" field.
func LessonPlanIDIn(vs ...int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldIn(FieldLessonPlanID, vs...))
}

// LessonPlanIDNotIn applies the NotIn predicate on the "lesson_plan_id" field.
func LessonPlanIDNotIn(vs ...int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNotIn(FieldLessonPlanID, vs...))
}

// GradeIDEQ applies the EQ predicate on the "grade_id" field.
func GradeIDEQ(v int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldEQ(FieldGradeID, v))
}

// GradeIDNEQ applies the NEQ predicate on the "grade_id" field.
func GradeIDNEQ(v int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNEQ(FieldGradeID, v))
}

// GradeIDIn applies the In predicate on the "grade_id" field.
func GradeIDIn(vs ...int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldIn(FieldGradeID, vs...))
}

// GradeIDNotIn applies the NotIn predicate on the "grade_id" field.
func GradeIDNotIn(vs ...int64) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.FieldNotIn(FieldGradeID, vs...))
}

// HasLessonPlan applies the HasEdge predicate on the "lesson_plan" edge.
func HasLessonPlan() predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, LessonPlanTable, LessonPlanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLessonPlanWith applies the HasEdge predicate on the "lesson_plan" edge with a given conditions (other predicates).
func HasLessonPlanWith(preds ...predicate.LessonPlan) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(func(s *sql.Selector) {
		step := newLessonPlanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGrade applies the HasEdge predicate on the "grade" edge.
func HasGrade() predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GradeTable, GradeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGradeWith applies the HasEdge predicate on the "grade" edge with a given conditions (other predicates).
func HasGradeWith(preds ...predicate.Grade) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(func(s *sql.Selector) {
		step := newGradeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LessonPlanGrade) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LessonPlanGrade) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LessonPlanGrade) predicate.LessonPlanGrade {
	return predicate.LessonPlanGrade(sql.NotPredicates(p))
}
