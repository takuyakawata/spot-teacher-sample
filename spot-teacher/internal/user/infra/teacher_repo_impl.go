package infra

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/user"
	schoolDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"time"
)

type TeacherRepoImpl struct {
	client *ent.Client
}

func NewTeacherRepoImpl(client *ent.Client) domain.TeacherRepository {
	return &TeacherRepoImpl{client: client}
}

func (r *TeacherRepoImpl) Create(ctx context.Context, t *domain.Teacher) error {
	createCmd := r.client.User.Create()
	createCmd.SetFamilyName(t.FamilyName.Value())
	createCmd.SetFirstName(t.FirstName.Value())
	createCmd.SetEmail(t.Email.Value())
	createCmd.SetPassword(t.Password.Value())
	//todo Phone Numberの追加
	createCmd.SetCreatedAt(time.Now())
	createCmd.SetUpdatedAt(time.Now())
	_, err := createCmd.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *TeacherRepoImpl) FindByID(ctx context.Context, id domain.TeacherID) (*domain.Teacher, error) {
	t, err := r.client.User.Get(ctx, int64(int(id.Value())))
	if err != nil {
		return nil, err
	}

	teacher, err := ToEntity(t)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (r *TeacherRepoImpl) FindByEmail(ctx context.Context, email sharedDomain.EmailAddress) (*domain.Teacher, error) {
	user, err := r.client.User.Query().Where(user.Email(email.Value())).Only(ctx)
	if err != nil {
		// Handle the case where no matching records are found or a general query error occurs
		return nil, err
	}

	teacher, err := ToEntity(user)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func ToEntity(user *ent.User) (*domain.Teacher, error) {
	var teacherPassword sharedDomain.Password
	if user.Password != nil {
		var err error
		teacherPassword, err = sharedDomain.NewPassword(*user.Password)
		if err != nil {
			return nil, err
		}
	}

	var teacherPhoneNumber *sharedDomain.PhoneNumber
	if user.PhoneNumber != "" {
		phoneNumber, err := sharedDomain.NewPhoneNumber(user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		teacherPhoneNumber = &phoneNumber
	}

	teacher := domain.Teacher{
		ID:          domain.TeacherID(user.ID),
		SchoolID:    schoolDomain.SchoolID(*user.SchoolID),
		FirstName:   domain.TeacherName(user.FirstName),
		FamilyName:  domain.TeacherName(user.FamilyName),
		Email:       sharedDomain.EmailAddress(user.Email),
		PhoneNumber: teacherPhoneNumber,
		Password:    teacherPassword,
	}
	return &teacher, nil
}
