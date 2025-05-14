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

// Use the setupTestDB function from the shared test package
func setupTestDB(t *testing.T) *ent.Client {
	// SQLite のインメモリ DSN - use unique database name for each test to ensure isolation
	dbName := "file:ent_" + t.Name() + "?mode=memory&_fk=1"
	drv, err := sql.Open("sqlite3", dbName)
	require.NoError(t, err)

	c := ent.NewClient(ent.Driver(drv))
	// スキーマ作成
	require.NoError(t, c.Schema.Create(context.Background()))
	return c
}

func TestGradeRepository_Create(t *testing.T) {
	entClient := setupTestDB(t)
	repo := infra.NewGradeRepository(entClient)

	// Create a valid grade
	grade, err := domain.NewGrade(domain.ElementaryOne)
	require.NoError(t, err)
	require.NotNil(t, grade)

	// Execute Create
	createdGrade, err := repo.Create(context.Background(), &grade)
	require.NoError(t, err)
	assert.NotNil(t, createdGrade)
	assert.Equal(t, domain.ElementaryOne, domain.GradeEnum(*createdGrade))

	// Test creating the same grade again (should return existing)
	duplicateGrade, err := repo.Create(context.Background(), &grade)
	require.NoError(t, err)
	assert.NotNil(t, duplicateGrade)
	assert.Equal(t, domain.ElementaryOne, domain.GradeEnum(*duplicateGrade))

	// Create another grade
	anotherGrade, err := domain.NewGrade(domain.JuniorHighTwo)
	require.NoError(t, err)
	require.NotNil(t, anotherGrade)

	// Execute Create for another grade
	createdAnotherGrade, err := repo.Create(context.Background(), &anotherGrade)
	require.NoError(t, err)
	assert.NotNil(t, createdAnotherGrade)
	assert.Equal(t, domain.JuniorHighTwo, domain.GradeEnum(*createdAnotherGrade))
}

func TestGradeRepository_RetrieveAll(t *testing.T) {
	entClient := setupTestDB(t)
	repo := infra.NewGradeRepository(entClient)

	// Initially, there should be no grades
	grades, err := repo.RetrieveAll(context.Background())
	require.NoError(t, err)
	assert.Empty(t, grades)

	// Create some grades
	gradeEnums := []domain.GradeEnum{domain.ElementaryOne, domain.JuniorHighTwo, domain.HighSchoolThree}
	for _, ge := range gradeEnums {
		grade, err := domain.NewGrade(ge)
		require.NoError(t, err)
		_, err = repo.Create(context.Background(), &grade)
		require.NoError(t, err)
	}

	// Retrieve all grades
	retrievedGrades, err := repo.RetrieveAll(context.Background())
	require.NoError(t, err)
	assert.Len(t, retrievedGrades, len(gradeEnums))

	// Verify all created grades are retrieved
	retrievedEnums := make([]domain.GradeEnum, 0, len(retrievedGrades))
	for _, g := range retrievedGrades {
		retrievedEnums = append(retrievedEnums, domain.GradeEnum(*g))
	}

	for _, expected := range gradeEnums {
		found := false
		for _, actual := range retrievedEnums {
			if expected == actual {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected grade %d not found in retrieved grades", expected)
	}
}
