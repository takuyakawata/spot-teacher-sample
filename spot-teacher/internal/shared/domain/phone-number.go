package domain

type PhoneNumber string

func NewPhoneNumber(value string) (PhoneNumber, error) {
	return PhoneNumber(value), nil
}

func (p PhoneNumber) Value() string {
	return string(p)
}
