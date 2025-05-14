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

func setupInMemoryClient(t *testing.T) *ent.Client {
	// SQLite のインメモリ DSN - use unique database name for each test to ensure isolation
	dbName := "file:ent_" + t.Name() + "?mode=memory&_fk=1"
	drv, err := sql.Open("sqlite3", dbName)
	require.NoError(t, err)

	c := ent.NewClient(ent.Driver(drv))
	// スキーマ作成
	require.NoError(t, c.Schema.Create(context.Background()))
	return c
}

func TestSubjectRepository_Create(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewSubjectRepository(entClient)

	// Create a valid subject
	subject, err := domain.NewSubject(domain.Math)
	require.NoError(t, err)
	require.NotNil(t, subject)

	// Execute Create
	createdSubject, err := repo.Create(context.Background(), &subject)
	require.NoError(t, err)
	assert.NotNil(t, createdSubject)
	assert.Equal(t, domain.Math, domain.SubjectEnum(*createdSubject))

	// Test creating the same subject again (should return existing)
	duplicateSubject, err := repo.Create(context.Background(), &subject)
	require.NoError(t, err)
	assert.NotNil(t, duplicateSubject)
	assert.Equal(t, domain.Math, domain.SubjectEnum(*duplicateSubject))

	// Create another subject
	anotherSubject, err := domain.NewSubject(domain.Japanese)
	require.NoError(t, err)
	require.NotNil(t, anotherSubject)

	// Execute Create for another subject
	createdAnotherSubject, err := repo.Create(context.Background(), &anotherSubject)
	require.NoError(t, err)
	assert.NotNil(t, createdAnotherSubject)
	assert.Equal(t, domain.Japanese, domain.SubjectEnum(*createdAnotherSubject))
}

func TestSubjectRepository_RetrieveAll(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewSubjectRepository(entClient)

	// Initially, there should be no subjects
	subjects, err := repo.RetrieveAll(context.Background())
	require.NoError(t, err)
	assert.Empty(t, subjects)

	// Create some subjects
	subjectEnums := []domain.SubjectEnum{domain.Math, domain.Japanese, domain.Science}
	for _, se := range subjectEnums {
		subject, err := domain.NewSubject(se)
		require.NoError(t, err)
		_, err = repo.Create(context.Background(), &subject)
		require.NoError(t, err)
	}

	// Retrieve all subjects
	retrievedSubjects, err := repo.RetrieveAll(context.Background())
	require.NoError(t, err)
	assert.Len(t, retrievedSubjects, len(subjectEnums))

	// Verify all created subjects are retrieved
	retrievedEnums := make([]domain.SubjectEnum, 0, len(retrievedSubjects))
	for _, s := range retrievedSubjects {
		retrievedEnums = append(retrievedEnums, domain.SubjectEnum(*s))
	}

	for _, expected := range subjectEnums {
		found := false
		for _, actual := range retrievedEnums {
			if expected == actual {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected subject %s not found in retrieved subjects", expected)
	}
}
