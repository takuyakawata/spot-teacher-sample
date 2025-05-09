package domain

import (
	"fmt"
)

type Day int

func NewDay(value int) (Day, error) {
	const minDay = 1
	const maxDay = 31
	if value < minDay || maxDay < value {
		return 0, fmt.Errorf("day must be between %d and %d", minDay, maxDay)
	}
	return Day(value), nil
}
func (d Day) Value() int { return int(d) }
