// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplan"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplansubject"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/subject"
)

// LessonPlanSubject is the model entity for the LessonPlanSubject schema.
type LessonPlanSubject struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LessonPlanID holds the value of the "lesson_plan_id" field.
	LessonPlanID int64 `json:"lesson_plan_id,omitempty"`
	// SubjectID holds the value of the "subject_id" field.
	SubjectID int64 `json:"subject_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LessonPlanSubjectQuery when eager-loading is set.
	Edges        LessonPlanSubjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LessonPlanSubjectEdges holds the relations/edges for other nodes in the graph.
type LessonPlanSubjectEdges struct {
	// LessonPlan holds the value of the lesson_plan edge.
	LessonPlan *LessonPlan `json:"lesson_plan,omitempty"`
	// Subject holds the value of the subject edge.
	Subject *Subject `json:"subject,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LessonPlanOrErr returns the LessonPlan value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LessonPlanSubjectEdges) LessonPlanOrErr() (*LessonPlan, error) {
	if e.LessonPlan != nil {
		return e.LessonPlan, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: lessonplan.Label}
	}
	return nil, &NotLoadedError{edge: "lesson_plan"}
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LessonPlanSubjectEdges) SubjectOrErr() (*Subject, error) {
	if e.Subject != nil {
		return e.Subject, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: subject.Label}
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LessonPlanSubject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lessonplansubject.FieldID, lessonplansubject.FieldLessonPlanID, lessonplansubject.FieldSubjectID:
			values[i] = new(sql.NullInt64)
		case lessonplansubject.FieldCreatedAt, lessonplansubject.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LessonPlanSubject fields.
func (lps *LessonPlanSubject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lessonplansubject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lps.ID = int64(value.Int64)
		case lessonplansubject.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				lps.CreatedAt = value.Time
			}
		case lessonplansubject.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				lps.UpdatedAt = value.Time
			}
		case lessonplansubject.FieldLessonPlanID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lesson_plan_id", values[i])
			} else if value.Valid {
				lps.LessonPlanID = value.Int64
			}
		case lessonplansubject.FieldSubjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field subject_id", values[i])
			} else if value.Valid {
				lps.SubjectID = value.Int64
			}
		default:
			lps.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LessonPlanSubject.
// This includes values selected through modifiers, order, etc.
func (lps *LessonPlanSubject) Value(name string) (ent.Value, error) {
	return lps.selectValues.Get(name)
}

// QueryLessonPlan queries the "lesson_plan" edge of the LessonPlanSubject entity.
func (lps *LessonPlanSubject) QueryLessonPlan() *LessonPlanQuery {
	return NewLessonPlanSubjectClient(lps.config).QueryLessonPlan(lps)
}

// QuerySubject queries the "subject" edge of the LessonPlanSubject entity.
func (lps *LessonPlanSubject) QuerySubject() *SubjectQuery {
	return NewLessonPlanSubjectClient(lps.config).QuerySubject(lps)
}

// Update returns a builder for updating this LessonPlanSubject.
// Note that you need to call LessonPlanSubject.Unwrap() before calling this method if this LessonPlanSubject
// was returned from a transaction, and the transaction was committed or rolled back.
func (lps *LessonPlanSubject) Update() *LessonPlanSubjectUpdateOne {
	return NewLessonPlanSubjectClient(lps.config).UpdateOne(lps)
}

// Unwrap unwraps the LessonPlanSubject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lps *LessonPlanSubject) Unwrap() *LessonPlanSubject {
	_tx, ok := lps.config.driver.(*txDriver)
	if !ok {
		panic("ent: LessonPlanSubject is not a transactional entity")
	}
	lps.config.driver = _tx.drv
	return lps
}

// String implements the fmt.Stringer.
func (lps *LessonPlanSubject) String() string {
	var builder strings.Builder
	builder.WriteString("LessonPlanSubject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", lps.ID))
	builder.WriteString("created_at=")
	builder.WriteString(lps.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(lps.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("lesson_plan_id=")
	builder.WriteString(fmt.Sprintf("%v", lps.LessonPlanID))
	builder.WriteString(", ")
	builder.WriteString("subject_id=")
	builder.WriteString(fmt.Sprintf("%v", lps.SubjectID))
	builder.WriteByte(')')
	return builder.String()
}

// LessonPlanSubjects is a parsable slice of LessonPlanSubject.
type LessonPlanSubjects []*LessonPlanSubject
