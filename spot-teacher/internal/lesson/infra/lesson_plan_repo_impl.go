package infra

import (
	"context"
	"errors"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/educationcategory"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/grade"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/lessonplan"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/subject"
	companyDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type lessonPlanRepository struct {
	client *ent.Client
}

func NewLessonPlanRepository(client *ent.Client) domain.LessonPlanRepository {
	return &lessonPlanRepository{client: client}
}

func (r *lessonPlanRepository) Create(ctx context.Context, lessonPlan *domain.LessonPlan) (*domain.LessonPlan, error) {
	createCmd := r.client.LessonPlan.Create()
	createCmd.SetCompanyID(int(lessonPlan.CompanyID))
	createCmd.SetTitle(lessonPlan.Title)
	if lessonPlan.Description != nil {
		createCmd.SetDescription(*lessonPlan.Description)
	}
	createCmd.SetStartMonth(int(lessonPlan.StartDate.Month()))
	createCmd.SetStartDay(lessonPlan.StartDate.Day().Value())
	createCmd.SetEndMonth(int(lessonPlan.EndDate.Month()))
	createCmd.SetEndDay(lessonPlan.EndDate.Day().Value())
	createCmd.SetLessonType(lessonplan.LessonType(lessonPlan.LessonType))
	createCmd.SetAnnualMaxExecutions(lessonPlan.AnnualMaxExecution)
	createCmd.SetStartTime(lessonPlan.StartTime)
	createCmd.SetEndTime(lessonPlan.EndTime)

	// Save the lesson plan first
	createdEntLessonPlan, err := createCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create lesson plan: %w", err)
	}

	// Add grades
	for _, gradeValue := range lessonPlan.Grade {
		entGrade, err := r.client.Grade.Query().
			Where(grade.Name(gradeToString(gradeValue))).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Create the grade if it doesn't exist
				entGrade, err = r.client.Grade.Create().
					SetName(gradeToString(gradeValue)).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("infra.ent: failed to create grade: %w", err)
				}
			} else {
				return nil, fmt.Errorf("infra.ent: failed to query grade: %w", err)
			}
		}
		_, err = r.client.LessonPlan.UpdateOneID(createdEntLessonPlan.ID).
			AddGrades(entGrade).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to add grade to lesson plan: %w", err)
		}
	}

	// Add subjects
	for _, subjectValue := range lessonPlan.Subject {
		entSubject, err := r.client.Subject.Query().
			Where(subject.Name(string(subjectValue))).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Create the subject if it doesn't exist
				entSubject, err = r.client.Subject.Create().
					SetName(string(subjectValue)).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("infra.ent: failed to create subject: %w", err)
				}
			} else {
				return nil, fmt.Errorf("infra.ent: failed to query subject: %w", err)
			}
		}
		_, err = r.client.LessonPlan.UpdateOneID(createdEntLessonPlan.ID).
			AddSubjects(entSubject).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to add subject to lesson plan: %w", err)
		}
	}

	// Add education categories
	for _, categoryValue := range lessonPlan.EducationCategory {
		entCategory, err := r.client.EducationCategory.Query().
			Where(educationcategory.Name(string(categoryValue))).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Create the category if it doesn't exist
				entCategory, err = r.client.EducationCategory.Create().
					SetName(string(categoryValue)).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("infra.ent: failed to create education category: %w", err)
				}
			} else {
				return nil, fmt.Errorf("infra.ent: failed to query education category: %w", err)
			}
		}
		_, err = r.client.LessonPlan.UpdateOneID(createdEntLessonPlan.ID).
			AddEducationCategories(entCategory).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to add education category to lesson plan: %w", err)
		}
	}

	// Fetch the updated lesson plan with all relationships
	updatedEntLessonPlan, err := r.client.LessonPlan.Query().
		Where(lessonplan.ID(createdEntLessonPlan.ID)).
		WithGrades().
		WithSubjects().
		WithEducationCategories().
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to fetch updated lesson plan: %w", err)
	}

	return mapEntLessonPlanToDomain(ctx, r.client, updatedEntLessonPlan)
}

func (r *lessonPlanRepository) Update(ctx context.Context, lessonPlan *domain.LessonPlan) (*domain.LessonPlan, error) {
	primitiveID := int(lessonPlan.ID)
	updateCmd := r.client.LessonPlan.UpdateOneID(primitiveID)
	updateCmd.SetCompanyID(int(lessonPlan.CompanyID))
	updateCmd.SetTitle(lessonPlan.Title)
	if lessonPlan.Description != nil {
		updateCmd.SetDescription(*lessonPlan.Description)
	}
	updateCmd.SetStartMonth(int(lessonPlan.StartDate.Month()))
	updateCmd.SetStartDay(int(lessonPlan.StartDate.Day().Value()))
	updateCmd.SetEndMonth(int(lessonPlan.EndDate.Month()))
	updateCmd.SetEndDay(int(lessonPlan.EndDate.Day().Value()))
	updateCmd.SetLessonType(lessonplan.LessonType(lessonPlan.LessonType))

	// Update the basic fields first
	_, err := updateCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to update lesson plan: %w", err)
	}

	// Clear existing relationships
	_, err = r.client.LessonPlan.UpdateOneID(primitiveID).
		ClearGrades().
		ClearSubjects().
		ClearEducationCategories().
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to clear relationships: %w", err)
	}

	// Add grades
	for _, gradeValue := range lessonPlan.Grade {
		entGrade, err := r.client.Grade.Query().
			Where(grade.Name(gradeToString(gradeValue))).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Create the grade if it doesn't exist
				entGrade, err = r.client.Grade.Create().
					SetName(gradeToString(gradeValue)).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("infra.ent: failed to create grade: %w", err)
				}
			} else {
				return nil, fmt.Errorf("infra.ent: failed to query grade: %w", err)
			}
		}
		_, err = r.client.LessonPlan.UpdateOneID(primitiveID).
			AddGrades(entGrade).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to add grade to lesson plan: %w", err)
		}
	}

	// Add subjects
	for _, subjectValue := range lessonPlan.Subject {
		entSubject, err := r.client.Subject.Query().
			Where(subject.Name(string(subjectValue))).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Create the subject if it doesn't exist
				entSubject, err = r.client.Subject.Create().
					SetName(string(subjectValue)).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("infra.ent: failed to create subject: %w", err)
				}
			} else {
				return nil, fmt.Errorf("infra.ent: failed to query subject: %w", err)
			}
		}
		_, err = r.client.LessonPlan.UpdateOneID(primitiveID).
			AddSubjects(entSubject).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to add subject to lesson plan: %w", err)
		}
	}

	// Add education categories
	for _, categoryValue := range lessonPlan.EducationCategory {
		entCategory, err := r.client.EducationCategory.Query().
			Where(educationcategory.Name(string(categoryValue))).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Create the category if it doesn't exist
				entCategory, err = r.client.EducationCategory.Create().
					SetName(string(categoryValue)).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("infra.ent: failed to create education category: %w", err)
				}
			} else {
				return nil, fmt.Errorf("infra.ent: failed to query education category: %w", err)
			}
		}
		_, err = r.client.LessonPlan.UpdateOneID(primitiveID).
			AddEducationCategories(entCategory).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to add education category to lesson plan: %w", err)
		}
	}

	// Fetch the updated lesson plan with all relationships
	finalEntLessonPlan, err := r.client.LessonPlan.Query().
		Where(lessonplan.ID(primitiveID)).
		WithGrades().
		WithSubjects().
		WithEducationCategories().
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to fetch updated lesson plan: %w", err)
	}

	return mapEntLessonPlanToDomain(ctx, r.client, finalEntLessonPlan)
}

func (r *lessonPlanRepository) FindByID(ctx context.Context, id domain.LessonPlanID) (*domain.LessonPlan, error) {
	entLessonPlan, err := r.client.LessonPlan.Query().
		Where(lessonplan.ID(int(id))).
		WithGrades().
		WithSubjects().
		WithEducationCategories().
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("infra.ent: lesson plan with id %d not found: %w", id, err)
		}
		return nil, fmt.Errorf("infra.ent: failed to find lesson plan by id %d: %w", id, err)
	}

	return mapEntLessonPlanToDomain(ctx, r.client, entLessonPlan)
}

func (r *lessonPlanRepository) FilterByCompanyID(ctx context.Context, companyID companyDomain.CompanyID) ([]*domain.LessonPlan, error) {
	entLessonPlans, err := r.client.LessonPlan.Query().
		Where(lessonplan.CompanyID(int(companyID))).
		WithGrades().
		WithSubjects().
		WithEducationCategories().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to filter lesson plans by company id %d: %w", companyID, err)
	}

	domainLessonPlans := make([]*domain.LessonPlan, 0, len(entLessonPlans))
	for _, entLP := range entLessonPlans {
		domainLP, mapErr := mapEntLessonPlanToDomain(ctx, r.client, entLP)
		if mapErr != nil {
			return nil, fmt.Errorf("failed to map lesson plan (ent ID: %v) in FilterByCompanyID: %w", entLP.ID, mapErr)
		}
		domainLessonPlans = append(domainLessonPlans, domainLP)
	}

	return domainLessonPlans, nil
}

func mapEntLessonPlanToDomain(ctx context.Context, client *ent.Client, entLP *ent.LessonPlan) (*domain.LessonPlan, error) {
	if entLP == nil {
		return nil, errors.New("infra.ent: cannot map nil ent.LessonPlan")
	}

	// Map ID
	domainID := domain.LessonPlanID(entLP.ID)

	// Map CompanyID
	companyID := companyDomain.CompanyID(entLP.CompanyID)

	// Map Description
	var description *string
	if entLP.Description != "" {
		desc := entLP.Description
		description = &desc
	}

	// Map StartDate and EndDate
	startDay, err := sharedDomain.NewDay(entLP.StartDay)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid start day: %w", err)
	}

	endDay, err := sharedDomain.NewDay(entLP.EndDay)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid end day: %w", err)
	}

	startDate := domain.NewLessonPlanDate(time.Month(entLP.StartMonth), startDay)
	endDate := domain.NewLessonPlanDate(time.Month(entLP.EndMonth), endDay)

	// Map Grades
	grades := make([]domain.Grade, 0)
	if entLP.Edges.Grades != nil {
		for _, entGrade := range entLP.Edges.Grades {
			grade := stringToGrade(entGrade.Name)
			grades = append(grades, grade)
		}
	}

	// Map Subjects
	subjects := make([]domain.Subject, 0)
	if entLP.Edges.Subjects != nil {
		for _, entSubject := range entLP.Edges.Subjects {
			subject := domain.Subject(entSubject.Name)
			subjects = append(subjects, subject)
		}
	}

	// Map EducationCategories
	educationCategories := make([]domain.EducationCategory, 0)
	if entLP.Edges.EducationCategories != nil {
		for _, entCategory := range entLP.Edges.EducationCategories {
			category := domain.EducationCategory(entCategory.Name)
			educationCategories = append(educationCategories, category)
		}
	}

	return &domain.LessonPlan{
		ID:                 domainID,
		CompanyID:          companyID,
		Title:              entLP.Title,
		Description:        description,
		Grade:              grades,
		Subject:            subjects,
		EducationCategory:  educationCategories,
		StartDate:          startDate,
		EndDate:            endDate,
		LessonType:         domain.LessonType(entLP.LessonType),
		AnnualMaxExecution: entLP.AnnualMaxExecutions,
		StartTime:          entLP.StartTime,
		EndTime:            entLP.EndTime,
	}, nil
}

// Helper function to convert Grade enum to string
func gradeToString(grade domain.Grade) string {
	switch grade {
	case domain.ElementaryOne:
		return "小学校1年生"
	case domain.ElementaryTwo:
		return "小学校2年生"
	case domain.ElementaryThree:
		return "小学校3年生"
	case domain.ElementaryFour:
		return "小学校4年生"
	case domain.ElementaryFive:
		return "小学校5年生"
	case domain.JuniorHighOne:
		return "中学校1年生"
	case domain.JuniorHighTwo:
		return "中学校2年生"
	case domain.JuniorHighThree:
		return "中学校3年生"
	case domain.HighSchoolOne:
		return "高校1年生"
	case domain.HighSchoolTwo:
		return "高校2年生"
	case domain.HighSchoolThree:
		return "高校3年生"
	default:
		return "不明"
	}
}

// Helper function to convert string to Grade enum
func stringToGrade(gradeStr string) domain.Grade {
	switch gradeStr {
	case "小学校1年生":
		return domain.ElementaryOne
	case "小学校2年生":
		return domain.ElementaryTwo
	case "小学校3年生":
		return domain.ElementaryThree
	case "小学校4年生":
		return domain.ElementaryFour
	case "小学校5年生":
		return domain.ElementaryFive
	case "中学校1年生":
		return domain.JuniorHighOne
	case "中学校2年生":
		return domain.JuniorHighTwo
	case "中学校3年生":
		return domain.JuniorHighThree
	case "高校1年生":
		return domain.HighSchoolOne
	case "高校2年生":
		return domain.HighSchoolTwo
	case "高校3年生":
		return domain.HighSchoolThree
	default:
		return domain.ElementaryOne // Default value
	}
}
