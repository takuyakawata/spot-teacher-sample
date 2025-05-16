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
	company "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
	lessonCategory "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	shared "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type lessonPlanRepository struct{ client *ent.Client }

func NewLessonPlanRepository(client *ent.Client) domain.LessonPlanRepository {
	return &lessonPlanRepository{client: client}
}

func (r *lessonPlanRepository) Create(ctx context.Context, lessonPlan *domain.LessonPlan) (*domain.LessonPlan, error) {
	createCmd := r.client.LessonPlan.Create()
	createCmd.SetCompanyID(lessonPlan.CompanyID.Value())
	createCmd.SetTitle(lessonPlan.Title)
	if lessonPlan.Description != nil {
		createCmd.SetDescription(*lessonPlan.Description)
	}
	createCmd.SetStartMonth(int64(lessonPlan.StartDate.Month()))
	createCmd.SetStartDay(int64(lessonPlan.StartDate.Day()))
	createCmd.SetEndMonth(int64(lessonPlan.EndDate.Month()))
	createCmd.SetEndDay(int64(lessonPlan.EndDate.Day().Value()))
	createCmd.SetLessonType(lessonplan.LessonType(lessonPlan.LessonType))
	createCmd.SetAnnualMaxExecutions(lessonPlan.AnnualMaxExecution)
	createCmd.SetStartTime(lessonPlan.StartTime)
	createCmd.SetEndTime(lessonPlan.EndTime)

	// Save the lesson plan first
	createdEntLessonPlan, err := createCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create lesson plan: %w", err)
	}

	gradeCodeNumbers := make([]int64, 0, len(lessonPlan.Grade))
	for _, g := range lessonPlan.Grade {
		gradeEnum := lessonCategory.GradeEnum(g)
		gradeCodeNumbers = append(gradeCodeNumbers, gradeEnum.Value())
	}

	subjectNames := make([]string, 0, len(lessonPlan.Subject))
	for _, s := range lessonPlan.Subject {
		subjectNames = append(subjectNames, string(s))
	}

	categoryCodes := make([]string, 0, len(lessonPlan.EducationCategory))
	for _, c := range lessonPlan.EducationCategory {
		categoryCodes = append(categoryCodes, string(c))
	}

	// ---- ② 一括取得 ----
	upd := r.client.LessonPlan.UpdateOneID(createdEntLessonPlan.ID)

	if len(gradeCodeNumbers) > 0 {
		grades, err := r.client.Grade.Query().
			Where(grade.CodeNumberIn(gradeCodeNumbers...)).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: query grades: %w", err)
		}
		upd = upd.AddGrades(grades...)
	}
	if len(subjectNames) > 0 {
		subjects, err := r.client.Subject.Query().
			Where(subject.NameIn(subjectNames...)).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: query subjects: %w", err)
		}
		upd = upd.AddSubjects(subjects...)
	}
	if len(categoryCodes) > 0 {
		categories, err := r.client.EducationCategory.Query().
			Where(educationcategory.CodeIn(categoryCodes...)).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: query education categories: %w", err)
		}
		upd = upd.AddEducationCategories(categories...)
	}

	if _, err = upd.Save(ctx); err != nil {
		return nil, fmt.Errorf("infra.ent: link grades/subjects/categories: %w", err)
	}

	//for _, subjectValue := range lessonPlan.Subject {
	//	entSubject, err := r.client.Subject.Query().Where(subject.Name(string(subjectValue))).First(ctx)
	//	if err == nil {
	//		return nil, fmt.Errorf("infra.ent: failed to create subject: %w", err)
	//	}
	//
	//	_, err = r.client.LessonPlan.UpdateOneID(createdEntLessonPlan.ID).AddSubjects(entSubject).Save(ctx)
	//	if err != nil {
	//		return nil, fmt.Errorf("infra.ent: failed to add subject to lesson plan: %w", err)
	//	}
	//}
	//
	//for _ = range lessonPlan.EducationCategory {
	//	_, err = r.client.LessonPlan.UpdateOneID(createdEntLessonPlan.ID).
	//		AddEducationCategories().
	//		Save(ctx)
	//	if err != nil {
	//		return nil, fmt.Errorf("infra.ent: failed to add education category to lesson plan: %w", err)
	//	}
	//}

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
	// lesson plan
	primitiveID := lessonPlan.ID.Value.Value()
	updateCmd := r.client.LessonPlan.UpdateOneID(primitiveID)
	updateCmd.SetCompanyID(lessonPlan.CompanyID.Value())
	updateCmd.SetTitle(lessonPlan.Title)
	if lessonPlan.Description != nil {
		updateCmd.SetDescription(*lessonPlan.Description)
	}
	updateCmd.SetStartMonth(int64(lessonPlan.StartDate.Month()))
	updateCmd.SetStartDay(int64(lessonPlan.StartDate.Day()))
	updateCmd.SetEndMonth(int64(lessonPlan.EndDate.Month()))
	updateCmd.SetEndDay(int64(lessonPlan.EndDate.Day().Value()))
	updateCmd.SetLessonType(lessonplan.LessonType(lessonPlan.LessonType))
	_, err := updateCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to update lesson plan: %w", err)
	}

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
		gradeEnum := lessonCategory.GradeEnum(gradeValue)
		entGrade, err := r.client.Grade.Query().
			Where(grade.CodeNumber(gradeEnum.Value())).
			First(ctx) // TODO ここで見るべきか？？
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to query grade: %w", err)
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
			return nil, fmt.Errorf("infra.ent: failed to create subject: %w", err)
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
			return nil, fmt.Errorf("infra.ent: failed to create education category: %w", err)
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
		Where(lessonplan.ID(id.Value.Value())).
		WithGrades().
		WithSubjects().
		WithEducationCategories().
		First(ctx)

	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to find lesson plan by id %d: %w", id, err)
	}

	return mapEntLessonPlanToDomain(ctx, r.client, entLessonPlan)
}

func (r *lessonPlanRepository) FilterByCompanyID(ctx context.Context, companyID company.CompanyID) ([]*domain.LessonPlan, error) {
	entLessonPlans, err := r.client.LessonPlan.Query().
		Where(lessonplan.CompanyID(companyID.Value())).
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
	domainID, err := domain.NewLessonPlanID(entLP.ID)
	companyID := company.CompanyID(entLP.CompanyID)

	// Map Description
	var description *string
	if entLP.Description != "" {
		desc := entLP.Description
		description = &desc
	}

	// Map StartDate and EndDate
	startDay, err := shared.NewDay(int(entLP.StartDay))
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid start day: %w", err)
	}

	endDay, err := shared.NewDay(int(entLP.EndDay))
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid end day: %w", err)
	}

	startDate := domain.NewLessonPlanDate(time.Month(entLP.StartMonth), startDay)
	endDate := domain.NewLessonPlanDate(time.Month(entLP.EndMonth), endDay)

	// Map Grades
	grades := make([]lessonCategory.Grade, 0)
	if entLP.Edges.Grades != nil {
		for _, entGrade := range entLP.Edges.Grades {
			gradeCode := entGrade.CodeNumber
			grades = append(grades, lessonCategory.Grade(gradeCode))
		}
	}

	// Map Subjects
	subjects := make([]lessonCategory.Subject, 0)
	if entLP.Edges.Subjects != nil {
		for _, entSubject := range entLP.Edges.Subjects {
			s := lessonCategory.Subject(entSubject.Name)
			subjects = append(subjects, s)
		}
	}

	// Map EducationCategories
	educationCategories := make([]lessonCategory.EducationCategory, 0)
	if entLP.Edges.EducationCategories != nil {
		for _, entCategory := range entLP.Edges.EducationCategories {
			category := lessonCategory.EducationCategory(entCategory.Name)
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
