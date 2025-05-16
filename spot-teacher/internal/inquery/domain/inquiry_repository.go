package domain

import "context"

type InquiryRepository interface {
	Create(ctx context.Context, inquiry *Inquiry) (*Inquiry, error)
}
