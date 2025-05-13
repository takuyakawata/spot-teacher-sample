package domain

import (
	compnayDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type LessonPlan struct {
	ID                LessonPlanID
	CompanyID         compnayDomain.CompanyID
	Title             string
	Description       *string
	Grade             []Grade
	Subject           []Subject
	EducationCategory []EducationCategory
	StartDate         LessonPlanDate
	EndDate           LessonPlanDate
}

type LessonPlanID int64
type LessonPlanDate struct {
	month time.Month
	day   sharedDomain.Day
}
