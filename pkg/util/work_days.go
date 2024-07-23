package util

import (
	"time"
)

// isWeekend checks if the given date is a weekend (Saturday or Sunday)
func IsWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// countBusinessDays calculates the number of business days between two dates, excluding weekends
func CountBusinessDays(start, end time.Time) int {
	businessDays := 0
	for date := start; !date.After(end); date = date.AddDate(0, 0, 1) {
		if !IsWeekend(date) {
			businessDays++
		}
	}
	return businessDays
}
