package domain

import (
	"errors"
	"fmt"
)

type Subject SubjectEnum

var ErrEmptySubjectValue = errors.New("subject value cannot be empty")
var ErrUnknownSubjectEnumValue = errors.New("unknown subject enum value")

func NewSubject(value SubjectEnum) (Subject, error) {
	if value == "" {
		return "", ErrEmptySubjectValue
	}

	isValid := false
	for _, definedEnum := range AllSubjectEnums {
		if value == definedEnum {
			isValid = true
			break
		}
	}

	if !isValid {
		return "", fmt.Errorf("%w: '%s'", ErrUnknownSubjectEnumValue, string(value))
	}

	return Subject(value), nil
}

type SubjectEnum string

const (
	Japanese          SubjectEnum = "JAPANESE"
	Math              SubjectEnum = "MATH"
	Society           SubjectEnum = "SOCIETY" // ユーザーの SocialStudy はこれに統合、または別途定義
	Science           SubjectEnum = "SCIENCE"
	English           SubjectEnum = "ENGLISH"
	Music             SubjectEnum = "MUSIC"
	ArtAndCrafts      SubjectEnum = "ART_AND_CRAFTS"
	HomeEconomics     SubjectEnum = "HOME_ECONOMICS"
	PhysicalEducation SubjectEnum = "PHYSICAL_EDUCATION"
	LifeSkills        SubjectEnum = "LIFE_SKILLS"
	Ethics            SubjectEnum = "ETHICS"
	IntegratedStudies SubjectEnum = "INTEGRATED_STUDIES"
	SpecialActivities SubjectEnum = "SPECIAL_ACTIVITIES"
	ClubActivities    SubjectEnum = "CLUB_ACTIVITIES"
)

var subjectEnumToLabel = make(map[SubjectEnum]string)
var AllSubjectEnums []SubjectEnum

func init() {
	labels := map[SubjectEnum]string{
		Japanese:          "国語",
		Math:              "算数",
		Society:           "社会",
		Science:           "理科",
		English:           "英語",
		Music:             "音楽",
		ArtAndCrafts:      "図工",
		HomeEconomics:     "家庭科",
		PhysicalEducation: "体育",
		LifeSkills:        "生活",
		Ethics:            "道徳",
		IntegratedStudies: "総合的な学習",
		SpecialActivities: "特別活動",
		ClubActivities:    "クラブ活動",
	}

	AllSubjectEnums = make([]SubjectEnum, 0, len(labels))
	for se, label := range labels {
		subjectEnumToLabel[se] = label
		AllSubjectEnums = append(AllSubjectEnums, se)
	}
}
