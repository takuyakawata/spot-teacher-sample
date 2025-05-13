package domain

import (
	"errors"
	companyDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

type CompanyMember struct {
	ID          CompanyMemberID
	CompanyID   companyDomain.CompanyID
	firstName   sharedDomain.UserName
	FamilyName  sharedDomain.UserName
	Email       sharedDomain.EmailAddress
	PhoneNumber *sharedDomain.PhoneNumber
}

type CompanyMemberID int64

func NewCompanyMemberID(value int64) (CompanyMemberID, error) {
	if value <= 0 {
		return 0, errors.New("product ID must be positive")
	}
	return CompanyMemberID(value), nil
}

func (p CompanyMemberID) Value() int64 {
	return int64(p)
}
