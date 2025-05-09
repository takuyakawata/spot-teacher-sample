package domain

import (
	"errors"
	"fmt"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"strings"
	"time"
)

type Schedule struct {
	ID        ScheduleID
	PlanID    PlanID
	Name      ScheduleName
	StartDate ScheduleDate
	EndDate   ScheduleDate
	StartTime TimeOfDay
	EndTime   TimeOfDay
}

type ScheduleDate struct {
	year  int
	month time.Month
	day   sharedDomain.Day
}

type TimeOfDay time.Time

// NewTimeOfDay は指定された時と分の TimeOfDay を作成します。
// 無効な時や分（例: 25時、70分）の場合はエラーを返します。
func NewTimeOfDay(hour, minute int) (TimeOfDay, error) {
	// 時と分のバリデーション
	if hour < 0 || hour > 23 {
		return TimeOfDay{}, fmt.Errorf("invalid hour: %d (must be 0-23)", hour)
	}
	if minute < 0 || minute > 59 {
		return TimeOfDay{}, fmt.Errorf("invalid minute: %d (must be 0-59)", minute)
	}
	t := time.Date(0, time.January, 1, hour, minute, 0, 0, time.UTC)

	return TimeOfDay(t), nil
}
func (t TimeOfDay) Value() time.Time {
	return time.Time(t)
}

type ScheduleID int64

func NewScheduleID(value int64) (ScheduleID, error) {
	if value <= 0 {
		return 0, errors.New("product ID must be positive")
	}
	return ScheduleID(value), nil
}
func (p ScheduleID) Value() int64 {
	return int64(p)
}

type ScheduleName string

func NewScheduleName(value string) (ScheduleName, error) {
	const maxLength = 50
	trimmedValue := strings.TrimSpace(value)
	if value == "" {
		return "", errors.New("product name cannot be empty or only whitespace")
	}
	if len(trimmedValue) > maxLength {
		return "", errors.New("product name cannot be longer than 50 characters")
	}
	return ScheduleName(value), nil
}

func (p ScheduleName) Value() string {
	return string(p)
}
