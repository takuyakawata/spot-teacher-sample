package domain

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type UserName string

func NewUserName(value string) (UserName, error) {
	const maxLength = 50
	trimmedValue := strings.TrimSpace(value)
	if value == "" {
		return "", errors.New("product name cannot be empty or only whitespace")
	}
	if utf8.RuneCountInString(trimmedValue) > maxLength {
		return "", fmt.Errorf("name cannot exceed %d characters", maxLength)
	}
	return UserName(value), nil
}

func (p UserName) Value() string {
	return string(p)
}
