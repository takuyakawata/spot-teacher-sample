// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// SchoolID applies equality check predicate on the "school_id" field. It's identical to SchoolIDEQ.
func SchoolID(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSchoolID, v))
}

// CompanyID applies equality check predicate on the "company_id" field. It's identical to CompanyIDEQ.
func CompanyID(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCompanyID, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// FamilyName applies equality check predicate on the "family_name" field. It's identical to FamilyNameEQ.
func FamilyName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFamilyName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhoneNumber, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserTypeEQ applies the EQ predicate on the "user_type" field.
func UserTypeEQ(v UserType) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserType, v))
}

// UserTypeNEQ applies the NEQ predicate on the "user_type" field.
func UserTypeNEQ(v UserType) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserType, v))
}

// UserTypeIn applies the In predicate on the "user_type" field.
func UserTypeIn(vs ...UserType) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserType, vs...))
}

// UserTypeNotIn applies the NotIn predicate on the "user_type" field.
func UserTypeNotIn(vs ...UserType) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserType, vs...))
}

// SchoolIDEQ applies the EQ predicate on the "school_id" field.
func SchoolIDEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSchoolID, v))
}

// SchoolIDNEQ applies the NEQ predicate on the "school_id" field.
func SchoolIDNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSchoolID, v))
}

// SchoolIDIn applies the In predicate on the "school_id" field.
func SchoolIDIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldSchoolID, vs...))
}

// SchoolIDNotIn applies the NotIn predicate on the "school_id" field.
func SchoolIDNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSchoolID, vs...))
}

// SchoolIDIsNil applies the IsNil predicate on the "school_id" field.
func SchoolIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldSchoolID))
}

// SchoolIDNotNil applies the NotNil predicate on the "school_id" field.
func SchoolIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldSchoolID))
}

// CompanyIDEQ applies the EQ predicate on the "company_id" field.
func CompanyIDEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCompanyID, v))
}

// CompanyIDNEQ applies the NEQ predicate on the "company_id" field.
func CompanyIDNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCompanyID, v))
}

// CompanyIDIn applies the In predicate on the "company_id" field.
func CompanyIDIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldCompanyID, vs...))
}

// CompanyIDNotIn applies the NotIn predicate on the "company_id" field.
func CompanyIDNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCompanyID, vs...))
}

// CompanyIDIsNil applies the IsNil predicate on the "company_id" field.
func CompanyIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCompanyID))
}

// CompanyIDNotNil applies the NotNil predicate on the "company_id" field.
func CompanyIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCompanyID))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFirstName, v))
}

// FamilyNameEQ applies the EQ predicate on the "family_name" field.
func FamilyNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFamilyName, v))
}

// FamilyNameNEQ applies the NEQ predicate on the "family_name" field.
func FamilyNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFamilyName, v))
}

// FamilyNameIn applies the In predicate on the "family_name" field.
func FamilyNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFamilyName, vs...))
}

// FamilyNameNotIn applies the NotIn predicate on the "family_name" field.
func FamilyNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFamilyName, vs...))
}

// FamilyNameGT applies the GT predicate on the "family_name" field.
func FamilyNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFamilyName, v))
}

// FamilyNameGTE applies the GTE predicate on the "family_name" field.
func FamilyNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFamilyName, v))
}

// FamilyNameLT applies the LT predicate on the "family_name" field.
func FamilyNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFamilyName, v))
}

// FamilyNameLTE applies the LTE predicate on the "family_name" field.
func FamilyNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFamilyName, v))
}

// FamilyNameContains applies the Contains predicate on the "family_name" field.
func FamilyNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFamilyName, v))
}

// FamilyNameHasPrefix applies the HasPrefix predicate on the "family_name" field.
func FamilyNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFamilyName, v))
}

// FamilyNameHasSuffix applies the HasSuffix predicate on the "family_name" field.
func FamilyNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFamilyName, v))
}

// FamilyNameEqualFold applies the EqualFold predicate on the "family_name" field.
func FamilyNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFamilyName, v))
}

// FamilyNameContainsFold applies the ContainsFold predicate on the "family_name" field.
func FamilyNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFamilyName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// HasSchool applies the HasEdge predicate on the "school" edge.
func HasSchool() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SchoolTable, SchoolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSchoolWith applies the HasEdge predicate on the "school" edge with a given conditions (other predicates).
func HasSchoolWith(preds ...predicate.School) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSchoolStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCompanyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInquiries applies the HasEdge predicate on the "inquiries" edge.
func HasInquiries() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InquiriesTable, InquiriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInquiriesWith applies the HasEdge predicate on the "inquiries" edge with a given conditions (other predicates).
func HasInquiriesWith(preds ...predicate.Inquiry) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newInquiriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLessonReservations applies the HasEdge predicate on the "lesson_reservations" edge.
func HasLessonReservations() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LessonReservationsTable, LessonReservationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLessonReservationsWith applies the HasEdge predicate on the "lesson_reservations" edge with a given conditions (other predicates).
func HasLessonReservationsWith(preds ...predicate.LessonReservation) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newLessonReservationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
