package domain

import (
	"errors"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"strings"
)

type school struct {
	ID          SchoolID
	SchoolType  SchoolType
	Name        SchoolName
	Email       *domain.EmailAddress
	PhoneNumber domain.PhoneNumber
	Address     domain.Address
	url         domain.URL
}

type SchoolID int64

func NewSchoolID(value int64) (SchoolID, error) {
	if value <= 0 {
		return 0, errors.New("product ID must be positive")
	}
	return SchoolID(value), nil
}
func (p SchoolID) Value() int64 {
	return int64(p)
}

type SchoolName string

func NewSchoolName(value string) (SchoolName, error) {
	const maxLength = 50
	trimmedValue := strings.TrimSpace(value)
	if value == "" {
		return "", errors.New("product name cannot be empty or only whitespace")
	}
	if len(trimmedValue) > maxLength {
		return "", errors.New("product name cannot be longer than 50 characters")
	}
	return SchoolName(value), nil
}
func (p SchoolName) Value() string {
	return string(p)
}

// SchoolType 学校の種類
type SchoolType string

const (
	elementary SchoolType = "elementary"
	juniorHigh SchoolType = "juniorHigh"
	highSchool SchoolType = "highSchool"
)

var schoolTypeNames = [...]string{
	"小学校", "中学校", "高校",
}
