package domain

import (
	"errors"
	compnay "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	lessonCategory "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	shared "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/util"
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
	AnnualMaxExecution int64
	StartTime          time.Time
	EndTime            time.Time
}

type LessonType string

const (
	LessonTypeOnline           LessonType = "online"
	LessonTypeOffline          LessonType = "offline"
	LessonTypeOnlineAndOffline LessonType = "online_and_offline"
)

type LessonPlanID struct {
	Value util.ValueObject[int64]
}

func NewLessonPlanID(value int64) (LessonPlanID, error) {
	if value <= 0 {
		return LessonPlanID{}, errors.New("LessonPlanID must be positive")
	}
	return LessonPlanID{util.NewValueObject[int64](value)}, nil
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
