package usecase

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
)

// SchoolUsecase defines the interface for school business logic operations
type SchoolUsecase interface {
	// ListSchools retrieves all schools
	ListSchools(ctx context.Context) ([]*domain.School, error)

	// GetSchool retrieves a school by ID
	GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error)

	// CreateSchool creates a new school
	CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error)

	// UpdateSchool updates an existing school
	UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error)

	// DeleteSchool deletes a school by ID
	// Returns an error if the school has associated teachers
	DeleteSchool(ctx context.Context, id domain.SchoolID) error
}
