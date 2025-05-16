package domain

import (
	domain2 "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	domain3 "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"time"
)

type Inquiry struct {
	ID               InquiryID
	SchoolID         domain.SchoolID
	TeacherID        domain3.TeacherID
	LessonScheduleID domain2.LessonScheduleID
	Detail           InquiryDetail
	Status           InquiryStatus
	Category         InquiryCategory
	CreatedAt        time.Time
}

type InquiryID int64

type InquiryStatus string

const (
	InquiryStatusPending  InquiryStatus = "pending"
	InquiryStatusAccepted InquiryStatus = "accepted"
	InquiryStatusRejected InquiryStatus = "rejected"
)

type InquiryCategory string

const (
	Lesson       InquiryCategory = "LESSON"
	Reservation  InquiryCategory = "RESERVATION"
	Cancellation InquiryCategory = "CANCELLATION"
	Other        InquiryCategory = "OTHER"
)

func (c InquiryCategory) Value() string {
	return string(c)
}

type InquiryDetail string

func (d InquiryDetail) Value() string {
	return string(d)
}
