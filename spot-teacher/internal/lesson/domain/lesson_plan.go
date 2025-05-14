package domain

import (
	"errors"
	compnay "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	lessonCategory "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	shared "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type LessonPlan struct {
	ID                 LessonPlanID
	CompanyID          compnay.CompanyID
	Title              string
	Description        *string
	Grade              []lessonCategory.Grade
	Subject            []lessonCategory.Subject
	EducationCategory  []lessonCategory.EducationCategory
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
	day   shared.Day
}

func NewLessonPlanDate(month time.Month, day shared.Day) LessonPlanDate {
	return LessonPlanDate{
		month: month,
		day:   day,
	}
}

func (d LessonPlanDate) Month() time.Month {
	return d.month
}

func (d LessonPlanDate) Day() shared.Day {
	return d.day
}
