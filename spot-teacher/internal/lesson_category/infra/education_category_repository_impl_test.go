package infra_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/infra"
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

// setupTestClient creates an in-memory SQLite database for testing
func setupTestClient(t *testing.T) *ent.Client {
	// SQLite のインメモリ DSN - use unique database name for each test to ensure isolation
	dbName := "file:ent_" + t.Name() + "?mode=memory&_fk=1"
	drv, err := sql.Open("sqlite3", dbName)
	require.NoError(t, err)

	c := ent.NewClient(ent.Driver(drv))
	// スキーマ作成
	require.NoError(t, c.Schema.Create(context.Background()))
	return c
}

func TestEducationCategoryRepository_Create(t *testing.T) {
	entClient := setupTestClient(t)
	repo := infra.NewEducationCategoryRepository(entClient)

	// Create a valid education category
	category, err := domain.NewEducationCategory(domain.SdgsEdu)
	require.NoError(t, err)
	require.NotNil(t, category)

	// Execute Create
	createdCategory, err := repo.Create(context.Background(), &category)
	require.NoError(t, err)
	assert.NotNil(t, createdCategory)
	assert.Equal(t, domain.SdgsEdu, domain.EducationCategoryEnum(*createdCategory))

	// Test creating the same education category again (should return existing)
	duplicateCategory, err := repo.Create(context.Background(), &category)
	require.NoError(t, err)
	assert.NotNil(t, duplicateCategory)
	assert.Equal(t, domain.SdgsEdu, domain.EducationCategoryEnum(*duplicateCategory))

	// Create another education category
	anotherCategory, err := domain.NewEducationCategory(domain.InfoEdu)
	require.NoError(t, err)
	require.NotNil(t, anotherCategory)

	// Execute Create for another education category
	createdAnotherCategory, err := repo.Create(context.Background(), &anotherCategory)
	require.NoError(t, err)
	assert.NotNil(t, createdAnotherCategory)
	assert.Equal(t, domain.InfoEdu, domain.EducationCategoryEnum(*createdAnotherCategory))
}

func TestEducationCategoryRepository_RetrieveAll(t *testing.T) {
	entClient := setupTestClient(t)
	repo := infra.NewEducationCategoryRepository(entClient)

	// Initially, there should be no education categories
	categories, err := repo.RetrieveAll(context.Background())
	require.NoError(t, err)
	assert.Empty(t, categories)

	// Create some education categories
	categoryEnums := []domain.EducationCategoryEnum{domain.SdgsEdu, domain.InfoEdu, domain.IctEdu}
	for _, ce := range categoryEnums {
		category, err := domain.NewEducationCategory(ce)
		require.NoError(t, err)
		_, err = repo.Create(context.Background(), &category)
		require.NoError(t, err)
	}

	// Retrieve all education categories
	retrievedCategories, err := repo.RetrieveAll(context.Background())
	require.NoError(t, err)
	assert.Len(t, retrievedCategories, len(categoryEnums))

	// Verify all created education categories are retrieved
	retrievedEnums := make([]domain.EducationCategoryEnum, 0, len(retrievedCategories))
	for _, c := range retrievedCategories {
		retrievedEnums = append(retrievedEnums, domain.EducationCategoryEnum(*c))
	}

	for _, expected := range categoryEnums {
		found := false
		for _, actual := range retrievedEnums {
			if expected == actual {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected education category %s not found in retrieved categories", expected)
	}
}
