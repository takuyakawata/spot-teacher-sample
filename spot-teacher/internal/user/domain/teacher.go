package domain

import (
	"errors"
	schoolDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

type Teacher struct {
	ID          TeacherID
	SchoolID    schoolDomain.SchoolID
	FirstName   sharedDomain.UserName
	FamilyName  sharedDomain.UserName
	Email       sharedDomain.EmailAddress
	PhoneNumber *sharedDomain.PhoneNumber
	Password    sharedDomain.Password
}

func NewTeacher(
	id TeacherID,
	schoolID schoolDomain.SchoolID,
	firstName sharedDomain.UserName,
	familyName sharedDomain.UserName,
	email sharedDomain.EmailAddress,
	phoneNumber *sharedDomain.PhoneNumber,
	password sharedDomain.Password) *Teacher {
	return &Teacher{
		ID:          id,
		SchoolID:    schoolID,
		FirstName:   firstName,
		FamilyName:  familyName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}

type TeacherID int64

func NewTeacherID(value int64) (TeacherID, error) {
	if value <= 0 {
		return 0, errors.New("product ID must be positive")
	}

	return TeacherID(value), nil
}

func (p TeacherID) Value() int64 {
	return int64(p)
}
