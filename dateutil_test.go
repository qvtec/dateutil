package dateutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddDays(t *testing.T) {
	t.Run("3/1", func(t *testing.T) {
		t1 := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		result := AddDays(t1, -1)
		assert.Equal(t, expected, result)
	})
	t.Run("3/1 (leap year)", func(t *testing.T) {
		t1 := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		result := AddDays(t1, -1)
		assert.Equal(t, expected, result)
	})
}

func TestAddMonths(t *testing.T) {
	t.Run("1/31 next month", func(t *testing.T) {
		t1 := time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, 1)
		assert.Equal(t, expected, result)
	})
	t.Run("1/30 next month", func(t *testing.T) {
		t1 := time.Date(2025, 1, 30, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, 1)
		assert.Equal(t, expected, result)
	})
	t.Run("1/30 next month (leep year)", func(t *testing.T) {
		t1 := time.Date(2024, 1, 30, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, 1)
		assert.Equal(t, expected, result)
	})
	t.Run("2/28 next month (end of month)", func(t *testing.T) {
		t1 := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 3, 28, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, 1)
		assert.Equal(t, expected, result)
	})

	t.Run("1/31 last month", func(t *testing.T) {
		t1 := time.Date(2025, 5, 31, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 4, 30, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, -1)
		assert.Equal(t, expected, result)
	})
	t.Run("1/30 last month", func(t *testing.T) {
		t1 := time.Date(2025, 3, 29, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, -1)
		assert.Equal(t, expected, result)
	})
	t.Run("1/30 last month (leep year)", func(t *testing.T) {
		t1 := time.Date(2024, 3, 30, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, -1)
		assert.Equal(t, expected, result)
	})
	t.Run("2/28 last month (end of month)", func(t *testing.T) {
		t1 := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 1, 28, 0, 0, 0, 0, time.UTC)
		result := AddMonths(t1, -1)
		assert.Equal(t, expected, result)
	})
}

func TestAddYears(t *testing.T) {
	t.Run("leap year", func(t *testing.T) {
		t1 := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)
		result := AddYears(t1, 1)
		assert.Equal(t, expected, result)
	})
}
