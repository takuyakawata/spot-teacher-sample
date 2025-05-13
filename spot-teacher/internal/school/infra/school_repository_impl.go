package infra

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	entSchool "github.com/takuyakawta/spot-teacher-sample/db/ent/school"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"time"
)

type SchoolRepositoryImpl struct {
	client *ent.Client
}

func NewSchoolRepositoryImpl(client *ent.Client) domain.SchoolRepository {
	return &SchoolRepositoryImpl{client: client}
}

// Create creates a new school in the database
func (r *SchoolRepositoryImpl) Create(ctx context.Context, s *domain.School) (*domain.School, error) {
	// Start a transaction
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Rollback in case of error
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	// Create the school
	createCmd := tx.School.Create()
	createCmd.SetSchoolType(entSchool.SchoolType(string(s.SchoolType)))
	createCmd.SetName(s.Name.Value())

	if s.Email != nil {
		createCmd.SetEmail(s.Email.Value())
	}

	createCmd.SetPhoneNumber(s.PhoneNumber.Value())

	// Set address fields
	createCmd.SetPrefecture(int(s.Address.Prefecture))
	createCmd.SetCity(s.Address.City)

	if s.Address.Street != nil {
		createCmd.SetStreet(*s.Address.Street)
	}

	createCmd.SetPostCode(s.Address.PostCode.Value())

	// URL is a value type, not a pointer, so we check if it's empty
	if s.URL.String() != "" {
		createCmd.SetURL(s.URL.String())
	}

	createCmd.SetCreatedAt(time.Now())
	createCmd.SetUpdatedAt(time.Now())

	// Save the school
	schoolEnt, err := createCmd.Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Convert to domain entity
	school, err := ToEntity(schoolEnt)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// Update updates an existing school in the database
func (r *SchoolRepositoryImpl) Update(ctx context.Context, s *domain.School) (*domain.School, error) {
	// Start a transaction
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Rollback in case of error
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	// Update the school
	updateCmd := tx.School.UpdateOneID(int(s.ID.Value()))
	updateCmd.SetSchoolType(entSchool.SchoolType(string(s.SchoolType)))
	updateCmd.SetName(s.Name.Value())

	if s.Email != nil {
		updateCmd.SetEmail(s.Email.Value())
	} else {
		updateCmd.ClearEmail()
	}

	updateCmd.SetPhoneNumber(s.PhoneNumber.Value())

	// Update address fields
	updateCmd.SetPrefecture(int(s.Address.Prefecture))
	updateCmd.SetCity(s.Address.City)

	if s.Address.Street != nil {
		updateCmd.SetStreet(*s.Address.Street)
	} else {
		updateCmd.ClearStreet()
	}

	updateCmd.SetPostCode(s.Address.PostCode.Value())

	// URL is a value type, not a pointer, so we check if it's empty
	if s.URL.String() != "" {
		updateCmd.SetURL(s.URL.String())
	} else {
		updateCmd.ClearURL()
	}

	updateCmd.SetUpdatedAt(time.Now())

	// Save the school
	schoolEnt, err := updateCmd.Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Convert to domain entity
	school, err := ToEntity(schoolEnt)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// Delete deletes a school by ID
func (r *SchoolRepositoryImpl) Delete(ctx context.Context, id domain.SchoolID) error {
	return r.client.School.DeleteOneID(int(id.Value())).Exec(ctx)
}

// FindByID finds a school by ID
func (r *SchoolRepositoryImpl) FindByID(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	schoolEnt, err := r.client.School.Get(ctx, int(id.Value()))
	if err != nil {
		return nil, err
	}

	school, err := ToEntity(schoolEnt)
	if err != nil {
		return nil, err
	}

	return school, nil
}

// FindAll retrieves all schools
func (r *SchoolRepositoryImpl) FindAll(ctx context.Context) ([]*domain.School, error) {
	schoolEnts, err := r.client.School.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	schools := make([]*domain.School, len(schoolEnts))
	for i, schoolEnt := range schoolEnts {
		school, err := ToEntity(schoolEnt)
		if err != nil {
			return nil, err
		}
		schools[i] = school
	}

	return schools, nil
}

func (r *SchoolRepositoryImpl) FindByName(ctx context.Context, name domain.SchoolName) (*domain.School, error) {
	schoolEnt, err := r.client.School.Query().Where(entSchool.NameEQ(name.Value())).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ToEntity(schoolEnt)
}

// ToEntity converts an ent.School to a domain.School
func ToEntity(schoolEnt *ent.School) (*domain.School, error) {
	// Create SchoolID
	schoolID, err := domain.NewSchoolID(int64(schoolEnt.ID))
	if err != nil {
		return nil, err
	}

	// Create SchoolName
	schoolName, err := domain.NewSchoolName(schoolEnt.Name)
	if err != nil {
		return nil, err
	}

	// Create EmailAddress if present
	var emailAddress *sharedDomain.EmailAddress
	if schoolEnt.Email != "" {
		email, err := sharedDomain.NewEmailAddress(schoolEnt.Email)
		if err != nil {
			return nil, err
		}
		emailAddress = &email
	}

	// Create PhoneNumber
	phoneNumber, err := sharedDomain.NewPhoneNumber(schoolEnt.PhoneNumber)
	if err != nil {
		return nil, err
	}

	// Create PostCode
	postCode, err := sharedDomain.NewPostCode(schoolEnt.PostCode)
	if err != nil {
		return nil, err
	}

	// Create Street pointer if present
	var street *string
	if schoolEnt.Street != "" {
		streetStr := schoolEnt.Street
		street = &streetStr
	}

	// Create Address
	address := sharedDomain.Address{
		Prefecture: sharedDomain.Prefecture(schoolEnt.Prefecture),
		City:       schoolEnt.City,
		Street:     street,
		PostCode:   postCode,
	}

	// Create URL
	urlPtr, err := sharedDomain.NewURL(schoolEnt.URL)
	if err != nil {
		return nil, err
	}

	// Convert URL pointer to value (or empty URL if nil)
	var urlValue sharedDomain.URL
	if urlPtr != nil {
		urlValue = *urlPtr
	}

	// Create School
	school, err := domain.NewSchool(
		schoolID,
		domain.SchoolType(schoolEnt.SchoolType),
		schoolName,
		emailAddress,
		phoneNumber,
		address,
		urlValue,
	)
	if err != nil {
		return nil, err
	}

	return school, nil
}
