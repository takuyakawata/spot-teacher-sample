package domain

import "errors"

type Password string

func NewPassword(value string) (Password, error) {
	const MinLength = 8
	if value == "" {
		return "", errors.New("password cannot be nil")
	}

	if len(value) > MinLength {
		return "", errors.New("password cannot be longer than characters")
	}

	return Password(value), nil
}

func (p Password) Value() string {
	return string(p)
}
