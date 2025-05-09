package domain

import (
	compnayDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type Plan struct {
	ID                PlanID
	CompanyID         compnayDomain.CompanyID
	Title             string
	Description       *string
	Grade             []Grade
	Subject           []Subject
	EducationCategory []EducationCategory
	StartDate         PlanDate
	EndDate           PlanDate
}

type PlanID int64
type PlanDate struct {
	month time.Month
	day   sharedDomain.Day
}
