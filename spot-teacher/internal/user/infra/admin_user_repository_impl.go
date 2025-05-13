package infra

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/user"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"time"
)

type AdminUserRepositoryImpl struct {
	client *ent.Client
}

func NewAdminUserRepositoryImpl(client *ent.Client) domain.AdminUserRepository {
	return &AdminUserRepositoryImpl{client: client}
}

func (r *AdminUserRepositoryImpl) Create(ctx context.Context, a *domain.AdminUser) error {
	createCmd := r.client.User.Create()
	createCmd.SetFamilyName(a.FamilyName.Value())
	createCmd.SetFirstName(a.FirstName.Value())
	createCmd.SetEmail(a.Email.Value())
	createCmd.SetPassword(a.Password.Value())
	createCmd.SetPhoneNumber("") // Set empty phone number as it's required
	createCmd.SetCreatedAt(time.Now())
	createCmd.SetUpdatedAt(time.Now())
	createCmd.SetUserType("admin")
	_, err := createCmd.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *AdminUserRepositoryImpl) FindByID(ctx context.Context, id domain.AdminUserID) (*domain.AdminUser, error) {
	u, err := r.client.User.Query().
		Where(user.ID(int(id.Value())), user.UserTypeEQ("admin")).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	adminUser, err := toAdminUserEntity(u)
	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

func (r *AdminUserRepositoryImpl) FindByEmail(ctx context.Context, email sharedDomain.EmailAddress) (*domain.AdminUser, error) {
	u, err := r.client.User.Query().
		Where(user.Email(email.Value()), user.UserTypeEQ("admin")).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	adminUser, err := toAdminUserEntity(u)
	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

func toAdminUserEntity(user *ent.User) (*domain.AdminUser, error) {
	adminUserID, err := domain.NewAdminUserID(int64(user.ID))
	if err != nil {
		return nil, err
	}

	firstName, err := sharedDomain.NewUserName(user.FirstName)
	if err != nil {
		return nil, err
	}

	familyName, err := sharedDomain.NewUserName(user.FamilyName)
	if err != nil {
		return nil, err
	}

	email, err := sharedDomain.NewEmailAddress(user.Email)
	if err != nil {
		return nil, err
	}

	var password sharedDomain.Password
	if user.Password != nil {
		password, err = sharedDomain.NewPassword(*user.Password)
		if err != nil {
			return nil, err
		}
	}

	adminUser := domain.AdminUser{
		ID:         adminUserID,
		FirstName:  firstName,
		FamilyName: familyName,
		Email:      email,
		Password:   password,
		CreatedAt:  user.CreatedAt,
	}

	return &adminUser, nil
}
