package infra_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/infra"
	lessonCategory "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/test/fixture"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
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

func TestLessonPlanRepository_Create(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewLessonPlanRepository(entClient)

	// Create a company in the database first
	company, err := fixture.CreateCompany(entClient)
	require.NoError(t, err)
	require.NotNil(t, company)
	companyID := company.ID

	// Use fixture to build a lesson plan
	lessonPlan, err := fixture.BuildLessonPlan(companyID)
	require.NoError(t, err)
	require.NotNil(t, lessonPlan)

	// Execute Create
	createdLessonPlan, err := repo.Create(context.Background(), lessonPlan)
	require.NoError(t, err)
	assert.NotNil(t, createdLessonPlan)
	assert.NotZero(t, createdLessonPlan.ID)
	assert.Equal(t, lessonPlan.Title, createdLessonPlan.Title)
	assert.Equal(t, *lessonPlan.Description, *createdLessonPlan.Description)
	assert.Equal(t, companyID, createdLessonPlan.CompanyID)
	assert.Equal(t, time.January, createdLessonPlan.StartDate.Month())
	assert.Equal(t, 1, createdLessonPlan.StartDate.Day().Value())
	assert.Equal(t, time.December, createdLessonPlan.EndDate.Month())
	assert.Equal(t, 31, createdLessonPlan.EndDate.Day().Value())

	// Check relationships
	assert.Len(t, createdLessonPlan.Grade, 2)
	assert.Len(t, createdLessonPlan.Subject, 2)
	assert.Len(t, createdLessonPlan.EducationCategory, 1)
}

func TestLessonPlanRepository_Update(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewLessonPlanRepository(entClient)

	// Create a company in the database first
	company, err := fixture.CreateCompany(entClient)
	require.NoError(t, err)
	require.NotNil(t, company)
	companyID := company.ID

	// Use fixture to build a lesson plan
	initialLessonPlan, err := fixture.BuildLessonPlan(companyID)
	require.NoError(t, err)
	require.NotNil(t, initialLessonPlan)

	// Create the initial lesson plan
	createdLessonPlan, err := repo.Create(context.Background(), initialLessonPlan)
	require.NoError(t, err)
	require.NotNil(t, createdLessonPlan)

	// Update data
	updatedTitle := "Updated Lesson Plan"
	updatedDescription := "This is an updated lesson plan"

	// Use fixture to build an updated lesson plan with custom data
	updatedLessonPlan, err := fixture.BuildLessonPlanWithCustomData(
		companyID,
		updatedTitle,
		&updatedDescription,
		[]lessonCategory.Grade{lessonCategory.Grade(lessonCategory.ElementaryTwo), lessonCategory.Grade(lessonCategory.ElementaryThree)},
		[]lessonCategory.Subject{lessonCategory.Subject("english"), lessonCategory.Subject("science")},
		[]lessonCategory.EducationCategory{lessonCategory.EducationCategory("elementary"), lessonCategory.EducationCategory("juniorHigh")},
		time.February,
		15,
		time.November,
		20,
		domain.LessonTypeOffline,
		5,
	)
	require.NoError(t, err)
	require.NotNil(t, updatedLessonPlan)

	// Set the ID from the created lesson plan
	updatedLessonPlan.ID = createdLessonPlan.ID

	// Execute Update
	result, err := repo.Update(context.Background(), updatedLessonPlan)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, updatedTitle, result.Title)
	assert.Equal(t, updatedDescription, *result.Description)
	assert.Equal(t, time.February, result.StartDate.Month())
	assert.Equal(t, 15, result.StartDate.Day().Value())
	assert.Equal(t, time.November, result.EndDate.Month())
	assert.Equal(t, 20, result.EndDate.Day().Value())

	// Check relationships
	assert.Len(t, result.Grade, 2)
	assert.Len(t, result.Subject, 2)
	assert.Len(t, result.EducationCategory, 2)
}

func TestLessonPlanRepository_FindByID(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewLessonPlanRepository(entClient)

	// Create a company in the database first
	company, err := fixture.CreateCompany(entClient)
	require.NoError(t, err)
	require.NotNil(t, company)
	companyID := company.ID

	// Use fixture to build a lesson plan
	lessonPlan, err := fixture.BuildLessonPlan(companyID)
	require.NoError(t, err)
	require.NotNil(t, lessonPlan)

	// Create the lesson plan
	createdLessonPlan, err := repo.Create(context.Background(), lessonPlan)
	require.NoError(t, err)
	require.NotNil(t, createdLessonPlan)

	// Execute FindByID
	foundLessonPlan, err := repo.FindByID(context.Background(), createdLessonPlan.ID)
	require.NoError(t, err)
	assert.NotNil(t, foundLessonPlan)
	assert.Equal(t, createdLessonPlan.ID, foundLessonPlan.ID)
	assert.Equal(t, lessonPlan.Title, foundLessonPlan.Title)
	assert.Equal(t, *lessonPlan.Description, *foundLessonPlan.Description)

	// Check relationships
	assert.Len(t, foundLessonPlan.Grade, 2)   // The fixture creates 2 grades
	assert.Len(t, foundLessonPlan.Subject, 2) // The fixture creates 2 subjects
	assert.Len(t, foundLessonPlan.EducationCategory, 1)
}

func TestLessonPlanRepository_FilterByCompanyID(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewLessonPlanRepository(entClient)

	// Create companies in the database first
	company1, err := fixture.CreateCompany(entClient)
	require.NoError(t, err)
	require.NotNil(t, company1)
	companyID1 := company1.ID

	// Create a second company
	company2, err := fixture.CreateCompany(entClient)
	require.NoError(t, err)
	require.NotNil(t, company2)
	companyID2 := company2.ID

	// Use fixture to build a lesson plan for company 1
	lessonPlan1, err := fixture.BuildLessonPlan(companyID1)
	require.NoError(t, err)
	require.NotNil(t, lessonPlan1)

	// Use fixture to build a custom lesson plan for company 2
	title2 := "Test Lesson Plan 2"
	description2 := "This is test lesson plan 2"
	lessonPlan2, err := fixture.BuildLessonPlanWithCustomData(
		companyID2,
		title2,
		&description2,
		[]lessonCategory.Grade{lessonCategory.Grade(lessonCategory.ElementaryTwo)},
		[]lessonCategory.Subject{lessonCategory.Subject("english")},
		[]lessonCategory.EducationCategory{lessonCategory.EducationCategory("elementary")},
		time.January,
		1,
		time.December,
		31,
		domain.LessonTypeOffline,
		5,
	)
	require.NoError(t, err)
	require.NotNil(t, lessonPlan2)

	// Create the lesson plans
	createdLessonPlan1, err := repo.Create(context.Background(), lessonPlan1)
	require.NoError(t, err)
	require.NotNil(t, createdLessonPlan1)

	createdLessonPlan2, err := repo.Create(context.Background(), lessonPlan2)
	require.NoError(t, err)
	require.NotNil(t, createdLessonPlan2)

	// Execute FilterByCompanyID for company 1
	lessonPlans1, err := repo.FilterByCompanyID(context.Background(), companyID1)
	require.NoError(t, err)
	assert.Len(t, lessonPlans1, 1)
	assert.Equal(t, lessonPlan1.Title, lessonPlans1[0].Title)

	// Execute FilterByCompanyID for company 2
	lessonPlans2, err := repo.FilterByCompanyID(context.Background(), companyID2)
	require.NoError(t, err)
	assert.Len(t, lessonPlans2, 1)
	assert.Equal(t, title2, lessonPlans2[0].Title)
}
