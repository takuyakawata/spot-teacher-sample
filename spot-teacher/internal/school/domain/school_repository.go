package domain

import (
	"context"
)

// SchoolRepository defines the interface for school data operations
type SchoolRepository interface {
	// Create creates a new school
	Create(ctx context.Context, school *School) (*School, error)

	// Update updates an existing school
	Update(ctx context.Context, school *School) (*School, error)

	// Delete deletes a school by ID
	Delete(ctx context.Context, id SchoolID) error

	// FindByID finds a school by ID
	FindByID(ctx context.Context, id SchoolID) (*School, error)

	// FindAll retrieves all schools
	FindAll(ctx context.Context) ([]*School, error)

	FindByName(ctx context.Context, name SchoolName) (*School, error)
}
