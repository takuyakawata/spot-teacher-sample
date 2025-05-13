package domain

import (
	"errors"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"strings"
	"unicode/utf8"
)

type Company struct {
	ID          CompanyID
	Name        CompanyName
	Address     domain.Address
	PhoneNumber domain.PhoneNumber
	URL         domain.URL
}

func NewCompany(
	id CompanyID,
	name CompanyName,
	address domain.Address,
	phoneNumber domain.PhoneNumber,
	url domain.URL,
) (*Company, error) {
	return &Company{
		ID:          id,
		Name:        name,
		Address:     address,
		PhoneNumber: phoneNumber,
		URL:         url,
	}, nil
}

type CompanyID int64

func NewCompanyID(value int64) (CompanyID, error) {
	if value <= 0 {
		return 0, errors.New("product ID must be positive")
	}
	return CompanyID(value), nil
}

func (p CompanyID) Value() int64 {
	return int64(p)
}

type CompanyName string

func NewCompanyName(value string) (CompanyName, error) {
	const maxLength = 50
	trimmedValue := strings.TrimSpace(value)
	if value == "" {
		return "", errors.New("product name cannot be empty or only whitespace")
	}
	if utf8.RuneCountInString(trimmedValue) > maxLength {
		return "", fmt.Errorf("product name cannot exceed %d characters", maxLength)
	}
	return CompanyName(value), nil
}

func (p CompanyName) Value() string {
	return string(p)
}
