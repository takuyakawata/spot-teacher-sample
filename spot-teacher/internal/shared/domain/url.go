package domain

import (
	"fmt"
	"net/url"
	"strings"
)

type URL string

func NewURL(value string) (*URL, error) {
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return nil, nil
	}

	_, err := url.ParseRequestURI(trimmedValue)
	if err != nil {
		return nil, fmt.Errorf("invalid URL format: %w", err)
	}

	u := URL(trimmedValue)
	return &u, nil
}

func (u *URL) String() string {
	if u == nil {
		return ""
	}
	return string(*u)
}
