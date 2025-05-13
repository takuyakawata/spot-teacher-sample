package domain

import (
	"errors"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type AdminUser struct {
	ID         AdminUserID
	FirstName  sharedDomain.UserName
	FamilyName sharedDomain.UserName
	Email      sharedDomain.EmailAddress
	Password   sharedDomain.Password
	CreatedAt  time.Time
}

func NewAdminUser(
	id AdminUserID,
	firstName sharedDomain.UserName,
	familyName sharedDomain.UserName,
	email sharedDomain.EmailAddress,
	password sharedDomain.Password) *AdminUser {
	return &AdminUser{
		ID:         id,
		FirstName:  firstName,
		FamilyName: familyName,
		Email:      email,
		Password:   password,
	}
}

type AdminUserID int64

func NewAdminUserID(value int64) (AdminUserID, error) {
	if value <= 0 {
		return 0, errors.New("admin user ID must be positive")
	}
	return AdminUserID(value), nil
}

func (p AdminUserID) Value() int64 {
	return int64(p)
}
