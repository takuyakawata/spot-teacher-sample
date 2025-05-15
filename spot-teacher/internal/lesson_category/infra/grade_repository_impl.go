package infra

import (
	"context"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	entgrade "github.com/takuyakawta/spot-teacher-sample/db/ent/grade"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
)

type gradeRepository struct{ client *ent.Client }

func NewGradeRepository(client *ent.Client) domain.GradeRepository {
	return &gradeRepository{client: client}
}

func (r *gradeRepository) Create(ctx context.Context, grade *domain.Grade) error {
	_, err := r.client.Grade.Query().Where(entgrade.CodeNumber(grade.Value().Value())).First(ctx)

	_, err = r.client.Grade.Create().
		SetName(grade.Value().Name()).
		SetCode(grade.Value().Name()). //　TODO いらんかも
		SetCodeNumber(grade.Value().Value()).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("infra.ent: failed to create grade: %w", err)
	}
	return nil
}

func (r *gradeRepository) RetrieveAll(ctx context.Context) ([]*domain.Grade, error) {
	entGrades, err := r.client.Grade.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to retrieve all grades: %w", err)
	}

	domainGrades := make([]*domain.Grade, 0, len(entGrades))
	for _, entGrade := range entGrades {
		domainGradeEnum := entGrade.CodeNumber
		domainGrade := domain.Grade(domainGradeEnum)
		domainGrades = append(domainGrades, &domainGrade)
	}

	return domainGrades, nil
}
