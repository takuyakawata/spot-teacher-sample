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

type TeacherRepositoryImpl struct {
	client *ent.Client
}

func NewTeacherRepositoryImpl(client *ent.Client) domain.TeacherRepository {
	return &TeacherRepositoryImpl{client: client}
}

func (r *TeacherRepositoryImpl) Create(ctx context.Context, t *domain.Teacher) error {
	createCmd := r.client.User.Create()
	createCmd.SetSchoolID(t.SchoolID.Value())
	createCmd.SetFamilyName(t.FamilyName.Value())
	createCmd.SetFirstName(t.FirstName.Value())
	createCmd.SetEmail(t.Email.Value())
	createCmd.SetPassword(t.Password.Value())
	if t.PhoneNumber != nil {
		createCmd.SetPhoneNumber(t.PhoneNumber.Value())
	}
	createCmd.SetCreatedAt(time.Now())
	createCmd.SetUpdatedAt(time.Now())
	_, err := createCmd.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *TeacherRepositoryImpl) FindByID(ctx context.Context, id domain.TeacherID) (*domain.Teacher, error) {
	t, err := r.client.User.Get(ctx, int(id.Value()))
	if err != nil {
		return nil, err
	}

	teacher, err := ToEntity(t)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (r *TeacherRepositoryImpl) FindByEmail(ctx context.Context, email sharedDomain.EmailAddress) (*domain.Teacher, error) {
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

func (r *TeacherRepositoryImpl) FindBySchoolID(ctx context.Context, schoolID schoolDomain.SchoolID) ([]*domain.Teacher, error) {
	users, err := r.client.User.Query().Where(user.SchoolIDEQ(int(schoolID.Value()))).All(ctx)
	if err != nil {
		return nil, err
	}

	teachers := make([]*domain.Teacher, 0, len(users))
	for _, u := range users {
		teacher, err := ToEntity(u)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
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

	var schoolID schoolDomain.SchoolID
	if user.SchoolID != nil {
		schoolID = schoolDomain.SchoolID(*user.SchoolID)
	}

	firstName, err := sharedDomain.NewUserName(user.FirstName)
	if err != nil {
		return nil, err
	}

	familyName, err := sharedDomain.NewUserName(user.FamilyName)
	if err != nil {
		return nil, err
	}

	teacher := domain.Teacher{
		ID:          domain.TeacherID(user.ID),
		SchoolID:    schoolID,
		FirstName:   firstName,
		FamilyName:  familyName,
		Email:       sharedDomain.EmailAddress(user.Email),
		PhoneNumber: teacherPhoneNumber,
		Password:    teacherPassword,
	}
	return &teacher, nil
}
