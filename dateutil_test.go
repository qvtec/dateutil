package dateutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddDays(t *testing.T) {
	t.Run("3/1", func(t *testing.T) {
		t1 := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-02-28 00:00:00 +0000 UTC", AddDays(t1, -1).String())
	})
	t.Run("3/1 (leap year)", func(t *testing.T) {
		t1 := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2024-02-29 00:00:00 +0000 UTC", AddDays(t1, -1).String())
	})
}

func TestAddMonths(t *testing.T) {
	t.Run("1/31 next month", func(t *testing.T) {
		t1 := time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-02-28 00:00:00 +0000 UTC", AddMonths(t1, 1).String())
		assert.Equal(t, "2025-03-03 00:00:00 +0000 UTC", t1.AddDate(0, 1, 0).String()) // AddDate rolls over
	})
	t.Run("1/30 next month", func(t *testing.T) {
		t1 := time.Date(2025, 1, 30, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-02-28 00:00:00 +0000 UTC", AddMonths(t1, 1).String())
		assert.Equal(t, "2025-03-02 00:00:00 +0000 UTC", t1.AddDate(0, 1, 0).String()) // AddDate rolls over
	})
	t.Run("1/30 next month (leep year)", func(t *testing.T) {
		t1 := time.Date(2024, 1, 30, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2024-02-29 00:00:00 +0000 UTC", AddMonths(t1, 1).String())
		assert.Equal(t, "2024-03-01 00:00:00 +0000 UTC", t1.AddDate(0, 1, 0).String()) // AddDate rolls over
	})
	t.Run("2/28 next month (end of month)", func(t *testing.T) {
		t1 := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-03-28 00:00:00 +0000 UTC", AddMonths(t1, 1).String())
	})

	t.Run("1/31 last month", func(t *testing.T) {
		t1 := time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-04-30 00:00:00 +0000 UTC", AddMonths(t1, -1).String())
		assert.Equal(t, "2025-05-01 00:00:00 +0000 UTC", t1.AddDate(0, -1, 0).String()) // AddDate rolls over
	})
	t.Run("1/30 last month", func(t *testing.T) {
		t1 := time.Date(2025, 3, 29, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-02-28 00:00:00 +0000 UTC", AddMonths(t1, -1).String())
		assert.Equal(t, "2025-03-01 00:00:00 +0000 UTC", t1.AddDate(0, -1, 0).String()) // AddDate rolls over
	})
	t.Run("1/30 last month (leep year)", func(t *testing.T) {
		t1 := time.Date(2024, 3, 30, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2024-02-29 00:00:00 +0000 UTC", AddMonths(t1, -1).String())
		assert.Equal(t, "2024-03-01 00:00:00 +0000 UTC", t1.AddDate(0, -1, 0).String()) // AddDate rolls over
	})
	t.Run("2/28 last month (end of month)", func(t *testing.T) {
		t1 := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-01-28 00:00:00 +0000 UTC", AddMonths(t1, -1).String())
	})
}

func TestAddYears(t *testing.T) {
	t.Run("leap year", func(t *testing.T) {
		t1 := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, "2025-02-28 00:00:00 +0000 UTC", AddYears(t1, 1).String())
		assert.Equal(t, "2025-03-01 00:00:00 +0000 UTC", t1.AddDate(1, 0, 0).String()) // AddDate rolls over
	})
}

func TestFirstOfMonth(t *testing.T) {
	t1 := time.Date(2024, 5, 15, 12, 0, 0, 0, time.UTC)
	assert.Equal(t, "2024-05-01 00:00:00 +0000 UTC", FirstOfMonth(t1).String())
}

func TestEndOfMonth(t *testing.T) {
	t1 := time.Date(2024, 2, 15, 12, 0, 0, 0, time.UTC) // Leap year
	assert.Equal(t, "2024-02-29 23:59:59 +0000 UTC", EndOfMonth(t1).String())

	t2 := time.Date(2023, 4, 15, 12, 0, 0, 0, time.UTC) // Normal month
	assert.Equal(t, "2023-04-30 23:59:59 +0000 UTC", EndOfMonth(t2).String())
}
