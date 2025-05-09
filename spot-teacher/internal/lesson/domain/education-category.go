package domain

type EducationCategory string

const (
	elementary EducationCategory = "elementary"
	juniorHigh EducationCategory = "juniorHigh"
	highSchool EducationCategory = "highSchool"
)

func (ec EducationCategory) String() string {
	return string(ec)
}
