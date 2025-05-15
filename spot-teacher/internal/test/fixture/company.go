package fixture

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/infra"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

// BuildCompany creates a test company domain instance
func BuildCompany() (*domain.Company, error) {
	// Create a company name
	companyName, err := domain.NewCompanyName("Sample Company")
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

	// Create a URL
	url, err := sharedDomain.NewURL("https://sample-company.com")
	if err != nil {
		return nil, err
	}

	return domain.NewCompany(
		domain.CompanyID(0), //回数増えたら増やしたい
		companyName,
		sharedDomain.Address{
			Prefecture: 1,
			City:       "Sample City",
			Street:     nil,
			PostCode:   postCode,
		},
		phoneNumber,
		*url,
	)
}

// CreateCompany creates a test company in the database
func CreateCompany(client *ent.Client) (*domain.Company, error) {
	// Create a company domain instance
	company, err := BuildCompany()
	if err != nil {
		return nil, err
	}

	// Create a repository
	repo := infra.NewCompanyRepository(client)

	// Save the company to the database
	return repo.Create(context.Background(), company)
}
