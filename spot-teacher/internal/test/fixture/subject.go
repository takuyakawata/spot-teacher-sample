package fixture

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/infra"
	"log"
)

func CreateAllSubject(client *ent.Client) error {
	subjectEnums := domain.AllSubjectEnums
	subjectRepo := infra.NewSubjectRepository(client)

	for _, enum := range subjectEnums {
		subject, err := domain.NewSubject(enum)
		_, err = subjectRepo.Create(context.Background(), &subject)
		if err != nil {
			log.Printf("Error seeding subject '%s' (ID %d)", enum, err)
			return err
		}
	}
	return nil
}

//// GetAllSubjects returns all available Subject enum values
//func GetAllSubjects() []domain.Subject {
//	subjects := make([]domain.Subject, 0, len(domain.AllSubjectEnums))
//	for _, enumValue := range domain.AllSubjectEnums {
//		subject, _ := domain.NewSubject(enumValue)
//		subjects = append(subjects, subject)
//	}
//	return subjects
//}
//
//// GetSubjectByValue returns a Subject enum value by its string value
//func GetSubjectByValue(value domain.SubjectEnum) (domain.Subject, error) {
//	return domain.NewSubject(value)
//}
//
//// GetSubjectsByValues returns multiple Subject enum values by their string values
//func GetSubjectsByValues(values []domain.SubjectEnum) ([]domain.Subject, error) {
//	subjects := make([]domain.Subject, 0, len(values))
//	for _, value := range values {
//		subject, err := domain.NewSubject(value)
//		if err != nil {
//			return nil, err
//		}
//		subjects = append(subjects, subject)
//	}
//	return subjects, nil
//}
//
//// GetCoreSubjects returns core academic subjects (Japanese, Math, Science, Society, English)
//func GetCoreSubjects() []domain.Subject {
//	coreValues := []domain.SubjectEnum{
//		domain.Japanese,
//		domain.Math,
//		domain.Science,
//		domain.Society,
//		domain.English,
//	}
//
//	subjects, _ := GetSubjectsByValues(coreValues)
//	return subjects
//}
//
//// GetArtsSubjects returns arts-related subjects (Music, ArtAndCrafts)
//func GetArtsSubjects() []domain.Subject {
//	artsValues := []domain.SubjectEnum{
//		domain.Music,
//		domain.ArtAndCrafts,
//	}
//
//	subjects, _ := GetSubjectsByValues(artsValues)
//	return subjects
//}
