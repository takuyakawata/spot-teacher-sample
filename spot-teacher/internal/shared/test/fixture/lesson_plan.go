package fixture

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	companyDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/infra"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

// BuildLessonPlan creates a test lesson plan domain instance
func BuildLessonPlan(companyID companyDomain.CompanyID) (*domain.LessonPlan, error) {
	// Create start and end dates
	startDay, err := sharedDomain.NewDay(1)
	if err != nil {
		return nil, err
	}
	endDay, err := sharedDomain.NewDay(31)
	if err != nil {
		return nil, err
	}

	startDate := domain.NewLessonPlanDate(time.January, startDay)
	endDate := domain.NewLessonPlanDate(time.December, endDay)

	// Create a description
	description := "This is a test lesson plan"

	return &domain.LessonPlan{
		CompanyID:          companyID,
		Title:              "Test Lesson Plan",
		Description:        &description,
		Grade:              []domain.Grade{domain.ElementaryOne, domain.ElementaryTwo},
		Subject:            []domain.Subject{domain.Subject("math"), domain.Subject("english")},
		EducationCategory:  []domain.EducationCategory{domain.EducationCategory("elementary")},
		StartDate:          startDate,
		EndDate:            endDate,
		LessonType:         domain.LessonTypeOnline,
		AnnualMaxExecution: 10,
	}, nil
}

// BuildLessonPlanWithCustomData creates a test lesson plan domain instance with custom data
func BuildLessonPlanWithCustomData(
	companyID companyDomain.CompanyID,
	title string,
	description *string,
	grades []domain.Grade,
	subjects []domain.Subject,
	educationCategories []domain.EducationCategory,
	startMonth time.Month,
	startDay int,
	endMonth time.Month,
	endDay int,
	lessonType domain.LessonType,
	annualMaxExecution int,
) (*domain.LessonPlan, error) {
	// Create start and end dates
	startDayObj, err := sharedDomain.NewDay(startDay)
	if err != nil {
		return nil, err
	}
	endDayObj, err := sharedDomain.NewDay(endDay)
	if err != nil {
		return nil, err
	}

	startDate := domain.NewLessonPlanDate(startMonth, startDayObj)
	endDate := domain.NewLessonPlanDate(endMonth, endDayObj)

	return &domain.LessonPlan{
		CompanyID:          companyID,
		Title:              title,
		Description:        description,
		Grade:              grades,
		Subject:            subjects,
		EducationCategory:  educationCategories,
		StartDate:          startDate,
		EndDate:            endDate,
		LessonType:         lessonType,
		AnnualMaxExecution: annualMaxExecution,
	}, nil
}

// CreateLessonPlan creates a test lesson plan in the database
func CreateLessonPlan(client *ent.Client, companyID companyDomain.CompanyID) (*domain.LessonPlan, error) {
	// Create a lesson plan domain instance
	lessonPlan, err := BuildLessonPlan(companyID)
	if err != nil {
		return nil, err
	}

	// Create a repository
	repo := infra.NewLessonPlanRepository(client)

	// Save the lesson plan to the database
	return repo.Create(context.Background(), lessonPlan)
}

// CreateLessonPlanWithCustomData creates a test lesson plan in the database with custom data
func CreateLessonPlanWithCustomData(
	client *ent.Client,
	companyID companyDomain.CompanyID,
	title string,
	description *string,
	grades []domain.Grade,
	subjects []domain.Subject,
	educationCategories []domain.EducationCategory,
	startMonth time.Month,
	startDay int,
	endMonth time.Month,
	endDay int,
	lessonType domain.LessonType,
	annualMaxExecution int,
) (*domain.LessonPlan, error) {
	// Create a lesson plan domain instance with custom data
	lessonPlan, err := BuildLessonPlanWithCustomData(
		companyID,
		title,
		description,
		grades,
		subjects,
		educationCategories,
		startMonth,
		startDay,
		endMonth,
		endDay,
		lessonType,
		annualMaxExecution,
	)
	if err != nil {
		return nil, err
	}

	// Create a repository
	repo := infra.NewLessonPlanRepository(client)

	// Save the lesson plan to the database
	return repo.Create(context.Background(), lessonPlan)
}
