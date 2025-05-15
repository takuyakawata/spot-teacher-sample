package fixture

import (
	company "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
	lessonCategory "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	shared "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

func BuildLessonPlan(companyID company.CompanyID) (*domain.LessonPlan, error) {
	// Create start and end dates
	startDay, err := shared.NewDay(1)
	if err != nil {
		return nil, err
	}
	endDay, err := shared.NewDay(31)
	if err != nil {
		return nil, err
	}

	startDate := domain.NewLessonPlanDate(time.January, startDay)
	endDate := domain.NewLessonPlanDate(time.December, endDay)
	description := "This is a test lesson plan"
	grades := []lessonCategory.Grade{
		lessonCategory.Grade(lessonCategory.ElementaryOne),
		lessonCategory.Grade(lessonCategory.JuniorHighTwo),
	}
	subjects := []lessonCategory.Subject{
		lessonCategory.Subject(lessonCategory.English),
		lessonCategory.Subject(lessonCategory.Science),
		lessonCategory.Subject(lessonCategory.Math),
	}
	educationCategories := []lessonCategory.EducationCategory{
		lessonCategory.EducationCategory(lessonCategory.SdgsEdu),
		lessonCategory.EducationCategory(lessonCategory.InfoEdu),
		lessonCategory.EducationCategory(lessonCategory.IctEdu),
	}

	return &domain.LessonPlan{
		CompanyID:          companyID,
		Title:              "Test Lesson Plan",
		Description:        &description,
		Grade:              grades,
		Subject:            subjects,
		EducationCategory:  educationCategories,
		StartDate:          startDate,
		EndDate:            endDate,
		LessonType:         domain.LessonTypeOnline,
		AnnualMaxExecution: 10,
	}, nil
}

// BuildLessonPlanWithCustomData creates a test lesson plan domain instance with custom data
//func BuildLessonPlanWithCustomData(
//	companyID company.CompanyID,
//	title string,
//	description *string,
//	grades []lessonCategory.Grade,
//	subjects []lessonCategory.Subject,
//	educationCategories []lessonCategory.EducationCategory,
//	startMonth time.Month,
//	startDay int,
//	endMonth time.Month,
//	endDay int,
//	lessonType domain.LessonType,
//	annualMaxExecution int,
//) (*domain.LessonPlan, error) {
//	// Create start and end dates
//	startDayObj, err := shared.NewDay(startDay)
//	if err != nil {
//		return nil, err
//	}
//	endDayObj, err := shared.NewDay(endDay)
//	if err != nil {
//		return nil, err
//	}
//
//	startDate := domain.NewLessonPlanDate(startMonth, startDayObj)
//	endDate := domain.NewLessonPlanDate(endMonth, endDayObj)
//
//	return &domain.LessonPlan{
//		CompanyID:          companyID,
//		Title:              title,
//		Description:        description,
//		Grade:              grades,
//		Subject:            subjects,
//		EducationCategory:  educationCategories,
//		StartDate:          startDate,
//		EndDate:            endDate,
//		LessonType:         lessonType,
//		AnnualMaxExecution: annualMaxExecution,
//	}, nil
//}

// CreateLessonPlan creates a test lesson plan in the database
//func CreateLessonPlan(client *ent.Client, companyID company.CompanyID) (*domain.LessonPlan, error) {
//	// Create a lesson plan domain instance
//	lessonPlan, err := BuildLessonPlan(companyID)
//	if err != nil {
//		return nil, err
//	}
//
//	// Create a repository
//	repo := infra.NewLessonPlanRepository(client)
//
//	// Save the lesson plan to the database
//	return repo.Create(context.Background(), lessonPlan)
//}
//
//// CreateLessonPlanWithCustomData creates a test lesson plan in the database with custom data
//func CreateLessonPlanWithCustomData(
//	client *ent.Client,
//	companyID company.CompanyID,
//	title string,
//	description *string,
//	grades []lessonCategory.Grade,
//	subjects []lessonCategory.Subject,
//	educationCategories []lessonCategory.EducationCategory,
//	startMonth time.Month,
//	startDay int,
//	endMonth time.Month,
//	endDay int,
//	lessonType domain.LessonType,
//	annualMaxExecution int,
//) (*domain.LessonPlan, error) {
//	// Create a lesson plan domain instance with custom data
//	lessonPlan, err := BuildLessonPlanWithCustomData(
//		companyID,
//		title,
//		description,
//		grades,
//		subjects,
//		educationCategories,
//		startMonth,
//		startDay,
//		endMonth,
//		endDay,
//		lessonType,
//		annualMaxExecution,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	// Create a repository
//	repo := infra.NewLessonPlanRepository(client)
//
//	// Save the lesson plan to the database
//	return repo.Create(context.Background(), lessonPlan)
//}
