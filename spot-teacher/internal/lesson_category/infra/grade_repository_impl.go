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
	codeNumberValue := grade.Value().Value()
	_, err := r.client.Grade.Query().Where(entgrade.CodeNumber(codeNumberValue)).First(ctx)
	// 検索結果のエラーをハンドリング
	if err != nil {
		if !ent.IsNotFound(err) {
			return fmt.Errorf("infra.ent: failed to query for existing grade with CodeNumber %d: %w", codeNumberValue, err)
		}
	} else {
		fmt.Printf("infra.ent: Grade with CodeNumber %d already exists, skipping creation.\n", codeNumberValue) // デバッグ用出力
		return nil
	}

	_, err = r.client.Grade.Create().
		SetCode(grade.Value().Name()). //　TODO いらんかも
		SetCodeNumber(codeNumberValue).
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
