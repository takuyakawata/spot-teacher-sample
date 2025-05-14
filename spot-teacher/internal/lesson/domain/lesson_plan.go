package domain

import (
	"errors"
	compnayDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type LessonPlan struct {
	ID                 LessonPlanID
	CompanyID          compnayDomain.CompanyID
	Title              string
	Description        *string
	Grade              []domain.Grade
	Subject            []domain.Subject
	EducationCategory  []domain.EducationCategory
	StartDate          LessonPlanDate
	EndDate            LessonPlanDate
	LessonType         LessonType
	AnnualMaxExecution int
	StartTime          time.Time
	EndTime            time.Time
}

type LessonType string

const (
	LessonTypeOnline           LessonType = "online"
	LessonTypeOffline          LessonType = "offline"
	LessonTypeOnlineAndOffline LessonType = "online_and_offline"
)

type LessonPlanID int64

func NewPlanID(value int64) (LessonPlanID, error) {
	if value <= 0 {
		return 0, errors.New("product ID must be positive")
	}
	return LessonPlanID(value), nil
}
func (p LessonPlanID) Value() int64 {
	return int64(p)
}

type LessonPlanDate struct {
	month time.Month
	day   sharedDomain.Day
}

func NewLessonPlanDate(month time.Month, day sharedDomain.Day) LessonPlanDate {
	return LessonPlanDate{
		month: month,
		day:   day,
	}
}

func (d LessonPlanDate) Month() time.Month {
	return d.month
}

func (d LessonPlanDate) Day() sharedDomain.Day {
	return d.day
}
