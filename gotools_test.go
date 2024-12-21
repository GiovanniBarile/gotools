package gotools

import (
	"testing"
	"time"
)

func TestGetRandomDateOfBirth(t *testing.T) {
	age := 25

	// Generate a random date of birth
	dob := GetRandomDateOfBirth(age)

	// Get the current year
	currentYear := time.Now().Year()

	// Check if the year of the generated date matches the expected year
	expectedYear := currentYear - age
	if dob.Year() != expectedYear {
		t.Errorf("Expected year %d, got %d", expectedYear, dob.Year())
	}

	// Check if the month is valid (between 1 and 12)
	if dob.Month() < 1 || dob.Month() > 12 {
		t.Errorf("Invalid month: %d", dob.Month())
	}

	// Check if the day is valid for the given month and year
	maxDays := daysInMonth(dob.Year(), int(dob.Month()))
	if dob.Day() < 1 || dob.Day() > maxDays {
		t.Errorf("Invalid day: %d for month: %d, year: %d", dob.Day(), dob.Month(), dob.Year())
	}
}

func TestDaysInMonth(t *testing.T) {
	// Test cases for the daysInMonth function
	tests := []struct {
		year     int
		month    int
		expected int
	}{
		{2024, 2, 29}, // Leap year
		{2023, 2, 28}, // Non-leap year
		{2023, 1, 31}, // January
		{2023, 4, 30}, // April
	}

	for _, tt := range tests {
		got := daysInMonth(tt.year, tt.month)
		if got != tt.expected {
			t.Errorf("daysInMonth(%d, %d) = %d; want %d", tt.year, tt.month, got, tt.expected)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	// Test cases for the isLeapYear function
	tests := []struct {
		year     int
		expected bool
	}{
		{2024, true},  // Leap year
		{2023, false}, // Non-leap year
		{2000, true},  // Divisible by 400
		{1900, false}, // Divisible by 100 but not 400
	}

	for _, tt := range tests {
		got := isLeapYear(tt.year)
		if got != tt.expected {
			t.Errorf("isLeapYear(%d) = %v; want %v", tt.year, got, tt.expected)
		}
	}
}

func TestGetRandomDateOfBirthMultiple(t *testing.T) {
	// Generate multiple random dates to ensure all are valid
	for i := 0; i < 1000; i++ {
		dob := GetRandomDateOfBirth(25)

		// Check if the month is valid (between 1 and 12)
		if dob.Month() < 1 || dob.Month() > 12 {
			t.Errorf("Invalid month: %d", dob.Month())
		}

		// Check if the day is valid for the given month and year
		maxDays := daysInMonth(dob.Year(), int(dob.Month()))
		if dob.Day() < 1 || dob.Day() > maxDays {
			t.Errorf("Invalid day: %d for month: %d, year: %d", dob.Day(), dob.Month(), dob.Year())
		}
	}
}

func BenchmarkGetRandomDateOfBirth(b *testing.B) {
	// Benchmark the GetRandomDateOfBirth function
	for i := 0; i < b.N; i++ {
		GetRandomDateOfBirth(25)
	}
}
