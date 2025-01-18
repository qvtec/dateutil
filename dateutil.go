package dateutil

import "time"

// AddDays adds the specified number of days to a given time.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths adds the specified number of months to a given time.
func AddMonths(t time.Time, months int) time.Time {
	newDate := t.AddDate(0, months, 0)
	// If the resulting date does not exist (e.g., adding 1 month to January 31),
	// it will be adjusted to the last day of the target month (e.g., February 28 or 29).
	//
	// Example 1: AddMonths("2025-01-30", 1) → "2025-02-28" (since February 30 doesn't exist)
	// Example 2: AddMonths("2025-01-31", 1) → "2025-02-28" (non-leap year)
	//
	// Unlike Go's AddDate, which rolls over to the next month (e.g., "2023-01-31" + 1 month → "2023-02-28"),
	// this function truncates the date to the last valid day of the target month.
	if d := newDate.Day(); d != t.Day() {
		return newDate.AddDate(0, 0, -d)
	}
	return newDate
}

// AddYears adds the specified number of years to a given time.
func AddYears(t time.Time, years int) time.Time {
	return AddMonths(t, years * 12)
}
