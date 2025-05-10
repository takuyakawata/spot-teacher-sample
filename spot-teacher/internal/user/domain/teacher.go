package domain

import (
	"errors"
	"fmt"
	schoolDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"strings"
	"unicode/utf8"
)

type Teacher struct {
	ID          TeacherID
	SchoolID    schoolDomain.SchoolID
	FirstName   TeacherName
	FamilyName  TeacherName
	Email       sharedDomain.EmailAddress
	PhoneNumber *sharedDomain.PhoneNumber
	Password    sharedDomain.Password
}

func NewTeacher(
	id TeacherID,
	schoolID schoolDomain.SchoolID,
	firstName TeacherName,
	familyName TeacherName,
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

type TeacherName string

func NewTeacherName(value string) (TeacherName, error) {
	const maxLength = 50
	trimmedValue := strings.TrimSpace(value)
	if value == "" {
		return "", errors.New("product name cannot be empty or only whitespace")
	}
	if utf8.RuneCountInString(trimmedValue) > maxLength {
		return "", fmt.Errorf("teacher name cannot exceed %d characters", maxLength)
	}
	return TeacherName(value), nil
}

func (p TeacherName) Value() string {
	return string(p)
}
