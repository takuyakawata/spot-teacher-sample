package fixture

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/infra"
)

//func GetAllEducationCategories() []domain.EducationCategory {
//	categories := make([]domain.EducationCategory, 0, len(domain.AllEducationCategoryEnums))
//	for _, enumValue := range domain.AllEducationCategoryEnums {
//		category, _ := domain.NewEducationCategory(enumValue)
//		categories = append(categories, category)
//	}
//	return categories
//}

//func GetEducationCategoriesByValues(values []domain.EducationCategoryEnum) ([]domain.EducationCategory, error) {
//	categories := make([]domain.EducationCategory, 0, len(values))
//	for _, value := range values {
//		category, err := domain.NewEducationCategory(value)
//		if err != nil {
//			return nil, err
//		}
//		categories = append(categories, category)
//	}
//	return categories, nil
//}

func CreateAllEducationCategories(client *ent.Client) error {
	categories := make([]domain.EducationCategory, 0, len(domain.AllEducationCategoryEnums))
	for _, enum := range categories {
		category := enum
		categories = append(categories, category)
		categoriesRepo := infra.NewEducationCategoryRepository(client)
		_, err := categoriesRepo.Create(context.Background(), &category)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
