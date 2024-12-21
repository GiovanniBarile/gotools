package gotools

import (
	"math/rand"
	"time"
)

// GetRandomDateOfBirth generates a random date of birth based on the provided age.
func GetRandomDateOfBirth(age int) time.Time {
	// Get the current date
	currentDate := time.Now()

	// Calculate the birth year
	birthYear := currentDate.Year() - age

	// Generate a random month (1 to 12)
	birthMonth := rand.Intn(12) + 1

	// Determine the maximum number of days in the generated month
	maxDays := daysInMonth(birthYear, birthMonth)

	// Generate a random day based on the max days
	birthDay := rand.Intn(maxDays) + 1

	// Create a time.Time object for the generated date
	return time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)
}

// daysInMonth returns the number of days in a specific month and year.
func daysInMonth(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		panic("invalid month")
	}
}

// isLeapYear checks if a year is a leap year.
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
