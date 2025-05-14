package infra

import (
	"context"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	entgrade "github.com/takuyakawta/spot-teacher-sample/db/ent/grade"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
)

type gradeRepository struct {
	client *ent.Client
}

func NewGradeRepository(client *ent.Client) domain.GradeRepository {
	return &gradeRepository{client: client}
}

func (r *gradeRepository) Create(ctx context.Context, grade *domain.Grade) (*domain.Grade, error) {
	// Convert domain Grade to string representation
	gradeEnum := domain.GradeEnum(*grade)
	gradeStr := gradeToString(gradeEnum)

	// Check if grade already exists
	existingGrade, err := r.client.Grade.Query().
		Where(entgrade.Name(gradeStr)).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("infra.ent: failed to query grade: %w", err)
	}

	// If grade already exists, return it
	if existingGrade != nil {
		domainGradeEnum := stringToGradeEnum(existingGrade.Name)
		domainGrade := domain.Grade(domainGradeEnum)
		return &domainGrade, nil
	}

	// Create new grade
	createdGrade, err := r.client.Grade.Create().
		SetName(gradeStr).
		SetCode(fmt.Sprintf("GRADE_%d", gradeEnum)). // Set code based on the grade enum value
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create grade: %w", err)
	}

	// Convert back to domain Grade
	domainGradeEnum := stringToGradeEnum(createdGrade.Name)
	domainGrade := domain.Grade(domainGradeEnum)
	return &domainGrade, nil
}

func (r *gradeRepository) RetrieveAll(ctx context.Context) ([]*domain.Grade, error) {
	entGrades, err := r.client.Grade.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to retrieve all grades: %w", err)
	}

	domainGrades := make([]*domain.Grade, 0, len(entGrades))
	for _, entGrade := range entGrades {
		domainGradeEnum := stringToGradeEnum(entGrade.Name)
		domainGrade := domain.Grade(domainGradeEnum)
		domainGrades = append(domainGrades, &domainGrade)
	}

	return domainGrades, nil
}

// Helper function to convert GradeEnum to string
func gradeToString(gradeEnum domain.GradeEnum) string {
	switch gradeEnum {
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
	case domain.ElementarySix:
		return "小学校6年生"
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

// Helper function to convert string to GradeEnum
func stringToGradeEnum(gradeStr string) domain.GradeEnum {
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
	case "小学校6年生":
		return domain.ElementarySix
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
