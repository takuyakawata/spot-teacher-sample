package infra

import (
	"context"
	"errors"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

type companyRepository struct {
	client *ent.Client
}

func NewCompanyRepository(client *ent.Client) domain.CompanyRepository {
	return &companyRepository{client: client}
}

func (r *companyRepository) Create(ctx context.Context, company *domain.Company) (*domain.Company, error) {
	createCmd := r.client.Company.Create()
	createCmd.SetName(company.Name.Value())

	// Address fields
	createCmd.SetPrefecture(int(company.Address.Prefecture))
	createCmd.SetCity(company.Address.City)
	createCmd.SetPostCode(company.Address.PostCode.Value())
	if company.Address.Street != nil {
		createCmd.SetStreet(*company.Address.Street)
	}

	// Phone number
	createCmd.SetPhoneNumber(company.PhoneNumber.Value())

	// URL (optional)
	urlStr := ""
	if company.URL != "" {
		urlStr = string(company.URL)
		createCmd.SetURL(urlStr)
	}

	createdEntCompany, err := createCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create company: %w", err)
	}
	return mapEntCompanyToDomain(createdEntCompany)
}

func (r *companyRepository) Update(ctx context.Context, company *domain.Company) (*domain.Company, error) {
	primitiveID := int(company.ID.Value())
	updateCmd := r.client.Company.UpdateOneID(primitiveID)
	updateCmd.SetName(company.Name.Value())

	// Address fields
	updateCmd.SetPrefecture(int(company.Address.Prefecture))
	updateCmd.SetCity(company.Address.City)
	updateCmd.SetPostCode(company.Address.PostCode.Value())
	if company.Address.Street != nil {
		updateCmd.SetStreet(*company.Address.Street)
	}

	// Phone number
	updateCmd.SetPhoneNumber(company.PhoneNumber.Value())

	// URL (optional)
	urlStr := ""
	if company.URL != "" {
		urlStr = string(company.URL)
		updateCmd.SetURL(urlStr)
	}

	updatedEntCompany, err := updateCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to update company: %w", err)
	}
	return mapEntCompanyToDomain(updatedEntCompany)
}

func (r *companyRepository) Delete(ctx context.Context, id domain.CompanyID) error {
	err := r.client.Company.DeleteOneID(int(id.Value())).Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("infra.ent: company to delete with id %v not found: %w", id, err)
		}
		return fmt.Errorf("infra.ent: failed to delete company with id %v: %w", id, err)
	}

	return nil
}

func (r *companyRepository) FindByID(ctx context.Context, id domain.CompanyID) (*domain.Company, error) {
	entCompany, err := r.client.Company.Get(ctx, int(id.Value()))
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to find company by id %d: %w", id, err)
	}
	return mapEntCompanyToDomain(entCompany)
}

func (r *companyRepository) FindAll(ctx context.Context) ([]*domain.Company, error) {
	companies, err := r.client.Company.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to find all companies: %w", err)
	}

	domainCompanies := make([]*domain.Company, 0, len(companies))
	for _, entC := range companies {
		domainC, mapErr := mapEntCompanyToDomain(entC)
		if mapErr != nil {
			return nil, fmt.Errorf("failed to map company (ent ID: %v) in FindAll: %w", entC.ID, mapErr)
		}
		domainCompanies = append(domainCompanies, domainC)
	}

	return domainCompanies, nil
}

func mapEntCompanyToDomain(entC *ent.Company) (*domain.Company, error) {
	if entC == nil {
		return nil, errors.New("infra.ent: cannot map nil ent.Company")
	}

	// ID の変換とバリデーション
	domainID, err := domain.NewCompanyID(int64(entC.ID))
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid id %d from db: %w", entC.ID, err)
	}

	// Name の変換とバリデーション
	domainName, err := domain.NewCompanyName(entC.Name)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid name '%s' from db (id: %d): %w", entC.Name, entC.ID, err)
	}

	// Address の作成
	postCode, err := sharedDomain.NewPostCode(entC.PostCode)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid post code '%s' from db (id: %d): %w", entC.PostCode, entC.ID, err)
	}

	var streetPtr *string
	if entC.Street != "" {
		streetValue := entC.Street
		streetPtr = &streetValue
	}

	address := sharedDomain.Address{
		Prefecture: sharedDomain.Prefecture(entC.Prefecture),
		City:       entC.City,
		Street:     streetPtr,
		PostCode:   postCode,
	}

	// PhoneNumber の作成
	phoneNumber := sharedDomain.PhoneNumber(entC.PhoneNumber)

	// URL の作成 (optional)
	var url sharedDomain.URL
	if entC.URL != "" {
		urlPtr, err := sharedDomain.NewURL(entC.URL)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: invalid URL '%s' from db (id: %d): %w", entC.URL, entC.ID, err)
		}
		if urlPtr != nil {
			url = *urlPtr
		}
	}

	return &domain.Company{
		ID:          domainID,
		Name:        domainName,
		Address:     address,
		PhoneNumber: phoneNumber,
		URL:         url,
	}, nil
}
