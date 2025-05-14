package domain

import (
	"errors"
	"fmt"
)

type EducationCategory EducationCategoryEnum

var ErrEmptyEducationCategoryValue = errors.New("education category value cannot be empty")
var ErrUnknownEducationCategoryEnumValue = errors.New("unknown education category enum value")

func NewEducationCategory(value EducationCategoryEnum) (EducationCategory, error) {
	if value == "" {
		return "", ErrEmptyEducationCategoryValue
	}

	isValid := false
	for _, definedEnum := range AllEducationCategoryEnums {
		if value == definedEnum {
			isValid = true
			break
		}
	}

	if !isValid {
		return "", fmt.Errorf("%w: '%s'", ErrUnknownEducationCategoryEnumValue, string(value))
	}

	return EducationCategory(value), nil
}

type EducationCategoryEnum string

// TODO 後で増やす
const (
	SdgsEdu                       EducationCategoryEnum = "SDGS_EDU"
	InfoEdu                       EducationCategoryEnum = "INFO_EDU"
	IctEdu                        EducationCategoryEnum = "ICT_EDU"
	CopyrightEdu                  EducationCategoryEnum = "COPYRIGHT_EDU"
	NetLiteracyEdu                EducationCategoryEnum = "NET_LITERACY_EDU"
	NetMoralEdu                   EducationCategoryEnum = "NET_MORAL_EDU"
	ProgrammingEdu                EducationCategoryEnum = "PROGRAMMING_EDU"
	CareerEdu                     EducationCategoryEnum = "CAREER_EDU"
	EntrepreneurshipEdu           EducationCategoryEnum = "ENTREPRENEURSHIP_EDU"
	FinancialEdu                  EducationCategoryEnum = "FINANCIAL_EDU"
	FoodEdu                       EducationCategoryEnum = "FOOD_EDU"
	HealthEdu                     EducationCategoryEnum = "HEALTH_EDU"
	DisasterEdu                   EducationCategoryEnum = "DISASTER_EDU"
	LgbtEdu                       EducationCategoryEnum = "LGBT_EDU"
	GenderDiversityEdu            EducationCategoryEnum = "GENDER_DIVERSITY_EDU"
	SexualViolenceEdu             EducationCategoryEnum = "SEXUAL_VIOLENCE_EDU"
	SexualityEdu                  EducationCategoryEnum = "SEXUALITY_EDU"
	CancerEdu                     EducationCategoryEnum = "CANCER_EDU"
	ForeignLanguageEdu            EducationCategoryEnum = "FOREIGN_LANGUAGE_EDU"
	InternationalUnderstandingEdu EducationCategoryEnum = "INTERNATIONAL_UNDERSTANDING_EDU"
	VolunteerEdu                  EducationCategoryEnum = "VOLUNTEER_EDU"
	MulticulturalCoexistenceEdu   EducationCategoryEnum = "MULTICULTURAL_COEXISTENCE_EDU"
	PeaceEdu                      EducationCategoryEnum = "PEACE_EDU"
	ConsumerEdu                   EducationCategoryEnum = "CONSUMER_EDU"
	LocalEdu                      EducationCategoryEnum = "LOCAL_EDU"
	HumanRightsEdu                EducationCategoryEnum = "HUMAN_RIGHTS_EDU"
	SovereignEdu                  EducationCategoryEnum = "SOVEREIGN_EDU"
	SchoolLinkEdu                 EducationCategoryEnum = "SCHOOL_LINK_EDU"
	ArtEdu                        EducationCategoryEnum = "ART_EDU"
	SpecialSupportEdu             EducationCategoryEnum = "SPECIAL_SUPPORT_EDU"
	UniversalDesignEdu            EducationCategoryEnum = "UNIVERSAL_DESIGN_EDU"
	InclusiveEdu                  EducationCategoryEnum = "INCLUSIVE_EDU"
	EnvironmentEdu                EducationCategoryEnum = "ENVIRONMENT_EDU"
	SafetyEdu                     EducationCategoryEnum = "SAFETY_EDU"
	TrafficSafetyEdu              EducationCategoryEnum = "TRAFFIC_SAFETY_EDU"
	NatureExperienceEdu           EducationCategoryEnum = "NATURE_EXPERIENCE_EDU"
	WelfareEdu                    EducationCategoryEnum = "WELFARE_EDU"
	NormativeConsciousnessEdu     EducationCategoryEnum = "NORMATIVE_CONSCIOUSNESS_EDU"
	MoralEdu                      EducationCategoryEnum = "MORAL_EDU"
	MindEdu                       EducationCategoryEnum = "MIND_EDU"
	MuseumEdu                     EducationCategoryEnum = "MUSEUM_EDU"
	AnimalProtectionEdu           EducationCategoryEnum = "ANIMAL_PROTECTION_EDU"
	LibraryUseEdu                 EducationCategoryEnum = "LIBRARY_USE_EDU"
	NieEdu                        EducationCategoryEnum = "NIE_EDU"
	PostingEdu                    EducationCategoryEnum = "POSTING_EDU"
)

var AllEducationCategoryEnums []EducationCategoryEnum
var educationCategoryEnumToLabel = make(map[EducationCategoryEnum]string)

var educationCategoryDisplayNames = make(map[EducationCategoryEnum]string)

func init() {
	labels := map[EducationCategoryEnum]string{
		SdgsEdu:                       "SDGs教育",
		InfoEdu:                       "情報教育",
		IctEdu:                        "ICT教育",
		CopyrightEdu:                  "著作権教育",
		NetLiteracyEdu:                "ネットリテラシー教育",
		NetMoralEdu:                   "ネットモラル教育",
		ProgrammingEdu:                "プログラミング教育",
		CareerEdu:                     "キャリア教育",
		EntrepreneurshipEdu:           "起業家精神教育",
		FinancialEdu:                  "金融教育",
		FoodEdu:                       "食育",
		HealthEdu:                     "健康教育",
		DisasterEdu:                   "防災教育",
		LgbtEdu:                       "LGBT教育",
		GenderDiversityEdu:            "性の多様性教育",
		SexualViolenceEdu:             "性暴力に関する教育",
		SexualityEdu:                  "性教育",
		CancerEdu:                     "癌教育",
		ForeignLanguageEdu:            "外国語教育",
		InternationalUnderstandingEdu: "国際理解教育",
		VolunteerEdu:                  "ボランティア教育",
		MulticulturalCoexistenceEdu:   "多文化共生教育",
		PeaceEdu:                      "平和教育",
		ConsumerEdu:                   "消費者教育",
		LocalEdu:                      "郷土教育",
		HumanRightsEdu:                "人権教育",
		SovereignEdu:                  "主権者教育",
		SchoolLinkEdu:                 "小中連携教育",
		ArtEdu:                        "芸術教育",
		SpecialSupportEdu:             "特別支援教育",
		UniversalDesignEdu:            "ユニバーサルデザイン教育",
		InclusiveEdu:                  "インクルーシブ教育",
		EnvironmentEdu:                "環境教育",
		SafetyEdu:                     "安全教育",
		TrafficSafetyEdu:              "交通安全教育",
		NatureExperienceEdu:           "自然体験教育",
		WelfareEdu:                    "福祉教育",
		NormativeConsciousnessEdu:     "規範意識教育",
		MoralEdu:                      "道徳教育",
		MindEdu:                       "心の教育",
		MuseumEdu:                     "博物館教育",
		AnimalProtectionEdu:           "動物愛護教育",
		LibraryUseEdu:                 "図書館活用教育",
		NieEdu:                        "ＮＩＥ教育",
		PostingEdu:                    "掲示教育",
	}

	AllEducationCategoryEnums = make([]EducationCategoryEnum, 0, len(labels))
	for ec, label := range labels {
		educationCategoryEnumToLabel[ec] = label
		AllEducationCategoryEnums = append(AllEducationCategoryEnums, ec)
	}
}
