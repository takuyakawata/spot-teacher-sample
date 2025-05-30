// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/school"
)

// School is the model entity for the School schema.
type School struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// SchoolType holds the value of the "school_type" field.
	SchoolType school.SchoolType `json:"school_type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Prefecture holds the value of the "prefecture" field.
	Prefecture int64 `json:"prefecture,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Street holds the value of the "street" field.
	Street string `json:"street,omitempty"`
	// PostCode holds the value of the "post_code" field.
	PostCode string `json:"post_code,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SchoolQuery when eager-loading is set.
	Edges        SchoolEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SchoolEdges holds the relations/edges for other nodes in the graph.
type SchoolEdges struct {
	// Teachers holds the value of the teachers edge.
	Teachers []*User `json:"teachers,omitempty"`
	// LessonReservations holds the value of the lesson_reservations edge.
	LessonReservations []*LessonReservation `json:"lesson_reservations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TeachersOrErr returns the Teachers value or an error if the edge
// was not loaded in eager-loading.
func (e SchoolEdges) TeachersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Teachers, nil
	}
	return nil, &NotLoadedError{edge: "teachers"}
}

// LessonReservationsOrErr returns the LessonReservations value or an error if the edge
// was not loaded in eager-loading.
func (e SchoolEdges) LessonReservationsOrErr() ([]*LessonReservation, error) {
	if e.loadedTypes[1] {
		return e.LessonReservations, nil
	}
	return nil, &NotLoadedError{edge: "lesson_reservations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*School) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case school.FieldID, school.FieldPrefecture:
			values[i] = new(sql.NullInt64)
		case school.FieldSchoolType, school.FieldName, school.FieldEmail, school.FieldPhoneNumber, school.FieldCity, school.FieldStreet, school.FieldPostCode, school.FieldURL:
			values[i] = new(sql.NullString)
		case school.FieldCreatedAt, school.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the School fields.
func (s *School) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case school.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case school.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case school.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case school.FieldSchoolType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field school_type", values[i])
			} else if value.Valid {
				s.SchoolType = school.SchoolType(value.String)
			}
		case school.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case school.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				s.Email = value.String
			}
		case school.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				s.PhoneNumber = value.String
			}
		case school.FieldPrefecture:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field prefecture", values[i])
			} else if value.Valid {
				s.Prefecture = value.Int64
			}
		case school.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				s.City = value.String
			}
		case school.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field street", values[i])
			} else if value.Valid {
				s.Street = value.String
			}
		case school.FieldPostCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_code", values[i])
			} else if value.Valid {
				s.PostCode = value.String
			}
		case school.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				s.URL = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the School.
// This includes values selected through modifiers, order, etc.
func (s *School) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryTeachers queries the "teachers" edge of the School entity.
func (s *School) QueryTeachers() *UserQuery {
	return NewSchoolClient(s.config).QueryTeachers(s)
}

// QueryLessonReservations queries the "lesson_reservations" edge of the School entity.
func (s *School) QueryLessonReservations() *LessonReservationQuery {
	return NewSchoolClient(s.config).QueryLessonReservations(s)
}

// Update returns a builder for updating this School.
// Note that you need to call School.Unwrap() before calling this method if this School
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *School) Update() *SchoolUpdateOne {
	return NewSchoolClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the School entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *School) Unwrap() *School {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: School is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *School) String() string {
	var builder strings.Builder
	builder.WriteString("School(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("school_type=")
	builder.WriteString(fmt.Sprintf("%v", s.SchoolType))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(s.Email)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(s.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("prefecture=")
	builder.WriteString(fmt.Sprintf("%v", s.Prefecture))
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(s.City)
	builder.WriteString(", ")
	builder.WriteString("street=")
	builder.WriteString(s.Street)
	builder.WriteString(", ")
	builder.WriteString("post_code=")
	builder.WriteString(s.PostCode)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(s.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Schools is a parsable slice of School.
type Schools []*School
