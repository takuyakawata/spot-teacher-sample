package infra

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/enttest"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

func setupTest(t *testing.T) (*ent.Client, domain.SchoolRepository) {
	// Create an in-memory SQLite client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// Create a new repository with the test client
	repo := NewSchoolRepositoryImpl(client)

	return client, repo
}

func createTestSchool(t *testing.T) *domain.School {
	// Create test data
	schoolID, err := domain.NewSchoolID(1)
	require.NoError(t, err)

	schoolName, err := domain.NewSchoolName("Test School")
	require.NoError(t, err)

	email, err := sharedDomain.NewEmailAddress("test@example.com")
	require.NoError(t, err)

	phoneNumber, err := sharedDomain.NewPhoneNumber("1234567890")
	require.NoError(t, err)

	postCode, err := sharedDomain.NewPostCode("1234567")
	require.NoError(t, err)

	street := "Test Street"

	address := sharedDomain.Address{
		Prefecture: sharedDomain.PrefectureTokyo,
		City:       "Test City",
		Street:     &street,
		PostCode:   postCode,
	}

	url, err := sharedDomain.NewURL("https://example.com")
	require.NoError(t, err)

	// Create a new school
	school, err := domain.NewSchool(
		schoolID,
		domain.SchoolType("elementary"),
		schoolName,
		&email,
		phoneNumber,
		address,
		*url,
	)
	require.NoError(t, err)

	return school
}

func TestSchoolRepoImpl_Create(t *testing.T) {
	// Setup
	client, repo := setupTest(t)
	defer client.Close()

	// Create test data
	school := createTestSchool(t)

	// Test Create
	ctx := context.Background()
	createdSchool, err := repo.Create(ctx, school)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, createdSchool)
	assert.Equal(t, school.Name.Value(), createdSchool.Name.Value())
	assert.Equal(t, school.SchoolType, createdSchool.SchoolType)
	assert.Equal(t, school.Email.Value(), createdSchool.Email.Value())
	assert.Equal(t, school.PhoneNumber.Value(), createdSchool.PhoneNumber.Value())
	assert.Equal(t, school.Address.Prefecture, createdSchool.Address.Prefecture)
	assert.Equal(t, school.Address.City, createdSchool.Address.City)
	assert.Equal(t, *school.Address.Street, *createdSchool.Address.Street)
	assert.Equal(t, school.Address.PostCode.Value(), createdSchool.Address.PostCode.Value())
	assert.Equal(t, school.URL.String(), createdSchool.URL.String())
}

func TestSchoolRepoImpl_FindByID(t *testing.T) {
	// Setup
	client, repo := setupTest(t)
	defer client.Close()

	// Create test data
	school := createTestSchool(t)

	// Create a school first
	ctx := context.Background()
	createdSchool, err := repo.Create(ctx, school)
	require.NoError(t, err)

	// Test FindByID
	foundSchool, err := repo.FindByID(ctx, createdSchool.ID)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, foundSchool)
	assert.Equal(t, createdSchool.ID.Value(), foundSchool.ID.Value())
	assert.Equal(t, createdSchool.Name.Value(), foundSchool.Name.Value())
}

func TestSchoolRepoImpl_Update(t *testing.T) {
	// Setup
	client, repo := setupTest(t)
	defer client.Close()

	// Create test data
	school := createTestSchool(t)

	// Create a school first
	ctx := context.Background()
	createdSchool, err := repo.Create(ctx, school)
	require.NoError(t, err)

	// Update the school
	updatedName, err := domain.NewSchoolName("Updated School")
	require.NoError(t, err)

	createdSchool.Name = updatedName

	// Test Update
	updatedSchool, err := repo.Update(ctx, createdSchool)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, updatedSchool)
	assert.Equal(t, updatedName.Value(), updatedSchool.Name.Value())
}

func TestSchoolRepoImpl_Delete(t *testing.T) {
	// Setup
	client, repo := setupTest(t)
	defer client.Close()

	// Create test data
	school := createTestSchool(t)

	// Create a school first
	ctx := context.Background()
	createdSchool, err := repo.Create(ctx, school)
	require.NoError(t, err)

	// Test Delete
	err = repo.Delete(ctx, createdSchool.ID)
	require.NoError(t, err)

	// Verify the school is deleted
	_, err = repo.FindByID(ctx, createdSchool.ID)
	assert.Error(t, err)
}

func TestSchoolRepoImpl_FindAll(t *testing.T) {
	// Setup
	client, repo := setupTest(t)
	defer client.Close()

	// Create test data
	school1 := createTestSchool(t)

	// Create a second school with different data
	schoolID, err := domain.NewSchoolID(2)
	require.NoError(t, err)

	schoolName, err := domain.NewSchoolName("Second School")
	require.NoError(t, err)

	email, err := sharedDomain.NewEmailAddress("second@example.com")
	require.NoError(t, err)

	phoneNumber, err := sharedDomain.NewPhoneNumber("9876543210")
	require.NoError(t, err)

	postCode, err := sharedDomain.NewPostCode("7654321")
	require.NoError(t, err)

	street := "Second Street"

	address := sharedDomain.Address{
		Prefecture: sharedDomain.PrefectureOsaka,
		City:       "Second City",
		Street:     &street,
		PostCode:   postCode,
	}

	url, err := sharedDomain.NewURL("https://second-example.com")
	require.NoError(t, err)

	school2, err := domain.NewSchool(
		schoolID,
		domain.SchoolType("highSchool"),
		schoolName,
		&email,
		phoneNumber,
		address,
		*url,
	)
	require.NoError(t, err)

	// Create schools
	ctx := context.Background()
	_, err = repo.Create(ctx, school1)
	require.NoError(t, err)

	_, err = repo.Create(ctx, school2)
	require.NoError(t, err)

	// Test FindAll
	schools, err := repo.FindAll(ctx)

	// Assert
	require.NoError(t, err)
	assert.Len(t, schools, 2)
}
