package infra

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/db/ent/inquiry"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/inquery/domain"
)

type InquiryRepoImpl struct {
	client *ent.Client
}

func NewInquiryRepoImpl(client *ent.Client) *InquiryRepoImpl {
	return &InquiryRepoImpl{client: client}
}

func (r *InquiryRepoImpl) Create(ctx context.Context, Inquiry *domain.Inquiry) (*domain.Inquiry, error) {
	createInquiryCmd := r.client.Inquiry.Create()
	createInquiryCmd.
		SetSchoolID(Inquiry.SchoolID.Value()).
		SetUserID(Inquiry.TeacherID.Value()).
		SetLessonScheduleID(Inquiry.LessonScheduleID.Value()).
		SetInquiryDetail(Inquiry.Detail.Value()).
		SetCategory(inquiry.Category(Inquiry.Category.Value()))
	_, err := createInquiryCmd.Save(ctx)
	if err != nil {
		return nil, err
	}
	return Inquiry, nil
}
