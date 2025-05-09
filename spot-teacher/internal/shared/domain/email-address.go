package domain

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type EmailAddress string

func NewEmailAddress(value string) (EmailAddress, error) {
	const maxLength = 200
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return "", errors.New("email address cannot be empty or only whitespace")
	}
	if utf8.RuneCountInString(trimmedValue) > maxLength {
		return "", fmt.Errorf("email address cannot exceed %d characters", maxLength)
	}
	return EmailAddress(value), nil
}

func (e EmailAddress) Value() string {
	return string(e)
}
