package fixture

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/infra"
)

func CreateAllGrades(client *ent.Client) error {
	grades := domain.AllGradeEnums()
	gradeRepo := infra.NewGradeRepository(client)
	for _, enum := range grades {
		gradeToCreate := domain.Grade(enum)
		err := gradeRepo.Create(context.Background(), &gradeToCreate)
		if err != nil {
			return err
		}
	}
	return nil
}
