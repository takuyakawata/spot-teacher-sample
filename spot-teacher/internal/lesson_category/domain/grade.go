package domain

import "fmt"

type Grade GradeEnum

func NewGrade(value GradeEnum) (Grade, error) {
	if value < ElementaryOne || value > HighSchoolThree {
		return 0, fmt.Errorf("invalid grade value: %d", value)
	}

	return Grade(value), nil
}

type GradeEnum int

const (
	ElementaryOne GradeEnum = iota
	ElementaryTwo
	ElementaryThree
	ElementaryFour
	ElementaryFive
	ElementarySix
	JuniorHighOne
	JuniorHighTwo
	JuniorHighThree
	HighSchoolOne
	HighSchoolTwo
	HighSchoolThree
)

// 使わないかも　frontで定義すべきかな
var gradeNames = [...]string{
	"小学校1年生",
	"小学校2年生",
	"小学校3年生",
	"小学校4年生",
	"小学校5年生",
	"小学校6年生",
	"中学校1年生",
	"中学校2年生",
	"中学校3年生",
	"高校1年生",
	"高校2年生",
	"高校3年生",
}
