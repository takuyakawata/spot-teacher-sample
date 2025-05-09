package domain

type Grade int

const (
	elementaryOne Grade = iota
	elementaryTwo
	elementaryThree
	elementaryFour
	elementaryFive
	juniorHighOne
	juniorHighTwo
	juniorHighThree
	highSchoolOne
	highSchoolTwo
	highSchoolThree
)

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
