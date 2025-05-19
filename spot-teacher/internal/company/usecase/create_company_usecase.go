package usecase

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
)

type CreateCompanyUseCase struct {
	repo domain.CompanyRepository
}

func NewCreateCompanyUseCase(repo domain.CompanyRepository) *CreateCompanyUseCase {
	return &CreateCompanyUseCase{repo: repo}
}

func (u *CreateCompanyUseCase) CreateCompany(ctx context.Context, company *domain.Company) error {
	// companyがすでにあったら？何で判定？name?電話番号？
	_, err := u.repo.Create(ctx, company)
	if err != nil {
		return err
	}
	return nil
}
