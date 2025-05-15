package fixture

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/infra"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

// BuildSchool creates a test school domain instance
func BuildSchool() (*domain.School, error) {
	// Create email address
	email, err := sharedDomain.NewEmailAddress("school@example.com")
	if err != nil {
		return nil, err
	}

	// Create URL
	url, err := sharedDomain.NewURL("https://sample-school.com")
	if err != nil {
		return nil, err
	}

	// Create a school name
	schoolName, err := domain.NewSchoolName("Sample School")
	if err != nil {
		return nil, err
	}

	// Create a phone number
	phoneNumber, err := sharedDomain.NewPhoneNumber("123-456-7890")
	if err != nil {
		return nil, err
	}

	// Create a post code
	postCode, err := sharedDomain.NewPostCode("1234567")
	if err != nil {
		return nil, err
	}

	return domain.NewSchool(
		domain.SchoolID(0),              // ID will be assigned by the database
		domain.SchoolType("elementary"), // Using the string value directly
		schoolName,
		&email,
		phoneNumber,
		sharedDomain.Address{
			Prefecture: 1,
			City:       "Sample City",
			Street:     nil,
			PostCode:   postCode,
		},
		*url,
	)
}

// CreateSchool creates a test school in the database
func CreateSchool(client *ent.Client) (*domain.School, error) {
	// Create a school domain instance
	school, err := BuildSchool()
	if err != nil {
		return nil, err
	}

	// Create a repository
	repo := infra.NewSchoolRepositoryImpl(client)

	// Save the school to the database
	return repo.Create(context.Background(), school)
}
