package infra

import (
	"context"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	entcategory "github.com/takuyakawta/spot-teacher-sample/db/ent/educationcategory"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
)

type educationCategoryRepository struct {
	client *ent.Client
}

func NewEducationCategoryRepository(client *ent.Client) domain.EducationCategoryRepository {
	return &educationCategoryRepository{client: client}
}

func (r *educationCategoryRepository) Create(ctx context.Context, educationCategory *domain.EducationCategory) (*domain.EducationCategory, error) {
	// Convert domain EducationCategory to string representation
	categoryEnum := domain.EducationCategoryEnum(*educationCategory)
	categoryStr := string(categoryEnum)

	// Check if education category already exists
	existingCategory, err := r.client.EducationCategory.Query().
		Where(entcategory.Name(categoryStr)).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("infra.ent: failed to query education category: %w", err)
	}

	// If education category already exists, return it
	if existingCategory != nil {
		domainCategoryEnum := domain.EducationCategoryEnum(existingCategory.Name)
		domainCategory, err := domain.NewEducationCategory(domainCategoryEnum)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to create domain education category from existing category: %w", err)
		}
		return &domainCategory, nil
	}

	// Create new education category
	createdCategory, err := r.client.EducationCategory.Create().
		SetName(categoryStr).
		SetCode(categoryStr). // Set code to the same value as name
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create education category: %w", err)
	}

	// Convert back to domain EducationCategory
	domainCategoryEnum := domain.EducationCategoryEnum(createdCategory.Name)
	domainCategory, err := domain.NewEducationCategory(domainCategoryEnum)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create domain education category from created category: %w", err)
	}
	return &domainCategory, nil
}

func (r *educationCategoryRepository) RetrieveAll(ctx context.Context) ([]*domain.EducationCategory, error) {
	entCategories, err := r.client.EducationCategory.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to retrieve all education categories: %w", err)
	}

	domainCategories := make([]*domain.EducationCategory, 0, len(entCategories))
	for _, entCategory := range entCategories {
		domainCategoryEnum := domain.EducationCategoryEnum(entCategory.Name)
		domainCategory, err := domain.NewEducationCategory(domainCategoryEnum)
		if err != nil {
			// Skip invalid categories but log the error
			fmt.Printf("infra.ent: failed to create domain education category from retrieved category: %v\n", err)
			continue
		}
		domainCategories = append(domainCategories, &domainCategory)
	}

	return domainCategories, nil
}
